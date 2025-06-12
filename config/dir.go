package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/flashbots/chain-monitor/utils"
)

type Dir struct {
	Persistent string `yaml:"persistent"`
}

var (
	errDirNotDirectory = errors.New("not a directory")
)

func (cfg *Dir) Validate() error {
	errs := make([]error, 0)

	if cfg.Persistent != "" { // persistent
		if info, err := os.Stat(cfg.Persistent); err != nil {
			if !os.IsNotExist(err) {
				if errMkdir := os.Mkdir(cfg.Persistent, 0640); errMkdir != nil {
					errs = append(errs, err, errMkdir)
				}
			} else {
				errs = append(errs, err)
			}
		} else {
			if !info.IsDir() {
				errs = append(errs, fmt.Errorf("%w: %s",
					errDirNotDirectory, cfg.Persistent,
				))
			}
		}
	}

	return utils.FlattenErrors(errs)
}
