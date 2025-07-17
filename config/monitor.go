package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flashbots/chain-monitor/utils"
)

type Monitor struct {
	PrivateKey string `yaml:"private_key"`

	MaxGasPerBlock uint64 `yaml:"max_gas_per_block"`
	MaxGasPrice    uint64 `yaml:"max_gas_price"`

	ResetInterval  time.Duration `yaml:"reset_interval"`
	ResetThreshold int64         `yaml:"reset_threshold"`

	TxGasLimit           uint64 `yaml:"tx_gas_limit"`
	TxGasPriceAdjustment int64  `yaml:"tx_gas_price_adjustment"`
	TxGasPriceCap        int64  `yaml:"tx_gas_price_cap"`

	TxReceipts bool `yaml:"tx_receipts"`
}

var (
	errL2InvalidMonitorPrivateKey = errors.New("invalid monitor's private key")
)

func (cfg *Monitor) Validate() error {
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
