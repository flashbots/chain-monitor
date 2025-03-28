package metrics

import (
	otelapi "go.opentelemetry.io/otel/metric"
)

var (
	BlocksLandedCount otelapi.Int64Gauge
	BlocksMissedCount otelapi.Int64Gauge
	BlocksSeenCount   otelapi.Int64Gauge

	BlockMissed otelapi.Int64Gauge

	ReorgsCount otelapi.Int64Counter
	ReorgDepth  otelapi.Int64Gauge

	WalletBalance otelapi.Float64ObservableGauge

	ProbesSentCount   otelapi.Int64Counter
	ProbesFailedCount otelapi.Int64Counter
	ProbesLandedCount otelapi.Int64Counter
	ProbesLatency     otelapi.Int64Histogram

	TxPerBlock  otelapi.Int64Histogram
	GasPerBlock otelapi.Int64Histogram
)
