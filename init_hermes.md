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
id = "chain-a"
type = "CosmosSdk"
rpc_addr = "http://localhost:26647"
grpc_addr = "http://localhost:9080"
rpc_timeout = "15s"
trusted_node = true
account_prefix = "cosmos"
key_name = "wallet"
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
url = "ws://localhost:26647/websocket"
batch_delay = "200ms"

[chains.trust_threshold]
numerator = "1"
denominator = "3"

[chains.gas_price]
price = 1.0
denom = "stake"

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

# [chains.trust_threshold]
# numerator = "2"
# denominator = "3"

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

# Add Gaia chain-a key (replace with your actual mnemonic)
echo "your_gaia_wallet_mnemonic_here" > ~/.hermes/keys/chain-a_key
hermes keys add --chain chain-a --mnemonic-file ~/.hermes/keys/chain-a_key
hermes --config hermes_config.toml keys list --chain chain-a


# Add SIX Protocol testnet key
echo "first educate action fee physical seek recipe hub anxiety best mango measure chimney sphere once cabbage strike dizzy near knock correct answer skin inside" > ~/.hermes/keys/alice_key
hermes --config hermes_config.toml keys add --chain testnet --mnemonic-file ~/.hermes/keys/alice_key
hermes --config hermes_config.toml keys list --chain testnet

# IMPORTANT: Fund the alice account on SIX Protocol testnet
# Address: 6x1kj9tl6pavd4825atx0e959unjt7g2l35knjtkm
sixd tx bank send alice 6x1kj9tl6pavd4825atx0e959unjt7g2l35knjtkm 10000000usix --keyring-backend test --chain-id testnet --gas 200000 --gas-prices 1.25usix --yes

# Create IBC clients
hermes --config hermes_config.toml create client --host-chain chain-a --reference-chain testnet
hermes --config hermes_config.toml create client --host-chain testnet --reference-chain chain-a

# Create IBC connection
hermes --config hermes_config.toml create connection --a-chain chain-a --b-chain testnet
# Result: chain-a connection-2 <-> testnet connection-0

# Create IBC channel for transfers  
hermes --config hermes_config.toml create channel --order unordered --a-chain chain-a --a-connection connection-2 --a-port transfer --b-port transfer
# Result: chain-a channel-1 <-> testnet channel-0

hermes --config hermes_config.toml start

hermes --config hermes_config.toml query channels --chain chain-a
hermes --config hermes_config.toml query channels --show-counterparty --chain chain-a
hermes --config hermes_config.toml query channels --show-counterparty --chain testnet
hermes --config hermes_config.toml query connections --chain testnet
hermes --config hermes_config.toml query connections --chain chain-a
hermes --config hermes_config.toml query packet acks --chain chain-a --port transfer --channel channel-1

## CHECK CHANNEL IS EXPIRE
hermes --config hermes_config.toml query channel end --chain testnet --port transfer --channel channel-0
hermes --config hermes_config.toml query connection end --chain testnet --connection connection-0
hermes --config hermes_config.toml query client state --chain testnet --client 07-tendermint-0
hermes --config hermes_config.toml query client status --chain testnet --client 07-tendermint-0

hermes --config hermes_config.toml query channel end --chain chain-a --port transfer --channel channel-1
hermes --config hermes_config.toml query connection end --chain chain-a --connection connection-2
hermes --config hermes_config.toml query client state --chain chain-a --client 07-tendermint-5
hermes --config hermes_config.toml query client status --chain chain-a --client 07-tendermint-5

## IBC TRANSFERS (UPDATED - WORKING CHANNELS)

# Current active connection and channels:
# chain-a connection-3 <-> testnet connection-1
# chain-a channel-2 <-> testnet channel-1

# Send tokens from chain-a (Gaia) to testnet (SIX Protocol)
hermes --config hermes_config.toml tx ft-transfer --dst-chain testnet --src-chain chain-a --src-port transfer --src-channel channel-2 --amount 1000000 --denom stake --timeout-height-offset 1000

# Send tokens from testnet (SIX Protocol) to chain-a (Gaia)  
hermes --config hermes_config.toml tx ft-transfer --dst-chain chain-a --src-chain testnet --src-port transfer --src-channel channel-1 --amount 1000000 --denom usix --timeout-height-offset 1000

## CLIENT REFRESH (if clients expire)
# If clients are too old, create new ones:
hermes --config hermes_config.toml create client --host-chain chain-a --reference-chain testnet
hermes --config hermes_config.toml create client --host-chain testnet --reference-chain chain-a
# Then create new connection and channel with the new client IDs
```
