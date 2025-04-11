package config

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

type Monitor struct {
	PrivateKey string `yaml:"private_key"`

	MaxGasPerBlock uint64 `yaml:"max_gas_per_block"`
	MaxGasPrice    uint64 `yaml:"max_gas_price"`

	TxGasLimit           uint64 `yaml:"tx_gas_limit"`
	TxGasPriceAdjustment int64  `yaml:"tx_gas_price_adjustment"`
	TxGasPriceCap        int64  `yaml:"tx_gas_price_cap"`
}

var (
	errL2InvalidMonitorPrivateKey = errors.New("invalid monitor's private key")
)

func (cfg *Monitor) Validate() error {
	if cfg.PrivateKey != "" {
		if _, err := crypto.HexToECDSA(cfg.PrivateKey); err != nil {
			return fmt.Errorf("%w: %w",
				errL2InvalidMonitorPrivateKey,
				err,
			)
		}
	}

	return nil
}
