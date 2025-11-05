package config

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flashbots/chain-monitor/utils"
)

type AuthorizeWorkloadIdTx struct {
	PrivateKey string `yaml:"private_key"`
}

var (
	errL2InvalidFlashtestaionsOwnerPrivateKey = errors.New("invalid flashtestations owner private key")
)

func (cfg *AuthorizeWorkloadIdTx) Validate() error {
	errs := make([]error, 0)

	if cfg.PrivateKey != "" {
		if _, err := crypto.HexToECDSA(cfg.PrivateKey); err != nil {
			errs = append(errs, fmt.Errorf("%w: %w",
				errL2InvalidMonitorPrivateKey,
				err,
			))
		}
	}

	return utils.FlattenErrors(errs)
}
