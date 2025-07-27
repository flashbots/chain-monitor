package server

import (
	"context"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/rpc"
	"github.com/flashbots/chain-monitor/utils"
	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type L1 struct {
	cfg *config.L1

	rpc *rpc.RPC

	chainID *big.Int
	wallets map[string]ethcommon.Address
}

func newL1(cfg *config.L1) (*L1, error) {
	l := zap.L()

	l1 := &L1{
		cfg:     cfg,
		wallets: make(map[string]ethcommon.Address, len(cfg.MonitorWalletAddresses)),
	}

	for name, addrStr := range cfg.MonitorWalletAddresses {
		var addr ethcommon.Address
		addrBytes, err := ethcommon.ParseHexOrString(addrStr)
		if err != nil {
			return nil, err
		}
		if len(addrBytes) != 20 {
			return nil, fmt.Errorf(
				"invalid length for the l1 wallet address (want 20, got %d)",
				len(addr),
			)
		}
		copy(addr[:], addrBytes)
		l1.wallets[name] = addr
	}

	{ // rpc
		rpc, err := rpc.New(cfg.NetworkID, cfg.Rpc, cfg.RpcFallback...)
		if err != nil {
			return nil, err
		}
		l1.rpc = rpc
	}

	{ // chainID
		chainID, err := l1.rpc.NetworkID(context.Background())
		if err != nil {
			l.Error("Failed to request network id",
				zap.Error(err),
				zap.String("kind", "l1"),
			)
			return nil, err
		}

		l1.chainID = chainID
	}

	return l1, nil
}

func (l1 *L1) run(_ context.Context) {
	// no-op
}

func (l1 *L1) stop() {
	// no-op
}

func (l1 *L1) observeWallets(ctx context.Context, o otelapi.Observer) error {
	l := logutils.LoggerFromContext(ctx)

	errs := make([]error, 0)

	for name, addr := range l1.wallets {
		_balance, err := l1.rpc.BalanceAt(ctx, addr)
		if err != nil {
			l.Error("Failed to request balance",
				zap.Error(err),
				zap.String("at", addr.String()),
				zap.String("kind", "l1"),
				zap.String("rpc", l1.cfg.Rpc),
			)
			errs = append(errs, err)
			continue
		}

		balance, _ := _balance.Float64()

		o.ObserveFloat64(metrics.WalletBalance, balance, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l1")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(l1.chainID.Int64())},
			attribute.KeyValue{Key: "wallet_address", Value: attribute.StringValue(addr.String())},
			attribute.KeyValue{Key: "wallet_name", Value: attribute.StringValue(name)},
		))
	}

	return utils.FlattenErrors(errs)
}
