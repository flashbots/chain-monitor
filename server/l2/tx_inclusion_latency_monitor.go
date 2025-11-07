package l2

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"math/big"
	"strings"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/rpc"
	"github.com/flashbots/chain-monitor/utils"
	"go.uber.org/zap"

	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

type TxInclusionLatencyMonitor struct {
	// parameters

	cfg *txInclusionLatencyMonitorConfig

	// actors

	blockTicker *time.Ticker
	resetTicker *time.Ticker
	rpc         *rpc.RPC

	// state

	monitorNonce uint64

	metrics *txInclusionLatencyMonitorMetrics
}

type txInclusionLatencyMonitorConfig struct {
	rpc string

	blockTime time.Duration
	chainID   *big.Int
	signer    ethtypes.EIP155Signer

	gasLimit           uint64
	gasPriceAdjustment int64
	gasPriceCap        int64

	monitorAddr ethcommon.Address
	monitorKey  *ecdsa.PrivateKey
}

type txInclusionLatencyMonitorMetrics struct {
	monitorProbesFailedCount int64
	monitorProbesSentCount   int64
}

func NewTxInclusionLatencyMonitor(cfg *config.L2) (*TxInclusionLatencyMonitor, error) {
	if cfg.ProbeTx.PrivateKey == "" {
		return nil, nil
	}

	l := zap.L()

	m := &TxInclusionLatencyMonitor{
		metrics: &txInclusionLatencyMonitorMetrics{},

		cfg: &txInclusionLatencyMonitorConfig{
			rpc: cfg.Rpc,

			blockTime: cfg.BlockTime,

			gasLimit:           cfg.ProbeTx.GasLimit,
			gasPriceAdjustment: cfg.ProbeTx.GasPriceAdjustment,
			gasPriceCap:        cfg.ProbeTx.GasPriceCap,
		},
	}

	{ // rpc
		rpc, err := rpc.New(cfg.NetworkID, cfg.Rpc, cfg.RpcFallback...)
		if err != nil {
			return nil, err
		}
		m.rpc = rpc
	}

	{ // chainID, signer
		chainID, err := m.rpc.NetworkID(context.Background())
		if err != nil {
			l.Error("Failed to request network id",
				zap.Error(err),
				zap.String("kind", "l2"),
			)
			return nil, err
		}
		m.cfg.chainID = chainID
		m.cfg.signer = ethtypes.NewEIP155Signer(chainID)
	}

	{ // monitor tx addr and key
		if cfg.ProbeTx.PrivateKey != "" {
			monitorKey, err := crypto.HexToECDSA(cfg.ProbeTx.PrivateKey)
			if err != nil {
				return nil, err
			}
			m.cfg.monitorKey = monitorKey
			m.cfg.monitorAddr = crypto.PubkeyToAddress(monitorKey.PublicKey)
		}
	}

	{ // block ticker
		now := time.Now()
		time.Sleep(now.Truncate(cfg.BlockTime).Add(cfg.BlockTime).Sub(now)) // align with block times
		m.blockTicker = time.NewTicker(cfg.BlockTime)
	}

	{ // reset ticker
		now := time.Now()
		time.Sleep(now.Truncate(cfg.BlockTime).Add(cfg.BlockTime).Sub(now)) // align with block times
		m.resetTicker = time.NewTicker(cfg.ProbeTx.ResetInterval)
	}

	return m, nil
}

func (m *TxInclusionLatencyMonitor) Run(ctx context.Context) {
	if m == nil {
		return
	}

	l := logutils.LoggerFromContext(ctx)
	processingContext := logutils.ContextWithLogger(context.Background(), l)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-m.blockTicker.C:
				m.sendProbeTx(processingContext)
			}
		}
	}()
}

func (m *TxInclusionLatencyMonitor) Stop() {
	if m == nil {
		return
	}

	m.blockTicker.Stop()
	m.rpc.Close()
}

func (m *TxInclusionLatencyMonitor) Observe(_ context.Context, o otelapi.Observer) error {
	if m == nil {
		return nil
	}

	o.ObserveInt64(metrics.ProbesFailedCount, m.metrics.monitorProbesFailedCount, otelapi.WithAttributes(
		attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
		attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(m.cfg.chainID.Int64())},
	))

	o.ObserveInt64(metrics.ProbesSentCount, m.metrics.monitorProbesSentCount, otelapi.WithAttributes(
		attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
		attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(m.cfg.chainID.Int64())},
	))

	return nil
}

func (l2 *TxInclusionLatencyMonitor) sendProbeTx(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	start := time.Now()

	var (
		data     = make([]byte, 8)
		gasPrice *big.Int
		err      error
	)

	{ // get the gas price
		gasPrice, err = l2.rpc.SuggestGasPrice(ctx)
		if err != nil {
			l.Warn("Failed to request suggested gas price",
				zap.Error(err),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.rpc),
			)
			l2.metrics.monitorProbesFailedCount++
			return
		}

		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(100+l2.cfg.gasPriceAdjustment))
		gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

		gasPrice = utils.MinBigInt(gasPrice, big.NewInt(l2.cfg.gasPriceCap))
	}

	if l2.monitorNonce == 0 { // get the nonce
		nonce, err := l2.rpc.NonceAt(ctx, l2.cfg.monitorAddr)
		if err != nil {
			l.Warn("Failed to request a nonce",
				zap.Error(err),
				zap.String("address", l2.cfg.monitorAddr.String()),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.rpc),
			)
			l2.metrics.monitorProbesFailedCount++
			return
		}
		l2.monitorNonce = nonce
	}

tryingNonces:
	for attempt := 1; attempt <= 8; attempt++ { // we don't want to get rate-limited
		if attempt > 1 && time.Since(start) > 500*time.Millisecond {
			return
		}

		thisBlock := time.Now().Add(-l2.cfg.blockTime / 2).Round(l2.cfg.blockTime)
		binary.BigEndian.PutUint64(data, uint64(thisBlock.Unix()))

		tx := ethtypes.NewTransaction(
			l2.monitorNonce,
			ethcommon.Address{},
			nil,
			l2.cfg.gasLimit,
			gasPrice,
			data,
		)

		signedTx, err := ethtypes.SignTx(tx, l2.cfg.signer, l2.cfg.monitorKey)
		if err != nil {
			l.Error("Failed to sign a transaction",
				zap.Error(err),
				zap.String("address", l2.cfg.monitorAddr.String()),
			)
			l2.metrics.monitorProbesFailedCount++
			return
		}

		err = l2.rpc.SendTransaction(ctx, signedTx)
		if err == nil {
			l2.metrics.monitorProbesSentCount++
			l2.monitorNonce++
			return
		}

		if ctxErr := ctx.Err(); ctxErr != nil {
			l.Error("Failed to send a transaction",
				zap.Error(errors.Join(err, ctxErr)),
				zap.String("address", l2.cfg.monitorAddr.String()),
				zap.String("to", tx.To().String()),
				zap.Uint64("nonce", l2.monitorNonce),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.rpc),
			)
			l2.metrics.monitorProbesFailedCount++
			return // irrecoverable error (for now, at least) => no point in trying other nonces
		}

		l.Error("Failed to send monitor transaction",
			zap.Error(err),
			zap.String("address", l2.cfg.monitorAddr.String()),
			zap.String("to", tx.To().String()),
			zap.Uint64("nonce", l2.monitorNonce),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.rpc),
		)

		switch {
		case strings.Contains(err.Error(), "insufficient funds"):
			l2.metrics.monitorProbesFailedCount++
			return // irrecoverable error

		case strings.Contains(err.Error(), "already known"):
			l2.monitorNonce++ // there's already a tx with this nonce => try next one
			continue tryingNonces

		case strings.Contains(err.Error(), "replacement transaction underpriced"):
			l2.monitorNonce++ // there's already a tx with this nonce => try next one
			continue tryingNonces

		case strings.Contains(err.Error(), "nonce too low"):
			nonce, err := l2.rpc.NonceAt(ctx, l2.cfg.monitorAddr)
			if err != nil {
				l.Warn("Failed to request a nonce",
					zap.Error(err),
					zap.String("address", l2.cfg.monitorAddr.String()),
					zap.String("kind", "l2"),
					zap.String("rpc", l2.cfg.rpc),
				)
				l2.metrics.monitorProbesFailedCount++
				return
			}
			l2.monitorNonce = nonce

			continue tryingNonces
		}

		l2.metrics.monitorProbesFailedCount++
		return
	}

	l2.metrics.monitorProbesFailedCount++
}
