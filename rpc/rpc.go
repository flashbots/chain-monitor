package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"time"

	"github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/flashbots/chain-monitor/utils"
)

type RPC struct {
	Main      *ethclient.Client
	networkID *big.Int
	Fallback  []*ethclient.Client
	url       rpcUrl

	timeout time.Duration
}

var (
	errFailedToDial     = errors.New("failed to dial rpc")
	errInvalidNetworkID = errors.New("invalid network id")
)

func New(networkID uint64, url string, fallback ...string) (*RPC, error) {
	cli, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w",
			errFailedToDial, url, err,
		)
	}

	rpc := &RPC{
		Main:     cli,
		Fallback: make([]*ethclient.Client, 0, len(fallback)),
		timeout:  time.Second,

		url: rpcUrl{
			main:     url,
			fallback: fallback,
		},
	}

	if networkID != 0 {
		rpc.networkID = new(big.Int).SetUint64(networkID)
	}

	for _, url := range fallback {
		cli, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("%w: %s: %w",
				errFailedToDial, url, err,
			)
		}
		rpc.Fallback = append(rpc.Fallback, cli)
	}

	return rpc, nil
}

func (rpc *RPC) Close() {
	rpc.Main.Close()
	for _, rpc := range rpc.Fallback {
		rpc.Close()
	}
}

func (rpc *RPC) NetworkID(ctx context.Context) (*big.Int, error) {
	networkIDs, err := callEveryoneWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (*big.Int, error) {
		return cli.NetworkID(ctx)
	})
	if len(networkIDs) == 0 {
		return nil, err
	}

	if rpc.networkID != nil {
		for _, networkID := range networkIDs {
			if rpc.networkID.Cmp(networkID) != 0 {
				return networkID, errors.Join(fmt.Errorf("invalid network id (want: %d, got: %d)",
					rpc.networkID.Uint64(), networkID.Uint64(),
				), err)
			}
		}
	} else {
		networkID := networkIDs[0]
		for _, _networkID := range networkIDs {
			if networkID.Cmp(_networkID) != 0 {
				return networkID, errors.Join(fmt.Errorf("mismatching network ids: %d vs. %d",
					networkID.Uint64(), _networkID.Uint64(),
				), err)
			}
		}
	}

	return networkIDs[0], nil
}

func (rpc *RPC) callCheckingNetworkID(
	ctx context.Context, cli *ethclient.Client, batchElem ethrpc.BatchElem,
) error {
	batch := make([]ethrpc.BatchElem, 0, 2)

	netVersion := ""
	if rpc.networkID != nil {
		batch = append(batch, ethrpc.BatchElem{Method: "net_version", Result: &netVersion})
	}

	batch = append(batch, batchElem)

	if err := cli.Client().BatchCallContext(ctx, batch); err != nil {
		return err
	}
	errs := make([]error, 0)
	for _, e := range batch {
		if e.Error != nil {
			errs = append(errs, e.Error)
		}
	}
	if len(errs) > 0 {
		return utils.FlattenErrors(errs)
	}

	if rpc.networkID != nil {
		networkID, ok := new(big.Int).SetString(netVersion, 0)
		if !ok {
			return fmt.Errorf("invalid net_version result %q", netVersion)
		}

		if rpc.networkID.Cmp(networkID) != 0 {
			return fmt.Errorf("%w: want %d, got %d",
				errInvalidNetworkID, rpc.networkID.Uint64(), networkID.Uint64(),
			)
		}
	}

	return nil
}

func (rpc *RPC) BlockNumber(ctx context.Context) (uint64, error) {
	blocks, err := callEveryoneWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (uint64, error) {
		var blockNumber hexutil.Uint64
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_blockNumber",
			Result: &blockNumber,
		})
		if err != nil {
			return 0, err
		}

		return uint64(blockNumber), nil
	})

	if len(blocks) == 0 {
		return 0, err
	}

	return slices.Max(blocks), nil
}

func (rpc *RPC) BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error) {

	return callMainThenFallbackWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (*ethtypes.Block, error) {
		var raw json.RawMessage
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []interface{}{hexutil.EncodeBig(number), true},
			Result: &raw,
		})
		if err != nil {
			return nil, err
		}

		{ // NOTE: this is a partial rip-off from ethereum's cli.BlockByNumber
			// Decode header and transactions.
			var head *ethtypes.Header
			if err := json.Unmarshal(raw, &head); err != nil {
				return nil, err
			}
			// When the block is not found, the API returns JSON null.
			if head == nil {
				return nil, ethereum.NotFound
			}

			var body rpcBlock
			if err := json.Unmarshal(raw, &body); err != nil {
				return nil, err
			}
			// Pending blocks don't return a block hash, compute it for sender caching.
			if body.Hash == nil {
				tmp := head.Hash()
				body.Hash = &tmp
			}

			if head.TxHash == ethtypes.EmptyTxsHash && len(body.Transactions) > 0 {
				return nil, errors.New("server returned non-empty transaction list but block header indicates no transactions")
			}
			if head.TxHash != ethtypes.EmptyTxsHash && len(body.Transactions) == 0 {
				return nil, errors.New("server returned empty transaction list but block header indicates transactions")
			}
			// Fill the sender cache of transactions in the block.
			txs := make([]*ethtypes.Transaction, len(body.Transactions))
			for i, tx := range body.Transactions {
				if tx.From != nil {
					_, _ = ethtypes.Sender(&senderFromServer{*tx.From, *body.Hash}, tx.tx)
				}
				txs[i] = tx.tx
			}

			return ethtypes.NewBlockWithHeader(head).
				WithBody(ethtypes.Body{
					Transactions: txs,
					Withdrawals:  body.Withdrawals,
				}), nil
		}
	})
}

func (rpc *RPC) BalanceAt(ctx context.Context, account ethcommon.Address) (*big.Int, error) {
	return callMainThenFallbackWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (*big.Int, error) {
		var balance hexutil.Big
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_getBalance",
			Args:   []interface{}{account, "latest"},
			Result: &balance,
		})
		if err != nil {
			return nil, err
		}
		return (*big.Int)(&balance), nil
	})
}

func (rpc *RPC) NonceAt(ctx context.Context, account ethcommon.Address) (uint64, error) {
	return callMainThenFallbackWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (uint64, error) {
		var nonce hexutil.Uint64
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_getTransactionCount",
			Args:   []interface{}{account, "latest"},
			Result: &nonce,
		})
		if err != nil {
			return 0, err
		}
		return uint64(nonce), nil
	})
}

func (rpc *RPC) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return callMainThenFallbackWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (*big.Int, error) {
		var gasPrice hexutil.Big
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_gasPrice",
			Result: &gasPrice,
		})
		if err != nil {
			return nil, err
		}
		return (*big.Int)(&gasPrice), nil
	})
}

func (rpc *RPC) SendTransaction(ctx context.Context, tx *ethtypes.Transaction) error {
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	rawTx := hexutil.Encode(data)
	return callMainThenFallback(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) error {
		var txhash hexutil.Bytes
		return rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_sendRawTransaction",
			Args:   []interface{}{rawTx},
			Result: &txhash,
		})
	})
}

func (rpc *RPC) TransactionReceipt(ctx context.Context, txHash ethcommon.Hash) (*ethtypes.Receipt, error) {
	return callFallbackThenMainWithResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (*ethtypes.Receipt, error) {
		var receipt *ethtypes.Receipt
		err := rpc.callCheckingNetworkID(ctx, cli, ethrpc.BatchElem{
			Method: "eth_sendRawTransaction",
			Args:   []interface{}{txHash},
			Result: &receipt,
		})
		if err != nil {
			return nil, err
		}
		return receipt, nil
	})
}
