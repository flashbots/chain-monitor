# chain-monitor

## TL;DR

```shell
go run github.com/flashbots/chain-monitor/cmd serve \
  --eth-block-time 1s \
  --eth-rpc http://127.0.0.1:8645 \
  --eth-builder-address 0xdD11751cdD3f6EFf01B1f6151B640685bfa5dB4a \
  --eth-monitor-wallets 0xdD11751cdD3f6EFf01B1f6151B640685bfa5dB4a
```

```shell
curl -sS 127.0.0.1:8080/metrics | grep -v -e "^#.*$" | sort
```

```text
chain_monitor_blocks_landed 22
chain_monitor_blocks_seen 22
chain_monitor_wallet_balance{address="0xdD11751cdD3f6EFf01B1f6151B640685bfa5dB4a"} 9.999999995811179e+20
```
