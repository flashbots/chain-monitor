package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/types"

	"go.uber.org/zap"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

type L2 struct {
	cfg *config.L2

	rpc    *ethclient.Client
	ticker *time.Ticker

	blockHeight  uint64
	blocks       *types.RingBuffer[bool]
	blocksLanded int64
	blocksSeen   int64
	builderAddr  ethcommon.Address
	reorgWindow  int
	wallets      map[string]ethcommon.Address
}

func newL2(cfg *config.L2) (*L2, error) {
	var (
		builderAddr ethcommon.Address
		wallets     = make(map[string]ethcommon.Address, len(cfg.WalletAddresses))
	)

	if cfg.BuilderAddress != "" { // builder address
		addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
		if err != nil {
			return nil, err
		}
		if len(addr) != 20 {
			return nil, fmt.Errorf(
				"invalid length for the builder address (want 20, got %d)",
				len(addr),
			)
		}
		copy(builderAddr[:], addr)
	}

	for name, addrStr := range cfg.WalletAddresses {
		var addr ethcommon.Address
		addrBytes, err := ethcommon.ParseHexOrString(addrStr)
		if err != nil {
			return nil, err
		}
		if len(addrBytes) != 20 {
			return nil, fmt.Errorf(
				"invalid length for the l2 wallet address (want 20, got %d)",
				len(addr),
			)
		}
		copy(addr[:], addrBytes)
		wallets[name] = addr
	}

	rpc, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, err
	}

	blockHeight, err := rpc.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return &L2{
		cfg: cfg,

		rpc:    rpc,
		ticker: time.NewTicker(cfg.BlockTime),

		blockHeight: blockHeight - 1,
		builderAddr: builderAddr,
		reorgWindow: int(cfg.ReorgWindow/cfg.BlockTime) + 1,
		wallets:     wallets,
	}, nil
}

func (l2 *L2) run(ctx context.Context) {
	if l2.builderAddr.Cmp(ethcommon.Address{}) != 0 {
		go func() {
			for {
				<-l2.ticker.C
				l2.processNewBlocks(ctx)
			}
		}()
	}
}

func (l2 *L2) stop() {
	l2.ticker.Stop()
}

func (l2 *L2) processNewBlocks(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	blockHeight, err := l2.rpc.BlockNumber(ctx)
	if err != nil {
		l.Error("Failed to get block height, skipping this round...",
			zap.Error(err),
			zap.String("rpc", l2.cfg.RPC),
		)
		return
	}

	if blockHeight == l2.blockHeight {
		return
	}

	if blockHeight < l2.blockHeight {
		l2.processReorg(ctx, blockHeight)
	}

	for b := l2.blockHeight + 1; b <= blockHeight; b++ {
		if err := l2.processBlock(ctx, b); err != nil {
			l.Error("Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block", blockHeight),
				zap.String("rpc", l2.cfg.RPC),
			)
			return
		}
		l2.blockHeight = b
	}
}

func (l2 *L2) processBlock(ctx context.Context, blockNumber uint64) error {
	l := logutils.LoggerFromContext(ctx)

	block, err := l2.rpc.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	if l2.blocks == nil {
		l2.blocks = types.NewRingBuffer[bool](int(blockNumber), l2.reorgWindow)
	}

	l2.blocksSeen++
	metrics.BlocksSeen.Record(ctx, l2.blocksSeen)

	if l2.hasBuilderTx(ctx, block) {
		l2.blocks.Push(true)
		l2.blocksLanded++
		metrics.BlocksLanded.Record(ctx, l2.blocksLanded)
	} else {
		l2.blocks.Push(false)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Uint64("block", blockNumber),
		)
	}

	if l2.blocks.Length() > l2.reorgWindow {
		_, _ = l2.blocks.Pop()
	}

	return nil
}

func (l2 *L2) processReorg(ctx context.Context, newBlockHeight uint64) {
	l := logutils.LoggerFromContext(ctx)

	depth := l2.blockHeight - newBlockHeight + 1

	if depth < uint64(l2.cfg.ReorgWindow) {
		l.Warn("Chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", l2.blockHeight),
			zap.Uint64("depth", depth),
		)
	} else {
		l.Warn("Super-deep chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", l2.blockHeight),
			zap.Uint64("depth", depth),
		)
	}

	metrics.ReorgCount.Add(ctx, 1)
	metrics.ReorgDepth.Record(ctx, int64(depth))

	adjustment := 0
	for b := newBlockHeight; b <= l2.blockHeight; b++ {
		if landed, _ := l2.blocks.At(int(b)); landed {
			adjustment++
		}
	}

	l2.blocksSeen -= int64(depth)
	metrics.BlocksSeen.Record(ctx, l2.blocksSeen)

	l2.blocksLanded -= int64(adjustment)
	metrics.BlocksLanded.Record(ctx, l2.blocksLanded)

	l2.blocks.Forget(int(depth))
	l2.blockHeight = newBlockHeight - 1
}

func (l2 *L2) hasBuilderTx(ctx context.Context, block *ethtypes.Block) bool {
	l := logutils.LoggerFromContext(ctx)

	expectedData := []byte(fmt.Sprintf("Block Number: %s", block.Number().String()))

	for _, tx := range block.Transactions() {
		if tx.To().Cmp(ethcommon.Address{}) != 0 {
			continue // builder's tx burns eth by sending it to zero address
		}

		if len(expectedData) != len(tx.Data()) {
			continue
		}
		for idx, b := range tx.Data() {
			if expectedData[idx] != b {
				continue
			}
		}

		from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			l.Warn("Failed to determine the sender for transaction",
				zap.Error(err),
				zap.String("tx", tx.Hash().Hex()),
				zap.String("block", block.Number().String()),
			)
			continue
		}

		if from.Cmp(l2.builderAddr) == 0 {
			return true
		}
	}

	return false
}

func (l2 *L2) observeWallets(ctx context.Context, o otelapi.Observer) error {
	errs := make([]error, 0)

	for name, addr := range l2.wallets {
		_balance, err := l2.rpc.BalanceAt(ctx, addr, nil)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		balance, _ := _balance.Float64()

		o.ObserveFloat64(metrics.WalletBalance, balance, otelapi.WithAttributes(
			attribute.KeyValue{Key: "wallet_address", Value: attribute.StringValue(addr.String())},
			attribute.KeyValue{Key: "wallet_name", Value: attribute.StringValue(name)},
		))
	}

	switch len(errs) {
	default:
		return errors.Join(errs...)
	case 1:
		return errs[0]
	case 0:
		return nil
	}
}
