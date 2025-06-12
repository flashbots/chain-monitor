package metrics

import (
	"context"

	"github.com/flashbots/chain-monitor/config"
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
	cfg *config.Monitor,
	observe func(ctx context.Context, o otelapi.Observer) error,
) error {
	for _, setup := range setups {
		if err := setup(ctx, cfg); err != nil {
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

func setupMeter(ctx context.Context, _ *config.Monitor) error {
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

func setupBlockMissed(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Gauge("block_missed",
		otelapi.WithDescription("height of the most recent missed block"),
	)
	if err != nil {
		return err
	}
	BlockMissed = m
	return nil
}

func setupBlocksLandedCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Gauge("blocks_landed_count",
		otelapi.WithDescription("blocks landed by our builder"),
	)
	if err != nil {
		return err
	}
	BlocksLandedCount = m
	return nil
}

func setupBlocksMissedCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Gauge("blocks_missed_count",
		otelapi.WithDescription("blocks missed by our builder"),
	)
	if err != nil {
		return err
	}
	BlocksMissedCount = m
	return nil
}

func setupBlocksSeenCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Gauge("blocks_seen_count",
		otelapi.WithDescription("blocks seen by the monitor"),
	)
	if err != nil {
		return err
	}
	BlocksSeenCount = m
	return nil
}

func setupReorgsCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Counter("reorgs_count",
		otelapi.WithDescription("chain reorgs count"),
	)
	if err != nil {
		return err
	}
	ReorgsCount = m
	return nil
}

func setupReorgDepth(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Gauge("reorg_depth",
		otelapi.WithDescription("depth of the most recent reorg"),
	)
	if err != nil {
		return err
	}
	ReorgDepth = m
	return nil
}

func setupWalletBalance(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Float64ObservableGauge("wallet_balance",
		otelapi.WithDescription("wallet balance"),
	)
	if err != nil {
		return err
	}
	WalletBalance = m
	return nil
}

func setupProbesSentCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Counter("probes_sent_count",
		otelapi.WithDescription("count of sent probe transactions"),
	)
	if err != nil {
		return err
	}
	ProbesSentCount = m
	return nil
}

func setupProbesFailedCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Counter("probes_failed_count",
		otelapi.WithDescription("count of probe transactions we failed to send"),
	)
	if err != nil {
		return err
	}
	ProbesFailedCount = m
	return nil
}

func setupProbesLandedCount(ctx context.Context, _ *config.Monitor) error {
	m, err := meter.Int64Counter("probes_landed_count",
		otelapi.WithDescription("count of landed probe transactions"),
	)
	if err != nil {
		return err
	}
	ProbesLandedCount = m
	return nil
}

func setupProbesLatency(ctx context.Context, _ *config.Monitor) error {
	m, err := NewInt64Candlestick("probes_latency_ohlc", "latency of landed probe transactions", "s")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	ProbesLatency = m

	return nil
}

func setupGasPerBlock(ctx context.Context, cfg *config.Monitor) error {
	m, err := NewInt64Candlestick("gas_per_block_ohlc", "gas per block", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	GasPerBlock = m
	return nil
}

func setupGasPerTx(ctx context.Context, cfg *config.Monitor) error {
	m, err := NewInt64Candlestick("gas_per_tx_ohlc", "gas per transaction", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	GasPerTx = m
	return nil
}

func setupGasPricePerTx(ctx context.Context, cfg *config.Monitor) error {
	m, err := NewInt64Candlestick("gas_price_per_tx_ohlc", "gas per transaction", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	GasPricePerTx = m
	return nil
}

func setupL1FeePerTx(ctx context.Context, cfg *config.Monitor) error {
	m, err := NewInt64Candlestick("l1_fee_per_tx_ohlc", "gas per transaction", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	L1FeePerTx = m
	return nil
}

func setupTxPerBlock(ctx context.Context, _ *config.Monitor) error {
	m, err := NewInt64Candlestick("tx_per_block_ohlc", "count of transactions in a block", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	TxPerBlock = m
	return nil
}

// TODO: get rid of below

func setupGasPrice(ctx context.Context, cfg *config.Monitor) error {
	m, err := NewInt64Candlestick("gas_price_ohlc", "gas price", "")
	if err != nil {
		return err
	}
	if _, err := m.registerCallback(meter); err != nil {
		return err
	}
	GasPrice = m
	return nil
}
