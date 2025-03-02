package config

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type L2 struct {
	BlockTime       time.Duration     `yaml:"block_time"`
	BuilderAddress  string            `yaml:"builder_address"`
	ReorgWindow     time.Duration     `yaml:"reorg_window"`
	RPC             string            `yaml:"rpc"`
	WalletAddresses map[string]string `yaml:"wallet_addresses"`
}

const (
	maxReorgWindow = 24 * time.Hour
)

var (
	errL2InvalidBuilderAddress = errors.New("invalid l2 builder address")
	errL2InvalidRPC            = errors.New("invalid l2 rpc url")
	errL2InvalidWalletAddress  = errors.New("invalid l2 wallet address")
	errL2ReorgWindowTooLarge   = errors.New("l2 reorg window is too large")
)

func (cfg *L2) Validate() error {
	if _, err := url.Parse(cfg.RPC); err != nil {
		return fmt.Errorf("%w: %s: %w",
			errL2InvalidRPC,
			cfg.RPC,
			err,
		)
	}

	if cfg.ReorgWindow > maxReorgWindow {
		return fmt.Errorf("%w (max %d): %d",
			errL2ReorgWindowTooLarge,
			maxReorgWindow,
			cfg.ReorgWindow,
		)
	}

	if cfg.BuilderAddress != "" {
		_addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
		if err != nil {
			return fmt.Errorf("%w: %s: %w",
				errL2InvalidBuilderAddress,
				cfg.BuilderAddress,
				err,
			)
		}
		if len(_addr) != 20 {
			return fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL2InvalidBuilderAddress,
				cfg.BuilderAddress,
				len(_addr),
			)
		}
	}

	for _, wa := range cfg.WalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			return fmt.Errorf("%w: %s: %w",
				errL2InvalidWalletAddress,
				wa,
				err,
			)
		}
		if len(_addr) != 20 {
			return fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL2InvalidWalletAddress,
				wa,
				len(wa),
			)
		}
	}

	return nil
}
