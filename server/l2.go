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
	blocksMissed int64
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
	blockHeight--
	blocks := types.NewRingBuffer[bool](int(blockHeight), int(cfg.ReorgWindow/cfg.BlockTime+1))

	return &L2{
		cfg: cfg,

		rpc:    rpc,
		ticker: time.NewTicker(cfg.BlockTime),

		blockHeight: blockHeight,
		blocks:      blocks,
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
		l.Debug("Still at the same height, skipping...",
			zap.Uint64("block_height", blockHeight),
		)
		return
	}

	if blockHeight < l2.blockHeight {
		l2.processReorg(ctx, blockHeight)
	}

	for b := l2.blockHeight + 1; b <= blockHeight; b++ {
		l.Debug("Processing new l2 block",
			zap.Uint64("block_height", b),
		)

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

	l2.blocksSeen++
	metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)

	if l2.hasBuilderTx(ctx, block) {
		l2.blocks.Push(true)
		l2.blocksLanded++
		metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
	} else {
		l2.blocks.Push(false)
		l2.blocksMissed++
		metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Uint64("block_number", blockNumber),
			zap.Int64("blocks_landed", l2.blocksLanded),
			zap.Int64("blocks_missed", l2.blocksMissed),
			zap.Int64("blocks_seen", l2.blocksSeen),
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

	metrics.ReorgsCount.Add(ctx, 1)
	metrics.ReorgDepth.Record(ctx, int64(depth))

	adjustLanded := 0
	adjustMissed := 0
	for b := newBlockHeight; b <= l2.blockHeight; b++ {
		if landed, ok := l2.blocks.At(int(b)); ok {
			if landed {
				adjustLanded++
			} else {
				adjustMissed++
			}
		}
	}

	l2.blocksSeen -= int64(depth)
	metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)

	l2.blocksLanded -= int64(adjustLanded)
	metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)

	l2.blocksMissed -= int64(adjustMissed)
	metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)

	l2.blocks.Forget(int(depth))
	l2.blockHeight = newBlockHeight - 1

	if depth < uint64(l2.cfg.ReorgWindow) {
		l.Warn("Chain reorg detected",
			zap.Int("adjust_landed", adjustLanded),
			zap.Int("adjust_missed", adjustMissed),
			zap.Uint64("adjust_seen", depth),
			zap.Int64("blocks_landed", l2.blocksLanded),
			zap.Int64("blocks_missed", l2.blocksMissed),
			zap.Int64("blocks_seen", l2.blocksSeen),
			zap.Uint64("block_height", newBlockHeight),
			zap.Uint64("reorg_depth", depth),
		)
	} else {
		l.Warn("Super-deep chain reorg detected",
			zap.Int("adjust_landed", adjustLanded),
			zap.Int("adjust_missed", adjustMissed),
			zap.Uint64("adjust_seen", depth),
			zap.Int64("blocks_landed", l2.blocksLanded),
			zap.Int64("blocks_missed", l2.blocksMissed),
			zap.Int64("blocks_seen", l2.blocksSeen),
			zap.Uint64("block_height", newBlockHeight),
			zap.Uint64("reorg_depth", depth),
		)
	}
}

func (l2 *L2) hasBuilderTx(ctx context.Context, block *ethtypes.Block) bool {
	l := logutils.LoggerFromContext(ctx)

	expectedData := []byte(fmt.Sprintf("Block Number: %s", block.Number().String()))

	for _, tx := range block.Transactions() {
		if tx == nil || tx.To() == nil || tx.To().Cmp(ethcommon.Address{}) != 0 {
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
