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
