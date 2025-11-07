package rpc

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/flashbots/chain-monitor/utils"
)

func callEveryoneWithResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
) ([]R, error) {
	var (
		mx sync.Mutex
		wg sync.WaitGroup

		res  = make([]R, 0, len(rpc.fallback)+1)
		errs = make([]error, 0, len(rpc.fallback)+1)
	)

	wg.Add(1)
	go func() {
		defer wg.Done()

		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		defer cancel()

		_res, err := call(_ctx, rpc.main)

		mx.Lock()
		defer mx.Unlock()

		if err == nil {
			res = append(res, _res)
		} else {
			errs = append(errs, fmt.Errorf("%s: %w",
				rpc.url.main, err,
			))
		}
	}()

	for idx, fallback := range rpc.fallback {
		wg.Add(1)
		go func() {
			defer wg.Done()

			_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
			defer cancel()

			_res, err := call(_ctx, fallback)

			mx.Lock()
			defer mx.Unlock()

			if err == nil {
				res = append(res, _res)
			} else {
				errs = append(errs, fmt.Errorf("%s: %w",
					rpc.url.fallback[idx], err,
				))
			}
		}()
	}

	wg.Wait()

	return res, utils.FlattenErrors(errs)
}

func callMainThenFallback(
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) error,
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

func callMainThenFallbackWithResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
) (R, error) {
	_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
	defer cancel()

	res, err := call(_ctx, rpc.main)
	if err == nil {
		return res, nil
	}

	errs := make([]error, 0, len(rpc.fallback)+1)
	errs = append(errs, fmt.Errorf("%s: %w",
		rpc.url.main, err,
	))

	for idx, fallback := range rpc.fallback {
		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		defer cancel()

		res, err := call(_ctx, fallback)
		if err == nil {
			return res, nil
		}

		errs = append(errs, fmt.Errorf("%s: %w",
			rpc.url.fallback[idx], err,
		))
	}

	var _nil R
	return _nil, utils.FlattenErrors(errs)
}

func callFallbackThenMainWithResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
) (R, error) {
	errs := make([]error, 0, len(rpc.fallback))

	for idx, fallback := range rpc.fallback {
		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		defer cancel()

		res, err := call(_ctx, fallback)
		if err == nil {
			return res, nil
		}

		errs = append(errs, fmt.Errorf("%s: %w",
			rpc.url.fallback[idx], err,
		))
	}

	_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
	defer cancel()

	res, err := call(_ctx, rpc.main)
	if err == nil {
		return res, nil
	}
	errs = append(errs, fmt.Errorf("%s: %w",
		rpc.url.main, err,
	))

	var _nil R
	return _nil, utils.FlattenErrors(errs)
}
