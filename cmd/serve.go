package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/flashbots/chain-monitor/config"
	"github.com/flashbots/chain-monitor/server"
)

const (
	categoryOpt    = "optimism"
	categoryServer = "server"
)

func CommandServe(cfg *config.Config) *cli.Command {
	walletAddresses := &cli.StringSlice{}

	ethFlags := []cli.Flag{
		&cli.DurationFlag{
			Category:    strings.ToUpper(categoryOpt),
			Destination: &cfg.Opt.BlockTime,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryOpt) + "_BLOCK_TIME"},
			Name:        categoryOpt + "-block-time",
			Usage:       "average `duration` between consecutive blocks",
			Value:       12 * time.Second,
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryOpt),
			Destination: &cfg.Opt.BuilderAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryOpt) + "_BUILDER_ADDRESS"},
			Name:        categoryOpt + "-builder-address",
			Required:    true,
			Usage:       "builder `address`",
		},

		&cli.DurationFlag{
			Category:    strings.ToUpper(categoryOpt),
			Destination: &cfg.Opt.ReorgWindow,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryOpt) + "_REORG_WINDOW"},
			Name:        categoryOpt + "-reorg-window",
			Usage:       "average `duration` between consecutive blocks",
			Value:       24 * time.Hour,
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryOpt),
			Destination: &cfg.Opt.RPC,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryOpt) + "_RPC"},
			Name:        categoryOpt + "-rpc",
			Usage:       "`url` of ethereum rpc endpoint",
			Value:       "http://127.0.0.1:8645",
		},

		&cli.StringSliceFlag{
			Category:    strings.ToUpper(categoryOpt),
			Destination: walletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryOpt) + "_MONITOR_WALLETS"},
			Name:        categoryOpt + "-monitor-wallets",
			Usage:       "`list` of wallet addresses to monitor the balances of",
		},
	}

	serverFlags := []cli.Flag{
		&cli.StringFlag{
			Category:    strings.ToUpper(categoryServer),
			Destination: &cfg.Server.ListenAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryServer) + "_LISTEN_ADDRESS"},
			Name:        categoryServer + "-listen-address",
			Usage:       "`host:port` for the server to listen on",
			Value:       "0.0.0.0:8080",
		},
	}

	flags := slices.Concat(
		ethFlags,
		serverFlags,
	)

	return &cli.Command{
		Name:  "serve",
		Usage: "run chain-monitor server",
		Flags: flags,

		Before: func(_ *cli.Context) error {
			_walletAddresses := make(map[string]string, len(walletAddresses.Value()))
			for _, wa := range walletAddresses.Value() {
				parts := strings.Split(wa, "=")
				if len(parts) != 2 {
					return fmt.Errorf("invalid wallet address (mush be like `name=0xNNNN`): %s", wa)
				}
				name := strings.TrimSpace(parts[0])
				addr := strings.TrimSpace(parts[1])
				_walletAddresses[name] = addr
			}

			cfg.Opt.WalletAddresses = _walletAddresses
			return cfg.Validate()
		},

		Action: func(_ *cli.Context) error {
			s, err := server.New(cfg)
			if err != nil {
				return err
			}
			return s.Run()
		},
	}
}
