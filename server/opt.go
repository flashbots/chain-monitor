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

	blockHeight, err := s.opt.BlockNumber(ctx)
	if err != nil {
		l.Error("Failed to get block height, skipping this round...",
			zap.Error(err),
			zap.String("rpc", s.cfg.Opt.RPC),
		)
		return
	}

	if blockHeight == s.optBlockHeight {
		return
	}

	if blockHeight < s.optBlockHeight {
		s.processReorg(ctx, blockHeight)
	}

	for b := s.optBlockHeight + 1; b <= blockHeight; b++ {
		if err := s.processBlock(ctx, b); err != nil {
			l.Error("Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block", blockHeight),
				zap.String("rpc", s.cfg.Opt.RPC),
			)
			return
		}
		s.optBlockHeight = b
	}
}

func (s *Server) processBlock(ctx context.Context, blockNumber uint64) error {
	l := logutils.LoggerFromContext(ctx)

	block, err := s.opt.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	if s.optBlocks == nil {
		s.optBlocks = types.NewRingBuffer[bool](int(blockNumber), s.optReorgWindow)
	}

	s.optBlocksSeen++
	metrics.BlocksSeen.Record(ctx, s.optBlocksSeen)

	if s.hasBuilderTx(ctx, block) {
		s.optBlocks.Push(true)
		s.optBlocksLanded++
		metrics.BlocksLanded.Record(ctx, s.optBlocksLanded)
	} else {
		s.optBlocks.Push(false)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Uint64("block", blockNumber),
		)
	}

	if s.optBlocks.Length() > s.optReorgWindow {
		_, _ = s.optBlocks.Pop()
	}

	return nil
}

func (s *Server) processReorg(ctx context.Context, newBlockHeight uint64) {
	l := logutils.LoggerFromContext(ctx)

	depth := s.optBlockHeight - newBlockHeight + 1

	if depth < uint64(s.cfg.Opt.ReorgWindow) {
		l.Warn("Chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", s.optBlockHeight),
			zap.Uint64("depth", depth),
		)
	} else {
		l.Warn("Super-deep chain reorg detected",
			zap.Uint64("current", newBlockHeight),
			zap.Uint64("seen", s.optBlockHeight),
			zap.Uint64("depth", depth),
		)
	}

	metrics.ReorgCount.Add(ctx, 1)
	metrics.ReorgDepth.Record(ctx, int64(depth))

	adjustment := 0
	for b := newBlockHeight; b <= s.optBlockHeight; b++ {
		if landed, _ := s.optBlocks.At(int(b)); landed {
			adjustment++
		}
	}

	s.optBlocksSeen -= int64(depth)
	metrics.BlocksSeen.Record(ctx, s.optBlocksSeen)

	s.optBlocksLanded -= int64(adjustment)
	metrics.BlocksLanded.Record(ctx, s.optBlocksLanded)
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

		if from.Cmp(s.optBuilderAddr) == 0 {
			return true
		}
	}

	return false
}

func (s *Server) observeWallets(ctx context.Context, o otelapi.Observer) error {
	errs := make([]error, 0)

	for name, addr := range s.optWallets {
		_balance, err := s.opt.BalanceAt(ctx, addr, nil)
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
