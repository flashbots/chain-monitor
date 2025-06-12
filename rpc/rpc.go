package rpc

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RPC struct {
	main     *ethclient.Client
	fallback []*ethclient.Client
	url      rpcUrl

	timeout time.Duration
}

var (
	errL2FailedToDial = errors.New("failed to dial rpc")
)

func New(url string, fallback ...string) (*RPC, error) {
	rpc, err := ethclient.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w",
			errL2FailedToDial, url, err,
		)
	}

	l2 := &RPC{
		main:     rpc,
		fallback: make([]*ethclient.Client, 0, len(fallback)),
		timeout:  time.Second,

		url: rpcUrl{
			main:     url,
			fallback: fallback,
		},
	}

	for _, url := range fallback {
		rpc, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("%w: %s: %w",
				errL2FailedToDial, url, err,
			)
		}
		l2.fallback = append(l2.fallback, rpc)
	}

	return l2, nil
}

func (rpc *RPC) Close() {
	rpc.main.Close()
	for _, rpc := range rpc.fallback {
		rpc.Close()
	}
}

func (rpc *RPC) NetworkID(ctx context.Context) (*big.Int, error) {
	return callWithFallbackAndResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (*big.Int, error) {
		return rpc.NetworkID(ctx)
	})
}

func (rpc *RPC) BlockNumber(ctx context.Context) (uint64, error) {
	blocks, err := callEveryoneWithResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (uint64, error) {
		return rpc.BlockNumber(ctx)
	})

	if len(blocks) == 0 {
		return 0, err
	}

	return slices.Max(blocks), nil
}

func (rpc *RPC) BlockByNumber(ctx context.Context, number *big.Int) (*ethtypes.Block, error) {
	return callWithFallbackAndResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (*ethtypes.Block, error) {
		return rpc.BlockByNumber(ctx, number)
	})
}

func (rpc *RPC) BalanceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (*big.Int, error) {
	return callWithFallbackAndResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (*big.Int, error) {
		return rpc.BalanceAt(ctx, account, blockNumber)
	})
}

func (rpc *RPC) NonceAt(ctx context.Context, account ethcommon.Address, blockNumber *big.Int) (uint64, error) {
	return callWithFallbackAndResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (uint64, error) {
		return rpc.NonceAt(ctx, account, blockNumber)
	})
}

func (rpc *RPC) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return callWithFallbackAndResult(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) (*big.Int, error) {
		return rpc.SuggestGasPrice(ctx)
	})
}

func (rpc *RPC) SendTransaction(ctx context.Context, tx *ethtypes.Transaction) error {
	return callWithFallback(ctx, rpc, func(ctx context.Context, rpc *ethclient.Client) error {
		return rpc.SendTransaction(ctx, tx)
	})
}
