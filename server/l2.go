package server

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/logutils"
	"github.com/flashbots/chain-monitor/metrics"
	"github.com/flashbots/chain-monitor/types"
	"github.com/flashbots/chain-monitor/utils"

	"go.uber.org/zap"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"go.opentelemetry.io/otel/attribute"
	otelapi "go.opentelemetry.io/otel/metric"
)

type L2 struct {
	cfg *config.L2

	rpc    *ethclient.Client
	ticker *time.Ticker

	builderAddr ethcommon.Address
	chainID     *big.Int
	monitorAddr ethcommon.Address
	monitorKey  *ecdsa.PrivateKey
	reorgWindow int
	signer      ethtypes.EIP155Signer
	wallets     map[string]ethcommon.Address

	blockHeight  uint64
	blocks       *types.RingBuffer[blockRecord]
	blocksLanded int64
	blocksMissed int64
	blocksSeen   int64

	processBlockFailuresCount uint

	unwindingByHash    bool
	unwindByHashHeight uint64
}

type blockRecord struct {
	number *big.Int
	hash   ethcommon.Hash
	landed bool
}

func newL2(cfg *config.L2) (*L2, error) {
	l := zap.L()

	l2 := &L2{
		cfg:         cfg,
		reorgWindow: int(cfg.ReorgWindow/cfg.BlockTime) + 1,
		wallets:     make(map[string]ethcommon.Address, len(cfg.WalletAddresses)),
	}

	{ // ticker
		now := time.Now()
		time.Sleep(now.Truncate(cfg.BlockTime).Add(cfg.BlockTime).Sub(now)) // align with block times
		l2.ticker = time.NewTicker(cfg.BlockTime)
	}

	if cfg.BuilderAddress != "" { // builderAddr
		addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
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

	if cfg.MonitorPrivateKey != "" { // monitorAddr, monitorKey
		monitorKey, err := crypto.HexToECDSA(cfg.MonitorPrivateKey)
		if err != nil {
			return nil, err
		}
		l2.monitorKey = monitorKey
		l2.monitorAddr = crypto.PubkeyToAddress(monitorKey.PublicKey)
	}

	for name, addrStr := range cfg.WalletAddresses { // wallets
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
		rpc, err := ethclient.Dial(cfg.RPC)
		if err != nil {
			return nil, err
		}
		l2.rpc = rpc
	}

	{ // chainID, signer
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		l.Debug("Requesting network id",
			zap.String("kind", "l2"),
			zap.String("rpc", cfg.RPC),
		)

		chainID, err := l2.rpc.NetworkID(ctx)
		if err != nil {
			return nil, err
		}
		l2.chainID = chainID
		l2.signer = ethtypes.NewEIP155Signer(chainID)
	}

	{ // blocks, blockHeight
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		l.Debug("Requesting block height",
			zap.String("kind", "l2"),
			zap.String("rpc", cfg.RPC),
		)

		blockHeight, err := l2.rpc.BlockNumber(ctx)
		if err != nil {
			return nil, err
		}
		if blockHeight > 0 {
			l2.blockHeight = blockHeight - 1
		}
		l2.blocks = types.NewRingBuffer[blockRecord](int(blockHeight), int(cfg.ReorgWindow/cfg.BlockTime+1))
	}

	return l2, nil
}

func (l2 *L2) run(ctx context.Context) {
	tick := func() {}

	if l2.builderAddr.Cmp(ethcommon.Address{}) != 0 {
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
	}

	go func() {
		for {
			<-l2.ticker.C
			tick()
		}
	}()
}

func (l2 *L2) stop() {
	l2.ticker.Stop()
}

func (l2 *L2) processNewBlocks(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	l.Debug("Requesting block number",
		zap.String("kind", "l2"),
		zap.String("rpc", l2.cfg.RPC),
	)

	blockHeight, err := l2.rpc.BlockNumber(ctx)
	if err != nil {
		l.Error("Failed to get block height, skipping this round...",
			zap.Error(err),
		)
		return
	}

	if blockHeight == l2.blockHeight {
		l.Debug("Still at the same height, skipping...",
			zap.Uint64("block_height", blockHeight),
		)
		return
	}

	if blockHeight < l2.blockHeight {
		l2.processReorgByHeight(ctx, blockHeight)
	}

	for b := l2.blockHeight + 1; b <= blockHeight; b++ {
		l.Debug("Processing new l2 block",
			zap.Uint64("block_height", b),
		)

		if err := l2.processBlock(ctx, b); err != nil {
			l2.processBlockFailuresCount++

			logLevel := zap.DebugLevel
			if l2.processBlockFailuresCount > 10 {
				logLevel = zap.WarnLevel
			}

			l.Log(logLevel, "Failed to process block, skipping this round...",
				zap.Error(err),
				zap.Uint64("block", blockHeight),
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
	l := logutils.LoggerFromContext(ctx)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	l.Debug("Requesting block by number",
		zap.Uint64("number", blockNumber),
		zap.String("kind", "l2"),
		zap.String("rpc", l2.cfg.RPC),
	)

	block, err := l2.rpc.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	metrics.TxPerBlock.Record(ctx, int64(len(block.Transactions())))
	metrics.GasPerBlock.Record(ctx, int64(block.GasUsed()))

	if blockNumber > 0 {
		if previous, ok := l2.blocks.At(int(blockNumber) - 1); ok {
			if previous.hash.Cmp(block.ParentHash()) != 0 {
				if !l2.unwindingByHash {
					l.Info("Chain reorg detected via hash mismatch, starting the unwind...",
						zap.Uint64("block_number", blockNumber),
						zap.String("new_parent_hash", block.ParentHash().String()),
						zap.String("old_parent_hash", previous.hash.String()),
					)
					l2.unwindingByHash = true
					l2.unwindByHashHeight = l2.blockHeight
					return l2.processReorgByHash(ctx)
				}

				l.Debug("Continuing the reorg unwind...",
					zap.Uint64("block_number", blockNumber),
				)
				return l2.processReorgByHash(ctx)
			}

			if l2.unwindingByHash {
				depth := l2.unwindByHashHeight - blockNumber

				metrics.ReorgsCount.Add(ctx, 1)
				metrics.ReorgDepth.Record(ctx, int64(depth))

				l.Warn("Chain reorg detected via hash mismatch",
					zap.Uint64("reorg_depth", depth),
					zap.Uint64("old_block_number", l2.unwindByHashHeight),
					zap.Uint64("new_block_number", blockNumber),
				)

				l2.unwindingByHash = false
				l2.unwindByHashHeight = 0
			}
		}
	}

	l2.blocksSeen++
	metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)

	hasBuilderTx := false
	expectedBuilderTxData := []byte(fmt.Sprintf("Block Number: %s", block.Number().String()))

	for _, tx := range block.Transactions() {
		if l2.isBuilderTx(ctx, block, tx, expectedBuilderTxData) {
			if !hasBuilderTx {
				hasBuilderTx = true
			} else {
				l.Warn("More than 1 builder transaction found in a block",
					zap.Uint64("block_number", blockNumber),
				)
			}
		}

		if l2.monitorKey != nil {
			if isProbeTx, sent, latency := l2.isProbeTx(ctx, block, tx); isProbeTx {
				l.Debug("Detected probe transaction",
					zap.Uint64("latency", latency),
					zap.Uint64("sent", sent),
					zap.Uint64("landed", block.Time()),
				)
				metrics.ProbesLatency.Record(ctx, int64(latency))
			}
		}
	}

	if hasBuilderTx {
		l2.blocks.Push(blockRecord{
			number: block.Number(),
			hash:   block.Hash(),
			landed: true,
		})
		l2.blocksLanded++
		metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
	} else {
		l2.blocks.Push(blockRecord{
			number: block.Number(),
			hash:   block.Hash(),
			landed: false,
		})
		l2.blocksMissed++
		metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)
		metrics.BlockMissed.Record(ctx, int64(blockNumber))
		l.Warn("Builder had missed a block",
			zap.Uint64("block_number", blockNumber),
			zap.Int64("blocks_landed", l2.blocksLanded),
			zap.Int64("blocks_missed", l2.blocksMissed),
			zap.Int64("blocks_seen", l2.blocksSeen),
		)
	}

	if l2.blocks.Length() > l2.reorgWindow {
		_, _ = l2.blocks.Pop()
	}

	return nil
}

func (l2 *L2) processReorgByHash(ctx context.Context) error {
	l := logutils.LoggerFromContext(ctx)

	defer func() {
		metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)
		metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
		metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)
	}()

	for l2.blocks.Length() > 0 {
		br, _ := l2.blocks.Pick()
		l2.blocksSeen--
		if br.landed {
			l2.blocksLanded--
		} else {
			l2.blocksMissed--
			l.Info("Missed block was reorgd (hash)",
				zap.Uint64("block_number", br.number.Uint64()),
				zap.Int64("blocks_landed", l2.blocksLanded),
				zap.Int64("blocks_missed", l2.blocksMissed),
				zap.Int64("blocks_seen", l2.blocksSeen),
			)
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		l.Debug("Requesting block by number",
			zap.String("number", br.number.String()),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.RPC),
		)

		block, err := l2.rpc.BlockByNumber(ctx, br.number)
		if err != nil {
			l.Error("Failed to unwind back to common root, skipping this round...",
				zap.Error(err),
				zap.Uint64("block", br.number.Uint64()),
			)
			return err
		}
		l2.blockHeight = br.number.Uint64() - 1

		if block.Hash().Cmp(br.hash) == 0 {
			return nil
		}
	}

	return nil
}

func (l2 *L2) processReorgByHeight(ctx context.Context, newBlockHeight uint64) {
	l := logutils.LoggerFromContext(ctx)

	if newBlockHeight == 0 {
		newBlockHeight = 1
	}

	depth := l2.blockHeight - newBlockHeight + 1

	adjustSeen := 0
	adjustLanded := 0
	adjustMissed := 0

	for b := l2.blockHeight; b >= newBlockHeight && l2.blocks.Length() > 0; b-- {
		if br, ok := l2.blocks.Pick(); ok {
			adjustSeen++
			if br.landed {
				adjustLanded++
			} else {
				adjustMissed++
				l.Info("Missed block was reorgd (height)",
					zap.Uint64("block_number", b),
					zap.Int64("blocks_landed", l2.blocksLanded-int64(adjustLanded)),
					zap.Int64("blocks_missed", l2.blocksMissed-int64(adjustMissed)),
					zap.Int64("blocks_seen", l2.blocksSeen-int64(adjustSeen)),
				)
			}
		}
	}

	l2.blocksSeen -= int64(adjustSeen)
	l2.blocksLanded -= int64(adjustLanded)
	l2.blocksMissed -= int64(adjustMissed)

	metrics.ReorgsCount.Add(ctx, 1)
	metrics.ReorgDepth.Record(ctx, int64(depth))

	metrics.BlocksSeenCount.Record(ctx, l2.blocksSeen)
	metrics.BlocksLandedCount.Record(ctx, l2.blocksLanded)
	metrics.BlocksMissedCount.Record(ctx, l2.blocksMissed)

	l.Warn("Chain reorg detected via block height",
		zap.Uint64("reorg_depth", depth),
		zap.Uint64("old_block_number", l2.blockHeight),
		zap.Uint64("new_block_number", newBlockHeight-1),
	)

	l2.blocks.Forget(adjustSeen)
	l2.blockHeight = newBlockHeight - 1
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

func (l2 *L2) isProbeTx(
	ctx context.Context,
	block *ethtypes.Block,
	tx *ethtypes.Transaction,
) (isProbeTx bool, sent, latency uint64) {
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

	blockTime := block.Time()
	sent = binary.BigEndian.Uint64(tx.Data())
	if blockTime < sent {
		l := logutils.LoggerFromContext(ctx)
		l.Warn("Monitoring transaction from future",
			zap.String("tx", tx.Hash().Hex()),
			zap.String("block", block.Number().String()),
		)
	}
	latency = blockTime - sent

	return true, sent, latency
}

func (l2 *L2) observeWallets(ctx context.Context, o otelapi.Observer) error {
	l := logutils.LoggerFromContext(ctx)

	errs := make([]error, 0)

	for name, addr := range l2.wallets {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		l.Debug("Requesting balance",
			zap.String("at", addr.String()),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.RPC),
		)

		_balance, err := l2.rpc.BalanceAt(ctx, addr, nil)
		if err != nil {
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

func (l2 *L2) sendProbeTx(ctx context.Context) {
	l := logutils.LoggerFromContext(ctx)

	var (
		data     = make([]byte, 8)
		gasPrice *big.Int
		nonce    uint64
		err      error
	)

	{ // get the gas price
		_ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		l.Debug("Requesting suggested gas price",
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.RPC),
		)

		gasPrice, err = l2.rpc.SuggestGasPrice(_ctx)
		if err != nil {
			l.Error("Failed to get suggested gas price for probe tx",
				zap.Error(err),
				zap.String("monitor_address", l2.monitorAddr.String()),
			)
			metrics.ProbesFailedCount.Add(ctx, 1)
			return
		}

		gasPrice = new(big.Int).Mul(gasPrice, big.NewInt(100+l2.cfg.MonitorTxGasPriceAdjustment))
		gasPrice = new(big.Int).Div(gasPrice, big.NewInt(100))

		gasPrice = utils.MinBigInt(gasPrice, big.NewInt(l2.cfg.MonitorTxGasPriceCap))
	}

	{ // get the nonce
		_ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		l.Debug("Requesting nonce",
			zap.String("at", l2.monitorAddr.String()),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.RPC),
		)

		nonce, err = l2.rpc.NonceAt(_ctx, l2.monitorAddr, nil)
		if err != nil {
			l.Error("Failed to get pending nonce for probe tx",
				zap.Error(err),
				zap.String("monitor_address", l2.monitorAddr.String()),
			)
			metrics.ProbesFailedCount.Add(ctx, 1)
			return
		}
	}

	errs := make([]error, 0, 8)

tryingNonces:
	for nonceIncrement := uint64(0); nonceIncrement < 8; nonceIncrement++ {
		thisBlock := time.Now().Add(-l2.cfg.BlockTime / 2).Round(l2.cfg.BlockTime)
		binary.BigEndian.PutUint64(data, uint64(thisBlock.Unix()))

		tx := ethtypes.NewTransaction(
			nonce+nonceIncrement,
			ethcommon.Address{},
			nil,
			l2.cfg.MonitorTxGasLimit,
			gasPrice,
			data,
		)

		signedTx, err := ethtypes.SignTx(tx, l2.signer, l2.monitorKey)
		if err != nil {
			l.Error("Failed to sign the probe tx",
				zap.Error(err),
				zap.String("monitor_address", l2.monitorAddr.String()),
			)
			metrics.ProbesFailedCount.Add(ctx, 1)
			return
		}

		_ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		l.Debug("Sending transaction",
			zap.String("from", l2.monitorAddr.String()),
			zap.String("to", tx.To().String()),
			zap.Uint64("nonce", nonce+nonceIncrement),
			zap.String("kind", "l2"),
			zap.String("rpc", l2.cfg.RPC),
		)

		if err := l2.rpc.SendTransaction(_ctx, signedTx); err != nil {
			errs = append(errs,
				fmt.Errorf("nonce %d (%d+%d): %w", nonce+nonceIncrement, nonce, nonceIncrement, err),
			)

			for _, msg := range []string{
				"replacement transaction underpriced",
				"nonce too low",
				"already known",
			} {
				if strings.HasPrefix(err.Error(), msg) {
					continue tryingNonces // perhaps the next nonce will be a success
				}
			}

			l.Error("Failed to send the probe tx",
				zap.Error(utils.FlattenErrors(errs)),
				zap.String("from", l2.monitorAddr.String()),
				zap.String("to", l2.builderAddr.String()),
				zap.String("data", hex.EncodeToString(data)),
				zap.String("gas_price", gasPrice.String()),
				zap.Uint64("gas_limit", l2.cfg.MonitorTxGasLimit),
			)
			metrics.ProbesFailedCount.Add(ctx, 1)

			return // irrecoverable error (for now, at least) => no point in trying other nonces
		}

		metrics.ProbesSentCount.Add(ctx, 1)

		l.Debug("Sent probe tx",
			zap.String("hash", tx.Hash().String()),
			zap.String("from", l2.monitorAddr.String()),
			zap.String("to", l2.builderAddr.String()),
			zap.String("data", hex.EncodeToString(data)),
			zap.String("gas_price", gasPrice.String()),
			zap.Uint64("gas_limit", l2.cfg.MonitorTxGasLimit),
		)

		return // yay, sent the probe tx
	}

	l.Error("Failed to send the probe tx",
		zap.Error(utils.FlattenErrors(errs)),
		zap.String("from", l2.monitorAddr.String()),
		zap.String("to", l2.builderAddr.String()),
		zap.String("data", hex.EncodeToString(data)),
		zap.String("gas_price", gasPrice.String()),
		zap.Uint64("gas_limit", l2.cfg.MonitorTxGasLimit),
	)
	metrics.ProbesFailedCount.Add(ctx, 1)
}
