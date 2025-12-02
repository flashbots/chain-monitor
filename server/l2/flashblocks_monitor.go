package l2

import (
	"bytes"
	"context"
	"encoding/json"

	"io"
	"time"

	"github.com/andybalholm/brotli"
	"github.com/coder/websocket"
	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/types"
	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"

	"go.uber.org/zap"
)

type FlashblocksMonitor struct {
	// parameters

	cfg *flashblocksMonitorConfig

	// actors

	stop context.CancelFunc

	// streams

	flashblocksPublic  chan *flashblockEvent
	flashblocksPrivate chan *flashblockEvent

	lastFlashblockPublic  *flashblockEvent
	lastFlashblockPrivate map[string]*flashblockEvent
}

type flashblocksMonitorConfig struct {
	flashblocksPerBlock int
	networkID           int64
	maxMessageSize      int64
	publicStream        string
	privateStreams      map[string]string
}

type flashblockEvent struct {
	stream     string
	timestamp  time.Time
	flashblock types.Flashblock
}

const (
	wsBackoffMin    = 200 * time.Millisecond
	wsBackoffMax    = time.Minute
	wsBackoffFactor = time.Duration(2)

	wsTimeout = 30 * time.Second
)

func NewFlashblocksMonitor(cfg *config.L2) (*FlashblocksMonitor, error) {
	if cfg.MonitorFlashblocksPublicStream == "" && len(cfg.MonitorFlashblocksPrivateStreams) == 0 {
		return nil, nil
	}

	fm := &FlashblocksMonitor{
		flashblocksPublic:     make(chan *flashblockEvent, 1),
		flashblocksPrivate:    make(chan *flashblockEvent, len(cfg.MonitorFlashblocksPublicStream)),
		lastFlashblockPrivate: make(map[string]*flashblockEvent, len(cfg.MonitorFlashblocksPublicStream)),

		cfg: &flashblocksMonitorConfig{
			networkID:      int64(cfg.NetworkID),
			maxMessageSize: cfg.MonitorFlashblocksMaxWsMessageSizeKb * 1024,
			publicStream:   cfg.MonitorFlashblocksPublicStream,
			privateStreams: cfg.MonitorFlashblocksPrivateStreams,
		},
	}

	return fm, nil
}

func (fm *FlashblocksMonitor) Run(ctx context.Context) *<-chan *flashblockEvent {
	if fm == nil {
		return nil
	}

	processingContext := logutils.ContextWithLogger(
		context.Background(),
		logutils.LoggerFromContext(ctx),
	)

	processingContext, cancel := context.WithCancel(processingContext)
	fm.stop = cancel

	if fm.cfg.publicStream != "" {
		fm.readStream(ctx, "public", fm.cfg.publicStream, fm.flashblocksPublic)
	}
	for stream, url := range fm.cfg.privateStreams {
		fm.readStream(ctx, stream, url, fm.flashblocksPrivate)
	}

	flashblocks := make(chan *flashblockEvent, 2*fm.cfg.flashblocksPerBlock)
	fm.processFlashblocks(processingContext, flashblocks)

	var output <-chan *flashblockEvent = flashblocks

	return &output
}

func (fm *FlashblocksMonitor) Stop() {
	if fm == nil {
		return
	}

	if fm.stop != nil {
		fm.stop()
	}
}

func (fm *FlashblocksMonitor) Observe(_ context.Context, o otelapi.Observer) error {
	if fm == nil {
		return nil
	}

	return nil
}

func (fm *FlashblocksMonitor) readStream(
	ctx context.Context,
	streamID, streamUrl string,
	flashblocks chan<- *flashblockEvent,
) {
	go func() {
		l := logutils.LoggerFromContext(ctx)

		backoff := wsBackoffMin

	redial:
		for ctx.Err() == nil {
			var (
				conn          *websocket.Conn
				doneReceiving context.CancelFunc
			)

			{ // dial
				l.Info("Connecting to flashblocks stream...",
					zap.String("stream", streamID),
					zap.String("url", streamUrl),
				)

				dialCtx, doneDialling := context.WithTimeout(ctx, wsTimeout)

				_conn, _, err := websocket.Dial(dialCtx, streamUrl, &websocket.DialOptions{
					CompressionMode: websocket.CompressionContextTakeover,
				})
				if err != nil {
					doneDialling()

					metrics.FlashblocksReceiveFailureCount.Add(ctx, 1, otelapi.WithAttributes(
						attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
						attribute.KeyValue{Key: "stream", Value: attribute.StringValue(streamID)},
						attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
					))
					l.Warn("Failed to connect to flashblocks stream",
						zap.Error(err),
						zap.String("stream", streamID),
						zap.String("url", streamUrl),
						zap.Duration("backoff", backoff),
					)
					time.Sleep(backoff)
					backoff = min(wsBackoffFactor*backoff, wsBackoffMax)
					continue redial
				}
				_conn.SetReadLimit(fm.cfg.maxMessageSize)

				backoff = wsBackoffMin
				conn = _conn
				doneReceiving = doneDialling
			}

			{ // receive
				for ctx.Err() == nil {
					readCtx, doneReading := context.WithTimeout(ctx, wsTimeout)
					mtype, mbytes, err := conn.Read(readCtx)
					if err != nil {
						metrics.FlashblocksReceiveFailureCount.Add(ctx, 1, otelapi.WithAttributes(
							attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
							attribute.KeyValue{Key: "stream", Value: attribute.StringValue(streamID)},
							attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
						))
						l.Warn("Failed to read message from flashblocks stream",
							zap.Error(err),
							zap.String("stream", streamID),
							zap.String("url", streamUrl),
						)
						doneReading()
						continue redial
					}
					doneReading()

					timestamp := time.Now()

					if mtype == websocket.MessageBinary { // binary means compressed text
						brotliReader := brotli.NewReader(bytes.NewReader(mbytes))
						dbytes, err := io.ReadAll(brotliReader)
						if err != nil {
							metrics.FlashblocksReceiveFailureCount.Add(ctx, 1, otelapi.WithAttributes(
								attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
								attribute.KeyValue{Key: "stream", Value: attribute.StringValue(streamID)},
								attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
							))
							l.Warn("Failed to decompress binary message from flashblocks stream, ignoring...",
								zap.Error(err),
								zap.String("stream", streamID),
								zap.String("url", streamUrl),
							)
							continue
						}
						mbytes = dbytes
					}

					event := flashblockEvent{
						stream:    streamID,
						timestamp: timestamp,
					}
					if err := json.Unmarshal(mbytes, &event.flashblock); err != nil {
						metrics.FlashblocksReceiveFailureCount.Add(ctx, 1, otelapi.WithAttributes(
							attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
							attribute.KeyValue{Key: "stream", Value: attribute.StringValue(streamID)},
							attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
						))
						l.Error("Failed to parse flashblock",
							zap.Error(err),
							zap.String("stream", streamID),
							zap.String("url", streamUrl),
							zap.String("message", string(mbytes)),
						)
						continue
					}

					flashblocks <- &event
				}

				doneReceiving()
			}
		}
	}()
}

func (fm *FlashblocksMonitor) processFlashblocks(
	ctx context.Context,
	output chan<- *flashblockEvent,
) {
	go func() {
		l := logutils.LoggerFromContext(ctx)

		for ctx.Err() == nil {
			select {
			case fb := <-fm.flashblocksPublic:
				metrics.FlashblocksReceiveSuccessCount.Add(ctx, 1, otelapi.WithAttributes(
					attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
					attribute.KeyValue{Key: "stream", Value: attribute.StringValue(fb.stream)},
					attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
				))
				l.Debug("Received flashblock on public stream",
					zap.Time("timestamp", fb.timestamp),
					zap.Any("flashblock", fb.flashblock),
				)
				if fm.lastFlashblockPublic == nil {
					fm.lastFlashblockPublic = fb
					continue // it's a first flashblock we've got
				}

				fm.processFlashblock(ctx, fb, fm.lastFlashblockPublic)
				fm.lastFlashblockPublic = fb
				fm.detectInconsistentFlashblocks(ctx, fb)

				select {
				case output <- fb:
					// no-op
				default:
					// we shouldn't block if there's no reader
				}

			case fb := <-fm.flashblocksPrivate:
				metrics.FlashblocksReceiveSuccessCount.Add(ctx, 1, otelapi.WithAttributes(
					attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
					attribute.KeyValue{Key: "stream", Value: attribute.StringValue(fb.stream)},
					attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
				))
				l.Debug("Received flashblock on private stream",
					zap.Time("timestamp", fb.timestamp),
					zap.String("stream", fb.stream),
					zap.Any("flashblock", fb.flashblock),
				)
				last, exists := fm.lastFlashblockPrivate[fb.stream]
				if !exists {
					fm.lastFlashblockPrivate[fb.stream] = fb
					continue // it's a first flashblock we've got
				}

				fm.processFlashblock(ctx, fb, last)
				fm.lastFlashblockPrivate[fb.stream] = fb
				fm.detectInconsistentFlashblocks(ctx, fb)
			}
		}
	}()
}

func (fm *FlashblocksMonitor) processFlashblock(ctx context.Context, this, last *flashblockEvent) {
	l := logutils.LoggerFromContext(ctx)

	if this.flashblock.Metadata.BlockNumber < last.flashblock.Metadata.BlockNumber {
		l.Warn("Received a flashblock with lower block number than the previous one (reorg?)",
			zap.String("stream", this.stream),
			zap.String("prev_payload_id", last.flashblock.PayloadId),
			zap.Uint64("prev_block", last.flashblock.Metadata.BlockNumber),
			zap.Int("prev_index", last.flashblock.Index),
			zap.String("this_payload_id", this.flashblock.PayloadId),
			zap.Uint64("this_block", this.flashblock.Metadata.BlockNumber),
			zap.Int("this_index", this.flashblock.Index),
		)
		return
	}

	if this.flashblock.Metadata.BlockNumber == last.flashblock.Metadata.BlockNumber &&
		this.flashblock.Index <= last.flashblock.Index {
		// ---
		l.Warn("Received a flashblock with lower index than the previous one (reorg?)",
			zap.String("stream", this.stream),
			zap.String("prev_payload_id", last.flashblock.PayloadId),
			zap.Uint64("prev_block", last.flashblock.Metadata.BlockNumber),
			zap.Int("prev_index", last.flashblock.Index),
			zap.String("this_payload_id", this.flashblock.PayloadId),
			zap.Uint64("this_block", this.flashblock.Metadata.BlockNumber),
			zap.Int("this_index", this.flashblock.Index),
		)
		return
	}

	if this.flashblock.Metadata.BlockNumber == last.flashblock.Metadata.BlockNumber {
		skippedFlashblocks := this.flashblock.Index - last.flashblock.Index - 1

		if skippedFlashblocks > 0 {
			metrics.FlashblocksSkipped.Add(ctx, int64(skippedFlashblocks), otelapi.WithAttributes(
				attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
				attribute.KeyValue{Key: "stream", Value: attribute.StringValue(this.stream)},
				attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
			))
			l.Warn("Flashblock(s) were skipped",
				zap.String("stream", this.stream),
				zap.Int("count", skippedFlashblocks),
				zap.Uint64("prev_block", last.flashblock.Metadata.BlockNumber),
				zap.Int("prev_index", last.flashblock.Index),
				zap.Uint64("this_block", this.flashblock.Metadata.BlockNumber),
				zap.Int("this_index", this.flashblock.Index),
			)
		}

		return
	}

	skippedBlocks := this.flashblock.Metadata.BlockNumber - last.flashblock.Metadata.BlockNumber - 1
	if skippedBlocks > 0 {
		skippedFlashblocks := int(skippedBlocks)*fm.cfg.flashblocksPerBlock + last.flashblock.Index

		metrics.FlashblocksSkipped.Add(ctx, int64(skippedFlashblocks), otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "stream", Value: attribute.StringValue(this.stream)},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
		))
		l.Warn("Flashblock(s) were skipped",
			zap.String("stream", this.stream),
			zap.Int("count", skippedFlashblocks),
			zap.Uint64("prev_block", last.flashblock.Metadata.BlockNumber),
			zap.Int("prev_index", last.flashblock.Index),
			zap.Uint64("this_block", this.flashblock.Metadata.BlockNumber),
			zap.Int("this_index", this.flashblock.Index),
		)
	}
}

func (fm *FlashblocksMonitor) detectInconsistentFlashblocks(ctx context.Context, this *flashblockEvent) {
	l := logutils.LoggerFromContext(ctx)

	compare := func(this, that *flashblockEvent) bool {
		if this == nil || that == nil {
			return true
		}

		if this.flashblock.PayloadId != that.flashblock.PayloadId {
			return true
		}
		if this.flashblock.Index != that.flashblock.Index {
			return true
		}

		if this.flashblock.Metadata.Equal(that.flashblock.Metadata) {
			return true
		}

		l.Warn("Mismatching flashblocks",
			zap.String("payload_id", this.flashblock.PayloadId),
			zap.Int("index", this.flashblock.Index),
			zap.Any("this", this),
			zap.Any("that", that),
		)
		return false
	}

	matches := compare(this, fm.lastFlashblockPublic)
	for _, that := range fm.lastFlashblockPrivate {
		matches = matches && compare(this, that)
	}

	if !matches {
		metrics.FlashblocksMismatched.Add(ctx, 1, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "stream", Value: attribute.StringValue(this.stream)},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(fm.cfg.networkID)},
		))
	}
}
