package wallet

import (
	"context"
	"fmt"
	"math/big"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/rpc"
	"github.com/flashbots/chain-monitor/utils"
	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
)

type Observer struct {
	rpc    *rpc.RPC
	rpcUrl string

	chainID *big.Int
	wallets map[string]ethcommon.Address
}

func NewObserver(networkID uint64, rpcUrl string, rpcFallbackUrls []string, wallets map[string]string) (*Observer, error) {
	if len(wallets) == 0 {
		return nil, nil
	}

	l := zap.L()

	obs := &Observer{
		rpcUrl:  rpcUrl,
		wallets: make(map[string]ethcommon.Address, len(wallets)),
	}

	{ // rpc
		rpc, err := rpc.New(networkID, rpcUrl, rpcFallbackUrls...)
		if err != nil {
			return nil, err
		}
		obs.rpc = rpc
	}

	{ // chainID
		chainID, err := obs.rpc.NetworkID(context.Background())
		if err != nil {
			l.Error("Failed to request network id",
				zap.Error(err),
				zap.String("kind", "l1"),
			)
			return nil, err
		}

		obs.chainID = chainID
	}

	{ // wallets
		for name, addrStr := range wallets {
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
			obs.wallets[name] = addr
		}
	}

	return obs, nil
}

func (obs *Observer) Run(_ context.Context) {
	// no-op
}

func (obs *Observer) Stop() {
	if obs == nil {
		return
	}

	obs.rpc.Close()
}

func (obs *Observer) Observe(ctx context.Context, o otelapi.Observer) error {
	if obs == nil {
		return nil
	}

	l := logutils.LoggerFromContext(ctx)

	errs := make([]error, 0)

	for name, addr := range obs.wallets {
		_balance, err := obs.rpc.BalanceAt(ctx, addr)
		if err != nil {
			l.Error("Failed to request balance",
				zap.Error(err),
				zap.String("at", addr.String()),
				zap.String("kind", "l1"),
				zap.String("rpc", obs.rpcUrl),
			)
			errs = append(errs, err)
			continue
		}

		balance, _ := _balance.Float64()

		o.ObserveFloat64(metrics.WalletBalance, balance, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l1")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(obs.chainID.Int64())},
			attribute.KeyValue{Key: "wallet_address", Value: attribute.StringValue(addr.String())},
			attribute.KeyValue{Key: "wallet_name", Value: attribute.StringValue(name)},
		))
	}

	return utils.FlattenErrors(errs)
}
