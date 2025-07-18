package config

import (
	"errors"
	"fmt"
	"net/url"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/chain-monitor/utils"
)

type L1 struct {
	Rpc                    string            `yaml:"rpc"`
	RpcFallback            []string          `yaml:"rpc_fallback"`
	MonitorWalletAddresses map[string]string `yaml:"monitor_wallet_addresses"`
}

var (
	errL1InvalidRpc           = errors.New("invalid l1 rpc url")
	errL1InvalidRpcFallback   = errors.New("invalid l1 fallback rpc url")
	errL1InvalidWalletAddress = errors.New("invalid l1 wallet address")
)

func (cfg *L1) Validate() error {
	errs := make([]error, 0)

	if _, err := url.Parse(cfg.Rpc); err != nil {
		errs = append(errs, fmt.Errorf("%w: %s: %w",
			errL1InvalidRpc,
			cfg.Rpc,
			err,
		))
	}

	for _, rpc := range cfg.RpcFallback {
		if _, err := url.Parse(rpc); err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL1InvalidRpcFallback,
				rpc,
				err,
			))
		}
	}

	for _, wa := range cfg.MonitorWalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL1InvalidWalletAddress,
				wa,
				err,
			))
		}
		if len(_addr) != 20 {
			errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL1InvalidWalletAddress,
				wa,
				len(wa),
			))
		}
	}

	return utils.FlattenErrors(errs)
}
