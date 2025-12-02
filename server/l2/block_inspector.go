package l2

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/rpc"
	"github.com/flashbots/chain-monitor/types"
	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type BlockInspector struct {
	// parameters

	cfg *blockInspectorConfig

	// actors

	blockTicker *time.Ticker
	done        chan struct{}
	rpc         *rpc.RPC

	// state

	blockHeight uint64
	blocks      *types.RingBuffer[blockRecord]

	flashblocks   map[uint64]map[string][]*flashblockEvent
	mxFlashblocks sync.Mutex

	metrics *blockInspectorMetrics

	unwinding         bool
	unwindBlockHeight uint64
}

type blockInspectorConfig struct {
	reorgWindow int
	rpc         string

	blockTime   time.Duration
	chainID     *big.Int
	genesisTime uint64
	signer      ethtypes.EIP155Signer

	builderAddr            ethcommon.Address
	builderAddrInitialised bool

	builderPolicyAddr                        ethcommon.Address
	builderPolicyAddrInitialised             bool
	builderPolicyAddWorkloadIdEventSignature ethcommon.Hash
	builderPolicyAddWorkloadIdSignature      [4]byte
	builderPolicySignature                   [4]byte

	flashtestationsRegistryAddr            ethcommon.Address
	flashtestationsRegistryAddrInitialised bool
	flashtestationsRegistryEventSignature  ethcommon.Hash
	flashtestationsRegistrySignature       [4]byte

	flashblockNumberAddr            ethcommon.Address
	flashblockNumberAddrInitialised bool
	flashblockNumberSignature       [4]byte

	flashblocksPerBlock     int64
	flashtestationsPerBlock int64

	monitorAddr            ethcommon.Address
	monitorAddrInitialised bool

	dirPersistent string
}

type blockInspectorMetrics struct {
	blocksLanded int64
	blocksSeen   int64
	blocksMissed int64

	flashblocksLanded int64
	flashblocksMissed int64

	flashtestationsLanded int64
	flashtestationsMissed int64

	addWorkloadSeen    int64
	addWorkloadError   int64
	registrationsSeen  int64
	registrationsError int64
	statsMu            sync.Mutex // protects addWorkload* and registrations* fields

	monitorProbesLandedCount int64

	processBlockFailuresCount uint
}

func NewBlockInspector(cfg *config.L2) (*BlockInspector, error) {
	l := zap.L()

	bi := &BlockInspector{
		metrics: &blockInspectorMetrics{},

		cfg: &blockInspectorConfig{
			blockTime:               cfg.BlockTime,
			dirPersistent:           cfg.Dir.Persistent,
			flashblocksPerBlock:     cfg.FlashblocksPerBlock,
			flashtestationsPerBlock: cfg.FlashtestationsPerBlock,
			genesisTime:             cfg.GenesisTime,
			reorgWindow:             int(cfg.ReorgWindow/cfg.BlockTime) + 1,
			rpc:                     cfg.Rpc,
		},
	}

	{ // rpc
		rpc, err := rpc.New(cfg.NetworkID, cfg.Rpc, cfg.RpcFallback...)
		if err != nil {
			return nil, err
		}
		bi.rpc = rpc
	}

	{ // chainID, signer
		chainID, err := bi.rpc.NetworkID(context.Background())
		if err != nil {
			l.Error("Failed to request network id",
				zap.Error(err),
				zap.String("kind", "l2"),
			)
			return nil, err
		}

		bi.cfg.chainID = chainID
		bi.cfg.signer = ethtypes.NewEIP155Signer(chainID)
	}

	{ // builder address
		if cfg.MonitorBuilderAddress != "" {
			addr, err := ethcommon.ParseHexOrString(cfg.MonitorBuilderAddress)
			if err != nil {
				return nil, err
			}
			if len(addr) != 20 {
				return nil, fmt.Errorf(
					"invalid length for the builder address (want 20, got %d)",
					len(addr),
				)
			}
			copy(bi.cfg.builderAddr[:], addr)
			bi.cfg.builderAddrInitialised = true
		}
	}

	{ // builder policy contract
		if cfg.MonitorBuilderPolicyContract != "" {
			addr, err := ethcommon.ParseHexOrString(cfg.MonitorBuilderPolicyContract)
			if err != nil {
				return nil, err
			}
			if len(addr) != 20 {
				return nil, fmt.Errorf(
					"invalid length for the builder policy contract address (want 20, got %d)",
					len(addr),
				)
			}
			copy(bi.cfg.builderPolicyAddr[:], addr)
			bi.cfg.builderPolicyAddrInitialised = true
		}
		if cfg.MonitorBuilderPolicyContractFunctionSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorBuilderPolicyContractFunctionSignature))
			copy(bi.cfg.builderPolicySignature[:], h[:4])
		}
		if cfg.MonitorBuilderPolicyAddWorkloadIdSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorBuilderPolicyAddWorkloadIdSignature))
			copy(bi.cfg.builderPolicyAddWorkloadIdSignature[:], h[:4])
		}

		if cfg.MonitorBuilderPolicyAddWorkloadIdEventSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorBuilderPolicyAddWorkloadIdEventSignature))
			copy(bi.cfg.builderPolicyAddWorkloadIdEventSignature[:], h[:])
		}
	}

	{ // flashblock number contract
		if cfg.MonitorFlashblockNumberContract != "" {
			addr, err := ethcommon.ParseHexOrString(cfg.MonitorFlashblockNumberContract)
			if err != nil {
				return nil, err
			}
			if len(addr) != 20 {
				return nil, fmt.Errorf(
					"invalid length for the builder policy contract address (want 20, got %d)",
					len(addr),
				)
			}
			copy(bi.cfg.flashblockNumberAddr[:], addr)
			bi.cfg.flashblockNumberAddrInitialised = true
		}
		if cfg.MonitorFlashblockNumberContractFunctionSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorFlashblockNumberContractFunctionSignature))
			copy(bi.cfg.flashblockNumberSignature[:], h[:4])
		}
	}

	{ // flashtestations registry contract
		if cfg.MonitorFlashtestationRegistryContract != "" {
			addr, err := ethcommon.ParseHexOrString(cfg.MonitorFlashtestationRegistryContract)
			if err != nil {
				return nil, err
			}
			if len(addr) != 20 {
				return nil, fmt.Errorf(
					"invalid length for the flashtestations registry contract address (want 20, got %d)",
					len(addr),
				)
			}
			copy(bi.cfg.flashtestationsRegistryAddr[:], addr)
			bi.cfg.flashtestationsRegistryAddrInitialised = true
		}
		if cfg.MonitorFlashtestationRegistryFunctionSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorFlashtestationRegistryFunctionSignature))
			copy(bi.cfg.flashtestationsRegistrySignature[:], h[:4])
		}
		if cfg.MonitorFlashtestationRegistryEventSignature != "" {
			h := crypto.Keccak256Hash([]byte(cfg.MonitorFlashtestationRegistryEventSignature))
			copy(bi.cfg.flashtestationsRegistryEventSignature[:], h[:])
		}
	}

	{ // monitor tx address
		if cfg.ProbeTx.PrivateKey != "" {
			monitorKey, err := crypto.HexToECDSA(cfg.ProbeTx.PrivateKey)
			if err != nil {
				return nil, err
			}
			bi.cfg.monitorAddr = crypto.PubkeyToAddress(monitorKey.PublicKey)
			bi.cfg.monitorAddrInitialised = true
		}
	}

	{ // ticker
		now := time.Now()
		time.Sleep(now.Truncate(cfg.BlockTime).Add(cfg.BlockTime).Sub(now)) // align with block times
		bi.blockTicker = time.NewTicker(cfg.BlockTime)
	}

	{ // blocks, blockHeight
		blockHeight, err := bi.rpc.BlockNumber(context.Background())
		if err != nil {
			l.Error("Failed to request block number",
				zap.Error(err),
				zap.String("kind", "l2"),
				zap.String("rpc", cfg.Rpc),
			)
			return nil, err
		}

		if blockHeight > 0 {
			bi.blockHeight = blockHeight - 1
		}
		if cfg.Dir.Persistent != "" {
			fname := filepath.Join(cfg.Dir.Persistent, "blocks.json")
			if f, err := os.Open(fname); err == nil {
				blocks := types.NewRingBuffer[blockRecord](0)
				if err := json.NewDecoder(f).Decode(&blocks); err == nil {
					bi.blocks = blocks
					if head, ok := blocks.Head(); ok {
						bi.blockHeight = head.Number.Uint64()
					}
					l.Info("Loaded the state",
						zap.String("file_name", fname),
						zap.Uint64("block_height", bi.blockHeight),
					)
				} else {
					legacyBlocks := types.NewRingBuffer[blockRecordLegacy](0)
					if legacyErr := json.NewDecoder(f).Decode(&legacyBlocks); legacyErr == nil {
						blocks := types.NewRingBuffer[blockRecord](legacyBlocks.Length())
						for legacyBlock, ok := legacyBlocks.Pop(); ok; legacyBlock, ok = legacyBlocks.Pop() {
							blocks.Push(blockRecord{
								Number:               legacyBlock.Number,
								Hash:                 legacyBlock.Hash,
								Landed:               legacyBlock.Landed,
								FlashblocksCount:     legacyBlock.FlashblocksCount,
								FlashtestationsCount: 0,
							})
						}
						bi.blocks = blocks
						if head, ok := blocks.Head(); ok {
							bi.blockHeight = head.Number.Uint64()
						}
						l.Info("Loaded the legacy state",
							zap.String("file_name", fname),
							zap.Uint64("block_height", bi.blockHeight),
						)
					} else {
						l.Error("Failed to load the state",
							zap.Error(errors.Join(err, legacyErr)),
							zap.String("file_name", fname),
						)
					}
				}
			}
		}
		if bi.blocks == nil {
			bi.blocks = types.NewRingBuffer[blockRecord](int(blockHeight), int(cfg.ReorgWindow/cfg.BlockTime+1))
		}
	}

	return bi, nil
}

func (bi *BlockInspector) Run(
	ctx context.Context,
	flashblocks *<-chan *flashblockEvent,
) {
	if bi == nil {
		return
	}

	processingContext := logutils.ContextWithLogger(
		context.Background(),
		logutils.LoggerFromContext(ctx),
	)
	bi.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-bi.done:
				return
			case <-bi.blockTicker.C:
				bi.processNewBlocks(processingContext)
			}
		}
	}()

	if flashblocks != nil {
		bi.flashblocks = make(map[uint64]map[string][]*flashblockEvent)

		go func() {
			for {
				select {
				case <-bi.done:
					return
				case evt := <-*flashblocks:
					bi.processFlashblock(ctx, evt)
				}
			}
		}()
	}
}

func (bi *BlockInspector) Stop() {
	if bi == nil {
		return
	}

	l := zap.L()

	bi.blockTicker.Stop()
	bi.done <- struct{}{}
	bi.rpc.Close()

	if err := bi.persist(); err != nil {
		l.Error("Failed to persist the state",
			zap.Error(err),
		)
	}
}

func (bi *BlockInspector) Observe(_ context.Context, o otelapi.Observer) error {
	if bi == nil {
		return nil
	}

	o.ObserveInt64(metrics.BlockHeight, int64(bi.blockHeight), otelapi.WithAttributes(
		attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
		attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
	))

	if bi.cfg.monitorAddrInitialised {
		o.ObserveInt64(metrics.ProbesLandedCount, bi.metrics.monitorProbesLandedCount, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))
	}

	return nil
}

func (bi *BlockInspector) persist() error {
	l := zap.L()

	if bi.cfg.dirPersistent == "" {
		return nil
	}

	if err := os.MkdirAll(bi.cfg.dirPersistent, 0750); err != nil {
		return err
	}

	fname := filepath.Join(bi.cfg.dirPersistent, "blocks.json")
	f, err := os.Create(fname)
	if err != nil {
		return errors.Join(err, f.Close())
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(bi.blocks); err != nil {
		return err
	}

	l.Info("Persisted the state",
		zap.String("file_name", fname),
	)

	return nil
}

func (bi *BlockInspector) processNewBlocks(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	var blockHeight uint64

	if bi.cfg.genesisTime == 0 {
		if heightAccordingToRpc, err := bi.rpc.BlockNumber(ctx); err == nil {
			blockHeight = heightAccordingToRpc
		} else {
			l.Warn("Failed to request block number, skipping this round...",
				zap.Error(err),
				zap.String("kind", "l2"),
				zap.String("rpc", bi.cfg.rpc),
			)
			return
		}
	} else {
		blockHeight = (uint64(time.Now().Unix()) - bi.cfg.genesisTime) / uint64(bi.cfg.blockTime.Seconds())
	}

	if blockHeight == bi.blockHeight {
		l.Debug("Still at the same height, skipping...",
			zap.Uint64("block_height", blockHeight),
		)
		return
	}

	for b := bi.blockHeight + 1; b <= blockHeight; b++ {
		if err := bi.processBlock(ctx, b); err != nil {
			bi.metrics.processBlockFailuresCount++

			logLevel := zap.DebugLevel
			if bi.metrics.processBlockFailuresCount > 10 {
				logLevel = zap.WarnLevel
			}

			l.Log(logLevel, "Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block_number", blockHeight),
				zap.Uint("failures_count", bi.metrics.processBlockFailuresCount),
			)
			return
		} else {
			bi.metrics.processBlockFailuresCount = 0
		}
		bi.blockHeight = b
	}
}

func (bi *BlockInspector) processBlock(ctx context.Context, blockNumber uint64) error {
	l := logutils.LoggerFromContext(ctx).With(
		zap.Uint64("block_number", blockNumber),
		zap.String("kind", "l2"),
	)
	ctx = logutils.ContextWithLogger(ctx, l)

	l.Debug("Processing new l2 block")

	block, err := bi.rpc.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		l.Warn("Failed to request block by number",
			zap.Error(err),
		)
		return err
	}

	if delay := time.Now().Unix() - int64(block.Time()); delay > 2 {
		l.Warn("Processing stale block",
			zap.Int64("delay", delay),
		)
	}

	metrics.TxPerBlock.Record(ctx, int64(len(block.Transactions())))
	metrics.GasPerBlock.Record(ctx, int64(block.GasUsed()))

	if blockNumber > 0 {
		if previous, ok := bi.blocks.At(int(blockNumber) - 1); ok {
			if previous.Hash.Cmp(block.ParentHash()) != 0 {
				if !bi.unwinding {
					l.Info("Chain reorg detected via hash mismatch, starting the unwind...",
						zap.String("parent_hash", block.ParentHash().String()),
						zap.String("old_parent_hash", previous.Hash.String()),
					)
					bi.unwinding = true
					bi.unwindBlockHeight = blockNumber
					return bi.processReorgUnwind(ctx)
				}

				l.Debug("Continuing the unwind...")
				return bi.processReorgUnwind(ctx)
			}
		}

		if bi.unwinding {
			var depth int64
			if bi.unwindBlockHeight > blockNumber {
				depth = int64(bi.unwindBlockHeight - blockNumber)
			} else {
				depth = int64(blockNumber - bi.unwindBlockHeight)
			}

			metrics.ReorgsCount.Add(ctx, 1, otelapi.WithAttributes(
				attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
				attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
			))

			metrics.ReorgDepth.Record(ctx, depth, otelapi.WithAttributes(
				attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
				attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
			))

			l.Info("Finished the unwind",
				zap.Int64("reorg_depth", depth),
				zap.Uint64("old_block_number", bi.unwindBlockHeight),
			)

			bi.unwinding = false
			bi.unwindBlockHeight = 0
		}
	}

	bi.metrics.blocksSeen++
	metrics.BlocksSeenCount.Record(ctx, bi.metrics.blocksSeen, otelapi.WithAttributes(
		attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
		attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
	))

	expectedBuilderTxData := []byte(fmt.Sprintf("Block Number: %s", block.Number().String()))

	var builderTxCount, failedTxCount, flashblockNumberTxCount, flashtestationsTxCount int64

	for _, tx := range block.Transactions() {
		if bi.cfg.builderAddrInitialised && bi.isBuilderTx(ctx, block, tx, expectedBuilderTxData) {
			builderTxCount++
		}

		if bi.cfg.builderPolicyAddrInitialised && bi.isBuilderPolicyBlockProofTx(tx) {
			flashtestationsTxCount++
			builderTxCount++
		}

		if bi.cfg.builderPolicyAddrInitialised && bi.isBuilderPolicyAddWorkloadIdTx(tx) {
			go func() {
				bi.handleAddWorkloadIdTx(ctx, tx.Hash())
				bi.metrics.statsMu.Lock()
				addWorkloadError := bi.metrics.addWorkloadError
				bi.metrics.statsMu.Unlock()
				metrics.WorkloadAddedToPolicyErrorCount.Record(ctx, addWorkloadError, otelapi.WithAttributes(
					attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
					attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
				))
			}()
		}

		if bi.cfg.flashblockNumberAddrInitialised && bi.isFlashblockNumberTx(ctx, block, tx) {
			flashblockNumberTxCount++
			builderTxCount++
		}

		if bi.cfg.flashtestationsRegistryAddrInitialised && bi.isFlashtestationsRegisterTx(tx) {
			go func() {
				bi.handleRegistrationTx(ctx, tx.Hash())
				bi.metrics.statsMu.Lock()
				registrationsError := bi.metrics.registrationsError
				bi.metrics.statsMu.Unlock()
				metrics.RegisteredFlashtestationsErrorCount.Record(ctx, registrationsError, otelapi.WithAttributes(
					attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
					attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
				))
			}()
		}

		if bi.cfg.monitorAddrInitialised {
			if isProbeTx, sent, latency := bi.isProbeTx(ctx, block, tx); isProbeTx {
				l.Debug("Detected probe transaction",
					zap.Uint64("latency", latency),
					zap.Uint64("sent", sent),
					zap.Uint64("landed", block.Time()),
				)
				bi.metrics.monitorProbesLandedCount++
				metrics.ProbesLatency.Record(ctx, int64(latency))
			}
		}

		if gasPrice := tx.GasPrice().Int64(); gasPrice > 0 {
			metrics.GasPrice.Record(ctx, gasPrice)
			metrics.GasPricePerTx.Record(ctx, gasPrice)
		}
	}

	switch builderTxCount {
	case 0:
		bi.blocks.Push(blockRecord{
			Number:           block.Number(),
			Hash:             block.Hash(),
			Landed:           false,
			FlashblocksCount: flashblockNumberTxCount,
		})
		bi.metrics.blocksMissed++

		metrics.BlocksMissedCount.Record(ctx, bi.metrics.blocksMissed, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		metrics.BlockMissed.Record(ctx, int64(blockNumber), otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		l.Warn("Builder had missed a block",
			zap.Int64("blocks_landed", bi.metrics.blocksLanded),
			zap.Int64("blocks_missed", bi.metrics.blocksMissed),
			zap.Int64("blocks_seen", bi.metrics.blocksSeen),
		)

	default:
		l.Debug("More than 1 builder transaction found in a block",
			zap.Int("count", int(builderTxCount)),
		)
		fallthrough

	case 1:
		bi.blocks.Push(blockRecord{
			Number:           block.Number(),
			Hash:             block.Hash(),
			Landed:           true,
			FlashblocksCount: flashblockNumberTxCount,
		})
		bi.metrics.blocksLanded++

		metrics.BlocksLandedCount.Record(ctx, bi.metrics.blocksLanded, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))
	}

	if bi.cfg.flashblockNumberAddrInitialised {
		bi.metrics.flashblocksLanded += flashblockNumberTxCount
		bi.metrics.flashblocksMissed += (bi.cfg.flashblocksPerBlock - flashblockNumberTxCount)

		metrics.FlashblocksLandedCount.Record(ctx, bi.metrics.flashblocksLanded, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))
		metrics.FlashblocksMissedCount.Record(ctx, bi.metrics.flashblocksMissed, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		if flashblockNumberTxCount < bi.cfg.flashblocksPerBlock {
			l.Warn("Builder missed flashblocks",
				zap.Int64("count", bi.cfg.flashblocksPerBlock-flashblockNumberTxCount),
				zap.Int64("flashblocks_landed", bi.metrics.flashblocksLanded),
				zap.Int64("flashblocks_missed", bi.metrics.flashblocksMissed),
			)
		}
	}

	if bi.cfg.builderPolicyAddrInitialised {
		bi.metrics.flashtestationsLanded += flashtestationsTxCount
		bi.metrics.flashtestationsMissed += (bi.cfg.flashtestationsPerBlock - flashtestationsTxCount)

		metrics.FlashtestationsLandedCount.Record(ctx, bi.metrics.flashtestationsLanded, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		metrics.FlashtestationsMissedCount.Record(ctx, bi.metrics.flashtestationsMissed, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		if flashtestationsTxCount < bi.cfg.flashtestationsPerBlock {
			l.Warn("Builder missed flashtestations",
				zap.Int64("count", bi.cfg.flashtestationsPerBlock-flashtestationsTxCount),
				zap.Int64("flashtestations_landed", bi.metrics.flashtestationsLanded),
				zap.Int64("flashtestations_missed", bi.metrics.flashtestationsMissed),
			)
		}
	}

	if bi.flashblocks != nil {
		bi.mxFlashblocks.Lock()

		blockHash := strings.TrimPrefix(strings.ToLower(block.Hash().String()), "0x")

		if payloads, seen := bi.flashblocks[blockNumber]; seen {
			matched := false
		scanningPayloads:
			for _, flashblocks := range payloads {
				for idx := len(flashblocks) - 1; idx >= 0; idx-- {
					if fb := flashblocks[idx]; fb != nil &&
						strings.TrimPrefix(strings.ToLower(fb.flashblock.Diff.BlockHash), "0x") == blockHash {
						// ---
						matched = true

						if fb.flashblock.Index < len(flashblocks)-1 {
							dropped := len(flashblocks) - fb.flashblock.Index - 1
							metrics.FlashblocksDropped.Add(ctx, int64(dropped), otelapi.WithAttributes(
								attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
								attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
							))
							for jdx := fb.flashblock.Index + 1; jdx < len(flashblocks); jdx++ {
								if dfb := flashblocks[jdx]; dfb != nil {
									l.Warn("Flashblock was dropped",
										zap.Any("flashblock", fb.flashblock),
									)
								}
							}
						}

						break scanningPayloads
					}
				}
			}

			if !matched {
				dropped := 0
				for _, flashblocks := range payloads {
					dropped += len(flashblocks)
					for _, fb := range flashblocks {
						if fb != nil {
							l.Warn("Flashblock was dropped",
								zap.Any("flashblock", fb.flashblock),
							)
						}
					}
				}
				metrics.FlashblocksDropped.Add(ctx, int64(dropped), otelapi.WithAttributes(
					attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
					attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
				))
			}
		}

		for fbBlockNumber := range bi.flashblocks {
			if fbBlockNumber <= blockNumber {
				delete(bi.flashblocks, fbBlockNumber)
			}
		}

		bi.mxFlashblocks.Unlock()
	}

	metrics.FailedTxPerBlock.Record(ctx, failedTxCount)

	if bi.blocks.Length() > bi.cfg.reorgWindow {
		_, _ = bi.blocks.Pop()
	}

	return nil
}

func (bi *BlockInspector) processFlashblock(ctx context.Context, evt *flashblockEvent) {
	l := logutils.LoggerFromContext(ctx).With(
		zap.Uint64("block_number", evt.flashblock.Metadata.BlockNumber),
		zap.String("kind", "l2"),
	)

	fb := evt.flashblock

	bi.mxFlashblocks.Lock()
	defer bi.mxFlashblocks.Unlock()

	if _, exists := bi.flashblocks[fb.Metadata.BlockNumber]; !exists {
		bi.flashblocks[fb.Metadata.BlockNumber] = make(map[string][]*flashblockEvent)
	}
	if _, exists := bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId]; !exists {
		bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId] = make([]*flashblockEvent, 0, bi.cfg.flashblocksPerBlock)
	}

	if len(bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId]) > fb.Index {
		existing := bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId][fb.Index]
		if existing == nil {
			bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId][fb.Index] = evt
			return // out-of-order delivery
		}
		if fb.Equal(existing.flashblock) {
			return // double-delivery of the same flashblock
		}
		l.Warn("Received different flashblocks for the same payload id and index",
			zap.Any("this", fb),
			zap.Any("that", existing),
		)
		bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId][fb.Index] = evt
	} else {
		// fill up the gaps (if any)
		for idx := len(bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId]); idx <= fb.Index; idx++ {
			bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId] = append(
				bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId],
				nil,
			)
		}
		// store
		bi.flashblocks[fb.Metadata.BlockNumber][fb.PayloadId][fb.Index] = evt
	}
}

func (bi *BlockInspector) processReorgUnwind(ctx context.Context) error {
	l := logutils.LoggerFromContext(ctx)

	defer func() {
		metrics.BlocksSeenCount.Record(ctx, bi.metrics.blocksSeen, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		metrics.BlocksLandedCount.Record(ctx, bi.metrics.blocksLanded, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))

		metrics.BlocksMissedCount.Record(ctx, bi.metrics.blocksMissed, otelapi.WithAttributes(
			attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
			attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		))
	}()

	for {
		br, ok := bi.blocks.Pick()
		if !ok {
			break
		}
		bi.blockHeight = br.Number.Uint64() - 1

		bi.metrics.blocksSeen--
		if br.Landed {
			bi.metrics.blocksLanded--
		} else {
			bi.metrics.blocksMissed--
			l.Info("Missed block was reorgd",
				zap.Uint64("block_number", br.Number.Uint64()),
				zap.Int64("blocks_landed", bi.metrics.blocksLanded),
				zap.Int64("blocks_missed", bi.metrics.blocksMissed),
				zap.Int64("blocks_seen", bi.metrics.blocksSeen),
			)
		}

		if bi.cfg.flashblockNumberAddrInitialised {
			bi.metrics.flashblocksLanded -= br.FlashblocksCount
			bi.metrics.flashblocksMissed -= (bi.cfg.flashblocksPerBlock - br.FlashblocksCount)

			if br.FlashblocksCount < bi.cfg.flashblocksPerBlock {
				l.Info("Missed flashblocks were reorgd",
					zap.Int64("count", bi.cfg.flashblocksPerBlock-br.FlashblocksCount),
					zap.Uint64("block_number", br.Number.Uint64()),
				)
			}
		}

		block, err := bi.rpc.BlockByNumber(ctx, br.Number)
		if err != nil {
			l.Warn("Failed to request block by number, skipping this round of unwind...",
				zap.Error(err),
				zap.String("number", br.Number.String()),
				zap.String("kind", "l2"),
				zap.String("rpc", bi.cfg.rpc),
			)
			return err
		}

		if block.Hash().Cmp(br.Hash) == 0 {
			return nil
		}

		l.Info("Unwinding...",
			zap.Uint64("block_number", bi.blockHeight),
		)
	}

	return nil
}

func (bi *BlockInspector) isBuilderTx(
	ctx context.Context,
	block *ethtypes.Block,
	tx *ethtypes.Transaction,
	expectedData []byte,
) bool {
	if tx == nil || tx.To() == nil || tx.To().Cmp(ethcommon.Address{}) != 0 {
		return false // builder's tx burns eth by sending 0 ETH to zero address
	}

	if slices.Compare(tx.Data(), expectedData) != 0 {
		return false
	}

	from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		l := logutils.LoggerFromContext(ctx)

		l.Warn("Failed to determine the sender for builder transaction",
			zap.Error(err),
			zap.String("tx", tx.Hash().Hex()),
			zap.String("block", block.Number().String()),
		)

		return false
	}

	return from.Cmp(bi.cfg.builderAddr) == 0
}

func (bi *BlockInspector) isBuilderPolicyBlockProofTx(
	tx *ethtypes.Transaction,
) bool {
	if tx == nil || tx.Rejected() {
		return false
	}

	if tx.To() == nil || tx.To().Cmp(bi.cfg.builderPolicyAddr) != 0 {
		return false
	}

	if len(tx.Data()) < len(bi.cfg.builderPolicySignature) {
		return false
	}

	return slices.Compare(tx.Data()[:4], bi.cfg.builderPolicySignature[:]) == 0
}

func (bi *BlockInspector) isBuilderPolicyAddWorkloadIdTx(
	tx *ethtypes.Transaction,
) bool {
	if tx == nil || tx.Rejected() {
		return false
	}

	if tx.To() == nil || tx.To().Cmp(bi.cfg.builderPolicyAddr) != 0 {
		return false
	}

	if len(tx.Data()) < len(bi.cfg.builderPolicyAddWorkloadIdSignature) {
		return false
	}

	return slices.Compare(tx.Data()[:4], bi.cfg.builderPolicyAddWorkloadIdSignature[:]) == 0
}

func (bi *BlockInspector) isFlashtestationsRegisterTx(
	tx *ethtypes.Transaction,
) bool {
	if tx == nil || tx.Rejected() {
		return false
	}

	if tx.To() == nil || tx.To().Cmp(bi.cfg.flashtestationsRegistryAddr) != 0 {
		return false
	}

	if len(tx.Data()) < len(bi.cfg.flashtestationsRegistrySignature) {
		return false
	}

	return slices.Compare(tx.Data()[:4], bi.cfg.flashtestationsRegistrySignature[:]) == 0
}

func (bi *BlockInspector) isFlashblockNumberTx(
	ctx context.Context,
	block *ethtypes.Block,
	tx *ethtypes.Transaction,
) bool {
	if tx == nil || tx.Rejected() {
		return false
	}

	if tx.To() == nil || tx.To().Cmp(bi.cfg.flashblockNumberAddr) != 0 {
		return false
	}

	if len(tx.Data()) < len(bi.cfg.flashblockNumberSignature) {
		return false
	}

	if slices.Compare(tx.Data()[:4], bi.cfg.flashblockNumberSignature[:]) != 0 {
		return false
	}

	from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		l := logutils.LoggerFromContext(ctx)

		l.Warn("Failed to determine the sender for flashblock number transaction",
			zap.Error(err),
			zap.String("tx", tx.Hash().Hex()),
			zap.String("block", block.Number().String()),
		)

		return false
	}

	if from.Cmp(bi.cfg.builderAddr) != 0 {
		return false
	}

	return true
}

func (bi *BlockInspector) isProbeTx(
	ctx context.Context,
	block *ethtypes.Block,
	tx *ethtypes.Transaction,
) (isProbeTx bool, txEpoch, latency uint64) {
	if tx == nil || tx.To() == nil || tx.To().Cmp(ethcommon.Address{}) != 0 {
		return false, 0, 0 // probe tx burns eth by sending 0 ETH to zero address
	}

	if len(tx.Data()) != 8 {
		return false, 0, 0
	}

	from, err := ethtypes.Sender(ethtypes.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		l := logutils.LoggerFromContext(ctx)

		l.Warn("Failed to determine the sender for probe transaction",
			zap.Error(err),
			zap.String("tx", tx.Hash().Hex()),
			zap.String("block", block.Number().String()),
		)

		return false, 0, 0
	}

	if from.Cmp(bi.cfg.monitorAddr) != 0 {
		return false, 0, 0
	}

	blockEpoch := block.Time()
	txEpoch = binary.BigEndian.Uint64(tx.Data())
	if blockEpoch < txEpoch {
		l := logutils.LoggerFromContext(ctx)
		l.Warn("Block time precedes the monitoring transaction's time",
			zap.String("block", block.Number().String()),
			zap.String("tx", tx.Hash().Hex()),
			zap.Uint64("block_epoch", blockEpoch),
			zap.Time("block_time", time.Unix(int64(blockEpoch), 0)),
			zap.Uint64("tx_epoch", txEpoch),
			zap.Time("tx_time", time.Unix(int64(txEpoch), 0)),
		)
	}
	latency = blockEpoch - txEpoch

	return true, txEpoch, latency
}

func (bi *BlockInspector) handleRegistrationTx(ctx context.Context, txHash ethcommon.Hash) {
	l := logutils.LoggerFromContext(ctx)

	teeAddress, rawQuote, err := bi.getTEEAddressAndQuoteFromTx(ctx, txHash)
	if err != nil {
		l.Warn("Failed to get register flashtestations transaction receipt",
			zap.Error(err),
			zap.String("tx", txHash.Hex()),
		)
		bi.metrics.statsMu.Lock()
		bi.metrics.registrationsError++
		bi.metrics.statsMu.Unlock()
		return
	}

	workloadId, err := ComputeWorkloadID(rawQuote)
	if err != nil {
		l.Warn("Failed to compute workload id",
			zap.Error(err),
			zap.String("tx", txHash.Hex()),
		)
		bi.metrics.statsMu.Lock()
		bi.metrics.registrationsError++
		bi.metrics.statsMu.Unlock()
		return
	}

	bi.metrics.statsMu.Lock()
	bi.metrics.registrationsSeen++
	registrationsSeen := bi.metrics.registrationsSeen
	bi.metrics.statsMu.Unlock()
	metrics.RegisteredFlashtestationsCount.Record(ctx, registrationsSeen, otelapi.WithAttributes(
		attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
		attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
		attribute.KeyValue{Key: "tee_address", Value: attribute.StringValue(teeAddress.Hex())},
		attribute.KeyValue{Key: "workload_id", Value: attribute.StringValue(hex.EncodeToString(workloadId[:]))},
	))

	l.Info("TEE service registered",
		zap.String("teeAddress", teeAddress.Hex()),
		zap.String("workloadId", hex.EncodeToString(workloadId[:])),
	)
}

func (bi *BlockInspector) handleAddWorkloadIdTx(ctx context.Context, txHash ethcommon.Hash) {
	l := logutils.LoggerFromContext(ctx)

	receipt, err := bi.rpc.TransactionReceipt(ctx, txHash)
	if err != nil {
		l.Warn("Failed to get add workload id transaction receipt",
			zap.Error(err),
			zap.String("tx", txHash.Hex()),
		)
		bi.metrics.statsMu.Lock()
		bi.metrics.addWorkloadError++
		bi.metrics.statsMu.Unlock()
		return
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		l.Warn("Add workload id transaction did not succeed",
			zap.String("tx", txHash.Hex()),
		)
		bi.metrics.statsMu.Lock()
		bi.metrics.addWorkloadError++
		bi.metrics.statsMu.Unlock()
		return
	}

	for _, log := range receipt.Logs {
		if len(log.Topics) > 1 && log.Topics[0] == bi.cfg.builderPolicyAddWorkloadIdEventSignature {
			// workloadId is bytes32 (32 bytes), stored directly in Topics[1]
			// log.Topics[1] is a common.Hash which is [32]byte
			workloadId := [32]byte(log.Topics[1])

			l.Info("Workload added to policy",
				zap.String("workloadId", hex.EncodeToString(workloadId[:])),
				zap.String("tx", txHash.Hex()),
			)
			bi.metrics.statsMu.Lock()
			bi.metrics.addWorkloadSeen++
			addWorkloadSeen := bi.metrics.addWorkloadSeen
			bi.metrics.statsMu.Unlock()
			metrics.WorkloadAddedToPolicyCount.Record(ctx, addWorkloadSeen, otelapi.WithAttributes(
				attribute.KeyValue{Key: "kind", Value: attribute.StringValue("l2")},
				attribute.KeyValue{Key: "network_id", Value: attribute.Int64Value(bi.cfg.chainID.Int64())},
				attribute.KeyValue{Key: "workload_id", Value: attribute.StringValue(hex.EncodeToString(workloadId[:]))},
			))
			return
		}
	}

	bi.metrics.statsMu.Lock()
	bi.metrics.addWorkloadError++
	bi.metrics.statsMu.Unlock()
	l.Warn("WorkloadAddedToPolicy event not found in transaction",
		zap.String("tx", txHash.Hex()),
	)
}

// Extract TEE address and raw quote from TEEServiceRegistered event
func (bi *BlockInspector) getTEEAddressAndQuoteFromTx(ctx context.Context, txHash ethcommon.Hash) (ethcommon.Address, []byte, error) {
	receipt, err := bi.rpc.TransactionReceipt(ctx, txHash)
	if err != nil {
		return ethcommon.Address{}, nil, err
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return ethcommon.Address{}, nil, fmt.Errorf("Register tee transaction did not succeed %s", txHash.Hex())
	}

	// Define the event arguments for decoding (non-indexed parameters only)
	// event TEEServiceRegistered(address indexed teeAddress, bytes rawQuote, bool alreadyExists);
	bytesType, _ := abi.NewType("bytes", "", nil)
	boolType, _ := abi.NewType("bool", "", nil)

	eventABI := abi.Arguments{
		{Type: bytesType}, // rawQuote
		{Type: boolType},  // alreadyExists
	}

	for _, log := range receipt.Logs {
		if len(log.Topics) > 1 && log.Topics[0] == bi.cfg.flashtestationsRegistryEventSignature {
			// TEE address is in Topics[1] (indexed parameter)
			teeAddress := ethcommon.BytesToAddress(log.Topics[1].Bytes())

			// Decode the data field to get rawQuote and alreadyExists
			decoded, err := eventABI.Unpack(log.Data)
			if err != nil {
				return ethcommon.Address{}, nil, fmt.Errorf("failed to decode event data: %w", err)
			}

			if len(decoded) < 2 {
				return ethcommon.Address{}, nil, fmt.Errorf("unexpected decoded data length: %d", len(decoded))
			}

			rawQuote, ok := decoded[0].([]byte)
			if !ok {
				return ethcommon.Address{}, nil, fmt.Errorf("failed to type assert rawQuote")
			}

			return teeAddress, rawQuote, nil
		}
	}

	return ethcommon.Address{}, nil, fmt.Errorf("TEEServiceRegistered event not found in tx %s", txHash.Hex())
}
