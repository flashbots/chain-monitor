package config

import (
	"errors"
	"fmt"
	"net"

	"github.com/flashbots/chain-monitor/utils"
)

type Server struct {
	EnablePprof   bool   `yaml:"enable_pprof"`
	ListenAddress string `yaml:"listen_address"`
}

var (
	errServerInvalidListenAddress = errors.New("invalid server listen address")
)

func (cfg *Server) Validate() error {
	errs := make([]error, 0)

	if _, err := net.ResolveTCPAddr("tcp", cfg.ListenAddress); err != nil {
		errs = append(errs, fmt.Errorf("%w: %w",
			errServerInvalidListenAddress, err,
		))
	}

	return utils.FlattenErrors(errs)
}
