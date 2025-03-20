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

OPTIONS:
   L1

   --l1-monitor-wallets list [ --l1-monitor-wallets list ]  list of l1 wallet addresses to monitor the balances of [$CHAIN_MONITOR_L1_MONITOR_WALLETS]
   --l1-rpc url                                             url of l1 rpc endpoint (default: "http://127.0.0.1:8545") [$CHAIN_MONITOR_L1_RPC]

   L2

   --l2-block-time duration                                 average duration between consecutive blocks on l2 (default: 2s) [$CHAIN_MONITOR_L2_BLOCK_TIME]
   --l2-builder-address address                             l2 builder address [$CHAIN_MONITOR_L2_BUILDER_ADDRESS]
   --l2-monitor-private-key key                             l2 private key to send tx inclusion latency probes with [$CHAIN_MONITOR_L2_MONITOR_PRIVATE_KEY]
   --l2-monitor-tx-gas-limit limit                          l2 gas limit for monitor transactions (default: 22000) [$CHAIN_MONITOR_L2_MONITOR_TX_GAS_LIMIT]
   --l2-monitor-tx-gas-price-adjustment percent             l2 gas price adjustment in percent for monitor transactions (default: 10) [$CHAIN_MONITOR_L2_MONITOR_TX_GAS_PRICE_ADJUSTMENT]
   --l2-monitor-wallets list [ --l2-monitor-wallets list ]  list of l2 wallet addresses to monitor the balances of [$CHAIN_MONITOR_L2_MONITOR_WALLETS]
   --l2-reorg-window duration                               max duration of block history to keep in memory for the l2 reorg adjustments (default: 24h0m0s) [$CHAIN_MONITOR_L2_REORG_WINDOW]
   --l2-rpc url                                             url of l2 rpc endpoint (default: "http://127.0.0.1:8645") [$CHAIN_MONITOR_L2_RPC]

   SERVER

   --server-enable-pprof              whether to enable pprof server (default: false) [$CHAIN_MONITOR_SERVER_ENABLE_PPROF]
   --server-listen-address host:port  host:port for the server to listen on (default: "0.0.0.0:8080") [$CHAIN_MONITOR_SERVER_LISTEN_ADDRESS]
```
