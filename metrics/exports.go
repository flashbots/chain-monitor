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
	ProbesLatency     *Int64Candlestick

	TxPerBlock  *Int64Candlestick
	GasPerBlock *Int64Candlestick
	GasPrice    *Int64Candlestick

	// TODO: get rid of this
	ProbesLatency_Old otelapi.Int64Histogram
	TxPerBlock_Old    otelapi.Int64Histogram
	GasPerBlock_Old   otelapi.Int64Histogram
	GasPrice_Old      otelapi.Int64Histogram
)
