package rpc

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flashbots/chain-monitor/utils"
)

func callWithFallback(
	ctx context.Context,
	rpc *RPC,
	call func(context.Context, *ethclient.Client) error,
) error {
	_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
	defer cancel()

	err := call(_ctx, rpc.main)
	if err == nil {
		return nil
	}

	errs := make([]error, 0, len(rpc.fallback)+1)
	errs = append(errs, fmt.Errorf("%s: %w",
		rpc.url.main, err,
	))

	for idx, fallback := range rpc.fallback {
		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		defer cancel()

		err := call(_ctx, fallback)
		if err == nil {
			return nil
		}

		errs = append(errs, fmt.Errorf("%s: %w",
			rpc.url.fallback[idx], err,
		))
	}

	return utils.FlattenErrors(errs)
}

func callWithFallbackAndResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(context.Context, *ethclient.Client) (R, error),
) (R, error) {
	_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
	defer cancel()

	res, err := call(_ctx, rpc.main)
	if err == nil {
		return res, nil
	}

	errs := make([]error, 0, len(rpc.fallback)+1)
	errs = append(errs, err)

	for _, fallback := range rpc.fallback {
		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		defer cancel()

		res, err := call(_ctx, fallback)
		if err == nil {
			return res, nil
		}

		errs = append(errs, err)
	}

	var _nil R
	return _nil, utils.FlattenErrors(errs)
}
