package server

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/types"

	"go.uber.org/zap"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

func (s *Server) processNewBlocks(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	blockHeight, err := s.eth.BlockNumber(ctx)
	if err != nil {
		l.Error("Failed to get block height, skipping this round...",
			zap.Error(err),
			zap.String("rpc", s.cfg.Eth.RPC),
		)
		return
	}

	if blockHeight == s.blockHeight {
		return
	}

	if blockHeight < s.blockHeight {
		s.processReorg(ctx, blockHeight)
	}

	for b := s.blockHeight + 1; b <= blockHeight; b++ {
		if err := s.processBlock(ctx, b); err != nil {
			l.Error("Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block", blockHeight),
				zap.String("rpc", s.cfg.Eth.RPC),
			)
			return
		}
		s.blockHeight = b
	}
}

func (s *Server) processBlock(ctx context.Context, blockNumber uint64) error {
	l := logutils.LoggerFromContext(ctx)

	block, err := s.eth.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	if s.blocks == nil {
		s.blocks = types.NewRingBuffer[bool](int(blockNumber), s.cfg.Eth.ReorgWindow)
	}

	s.blocksSeen++
	metrics.BlocksSeen.Record(ctx, s.blocksSeen)

	if s.hasBuilderTx(ctx, block) {
		s.blocks.Push(true)
		s.blocksLanded++
		metrics.BlocksLanded.Record(ctx, s.blocksLanded)
	} else {
		s.blocks.Push(false)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Uint64("block", blockNumber),
		)
	}

	if s.blocks.Length() > s.cfg.Eth.ReorgWindow {
		_, _ = s.blocks.Pop()
	}

	return nil
}

func (s *Server) processReorg(ctx context.Context, newBlockHeight uint64) {
	l := logutils.LoggerFromContext(ctx)

	depth := s.blockHeight - newBlockHeight + 1

	if depth < uint64(s.cfg.Eth.ReorgWindow) {
		l.Warn("Chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", s.blockHeight),
			zap.Uint64("depth", depth),
		)
	} else {
		l.Warn("Super-deep chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", s.blockHeight),
			zap.Uint64("depth", depth),
		)
	}

	metrics.ReorgCount.Add(ctx, 1)
	metrics.ReorgDepth.Record(ctx, int64(depth))

	adjustment := 0
	for b := newBlockHeight; b <= s.blockHeight; b++ {
		if landed, _ := s.blocks.At(int(b)); landed {
			adjustment++
		}
	}

	s.blocksSeen -= int64(depth)
	metrics.BlocksSeen.Record(ctx, s.blocksSeen)

	s.blocksLanded -= int64(adjustment)
	metrics.BlocksLanded.Record(ctx, s.blocksLanded)
}

func (s *Server) hasBuilderTx(ctx context.Context, block *ethtypes.Block) bool {
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

		if from.Cmp(s.builderAddr) == 0 {
			return true
		}
	}

	return false
}

func (s *Server) observeWallets(ctx context.Context, o otelapi.Observer) error {
	errs := make([]error, 0)

	for name, addr := range s.wallets {
		_balance, err := s.eth.BalanceAt(ctx, addr, nil)
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
