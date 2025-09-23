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
	categoryDir    = "dir"
	categoryL1     = "l1"
	categoryL2     = "l2"
	categoryServer = "server"
)

func CommandServe(cfg *config.Config) *cli.Command {
	l1RpcFallback := &cli.StringSlice{}
	l1WalletAddresses := &cli.StringSlice{}
	l2RpcFallback := &cli.StringSlice{}
	l2WalletAddresses := &cli.StringSlice{}

	dirFlags := []cli.Flag{
		&cli.StringFlag{ // --dir-persistent
			Category:    strings.ToUpper(categoryDir),
			DefaultText: "disabled",
			Destination: &cfg.Dir.Persistent,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryDir) + "_PERSISTENT"},
			Name:        categoryDir + "-persistent",
			Usage:       "`path` to the directory where chain-monitor will store its state b/w restarts",
		},
	}

	l1Flags := []cli.Flag{
		&cli.StringSliceFlag{ // --l1-monitor-wallet
			Category:    strings.ToUpper(categoryL1),
			Destination: l1WalletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL1) + "_MONITOR_WALLET"},
			Name:        categoryL1 + "-monitor-wallet",
			Usage:       "list of l1 wallet `label=address` to monitor the balances of",
		},

		&cli.Uint64Flag{
			Category:    strings.ToUpper(categoryL1),
			Destination: &cfg.L1.NetworkID,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL1) + "_NETWORK_ID"},
			Name:        categoryL1 + "-network-id",
			Usage:       "on every rpc call, verify that network id matches this `number`",
			Value:       0,
			DefaultText: "do not check",
		},

		&cli.StringFlag{ // --l1-rpc
			Category:    strings.ToUpper(categoryL1),
			Destination: &cfg.L1.Rpc,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL1) + "_RPC"},
			Name:        categoryL1 + "-rpc",
			Usage:       "`url` of l1 rpc endpoint",
		},

		&cli.StringSliceFlag{ // --l1-rpc-fallback
			Category:    strings.ToUpper(categoryL1),
			Destination: l1RpcFallback,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL1) + "_RPC_FALLBACK"},
			Name:        categoryL1 + "-rpc-fallback",
			Usage:       "`url`s of fallback l1 rpc endpoints",
		},
	}

	l2Flags := []cli.Flag{
		&cli.DurationFlag{ // --l2-block-time
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.BlockTime,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_BLOCK_TIME"},
			Name:        categoryL2 + "-block-time",
			Usage:       "average `duration` between consecutive blocks on l2",
			Value:       2 * time.Second,
		},

		&cli.Uint64Flag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.GenesisTime,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_GENESIS_TIME"},
			Name:        categoryL2 + "-genesis-time",
			Usage:       "genesis time of the chain (used to determine current height)",
			Value:       0,
		},

		&cli.StringFlag{ // --l2-monitor-builder-address
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorBuilderAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_BUILDER_ADDRESS"},
			Name:        categoryL2 + "-monitor-builder-address",
			Usage:       "l2 builder `address` to monitor",
		},

		&cli.StringFlag{ // --l2-monitor-builder-policy-contract
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorBuilderPolicyContract,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_BUILDER_POLICY_CONTRACT"},
			Name:        categoryL2 + "-monitor-builder-policy-contract",
			Usage:       "l2 builder flashtestations policy contract `address` to monitor",
		},

		&cli.StringFlag{ // --l2-monitor-builder-policy-contract-function-signature
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorBuilderPolicyContractFunctionSignature,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_BUILDER_POLICY_CONTRACT_FUNCTION_SIGNATURE"},
			Name:        categoryL2 + "-monitor-builder-policy-contract-function-signature",
			Usage:       "l2 builder flashtestations policy contract function `signature` to monitor",
			Value:       "verifyBlockBuilderProof(uint8,bytes32)",
		},

		&cli.BoolFlag{ // --l2-monitor-tx-receipts
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.MonitorTxReceipts,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_TX_RECEIPTS"},
			Name:        categoryL2 + "-monitor-tx-receipts",
			Usage:       "l2 monitor transactions receipts (can be slow on busy chains)",
			Value:       false,
		},

		&cli.StringSliceFlag{ // --l2-monitor-wallet
			Category:    strings.ToUpper(categoryL2),
			Destination: l2WalletAddresses,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_MONITOR_WALLET"},
			Name:        categoryL2 + "-monitor-wallet",
			Usage:       "list of l2 wallet `label=address` to monitor the balances of",
		},

		&cli.Uint64Flag{
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.NetworkID,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_NETWORK_ID"},
			Name:        categoryL2 + "-network-id",
			Usage:       "on every rpc call, verify that network id matches this `number`",
			Value:       0,
			DefaultText: "do not check",
		},

		&cli.Uint64Flag{ // --l2-probe-tx-gas-limit
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.GasLimit,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_PROBE_TX_GAS_LIMIT"},
			Name:        categoryL2 + "-probe-tx-gas-limit",
			Usage:       "l2 probe transaction gas `limit`",
			Value:       1000000,
		},

		&cli.Int64Flag{ // --l2-probe-tx-gas-price-adjustment
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.GasPriceAdjustment,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_PROBE_TX_GAS_PRICE_ADJUSTMENT"},
			Name:        categoryL2 + "-probe-tx-gas-price-adjustment",
			Usage:       "l2 probe transaction gas price adjustment in `percent`",
			Value:       10,
		},

		&cli.Int64Flag{ // --l2-probe-tx-gas-price-cap
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.GasPriceCap,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_PROBE_TX_GAS_PRICE_CAP"},
			Name:        categoryL2 + "-probe-tx-gas-price-cap",
			Usage:       "l2 probe transaction gas price cap in `wei`",
			Value:       10,
		},

		&cli.DurationFlag{ // --l2-probe-tx-nonce-reset-interval
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.ResetInterval,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_PROBE_TX_NONCE_RESET_INTERVAL"},
			Name:        categoryL2 + "-probe-tx-nonce-reset-interval",
			Usage:       "`interval` at which to conditionally reset l2 probe tx nonce",
			Value:       10 * time.Minute,
		},

		&cli.Int64Flag{ // --l2-probe-tx-nonce-reset-threshold
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.ResetThreshold,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "PROBE_TX_RESET_THRESHOLD"},
			Name:        categoryL2 + "-probe-tx-nonce-reset-threshold",
			Usage:       "`difference` probeTxSent-probeTxLanded that should trigger probe tx nonce reset",
			Value:       60,
		},

		&cli.StringFlag{ // --l2-probe-tx-private-key
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ProbeTx.PrivateKey,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_PROBE_TX_PRIVATE_KEY"},
			Name:        categoryL2 + "-probe-tx-private-key",
			Usage:       "l2 private `key` to send tx inclusion latency probes with",
		},

		&cli.DurationFlag{ // --l2-reorg-window
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.ReorgWindow,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_REORG_WINDOW"},
			Name:        categoryL2 + "-reorg-window",
			Usage:       "max `duration` of block history to keep in memory for the l2 reorg adjustments",
			Value:       24 * time.Hour,
		},

		&cli.StringFlag{ // --l2-rpc
			Category:    strings.ToUpper(categoryL2),
			Destination: &cfg.L2.Rpc,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_RPC"},
			Name:        categoryL2 + "-rpc",
			Usage:       "`url` of l2 rpc endpoint",
			Value:       "http://127.0.0.1:8645",
		},

		&cli.StringSliceFlag{ // --l2-rpc-fallback
			Category:    strings.ToUpper(categoryL2),
			Destination: l2RpcFallback,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryL2) + "_RPC_FALLBACK"},
			Name:        categoryL2 + "-rpc-fallback",
			Usage:       "`url`s of fallback l2 rpc endpoints",
		},
	}

	serverFlags := []cli.Flag{
		&cli.BoolFlag{ // --server-enable-pprof
			Category:    strings.ToUpper(categoryServer),
			Destination: &cfg.Server.EnablePprof,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryServer) + "_ENABLE_PPROF"},
			Name:        categoryServer + "-enable-pprof",
			Usage:       "whether to enable pprof server",
			Value:       false,
		},

		&cli.StringFlag{ // --server-enable-pprof
			Category:    strings.ToUpper(categoryServer),
			Destination: &cfg.Server.ListenAddress,
			EnvVars:     []string{envPrefix + strings.ToUpper(categoryServer) + "_LISTEN_ADDRESS"},
			Name:        categoryServer + "-listen-address",
			Usage:       "`host:port` for the server to listen on",
			Value:       "0.0.0.0:8080",
		},
	}

	flags := slices.Concat(
		dirFlags,
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
				cfg.L1.RpcFallback = l1RpcFallback.Value()

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
				cfg.L1.MonitorWalletAddresses = _walletAddresses
			}

			{
				cfg.L2.RpcFallback = l2RpcFallback.Value()

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
				cfg.L2.MonitorWalletAddresses = _walletAddresses
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
