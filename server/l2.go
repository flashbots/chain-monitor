package server

import (
	"context"
	"fmt"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/server/l2"
	"github.com/flashbots/chain-monitor/server/wallet"
	"github.com/flashbots/chain-monitor/utils"

	otelapi "go.opentelemetry.io/otel/metric"
)

type L2 struct {
	blockInspector            *l2.BlockInspector
	flashblocksMonitor        *l2.FlashblocksMonitor
	txInclusionLatencyMonitor *l2.TxInclusionLatencyMonitor
	walletObserver            *wallet.Observer

	canceller context.CancelFunc
}

func newL2(cfg *config.L2) (*L2, error) {
	blockInspector, err := l2.NewBlockInspector(cfg)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialise block inspector: %w",
			err,
		)
	}

	flashblocksMonitor, err := l2.NewFlashblocksMonitor(cfg)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialise flashblocks monitor: %w",
			err,
		)
	}

	txInclusionLatencyMonitor, err := l2.NewTxInclusionLatencyMonitor(cfg)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialise tx-inclusion-latency monitor: %w",
			err,
		)
	}

	walletObserver, err := wallet.NewObserver(cfg.NetworkID, cfg.Rpc, cfg.RpcFallback, cfg.MonitorWalletAddresses)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to initialise wallet observer: %w",
			err,
		)
	}

	return &L2{
		blockInspector:            blockInspector,
		flashblocksMonitor:        flashblocksMonitor,
		txInclusionLatencyMonitor: txInclusionLatencyMonitor,
		walletObserver:            walletObserver,
	}, nil
}

func (l2 *L2) run(ctx context.Context) {
	if l2 == nil {
		return
	}

	ctx, cancel := context.WithCancel(ctx)
	l2.canceller = cancel

	flashblocks := l2.flashblocksMonitor.Run(ctx)
	l2.blockInspector.Run(ctx, flashblocks)
	l2.txInclusionLatencyMonitor.Run(ctx)
	l2.walletObserver.Run(ctx)
}

func (l2 *L2) stop() {
	if l2 == nil {
		return
	}

	l2.walletObserver.Stop()
	l2.txInclusionLatencyMonitor.Stop()
	l2.flashblocksMonitor.Stop()
	l2.blockInspector.Stop()
}

func (l2 *L2) observe(ctx context.Context, o otelapi.Observer) error {
	errs := make([]error, 0)

	if err := l2.blockInspector.Observe(ctx, o); err != nil {
		errs = append(errs, err)
	}

	if err := l2.flashblocksMonitor.Observe(ctx, o); err != nil {
		errs = append(errs, err)
	}

	if err := l2.txInclusionLatencyMonitor.Observe(ctx, o); err != nil {
		errs = append(errs, err)
	}

	if err := l2.walletObserver.Observe(ctx, o); err != nil {
		errs = append(errs, err)
	}

	return utils.FlattenErrors(errs)
}
