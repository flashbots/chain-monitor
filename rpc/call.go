package rpc

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

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

// callMainThenFallbackWithBackoff calls the main RPC then fallback RPCs with exponential backoff
func callMainThenFallbackWithBackoff(
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) error,
	maxRetries int,
	initialDelay time.Duration,
) error {
	_, err := callMainThenFallbackWithBackoffAndResult(ctx, rpc, func(ctx context.Context, cli *ethclient.Client) (struct{}, error) {
		return struct{}{}, call(ctx, cli)
	}, maxRetries, initialDelay)
	return err
}

// callMainThenFallbackWithBackoffAndResult calls the main RPC then fallback RPCs with exponential backoff
func callMainThenFallbackWithBackoffAndResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
	maxRetries int,
	initialDelay time.Duration,
) (R, error) {
	errs := make([]error, 0, len(rpc.fallback)+1)

	// Try main with exponential backoff
	if res, err := retryWithBackoff(ctx, rpc, rpc.main, call, maxRetries, initialDelay); err == nil {
		return res, nil
	} else {
		errs = append(errs, fmt.Errorf("%s: %w", rpc.url.main, err))
	}

	// Try fallback RPCs with exponential backoff
	for idx, fallback := range rpc.fallback {
		if res, err := retryWithBackoff(ctx, rpc, fallback, call, maxRetries, initialDelay); err == nil {
			return res, nil
		} else {
			errs = append(errs, fmt.Errorf("%s: %w", rpc.url.fallback[idx], err))
		}
	}

	var _nil R
	return _nil, utils.FlattenErrors(errs)
}

// callFallbackThenMainWithBackoffAndResult calls fallback RPCs then main RPC with exponential backoff
func callFallbackThenMainWithBackoffAndResult[R any](
	ctx context.Context,
	rpc *RPC,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
	maxRetries int,
	initialDelay time.Duration,
) (R, error) {
	errs := make([]error, 0, len(rpc.fallback)+1)

	// Try fallback RPCs with exponential backoff
	for idx, fallback := range rpc.fallback {
		if res, err := retryWithBackoff(ctx, rpc, fallback, call, maxRetries, initialDelay); err == nil {
			return res, nil
		} else {
			errs = append(errs, fmt.Errorf("%s: %w", rpc.url.fallback[idx], err))
		}
	}

	// Try main with exponential backoff
	if res, err := retryWithBackoff(ctx, rpc, rpc.main, call, maxRetries, initialDelay); err == nil {
		return res, nil
	} else {
		errs = append(errs, fmt.Errorf("%s: %w", rpc.url.main, err))
	}

	var _nil R
	return _nil, utils.FlattenErrors(errs)
}

// retryWithBackoff retries a call with exponential backoff
func retryWithBackoff[R any](
	ctx context.Context,
	rpc *RPC,
	client *ethclient.Client,
	call func(ctx context.Context, rpc *ethclient.Client) (R, error),
	maxRetries int,
	initialDelay time.Duration,
) (R, error) {
	var lastErr error
	var result R

	for attempt := 0; attempt <= maxRetries; attempt++ {
		_ctx, cancel := context.WithTimeout(ctx, rpc.timeout)
		res, err := call(_ctx, client)
		cancel()

		if err == nil {
			return res, nil
		}

		lastErr = err

		// Don't sleep after the last attempt
		if attempt < maxRetries {
			// Calculate exponential backoff delay: initialDelay * 2^attempt
			delay := time.Duration(float64(initialDelay) * math.Pow(2, float64(attempt)))

			// Wait for the backoff delay or context cancellation
			select {
			case <-ctx.Done():
				var _nil R
				return _nil, ctx.Err()
			case <-time.After(delay):
			}
		}
	}

	return result, lastErr
}
