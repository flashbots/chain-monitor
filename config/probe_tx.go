package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flashbots/chain-monitor/utils"
)

type ProbeTx struct {
	PrivateKey string `yaml:"private_key"`

	ResetInterval  time.Duration `yaml:"reset_interval"`
	ResetThreshold int64         `yaml:"reset_threshold"`

	GasLimit           uint64 `yaml:"gas_limit"`
	GasPriceAdjustment int64  `yaml:"gas_price_adjustment"`
	GasPriceCap        int64  `yaml:"tx_gas_price_cap"`
}

var (
	errL2InvalidMonitorPrivateKey = errors.New("invalid monitor's private key")
)

func (cfg *ProbeTx) Validate() error {
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
