package metrics

import (
	otelapi "go.opentelemetry.io/otel/metric"
)

var (
	BlocksLanded otelapi.Int64Gauge
	BlocksSeen   otelapi.Int64Gauge

	BlockMissed otelapi.Int64Gauge

	ReorgCount otelapi.Int64Counter
	ReorgDepth otelapi.Int64Gauge

	WalletBalance otelapi.Float64ObservableGauge
)
