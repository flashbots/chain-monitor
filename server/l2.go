package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/rpc"
	"github.com/flashbots/chain-monitor/types"
	"github.com/flashbots/chain-monitor/utils"

	"go.uber.org/zap"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

type L2 struct {
	cfg *config.L2

	rpc    *rpc.RPC
	ticker *time.Ticker

	builderAddr            ethcommon.Address
	builderPolicyAddr      ethcommon.Address
	builderPolicySignature [4]byte
	chainID                *big.Int
	monitorAddr            ethcommon.Address
	monitorKey             *ecdsa.PrivateKey
	reorgWindow            int
	signer                 ethtypes.EIP155Signer
	wallets                map[string]ethcommon.Address

	blockHeight uint64
	blocks      *types.RingBuffer[blockRecord]

	blocksLanded int64
	blocksMissed int64
	blocksSeen   int64

	processBlockFailuresCount uint

	monitorNonce       uint64
	monitorResetTicker *time.Ticker

	monitorProbesFailedCount int64
	monitorProbesLandedCount int64
	monitorProbesSentCount   int64

	unwinding         bool
	unwindBlockHeight uint64
}

func newL2(cfg *config.L2) (*L2, error) {
	l := zap.L()

	l2 := &L2{
		cfg:         cfg,
		reorgWindow: int(cfg.ReorgWindow/cfg.BlockTime) + 1,
		wallets:     make(map[string]ethcommon.Address, len(cfg.MonitorWalletAddresses)),
	}

	{ // ticker
		now := time.Now()
		time.Sleep(now.Truncate(cfg.BlockTime).Add(cfg.BlockTime).Sub(now)) // align with block times
		l2.ticker = time.NewTicker(cfg.BlockTime)
	}

	if cfg.MonitorBuilderAddress != "" { // builderAddr
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
		copy(l2.builderAddr[:], addr)
	}

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
		copy(l2.builderPolicyAddr[:], addr)
	}

	if cfg.MonitorBuilderPolicyContractFunctionSignature != "" {
		h := crypto.Keccak256Hash([]byte(cfg.MonitorBuilderPolicyContractFunctionSignature))
		copy(l2.builderPolicySignature[:], h[:4])
	}

	if cfg.ProbeTx.PrivateKey != "" { // monitorAddr, monitorKey
		monitorKey, err := crypto.HexToECDSA(cfg.ProbeTx.PrivateKey)
		if err != nil {
			return nil, err
		}
		l2.monitorKey = monitorKey
		l2.monitorAddr = crypto.PubkeyToAddress(monitorKey.PublicKey)

		l2.monitorResetTicker = time.NewTicker(cfg.ProbeTx.ResetInterval)
	}

	for name, addrStr := range cfg.MonitorWalletAddresses { // wallets
		var addr ethcommon.Address
		addrBytes, err := ethcommon.ParseHexOrString(addrStr)
		if err != nil {
			return nil, err
		}
		if len(addrBytes) != 20 {
			return nil, fmt.Errorf(
				"invalid length for the l2 wallet address (want 20, got %d)",
				len(addr),
			)
		}
		copy(addr[:], addrBytes)
		l2.wallets[name] = addr
	}

	{ // rpc
		rpc, err := rpc.New(cfg.Rpc, cfg.RpcFallback...)
		if err != nil {
			return nil, err
		}
		l2.rpc = rpc
	}

	{ // chainID, signer
		chainID, err := l2.rpc.NetworkID(context.Background())
		if err != nil {
			l.Error("Failed to request network id",
				zap.Error(err),
				zap.String("kind", "l2"),
			)
			return nil, err
		}

		l2.chainID = chainID
		l2.signer = ethtypes.NewEIP155Signer(chainID)
	}

	{ // blocks, blockHeight
		blockHeight, err := l2.rpc.BlockNumber(context.Background())
		if err != nil {
			l.Error("Failed to request block number",
				zap.Error(err),
				zap.String("kind", "l2"),
				zap.String("rpc", cfg.Rpc),
			)
			return nil, err
		}

		if blockHeight > 0 {
			l2.blockHeight = blockHeight - 1
		}
		if cfg.Dir.Persistent != "" {
			fname := filepath.Join(cfg.Dir.Persistent, "blocks.json")
			if f, err := os.Open(fname); err == nil {
				blocks := types.NewRingBuffer[blockRecord](0)
				if err := json.NewDecoder(f).Decode(&blocks); err == nil {
					l2.blocks = blocks
					if head, ok := blocks.Head(); ok {
						l2.blockHeight = head.Number.Uint64()
					}
					l.Info("Loaded the state",
						zap.String("file_name", fname),
						zap.Uint64("block_height", l2.blockHeight),
					)
				} else {
					l.Error("Failed to load the state",
						zap.Error(err),
						zap.String("file_name", fname),
					)
				}
			}
		}
		if l2.blocks == nil {
			l2.blocks = types.NewRingBuffer[blockRecord](int(blockHeight), int(cfg.ReorgWindow/cfg.BlockTime+1))
		}
	}

	return l2, nil
}

func (l2 *L2) run(ctx context.Context) {
	tick := func() {}

	if l2.monitorKey != nil {
		tick = func() {
			l2.processNewBlocks(ctx)
			l2.sendProbeTx(ctx)
		}
	} else {
		tick = func() {
			l2.processNewBlocks(ctx)
		}
	}

	go func() {
		for {
			<-l2.ticker.C
			tick()
		}
	}()

	if l2.monitorResetTicker != nil {
		go func() {
			for {
				<-l2.monitorResetTicker.C
				l2.checkAndResetProbeTxNonce(ctx)
			}
		}()
	}
}

func (l2 *L2) persist() error {
	l := zap.L()

	if l2.cfg.Dir.Persistent == "" {
		return nil
	}

	if err := os.MkdirAll(l2.cfg.Dir.Persistent, 0750); err != nil {
		return err
	}

	fname := filepath.Join(l2.cfg.Dir.Persistent, "blocks.json")
	f, err := os.Create(fname)
	if err != nil {
		return errors.Join(err, f.Close())
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(l2.blocks); err != nil {
		return err
	}

	l.Info("Persisted the state",
		zap.String("file_name", fname),
	)

	return nil
}

func (l2 *L2) stop() {
	l := zap.L()

	if err := l2.persist(); err != nil {
		l.Error("Failed to persist the state",
			zap.Error(err),
		)
	}

	l2.ticker.Stop()
	if l2.monitorResetTicker != nil {
		l2.monitorResetTicker.Stop()
	}
}

func (l2 *L2) processNewBlocks(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	chainID, err := l2.rpc.NetworkID(context.Background())
	if err != nil {
		l.Warn("Failed to request chain id, skipping this round...",
			zap.Error(err),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.Rpc),
		)
		return
	}
	if chainID.Cmp(l2.chainID) != 0 {
		l.Warn("Unexpected chain id, skipping this round...",
			zap.Uint64("expected", l2.chainID.Uint64()),
			zap.Uint64("got", chainID.Uint64()),
		)
	}

	blockHeight, err := l2.rpc.BlockNumber(ctx)
	if err != nil {
		l.Warn("Failed to request block number, skipping this round...",
			zap.Error(err),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.Rpc),
		)
		return
	}

	if blockHeight == l2.blockHeight {
		l.Debug("Still at the same height, skipping...",
			zap.Uint64("block_height", blockHeight),
		)
		return
	}

	for b := l2.blockHeight + 1; b <= blockHeight; b++ {
		if err := l2.processBlock(ctx, b); err != nil {
			l2.processBlockFailuresCount++

			logLevel := zap.DebugLevel
			if l2.processBlockFailuresCount > 10 {
				logLevel = zap.WarnLevel
			}

			l.Log(logLevel, "Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block_number", blockHeight),
				zap.Uint("failures_count", l2.processBlockFailuresCount),
			)
			return
		} else {
			l2.processBlockFailuresCount = 0
		}
		l2.blockHeight = b
	}
}

func (l2 *L2) processBlock(ctx context.Context, blockNumber uint64) error {
	l := logutils.LoggerFromContext(ctx).With(
		zap.Uint64("block_number", blockNumber),
		zap.String("kind", "l2"),
	)
	ctx = logutils.ContextWithLogger(ctx, l)

	l.Debug("Processing new l2 block")

	block, err := l2.rpc.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
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
		if previous, ok := l2.blocks.At(int(blockNumber) - 1); ok {
			if previous.Hash.Cmp(block.ParentHash()) != 0 {
				if !l2.unwinding {
					l.Info("Chain reorg detected via hash mismatch, starting the unwind...",
						zap.String("parent_hash", block.ParentHash().String()),
						zap.String("old_parent_hash", previous.Hash.String()),
					)
					l2.unwinding = true
					l2.unwindBlockHeight = blockNumber
					return l2.processReorgUnwind(ctx)
				}

				l.Debug("Continuing the unwind...")
				return l2.processReorgUnwind(ctx)
			}
		}

		if l2.unwinding {
			depth := l2.unwindBlockHeight - blockNumber

			metrics.ReorgsCount.Add(ctx, 1)
			metrics.ReorgDepth.Record(ctx, int64(depth))

			l.Info("Finished the unwind",
				zap.Uint64("reorg_depth", depth),
				zap.Uint64("old_block_number", l2.unwindBlockHeight),
			)

			l2.unwinding = false
			l2.unwindBlockHeight = 0
		}
	}

	l2.blocksSeen++
	metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)

	expectedBuilderTxData := []byte(fmt.Sprintf("Block Number: %s", block.Number().String()))

	var builderTxCount, failedTxCount int64
	for _, tx := range block.Transactions() {
		if l2.cfg.MonitorBuilderAddress != "" && l2.isBuilderTx(ctx, block, tx, expectedBuilderTxData) {
			builderTxCount++
		}

		if l2.cfg.MonitorBuilderPolicyContract != "" && l2.isBuilderPolicyTx(tx) {
			builderTxCount++
		}

		if l2.monitorKey != nil {
			if isProbeTx, sent, latency := l2.isProbeTx(ctx, block, tx); isProbeTx {
				l.Debug("Detected probe transaction",
					zap.Uint64("latency", latency),
					zap.Uint64("sent", sent),
					zap.Uint64("landed", block.Time()),
				)
				l2.monitorProbesLandedCount++
				metrics.ProbesLatency.Record(ctx, int64(latency))
			}
		}

		if gasPrice := tx.GasPrice().Int64(); gasPrice > 0 {
			metrics.GasPrice.Record(ctx, gasPrice)
			metrics.GasPricePerTx.Record(ctx, gasPrice)
		}

		if l2.cfg.MonitorTxReceipts {
			if receipt, err := l2.rpc.TransactionReceipt(ctx, tx.Hash()); err == nil {
				if receipt != nil {
					metrics.GasPerTx.Record(ctx, int64(receipt.GasUsed))
					if receipt.Status == ethtypes.ReceiptStatusFailed {
						failedTxCount++
					}
				}
			} else {
				l.Warn("Failed to get transaction receipt",
					zap.Error(err),
					zap.String("tx", tx.Hash().Hex()),
				)
			}
		}
	}

	switch builderTxCount {
	case 0:
		l2.blocks.Push(blockRecord{
			Number: block.Number(),
			Hash:   block.Hash(),
			Landed: false,
		})
		l2.blocksMissed++
		metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Int64("blocks_landed", l2.blocksLanded),
			zap.Int64("blocks_missed", l2.blocksMissed),
			zap.Int64("blocks_seen", l2.blocksSeen),
		)

	default:
		l.Debug("More than 1 builder transaction found in a block",
			zap.Int("count", int(builderTxCount)),
		)
		fallthrough

	case 1:
		l2.blocks.Push(blockRecord{
			Number: block.Number(),
			Hash:   block.Hash(),
			Landed: true,
		})
		l2.blocksLanded++
		metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
	}

	metrics.FailedTxPerBlock.Record(ctx, failedTxCount)

	if l2.blocks.Length() > l2.reorgWindow {
		_, _ = l2.blocks.Pop()
	}

	return nil
}

func (l2 *L2) processReorgUnwind(ctx context.Context) error {
	l := logutils.LoggerFromContext(ctx)

	defer func() {
		metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)
		metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
		metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)
	}()

	for br, ok := l2.blocks.Pick(); ok; {
		l2.blockHeight = br.Number.Uint64() - 1

		l2.blocksSeen--
		if br.Landed {
			l2.blocksLanded--
		} else {
			l2.blocksMissed--
			l.Info("Missed block was reorgd (hash)",
				zap.Uint64("block_number", br.Number.Uint64()),
				zap.Int64("blocks_landed", l2.blocksLanded),
				zap.Int64("blocks_missed", l2.blocksMissed),
				zap.Int64("blocks_seen", l2.blocksSeen),
			)
		}

		block, err := l2.rpc.BlockByNumber(ctx, br.Number)
		if err != nil {
			l.Warn("Failed to request block by number, skipping this round of unwind...",
				zap.Error(err),
				zap.String("number", br.Number.String()),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.Rpc),
			)
			return err
		}

		if block.Hash().Cmp(br.Hash) == 0 {
			return nil
		}

		l.Info("Unwinding...",
			zap.Uint64("block_number", l2.blockHeight),
		)
	}

	return nil
}

func (l2 *L2) isBuilderTx(
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

	return from.Cmp(l2.builderAddr) == 0
}

func (l2 *L2) isBuilderPolicyTx(
	tx *ethtypes.Transaction,
) bool {
	if tx == nil || tx.Rejected() {
		return false
	}

	if tx.To() == nil || tx.To().Cmp(l2.builderPolicyAddr) != 0 {
		return false
	}

	if len(tx.Data()) < len(l2.builderPolicySignature) {
		return false
	}

	return slices.Compare(tx.Data()[:4], l2.builderPolicySignature[:]) == 0
}

func (l2 *L2) isProbeTx(
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

	if from.Cmp(l2.monitorAddr) != 0 {
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

func (l2 *L2) observeWallets(ctx context.Context, o otelapi.Observer) error {
	l := logutils.LoggerFromContext(ctx)

	errs := make([]error, 0)

	for name, addr := range l2.wallets {
		_balance, err := l2.rpc.BalanceAt(ctx, addr, nil)
		if err != nil {
			l.Warn("Failed to request balance",
				zap.Error(err),
				zap.String("at", addr.String()),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.Rpc),
			)
			errs = append(errs, err)
			continue
		}

		balance, _ := _balance.Float64()

		o.ObserveFloat64(metrics.WalletBalance, balance, otelapi.WithAttributes(
			attribute.KeyValue{Key: "wallet_address", Value: attribute.StringValue(addr.String())},
			attribute.KeyValue{Key: "wallet_name", Value: attribute.StringValue(name)},
		))
	}

	return utils.FlattenErrors(errs)
}

func (l2 *L2) observerProbes(_ context.Context, o otelapi.Observer) error {
	if l2.cfg.ProbeTx.PrivateKey == "" {
		return nil
	}

	o.ObserveInt64(metrics.ProbesFailedCount, l2.monitorProbesFailedCount)
	o.ObserveInt64(metrics.ProbesLandedCount, l2.monitorProbesLandedCount)
	o.ObserveInt64(metrics.ProbesSentCount, l2.monitorProbesSentCount)

	return nil
}

func (l2 *L2) sendProbeTx(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	var (
		data     = make([]byte, 8)
		gasPrice *big.Int
		err      error
	)

	{ // get the gas price
		gasPrice, err = l2.rpc.SuggestGasPrice(ctx)
		if err != nil {
			l.Warn("Failed to request suggested gas price",
				zap.Error(err),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.Rpc),
			)
			l2.monitorProbesFailedCount++
			return
		}

		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(100+l2.cfg.ProbeTx.GasPriceAdjustment))
		gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

		gasPrice = utils.MinBigInt(gasPrice, big.NewInt(l2.cfg.ProbeTx.GasPriceCap))
	}

	if l2.monitorNonce == 0 { // get the nonce
		nonce, err := l2.rpc.NonceAt(ctx, l2.monitorAddr, nil)
		if err != nil {
			l.Warn("Failed to request a nonce",
				zap.Error(err),
				zap.String("address", l2.monitorAddr.String()),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.Rpc),
			)
			l2.monitorProbesFailedCount++
			return
		}
		l2.monitorNonce = nonce
	}

tryingNonces:
	for attempt := 1; attempt <= 8; attempt++ { // we don't want to get rate-limited
		thisBlock := time.Now().Add(-l2.cfg.BlockTime / 2).Round(l2.cfg.BlockTime)
		binary.BigEndian.PutUint64(data, uint64(thisBlock.Unix()))

		tx := ethtypes.NewTransaction(
			l2.monitorNonce,
			ethcommon.Address{},
			nil,
			l2.cfg.ProbeTx.GasLimit,
			gasPrice,
			data,
		)

		signedTx, err := ethtypes.SignTx(tx, l2.signer, l2.monitorKey)
		if err != nil {
			l.Error("Failed to sign a transaction",
				zap.Error(err),
				zap.String("address", l2.monitorAddr.String()),
			)
			l2.monitorProbesFailedCount++
			return
		}

		err = l2.rpc.SendTransaction(ctx, signedTx)
		if err == nil {
			l2.monitorProbesSentCount++
			l2.monitorNonce++
			return
		}

		if ctxErr := ctx.Err(); ctxErr != nil {
			l.Error("Failed to send a transaction",
				zap.Error(errors.Join(err, ctxErr)),
				zap.String("address", l2.monitorAddr.String()),
				zap.String("to", tx.To().String()),
				zap.Uint64("nonce", l2.monitorNonce),
				zap.String("kind", "l2"),
				zap.String("rpc", l2.cfg.Rpc),
			)
			l2.monitorProbesFailedCount++
			return // irrecoverable error (for now, at least) => no point in trying other nonces
		}

		l.Error("Failed to send monitor transaction",
			zap.Error(err),
			zap.String("address", l2.monitorAddr.String()),
			zap.String("to", tx.To().String()),
			zap.Uint64("nonce", l2.monitorNonce),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.Rpc),
		)

		switch {
		case strings.Contains(err.Error(), "insufficient funds"):
			l2.monitorProbesFailedCount++
			return // irrecoverable error

		case strings.Contains(err.Error(), "already known"):
			l2.monitorNonce++ // there's already a tx with this nonce => try next one
			continue tryingNonces

		case strings.Contains(err.Error(), "replacement transaction underpriced"):
			l2.monitorNonce++ // there's already a tx with this nonce => try next one
			continue tryingNonces

		case strings.Contains(err.Error(), "nonce too low"):
			nonce, err := l2.rpc.NonceAt(ctx, l2.monitorAddr, nil)
			if err != nil {
				l.Warn("Failed to request a nonce",
					zap.Error(err),
					zap.String("address", l2.monitorAddr.String()),
					zap.String("kind", "l2"),
					zap.String("rpc", l2.cfg.Rpc),
				)
				l2.monitorProbesFailedCount++
				return
			}
			l2.monitorNonce = nonce

			continue tryingNonces
		}

		l2.monitorProbesFailedCount++
		return
	}

	l2.monitorProbesFailedCount++
}

func (l2 *L2) checkAndResetProbeTxNonce(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	inFlight := l2.monitorProbesSentCount - l2.monitorProbesLandedCount
	if inFlight > l2.cfg.ProbeTx.ResetThreshold {
		l.Warn("In-flight probe transaction count is above threshold, resetting the nonce",
			zap.Int64("count", inFlight),
			zap.Int64("threshold", l2.monitorProbesLandedCount),
		)

		l2.monitorNonce = 0
		l2.monitorProbesSentCount = 0
		l2.monitorProbesLandedCount = 0
	}
}
