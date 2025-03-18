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
	categoryEth    = "l1"
	categoryL2     = "l2"
	categoryServer = "server"
)

func CommandServe(cfg *config.Config) *cli.Command {
	l1WalletAddresses := &cli.StringSlice{}
	l2WalletAddresses := &cli.StringSlice{}

	l1Flags := []cli.Flag{
		&cli.StringFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: &cfg.L1.RPC,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_RPC"},
			Name:        categoryEth + "-rpc",
			Usage:       "`url` of l1 rpc endpoint",
			Value:       "http://127.0.0.1:8545",
		},

		&cli.StringSliceFlag{
			Category:    strings.ToUpper(categoryEth),
			Destination: l1WalletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryEth) + "_MONITOR_WALLETS"},
			Name:        categoryEth + "-monitor-wallets",
			Usage:       "`list` of l1 wallet addresses to monitor the balances of",
		},
	}

	l2Flags := []cli.Flag{
		&cli.DurationFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.BlockTime,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_BLOCK_TIME"},
			Name:        categoryL2 + "-block-time",
			Usage:       "average `duration` between consecutive blocks on l2",
			Value:       2 * time.Second,
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.BuilderAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_BUILDER_ADDRESS"},
			Name:        categoryL2 + "-builder-address",
			Usage:       "l2 builder `address`",
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorPrivateKey,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_PRIVATE_KEY"},
			Name:        categoryL2 + "-monitor-private-key",
			Usage:       "l2 private `key` to send tx inclusion latency probes with",
		},

		&cli.Uint64Flag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorTxGasLimit,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_TX_GAS_LIMIT"},
			Name:        categoryL2 + "-monitor-tx-gas-limit",
			Usage:       "l2 gas `limit` for monitor transactions",
			Value:       22000,
		},

		&cli.Int64Flag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorTxGasPriceAdjustment,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_TX_GAS_PRICE_ADJUSTMENT"},
			Name:        categoryL2 + "-monitor-tx-gas-price-adjustment",
			Usage:       "l2 gas price adjustment in `percent` for monitor transactions",
			Value:       10,
		},

		&cli.DurationFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ReorgWindow,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_REORG_WINDOW"},
			Name:        categoryL2 + "-reorg-window",
			Usage:       "max `duration` of block history to keep in memory for the l2 reorg adjustments",
			Value:       24 * time.Hour,
		},

		&cli.StringFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.RPC,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_RPC"},
			Name:        categoryL2 + "-rpc",
			Usage:       "`url` of l2 rpc endpoint",
			Value:       "http://127.0.0.1:8645",
		},

		&cli.StringSliceFlag{
			Category:    strings.ToUpper(categoryL2),
			Destination: l2WalletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_WALLETS"},
			Name:        categoryL2 + "-monitor-wallets",
			Usage:       "`list` of l2 wallet addresses to monitor the balances of",
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
		l1Flags,
		l2Flags,
		serverFlags,
	)

	return &cli.Command{
		Name:  "serve",
		Usage: "run chain-monitor server",
		Flags: flags,

		Before: func(_ *cli.Context) error {
			{
				_walletAddresses := make(map[string]string, len(l1WalletAddresses.Value()))
				for _, wa := range l1WalletAddresses.Value() {
					parts := strings.Split(wa, "=")
					if len(parts) != 2 {
						return fmt.Errorf("invalid wallet address (mush be like `name=0xNNNN`): %s", wa)
					}
					name := strings.TrimSpace(parts[0])
					addr := strings.TrimSpace(parts[1])
					_walletAddresses[name] = addr
				}

				cfg.L1.WalletAddresses = _walletAddresses
			}

			{
				_walletAddresses := make(map[string]string, len(l2WalletAddresses.Value()))
				for _, wa := range l2WalletAddresses.Value() {
					parts := strings.Split(wa, "=")
					if len(parts) != 2 {
						return fmt.Errorf("invalid wallet address (mush be like `name=0xNNNN`): %s", wa)
					}
					name := strings.TrimSpace(parts[0])
					addr := strings.TrimSpace(parts[1])
					_walletAddresses[name] = addr
				}

				cfg.L2.WalletAddresses = _walletAddresses
			}

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
