# chain-monitor

## TL;DR

```shell
go run github.com/flashbots/chain-monitor/cmd serve \
  --l1-rpc http://127.0.0.1:8545 \
  --l1-monitor-wallets batcher=0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN,proposer=0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN \
  --l2-block-time 1s \
  --l2-builder-address 0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN \
  --l2-monitor-wallets builder=0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN \
  --l2-rpc http://127.0.0.1:8645
```

```shell
curl -sS 127.0.0.1:8080/metrics | grep -v -e "^#.*$" | sort
```

```text
chain_monitor_block_missed nnn
chain_monitor_blocks_landed xxx
chain_monitor_blocks_missed yyy
chain_monitor_blocks_seen zzz
chain_monitor_wallet_balance{wallet_address="0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",wallet_name="batcher"} $$$
chain_monitor_wallet_balance{wallet_address="0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",wallet_name="proposer"} $$$
chain_monitor_wallet_balance{wallet_address="0xNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNNN",wallet_name="builder"} $$$
```

## Usage

```text
chain-monitor serve [command options]

NAME:
   chain-monitor serve - run chain-monitor server

USAGE:
   chain-monitor serve [command options]

OPTIONS:
   DIR

   --dir-persistent path  path to the directory where chain-monitor will store its state b/w restarts (default: disabled) [$CHAIN_MONITOR_DIR_PERSISTENT]

   L1

   --l1-monitor-wallet label=address [ --l1-monitor-wallet label=address ]  list of l1 wallet label=address to monitor the balances of [$CHAIN_MONITOR_L1_MONITOR_WALLET]
   --l1-network-id number                                                   on every rpc call, verify that network id matches this number (default: do not check) [$CHAIN_MONITOR_L1_NETWORK_ID]
   --l1-rpc url                                                             url of l1 rpc endpoint [$CHAIN_MONITOR_L1_RPC]
   --l1-rpc-fallback url [ --l1-rpc-fallback url ]                          urls of fallback l1 rpc endpoints [$CHAIN_MONITOR_L1_RPC_FALLBACK]

   L2

   --l2-block-time duration                                                                         average duration between consecutive blocks on l2 (default: 2s) [$CHAIN_MONITOR_L2_BLOCK_TIME]
   --l2-flashblocks-per-block value                                                                 expected count of non-deposit flashblocks per block on l2 (default: 4) [$CHAIN_MONITOR_L2_FLASHBLOCKS_PER_BLOCK]
   --l2-flashtestations-per-block value                                                             expected count of flashtestations per block on l2 (default: 1) [$CHAIN_MONITOR_L2_FLASHTESTATIONS_PER_BLOCK]
   --l2-genesis-time value                                                                          genesis time of the chain (used to determine current height) (default: 0) [$CHAIN_MONITOR_L2_GENESIS_TIME]
   --l2-monitor-builder-address address                                                             l2 builder address to monitor [$CHAIN_MONITOR_L2_MONITOR_BUILDER_ADDRESS]
   --l2-monitor-builder-policy-add-workload-id-event-signature signature                            l2 builder policy event signature to add workload id (default: "WorkloadAddedToPolicy(bytes32)") [$CHAIN_MONITOR_L2_MONITOR_BUILDER_POLICY_ADD_WORKLOAD_ID_EVENT_SIGNATURE]
   --l2-monitor-builder-policy-add-workload-id-signature signature                                  l2 builder policy function signature to add workload id (default: "addWorkloadToPolicy(bytes32,string,string[])") [$CHAIN_MONITOR_L2_MONITOR_BUILDER_POLICY_ADD_WORKLOAD_ID_SIGNATURE]
   --l2-monitor-builder-policy-contract address                                                     l2 builder flashtestations policy contract address to monitor [$CHAIN_MONITOR_L2_MONITOR_BUILDER_POLICY_CONTRACT]
   --l2-monitor-builder-policy-contract-function-signature signature                                l2 builder flashtestations policy contract function signature to monitor (default: "permitVerifyBlockBuilderProof(uint8,bytes32,uint256,bytes)") [$CHAIN_MONITOR_L2_MONITOR_BUILDER_POLICY_CONTRACT_FUNCTION_SIGNATURE]
   --l2-monitor-flashblock-number-contract address                                                  l2 builder flashblock number contract address to monitor [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCK_NUMBER_CONTRACT]
   --l2-monitor-flashblock-number-contract-function-signature signature                             l2 builder flashblock number contract function signature to monitor (default: "incrementFlashblockNumber()") [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCK_NUMBER_CONTRACT_FUNCTION_SIGNATURE]
   --l2-monitor-flashblocks-main-public-stream value                                                the name of the main public l2 flashblocks stream [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCKS_MAIN_PUBLIC_STREAM]
   --l2-monitor-flashblocks-max-ws-message-size-kb value                                            max size (in kb) of l2 builder flashblocks ws messages (default: 256) [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCKS_MAX_WS_MESSAGE_SIZE_KB]
   --l2-monitor-flashblocks-private-stream value [ --l2-monitor-flashblocks-private-stream value ]  private websocket stream(s) of l2 flashblocks [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCKS_PRIVATE_STREAMS]
   --l2-monitor-flashblocks-public-stream value [ --l2-monitor-flashblocks-public-stream value ]    public websocket stream(s) of l2 flashblocks [$CHAIN_MONITOR_L2_MONITOR_FLASHBLOCKS_PUBLIC_STREAMS]
   --l2-monitor-flashtestations-registry-contract address                                           l2 builder flashtestations registry contract address to monitor [$CHAIN_MONITOR_L2_MONITOR_FLASHTESTATIONS_REGISTRY_CONTRACT]
   --l2-monitor-flashtestations-registry-contract-function-signature signature                      l2 builder flashtestations registry contract function signature to monitor (default: "permitRegisterTEEService(bytes,bytes,uint256,uint256,bytes)") [$CHAIN_MONITOR_L2_MONITOR_FLASHTESTATIONS_REGISTRY_CONTRACT_FUNCTION_SIGNATURE]
   --l2-monitor-flashtestations-registry-event-signature signature                                  l2 builder flashtestations registry contract event signature to monitor (default: "TEEServiceRegistered(address,bytes,bool)") [$CHAIN_MONITOR_L2_MONITOR_FLASHTESTATIONS_REGISTRY_EVENT_SIGNATURE]
   --l2-monitor-tx-receipts                                                                         l2 monitor transactions receipts (can be slow on busy chains) (default: false) [$CHAIN_MONITOR_L2_MONITOR_TX_RECEIPTS]
   --l2-monitor-wallet label=address [ --l2-monitor-wallet label=address ]                          list of l2 wallet label=address to monitor the balances of [$CHAIN_MONITOR_L2_MONITOR_WALLET]
   --l2-network-id number                                                                           on every rpc call, verify that network id matches this number (default: do not check) [$CHAIN_MONITOR_L2_NETWORK_ID]
   --l2-probe-tx-gas-limit limit                                                                    l2 probe transaction gas limit (default: 1000000) [$CHAIN_MONITOR_L2_PROBE_TX_GAS_LIMIT]
   --l2-probe-tx-gas-price-adjustment percent                                                       l2 probe transaction gas price adjustment in percent (default: 10) [$CHAIN_MONITOR_L2_PROBE_TX_GAS_PRICE_ADJUSTMENT]
   --l2-probe-tx-gas-price-cap wei                                                                  l2 probe transaction gas price cap in wei (default: 10) [$CHAIN_MONITOR_L2_PROBE_TX_GAS_PRICE_CAP]
   --l2-probe-tx-nonce-reset-interval interval                                                      interval at which to conditionally reset l2 probe tx nonce (default: 10m0s) [$CHAIN_MONITOR_L2_PROBE_TX_NONCE_RESET_INTERVAL]
   --l2-probe-tx-nonce-reset-threshold difference                                                   difference probeTxSent-probeTxLanded that should trigger probe tx nonce reset (default: 60) [$CHAIN_MONITOR_L2PROBE_TX_RESET_THRESHOLD]
   --l2-probe-tx-private-key key                                                                    l2 private key to send tx inclusion latency probes with [$CHAIN_MONITOR_L2_PROBE_TX_PRIVATE_KEY]
   --l2-reorg-window duration                                                                       max duration of block history to keep in memory for the l2 reorg adjustments (default: 24h0m0s) [$CHAIN_MONITOR_L2_REORG_WINDOW]
   --l2-rpc url                                                                                     url of l2 rpc endpoint (default: "http://127.0.0.1:8645") [$CHAIN_MONITOR_L2_RPC]
   --l2-rpc-fallback url [ --l2-rpc-fallback url ]                                                  urls of fallback l2 rpc endpoints [$CHAIN_MONITOR_L2_RPC_FALLBACK]

   SERVER

   --server-enable-pprof              whether to enable pprof server (default: false) [$CHAIN_MONITOR_SERVER_ENABLE_PPROF]
   --server-listen-address host:port  host:port for the server to listen on (default: "0.0.0.0:8080") [$CHAIN_MONITOR_SERVER_LISTEN_ADDRESS]
```
