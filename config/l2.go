package config

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/flashbots/chain-monitor/utils"
)

type L2 struct {
	Dir *Dir `yaml:"-"`

	BlockTime           time.Duration `yaml:"block_time"`
	FlashblocksPerBlock int64         `yaml:"flashblocks_per_block"`
	GenesisTime         uint64        `yaml:"genesis_time"`
	NetworkID           uint64        `yaml:"network_id"`
	ReorgWindow         time.Duration `yaml:"reorg_window"`
	Rpc                 string        `yaml:"rpc"`
	RpcFallback         []string      `yaml:"rpc_fallback"`

	MonitorBuilderAddress                            string            `yaml:"monitor_builder_address"`
	MonitorBuilderPolicyContract                     string            `yaml:"monitor_builder_policy_contract"`
	MonitorBuilderPolicyContractFunctionSignature    string            `yaml:"monitor_builder_policy_contract_function_signature"`
	MonitorFlashblockNumberContract                  string            `yaml:"monitor_builder_flashblock_number_contract"`
	MonitorFlashblockNumberContractFunctionSignature string            `yaml:"monitor_builder_flashblock_number_contract_function_signature"`
	MonitorTxReceipts                                bool              `yaml:"monitor_tx_receipts"`
	MonitorWalletAddresses                           map[string]string `yaml:"monitor_wallet_addresses"`

	ProbeTx *ProbeTx `yaml:"probe"`
}

const (
	maxReorgWindow = 24 * time.Hour
)

var (
	errL2InvalidBuilderAddress          = errors.New("invalid l2 builder address")
	errL2InvalidBuilderPolicyContact    = errors.New("invalid l2 builder policy contract address")
	errL2InvalidFlashblockNumberContact = errors.New("invalid l2 flashblocks number contract address")
	errL2InvalidRpc                     = errors.New("invalid l2 rpc url")
	errL2InvalidRpcFallback             = errors.New("invalid l2 fallback rpc url")
	errL2InvalidWalletAddress           = errors.New("invalid l2 wallet address")
	errL2ReorgWindowTooLarge            = errors.New("l2 reorg window is too large")
)

func (cfg *L2) Validate() error {
	errs := make([]error, 0)

	{ // reorg_window
		if cfg.ReorgWindow > maxReorgWindow {
			errs = append(errs, fmt.Errorf("%w (max %d): %d",
				errL2ReorgWindowTooLarge,
				maxReorgWindow,
				cfg.ReorgWindow,
			))
		}
	}

	{ // rpc
		if _, err := url.Parse(cfg.Rpc); err != nil {
			errs = append(errs, fmt.Errorf("%w: %s: %w",
				errL2InvalidRpc,
				cfg.Rpc,
				err,
			))
		}
	}

	{ // rpc_fallback
		for _, rpc := range cfg.RpcFallback {
			if _, err := url.Parse(rpc); err != nil {
				errs = append(errs, fmt.Errorf("%w: %s: %w",
					errL2InvalidRpcFallback,
					rpc,
					err,
				))
			}
		}
	}

	{ // monitor_builder_address
		if cfg.MonitorBuilderAddress != "" {
			_addr, err := ethcommon.ParseHexOrString(cfg.MonitorBuilderAddress)
			if err != nil {
				errs = append(errs, fmt.Errorf("%w: %s: %w",
					errL2InvalidBuilderAddress,
					cfg.MonitorBuilderAddress,
					err,
				))
			}
			if len(_addr) != 20 {
				errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
					errL2InvalidBuilderAddress,
					cfg.MonitorBuilderAddress,
					len(_addr),
				))
			}
		}
	}

	{ // monitor_builder_policy_contract
		if cfg.MonitorBuilderPolicyContract != "" {
			_addr, err := ethcommon.ParseHexOrString(cfg.MonitorBuilderPolicyContract)
			if err != nil {
				errs = append(errs, fmt.Errorf("%w: %s: %w",
					errL2InvalidBuilderPolicyContact,
					cfg.MonitorBuilderPolicyContract,
					err,
				))
			}
			if len(_addr) != 20 {
				errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
					errL2InvalidBuilderPolicyContact,
					cfg.MonitorBuilderPolicyContract,
					len(_addr),
				))
			}
		}
	}

	{ // monitor_builder_flashblock_number_contract
		if cfg.MonitorFlashblockNumberContract != "" {
			_addr, err := ethcommon.ParseHexOrString(cfg.MonitorFlashblockNumberContract)
			if err != nil {
				errs = append(errs, fmt.Errorf("%w: %s: %w",
					errL2InvalidFlashblockNumberContact,
					cfg.MonitorFlashblockNumberContract,
					err,
				))
			}
			if len(_addr) != 20 {
				errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
					errL2InvalidFlashblockNumberContact,
					cfg.MonitorFlashblockNumberContract,
					len(_addr),
				))
			}
		}
	}

	{ // monitor_wallet_address
		for _, wa := range cfg.MonitorWalletAddresses {
			_addr, err := ethcommon.ParseHexOrString(wa)
			if err != nil {
				errs = append(errs, fmt.Errorf("%w: %s: %w",
					errL2InvalidWalletAddress,
					wa,
					err,
				))
			}
			if len(_addr) != 20 {
				errs = append(errs, fmt.Errorf("%w: %s: invalid length (want 20, got %d)",
					errL2InvalidWalletAddress,
					wa,
					len(wa),
				))
			}
		}
	}

	return utils.FlattenErrors(errs)
}
