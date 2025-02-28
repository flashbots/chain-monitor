package metrics

import (
	"context"

	"go.opentelemetry.io/otel/exporters/prometheus"
	otelapi "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	metricsNamespace = "chain-monitor"
)

var (
	meter otelapi.Meter
)

func Setup(
	ctx context.Context,
	observe func(ctx context.Context, o otelapi.Observer) error,
) error {
	for _, setup := range []func(context.Context) error{
		setupMeter, // must come first
		setupBlocksLanded,
		setupBlocksSeen,
		setupBlockMissed,
		setupReorgCount,
		setupReorgDepth,
		setupWalletBalance,
	} {
		if err := setup(ctx); err != nil {
			return err
		}
	}

	_, err := meter.RegisterCallback(observe,
		WalletBalance,
	)
	if err != nil {
		return err
	}

	return nil
}

func setupMeter(ctx context.Context) error {
	res, err := resource.New(ctx)
	if err != nil {
		return err
	}

	exporter, err := prometheus.New(
		prometheus.WithNamespace(metricsNamespace),
		prometheus.WithoutScopeInfo(),
	)
	if err != nil {
		return err
	}

	provider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(res),
	)

	meter = provider.Meter(metricsNamespace)

	return nil
}

func setupBlocksLanded(ctx context.Context) error {
	blocksLanded, err := meter.Int64Gauge("blocks_landed",
		otelapi.WithDescription("blocks landed by our builder"),
	)
	if err != nil {
		return err
	}
	BlocksLanded = blocksLanded
	return nil
}

func setupBlocksSeen(ctx context.Context) error {
	blocksSeen, err := meter.Int64Gauge("blocks_seen",
		otelapi.WithDescription("blocks seen by the monitor"),
	)
	if err != nil {
		return err
	}
	BlocksSeen = blocksSeen
	return nil
}

func setupBlockMissed(ctx context.Context) error {
	blockMissed, err := meter.Int64Gauge("block_missed",
		otelapi.WithDescription("height of the most recent missed block"),
	)
	if err != nil {
		return err
	}
	BlockMissed = blockMissed
	return nil
}

func setupReorgCount(ctx context.Context) error {
	reorgCount, err := meter.Int64Counter("reorg_count",
		otelapi.WithDescription("chain reorgs count"),
	)
	if err != nil {
		return err
	}
	ReorgCount = reorgCount
	return nil
}

func setupReorgDepth(ctx context.Context) error {
	reorgDepth, err := meter.Int64Gauge("reorg_depth",
		otelapi.WithDescription("depth of the most recent reorg"),
	)
	if err != nil {
		return err
	}
	ReorgDepth = reorgDepth
	return nil
}

func setupWalletBalance(ctx context.Context) error {
	walletBalance, err := meter.Float64ObservableGauge("wallet_balance",
		otelapi.WithDescription("wallet balance"),
	)
	if err != nil {
		return err
	}
	WalletBalance = walletBalance
	return nil
}
