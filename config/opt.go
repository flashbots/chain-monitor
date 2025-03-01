package config

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type Opt struct {
	BlockTime       time.Duration     `yaml:"block_time"`
	BuilderAddress  string            `yaml:"builder_address"`
	ReorgWindow     time.Duration     `yaml:"reorg_window"`
	RPC             string            `yaml:"rpc"`
	WalletAddresses map[string]string `yaml:"wallet_addresses"`
}

const (
	maxReorgWindow = 86400
)

var (
	errEthInvalidBuilderAddress = errors.New("invalid builder address")
	errEthInvalidRPC            = errors.New("invalid rpc url")
	errEthInvalidWalletAddress  = errors.New("invalid wallet address")
	errEthReorgWindowTooLarge   = errors.New("reorg window is too large")
)

func (cfg *Opt) Validate() error {
	if _, err := url.Parse(cfg.RPC); err != nil {
		return fmt.Errorf("%w: %s: %w",
			errEthInvalidRPC,
			cfg.RPC,
			err,
		)
	}

	if cfg.ReorgWindow > maxReorgWindow {
		return fmt.Errorf("%w (max %d): %d",
			errEthReorgWindowTooLarge,
			maxReorgWindow,
			cfg.ReorgWindow,
		)
	}

	_addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
	if err != nil {
		return fmt.Errorf("%w: %s: %w",
			errEthInvalidBuilderAddress,
			cfg.BuilderAddress,
			err,
		)
	}
	if len(_addr) != 20 {
		return fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
			errEthInvalidBuilderAddress,
			cfg.BuilderAddress,
			len(_addr),
		)
	}

	for _, wa := range cfg.WalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(cfg.BuilderAddress)
		if err != nil {
			return fmt.Errorf("%w: %s: %w",
				errEthInvalidWalletAddress,
				wa,
				err,
			)
		}
		if len(_addr) != 20 {
			return fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errEthInvalidWalletAddress,
				cfg.BuilderAddress,
				len(wa),
			)
		}
	}

	return nil
}
