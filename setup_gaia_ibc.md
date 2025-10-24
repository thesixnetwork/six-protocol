# Setup Gaia Chain-A with SIX Protocol Testnet IBC Connection

## Step 1: Start Gaia Chain-A Locally

### Prerequisites
```bash
# Install Gaia (Cosmos Hub)
git clone https://github.com/cosmos/gaia.git
cd gaia
git checkout v15.2.0  # or latest stable version
make install
```

### Initialize Chain-A (Gaia)
```bash
# Initialize chain-a
gaiad init chain-a --chain-id chain-a
cd ~/.gaia

# Create validator key
gaiad keys add validator --keyring-backend test

# Add genesis account
gaiad genesis add-genesis-account validator 100000000000stake --keyring-backend test

# Create genesis transaction
gaiad genesis gentx validator 1000000stake --chain-id chain-a --keyring-backend test

# Collect genesis transactions
gaiad genesis collect-gentxs

# Start chain-a
gaiad start --rpc.laddr tcp://0.0.0.0:26657 --grpc.address 0.0.0.0:9090
```

## Step 2: Update Hermes Configuration

Replace the `pepe_555555-1` chain configuration with Gaia chain-a:

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

# Chain-A (Gaia Local)
[[chains]]
id = "chain-a"
type = "CosmosSdk"
rpc_addr = "http://127.0.0.1:26657/"
grpc_addr = "http://127.0.0.1:9090/"
rpc_timeout = "10s"
trusted_node = false
account_prefix = "cosmos"
key_name = "relayer"
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
url = "ws://127.0.0.1:26657/websocket"
batch_delay = "500ms"

[chains.trust_threshold]
numerator = "1"
denominator = "3"

[chains.gas_price]
price = 0.025
denom = "stake"

[chains.packet_filter]
policy = "allowall"

[chains.address_type]
derivation = "cosmos"

# SIX Protocol Testnet (keep existing)
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

[chains.address_type]
derivation = "cosmos"

[tracing_server]
enabled = false
port = 5555
```

## Step 3: Setup Hermes Keys and Connection

```bash
# Create relayer account on Gaia chain-a
gaiad keys add relayer --keyring-backend test

# Fund the relayer account
gaiad tx bank send validator $(gaiad keys show relayer -a --keyring-backend test) 1000000stake --chain-id chain-a --keyring-backend test --yes

# Export mnemonic for Hermes (replace with actual mnemonic)
echo "your_gaia_relayer_mnemonic_here" > ~/.hermes/keys/chain-a_key
hermes keys add --chain chain-a --mnemonic-file ~/.hermes/keys/chain-a_key

# Keep existing SIX Protocol key
echo "first educate action fee physical seek recipe hub anxiety best mango measure chimney sphere once cabbage strike dizzy near knock correct answer skin inside" > ~/.hermes/keys/testnet_key
hermes keys add --chain testnet --mnemonic-file ~/.hermes/keys/testnet_key

# Verify keys
hermes keys list --chain chain-a
hermes keys list --chain testnet
```

## Step 4: Create IBC Connection

```bash
# Create clients
hermes create client --host-chain chain-a --reference-chain testnet
hermes create client --host-chain testnet --reference-chain chain-a

# Create connection
hermes create connection --a-chain chain-a --b-chain testnet

# Create transfer channel
hermes create channel --order unordered --a-chain chain-a --a-connection connection-0 --a-port transfer --b-port transfer

# Start relayer
hermes start
```

## Step 5: Test IBC Transfer

```bash
# Transfer from chain-a to testnet
hermes tx ft-transfer \
    --dst-chain testnet \
    --src-chain chain-a \
    --src-port transfer \
    --src-channel channel-0 \
    --amount 1000000 \
    --denom stake

# Transfer from testnet to chain-a
hermes tx ft-transfer \
    --dst-chain chain-a \
    --src-chain testnet \
    --src-port transfer \
    --src-channel channel-0 \
    --amount 1000000 \
    --denom usix
```

## Step 6: Query and Monitor

```bash
# Query channels
hermes query channels --chain chain-a
hermes query channels --chain testnet

# Query connections
hermes query connections --chain chain-a
hermes query connections --chain testnet

# Check client status
hermes query client status --chain chain-a --client 07-tendermint-0
hermes query client status --chain testnet --client 07-tendermint-0

# Monitor packets
hermes query packet acks --chain chain-a --port transfer --channel channel-0
```

## Notes

1. **Port Configuration**: Make sure SIX Protocol testnet uses different ports than Gaia (26657/9090 vs 26658/9091 if running locally)

2. **Chain IDs**: 
   - Gaia local: `chain-a`
   - SIX Protocol: `testnet`

3. **Account Prefixes**:
   - Gaia: `cosmos`
   - SIX Protocol: `6x`

4. **Native Tokens**:
   - Gaia: `stake`
   - SIX Protocol: `usix`

5. **Funding**: Ensure both relayer accounts have sufficient funds for gas fees

This setup will create an IBC connection between your local Gaia chain-a and the SIX Protocol testnet.
