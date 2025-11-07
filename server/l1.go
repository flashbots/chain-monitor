package server

import (
	"context"
	"fmt"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/server/wallet"
	otelapi "go.opentelemetry.io/otel/metric"
)

type L1 struct {
	walletObserver *wallet.Observer
}

func newL1(cfg *config.L1) (*L1, error) {
	walletObserver, err := wallet.NewObserver(cfg.NetworkID, cfg.Rpc, cfg.RpcFallback, cfg.MonitorWalletAddresses)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialise wallet observer: %w",
			err,
		)
	}

	return &L1{
		walletObserver: walletObserver,
	}, nil
}

func (l1 *L1) run(_ context.Context) {
	// no-op
}

func (l1 *L1) stop() {
	if l1 == nil {
		return
	}

	l1.walletObserver.Stop()
}

func (l1 *L1) observeWallets(ctx context.Context, o otelapi.Observer) error {
	return l1.walletObserver.Observe(ctx, o)
}
