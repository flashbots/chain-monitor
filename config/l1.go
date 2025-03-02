package config

import (
	"errors"
	"fmt"
	"net/url"

	ethcommon "github.com/ethereum/go-ethereum/common"
)

type L1 struct {
	RPC             string            `yaml:"rpc"`
	WalletAddresses map[string]string `yaml:"wallet_addresses"`
}

var (
	errL1InvalidRPC           = errors.New("invalid l1 rpc url")
	errL1InvalidWalletAddress = errors.New("invalid l1 wallet address")
)

func (cfg *L1) Validate() error {
	if _, err := url.Parse(cfg.RPC); err != nil {
		return fmt.Errorf("%w: %s: %w",
			errL1InvalidRPC,
			cfg.RPC,
			err,
		)
	}

	for _, wa := range cfg.WalletAddresses {
		_addr, err := ethcommon.ParseHexOrString(wa)
		if err != nil {
			return fmt.Errorf("%w: %s: %w",
				errL1InvalidWalletAddress,
				wa,
				err,
			)
		}
		if len(_addr) != 20 {
			return fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
				errL1InvalidWalletAddress,
				wa,
				len(wa),
			)
		}
	}

	return nil
}
