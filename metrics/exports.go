package metrics

import (
	"context"

	"github.com/flashbots/chain-monitor/config"
	otelapi "go.opentelemetry.io/otel/metric"
)

var (
	BlockHeight otelapi.Int64ObservableGauge
	BlockMissed otelapi.Int64Gauge

	BlocksLandedCount otelapi.Int64Gauge
	BlocksMissedCount otelapi.Int64Gauge
	BlocksSeenCount   otelapi.Int64Gauge

	FlashblocksLandedCount otelapi.Int64Gauge
	FlashblocksMissedCount otelapi.Int64Gauge

	FlashtestationsLandedCount     otelapi.Int64Gauge
	FlashtestationsMissedCount     otelapi.Int64Gauge
	RegisteredFlashtestationsCount otelapi.Int64Gauge

	ReorgsCount otelapi.Int64Counter
	ReorgDepth  otelapi.Int64Gauge

	WalletBalance otelapi.Float64ObservableGauge

	ProbesSentCount   otelapi.Int64ObservableCounter
	ProbesFailedCount otelapi.Int64ObservableCounter
	ProbesLandedCount otelapi.Int64ObservableCounter
	ProbesLatency     *Int64Candlestick

	FailedTxPerBlock *Int64Candlestick
	GasPerBlock      *Int64Candlestick
	GasPerTx         *Int64Candlestick
	GasPricePerTx    *Int64Candlestick
	TxPerBlock       *Int64Candlestick

	// TODO: get rid of this
	GasPrice *Int64Candlestick
)

var (
	setups = []func(context.Context, *config.ProbeTx) error{
		setupMeter, // must come first

		setupBlockHeight,
		setupBlockMissed,

		setupBlocksLandedCount,
		setupBlocksMissedCount,
		setupBlocksSeenCount,

		setupFlashblocksLandedCount,
		setupFlashblocksMissedCount,

		setupFlashtestationsLandedCount,
		setupFlashtestationsMissedCount,
		setupRegisteredFlashtestationsCount,

		setupReorgsCount,
		setupReorgDepth,

		setupWalletBalance,

		setupProbesSentCount,
		setupProbesFailedCount,
		setupProbesLandedCount,
		setupProbesLatency,

		setupFailedTxPerBlock,
		setupGasPerBlock,
		setupGasPerTx,
		setupGasPricePerTx,
		setupTxPerBlock,

		setupGasPrice,
	}
)
