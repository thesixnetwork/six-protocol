# INIT HERMES RELAYER ON LOCAL

```toml
[global]
log_level = "info"

[mode.clients]
enabled = true
refresh = true
misbehaviour = true

[mode.connections]
enabled = false

[mode.channels]
enabled = false

[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = false
auto_register_counterparty_payee = false

[rest]
enabled = false
host = "127.0.0.1"
port = 3000

[telemetry]
enabled = false
host = "127.0.0.1"
port = 3001

[telemetry.buckets.latency_submitted]
start = 500
end = 20000
buckets = 10

[telemetry.buckets.latency_confirmed]
start = 1000
end = 30000
buckets = 10

[[chains]]
id = "pepe_555555-1"
type = "CosmosSdk"
rpc_addr = "http://0.0.0.0:26659/"
grpc_addr = "http://0.0.0.0:9092/"
rpc_timeout = "10s"
trusted_node = false
account_prefix = "lol"
key_name = "alice"
key_store_type = "Test"
store_prefix = "ibc"
default_gas = 100000
max_gas = 400000
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 180000
max_grpc_decoding_size = 33554432
clock_drift = "5s"
max_block_time = "30s"
ccv_consumer_chain = false
memo_prefix = ""
sequential_batch_tx = false

[chains.event_source]
mode = "push"
url = "ws://0.0.0.0:26659/websocket"
batch_delay = "500ms"

[chains.trust_threshold]
numerator = "1"
denominator = "3"

[chains.gas_price]
price = 0.025
denom = "apepe"

[chains.packet_filter]
policy = "allowall"

[chains.packet_filter.min_fees]

[chains.address_type]
derivation = "cosmos"

[[chains]]
id = "testnet"
type = "CosmosSdk"
rpc_addr = "http://0.0.0.0:26657/"
grpc_addr = "http://0.0.0.0:9090/"
rpc_timeout = "10s"
trusted_node = false
account_prefix = "6x"
key_name = "alice"
key_store_type = "Test"
store_prefix = "ibc"
default_gas = 100000
max_gas = 400000
gas_multiplier = 1.1
max_msg_num = 30
max_tx_size = 180000
max_grpc_decoding_size = 33554432
clock_drift = "5s"
max_block_time = "30s"
ccv_consumer_chain = false
memo_prefix = ""
sequential_batch_tx = false

[chains.event_source]
mode = "push"
url = "ws://0.0.0.0:26657/websocket"
batch_delay = "500ms"

[chains.trust_threshold]
numerator = "1"
denominator = "3"

[chains.gas_price]
price = 1.25
denom = "usix"

[chains.packet_filter]
policy = "allowall"

[chains.packet_filter.min_fees]

[chains.address_type]
derivation = "cosmos"


[tracing_server]
enabled = false
port = 5555
```

```bash

echo word word word word word word word word word word word word word word word word word word word word word word word word >> pepe_key
hermes keys add --chain <chain_id_a> --mnemonic-file pepe_key
hermes keys list --chain pepe_555555-1
echo word word word word word word word word word word word word word word word word word word word word word word word word >> testnet_key
hermes keys add --chain testnet --mnemonic-file testnet_key
hermes create client --host-chain pepe_555555-1 --reference-chain testnet
hermes create client --host-chain testnet --reference-chain pepe_555555-1
hermes create connection --a-chain testnet --b-chain pepe_555555-1
hermes create connection --a-chain pepe_555555-1 --b-chain testnet
hermes create channel --order unordered --a-chain testnet --a-connection connection-0 --a-port transfer --b-port transfer

hermes query channels --chain pepe_555555-1
hermes query channels --show-counterparty --chain pepe_555555-1
hermes query channels --show-counterparty --chain testnet
hermes query connections --chain testnet
hermes query connections --chain pepe_555555-1
hermes query packet acks --chain pepe_555555-1 --port transfer --channel channel-0
```
