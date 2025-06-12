package config

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/chain-monitor/utils"
)

type L2 struct {
	Dir *Dir `yaml:"-"`

	BlockTime   time.Duration `yaml:"block_time"`
	ReorgWindow time.Duration `yaml:"reorg_window"`
	Rpc         string        `yaml:"rpc"`
	RpcFallback []string      `yaml:"rpc_fallback"`

	BuilderAddress  string            `yaml:"builder_address"`
	WalletAddresses map[string]string `yaml:"wallet_addresses"`

	Monitor *Monitor `yaml:"monitor"`
}

const (
	maxReorgWindow = 24 * time.Hour
)

var (
	errL2InvalidBuilderAddress = errors.New("invalid l2 builder address")
	errL2InvalidRpc            = errors.New("invalid l2 rpc url")
	errL2InvalidRpcFallback    = errors.New("invalid l2 fallback rpc url")
	errL2InvalidWalletAddress  = errors.New("invalid l2 wallet address")
	errL2ReorgWindowTooLarge   = errors.New("l2 reorg window is too large")
)

func (cfg *L2) Validate() error {
	errs := make([]error, 0)

	if _, err := url.Parse(cfg.Rpc); err != nil {
		errs = append(errs, fmt.Errorf("%w: %s: %w",
			errL2InvalidRpc,
			cfg.Rpc,
			err,
		))
	}

	for _, rpc := range cfg.RpcFallback {
		if _, err := url.Parse(rpc); err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL2InvalidRpcFallback,
				rpc,
				err,
			))
		}
	}

	if cfg.BuilderAddress != "" {
		_addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
		if err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL2InvalidBuilderAddress,
				cfg.BuilderAddress,
				err,
			))
		}
		if len(_addr) != 20 {
			errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL2InvalidBuilderAddress,
				cfg.BuilderAddress,
				len(_addr),
			))
		}
	}

	if cfg.ReorgWindow > maxReorgWindow {
		errs = append(errs, fmt.Errorf("%w (max %d): %d",
			errL2ReorgWindowTooLarge,
			maxReorgWindow,
			cfg.ReorgWindow,
		))
	}

	for _, wa := range cfg.WalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL2InvalidWalletAddress,
				wa,
				err,
			))
		}
		if len(_addr) != 20 {
			errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL2InvalidWalletAddress,
				wa,
				len(wa),
			))
		}
	}

	return utils.FlattenErrors(errs)
}
