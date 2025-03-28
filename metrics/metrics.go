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
		setupBlockMissed,
		setupBlocksLandedCount,
		setupBlocksMissedCount,
		setupBlocksSeenCount,
		setupReorgsCount,
		setupReorgDepth,
		setupWalletBalance,
		setupProbesSent,
		setupProbesFailed,
		setupProbesLanded,
		setupProbesLatency,
		setupTxPerBlock,
		setupGasPerBlock,
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

func setupBlockMissed(ctx context.Context) error {
	m, err := meter.Int64Gauge("block_missed",
		otelapi.WithDescription("height of the most recent missed block"),
	)
	if err != nil {
		return err
	}
	BlockMissed = m
	return nil
}

func setupBlocksLandedCount(ctx context.Context) error {
	m, err := meter.Int64Gauge("blocks_landed_count",
		otelapi.WithDescription("blocks landed by our builder"),
	)
	if err != nil {
		return err
	}
	BlocksLandedCount = m
	return nil
}

func setupBlocksMissedCount(ctx context.Context) error {
	m, err := meter.Int64Gauge("blocks_missed_count",
		otelapi.WithDescription("blocks missed by our builder"),
	)
	if err != nil {
		return err
	}
	BlocksMissedCount = m
	return nil
}

func setupBlocksSeenCount(ctx context.Context) error {
	m, err := meter.Int64Gauge("blocks_seen_count",
		otelapi.WithDescription("blocks seen by the monitor"),
	)
	if err != nil {
		return err
	}
	BlocksSeenCount = m
	return nil
}

func setupReorgsCount(ctx context.Context) error {
	m, err := meter.Int64Counter("reorgs_count",
		otelapi.WithDescription("chain reorgs count"),
	)
	if err != nil {
		return err
	}
	ReorgsCount = m
	return nil
}

func setupReorgDepth(ctx context.Context) error {
	m, err := meter.Int64Gauge("reorg_depth",
		otelapi.WithDescription("depth of the most recent reorg"),
	)
	if err != nil {
		return err
	}
	ReorgDepth = m
	return nil
}

func setupWalletBalance(ctx context.Context) error {
	m, err := meter.Float64ObservableGauge("wallet_balance",
		otelapi.WithDescription("wallet balance"),
	)
	if err != nil {
		return err
	}
	WalletBalance = m
	return nil
}

func setupProbesSent(ctx context.Context) error {
	m, err := meter.Int64Counter("probes_sent_count",
		otelapi.WithDescription("count of sent probe transactions"),
	)
	if err != nil {
		return err
	}
	ProbesSentCount = m
	return nil
}

func setupProbesFailed(ctx context.Context) error {
	m, err := meter.Int64Counter("probes_failed_count",
		otelapi.WithDescription("count of probe transactions we failed to send"),
	)
	if err != nil {
		return err
	}
	ProbesFailedCount = m
	return nil
}

func setupProbesLanded(ctx context.Context) error {
	m, err := meter.Int64Counter("probes_landed_count",
		otelapi.WithDescription("count of landed probe transactions"),
	)
	if err != nil {
		return err
	}
	ProbesLandedCount = m
	return nil
}

func setupProbesLatency(ctx context.Context) error {
	m, err := meter.Int64Histogram("probes_latency",
		otelapi.WithDescription("latency of landed probe transactions"),
		otelapi.WithUnit("s"),
		otelapi.WithExplicitBucketBoundaries(0, 1, 4, 16, 64, 256),
	)
	if err != nil {
		return err
	}
	ProbesLatency = m
	return nil
}

func setupTxPerBlock(ctx context.Context) error {
	m, err := meter.Int64Histogram("tx_per_block",
		otelapi.WithDescription("count of transactions in a block"),
		otelapi.WithExplicitBucketBoundaries(0, 1, 2, 4, 8, 16, 32, 64, 128, 256),
	)
	if err != nil {
		return err
	}
	TxPerBlock = m
	return nil
}

func setupGasPerBlock(ctx context.Context) error {
	m, err := meter.Int64Histogram("gas_per_block",
		otelapi.WithDescription("gas per a block"),
		otelapi.WithExplicitBucketBoundaries(0, 1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000),
	)
	if err != nil {
		return err
	}
	GasPerBlock = m
	return nil
}
