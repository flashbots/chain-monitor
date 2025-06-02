package config

import (
	"errors"
	"fmt"

	"github.com/flashbots/chain-monitor/utils"
	"go.uber.org/zap"
)

type Log struct {
	Level string `yaml:"level"`
	Mode  string `yaml:"mode"`
}

var (
	errLogInvalidLevel = errors.New("invalid log-level")
	errLogInvalidMode  = errors.New("invalid log-mode")
)

func (cfg *Log) Validate() error {
	errs := make([]error, 0)

	if cfg.Mode != "dev" && cfg.Mode != "prod" {
		errs = append(errs, fmt.Errorf("%w: %s",
			errLogInvalidMode, cfg.Mode,
		))
	}

	if _, err := zap.ParseAtomicLevel(cfg.Level); err != nil {
		errs = append(errs, fmt.Errorf("%w: %s: %w",
			errLogInvalidLevel, cfg.Level, err,
		))
	}

	return utils.FlattenErrors(errs)
}
