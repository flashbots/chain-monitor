package server

import (
	"context"
	"errors"
	"fmt"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/metrics"
	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

type L1 struct {
	cfg     *config.L1
	rpc     *ethclient.Client
	wallets map[string]ethcommon.Address
}

func newL1(cfg *config.L1) (*L1, error) {
	var (
		wallets = make(map[string]ethcommon.Address, len(cfg.WalletAddresses))
	)

	for name, addrStr := range cfg.WalletAddresses {
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
		wallets[name] = addr
	}

	rpc, err := ethclient.Dial(cfg.RPC)
	if err != nil {
		return nil, err
	}

	return &L1{
		cfg:     cfg,
		rpc:     rpc,
		wallets: wallets,
	}, nil
}

func (l1 *L1) run(_ context.Context) {
	// no-op
}

func (l1 *L1) stop() {
	// no-op
}

func (l1 *L1) observeWallets(ctx context.Context, o otelapi.Observer) error {
	errs := make([]error, 0)

	for name, addr := range l1.wallets {
		_balance, err := l1.rpc.BalanceAt(ctx, addr, nil)
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
