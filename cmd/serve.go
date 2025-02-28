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
	categoryEth    = "eth"
	categoryServer = "server"
)

func CommandServe(cfg *config.Config) *cli.Command {
	walletAddresses := &cli.StringSlice{}

	ethFlags := []cli.Flag{
		&cli.DurationFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: &cfg.Eth.BlockTime,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_BLOCK_TIME"},
			Name:        categoryEth + "-block-time",
			Usage:       "average `duration` between consecutive blocks",
			Value:       12 * time.Second,
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: &cfg.Eth.BuilderAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_BUILDER_ADDRESS"},
			Name:        categoryEth + "-builder-address",
			Required:    true,
			Usage:       "builder `address`",
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: &cfg.Eth.RPC,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_RPC"},
			Name:        categoryEth + "-rpc",
			Usage:       "`url` of ethereum rpc endpoint",
			Value:       "http://127.0.0.1:8645",
		},

		&cli.StringSliceFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: walletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_MONITOR_WALLETS"},
			Name:        categoryEth + "-monitor-wallets",
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

			cfg.Eth.WalletAddresses = _walletAddresses
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
