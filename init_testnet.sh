#!/bin/bash
set -e # Exit on error

# =====================================================
# CONFIGURATION SECTION - Easy to modify parameters
# =====================================================

# Chain configuration
CHAINID="testnet"
MONIKER="${1:-mynode}"
KEYRING="test"
KEYALGO="secp256k1"
SIX_HOME=~/.six
LOGLEVEL="info"
TRACE="" # Set to "--trace" for tracing

# Token denominations
STAKING_TOKEN="usix"
EVM_TOKEN="asix"

# =====================================================
# KEY ADDRESS MAPPING - Important for matching config.yml
# =====================================================

# These are the key addresses from the working genesis
ALICE_ADDRESS="6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq"
BOB_ADDRESS="6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp"
SUPER_ADMIN_ADDRESS="6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"
SPECIAL_EVM_ADDRESS="6x18743s33zmsvmvyynfxu5sy2f80e2g5mz8dk65g"

# =====================================================
# MNEMONICS SECTION - From config.yml only
# =====================================================

# Mnemonics from config.yml
ALICE_MNEMONIC="history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
BOB_MNEMONIC="limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"

# =====================================================
# VALIDATION SECTION
# =====================================================
echo "Starting initialization of $CHAINID testnet with validator bob..."

# Validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# =====================================================
# SETUP SECTION
# =====================================================
echo "Setting up environment..."

# Reinstall daemon
rm -rf ${SIX_HOME}
rm go.sum && touch go.sum
go mod tidy
make install

# Set client config
sixd config keyring-backend $KEYRING --home ${SIX_HOME}
sixd config chain-id $CHAINID --home ${SIX_HOME}

# =====================================================
# KEY MANAGEMENT SECTION
# =====================================================
echo "Importing keys from config.yml..."

# Import keys from config.yml only
echo $ALICE_MNEMONIC | sixd keys add alice --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $BOB_MNEMONIC | sixd keys add bob --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}

# =====================================================
# CHAIN INITIALIZATION SECTION
# =====================================================
echo "Initializing chain with moniker: $MONIKER and chain-id: $CHAINID"
sixd init $MONIKER --chain-id $CHAINID --home ${SIX_HOME}

# =====================================================
# GENESIS CONFIGURATION SECTION
# =====================================================
echo "Configuring genesis..."

# Function to update genesis using jq
update_genesis() {
    cat ${SIX_HOME}/config/genesis.json | jq "$1" > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
}

# Change parameter token denominations from stake to usix
update_genesis '.app_state["staking"]["params"]["bond_denom"]="'$STAKING_TOKEN'"'
update_genesis '.app_state["crisis"]["constant_fee"]["denom"]="'$STAKING_TOKEN'"'
update_genesis '.app_state["crisis"]["constant_fee"]["amount"]="1000"'
update_genesis '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="'$STAKING_TOKEN'"'
update_genesis '.app_state["gov"]["deposit_params"]["min_deposit"][0]["amount"]="1000000"'
update_genesis '.app_state["evm"]["params"]["evm_denom"]="'$EVM_TOKEN'"'
update_genesis '.app_state["evm"]["params"]["allow_unprotected_txs"]=true'  # Match working genesis
update_genesis '.app_state["inflation"]["params"]["mint_denom"]="'$STAKING_TOKEN'"'
update_genesis '.app_state["mint"]["params"]["mint_denom"]="'$STAKING_TOKEN'"'

# Bank configuration from config.yml
update_genesis '.app_state.bank.params.send_enabled = []'
update_genesis '.app_state.bank.params.default_send_enabled = true'

# Token metadata - identical to working genesis
update_genesis '.app_state.bank.denom_metadata[0] = {
  "description": "The native staking token of the SIX Protocol.",
  "denom_units": [
    {"denom": "usix","exponent": 0,"aliases": ["microsix"]},
    {"denom": "msix","exponent": 3,"aliases": ["millisix"]},
    {"denom": "six","exponent": 6,"aliases": []}
  ],
  "base": "usix",
  "display": "six",
  "name": "Six token",
  "symbol": "six"
}'

update_genesis '.app_state.bank.denom_metadata[1] = {
  "description": "The native evm token of the SIX Protocol.",
  "denom_units": [
    {"denom": "asix","exponent": 0,"aliases": ["attosix"]},
    {"denom": "usix","exponent": 12,"aliases": ["microsix"]},
    {"denom": "msix","exponent": 15,"aliases": ["millisix"]},
    {"denom": "six","exponent": 18,"aliases": []}
  ],
  "base": "asix",
  "display": "six",  
  "name": "eSix token",
  "symbol": "asix"
}'

# NFT Admin configuration
update_genesis '.app_state.nftadmin.authorization = {
  "root_admin": "'$SUPER_ADMIN_ADDRESS'",
  "permissions": {
    "permissions": [
      {
        "name": "nft_fee_admin",
        "addresses": {
          "addresses": ["'$SUPER_ADMIN_ADDRESS'"]
        }
      }
    ]
  }
}'

# NFT Manager configuration
update_genesis '.app_state.nftmngr.nft_fee_config = {
  "schema_fee": {
    "fee_amount": "200000000usix",
    "fee_distributions": [
      {"method": "BURN", "portion": 0.5},
      {"method": "REWARD_POOL", "portion": 0.5}
    ]
  }
}'

# NFT Oracle configuration
update_genesis '.app_state.nftoracle.params = {
  "action_request_active_duration": "120s",
  "mint_request_active_duration": "120s",
  "verify_request_active_duration": "120s", 
  "action_signer_active_duration": "2592000s",
  "sync_action_signer_active_duration": "300s"
}'

update_genesis '.app_state.nftoracle.oracle_config = {"minimum_confirmation": 4}'

# Protocol admin configuration
update_genesis '.app_state.protocoladmin.adminList[0] |= . + {
  "admin": "'$SUPER_ADMIN_ADDRESS'",
  "group": "super.admin"
}'

update_genesis '.app_state.protocoladmin.adminList[1] |= . + {
  "admin": "'$SUPER_ADMIN_ADDRESS'",
  "group": "token.admin"
}'

update_genesis '.app_state.protocoladmin.groupList[0] |= . + {
  "name": "super.admin",
  "owner": "'$SUPER_ADMIN_ADDRESS'"
}'

update_genesis '.app_state.protocoladmin.groupList[1] |= . + {
  "name": "token.admin",
  "owner": "'$SUPER_ADMIN_ADDRESS'"
}'

# Validator approval configuration
update_genesis '.app_state.staking.validator_approval = {
  "approver_address": "'$SUPER_ADMIN_ADDRESS'",
  "enabled": false
}'

update_genesis '.app_state.staking.params.max_validators = 3'
update_genesis '.app_state.staking.params.unbonding_time = "300s"'

# Token manager configuration
update_genesis '.app_state.tokenmngr.mintpermList[0] |= . + {
  "address": "'$ALICE_ADDRESS'",
  "creator": "'$SUPER_ADMIN_ADDRESS'",
  "token": "usix"
}'

update_genesis '.app_state.tokenmngr.mintpermList[1] |= . + {
  "address": "'$ALICE_ADDRESS'",
  "creator": "'$SUPER_ADMIN_ADDRESS'",
  "token": "asix"
}'

update_genesis '.app_state.tokenmngr.options = {
  "defaultMintee": "'$SUPER_ADMIN_ADDRESS'"
}'

update_genesis '.app_state.tokenmngr.tokenList[0] |= . + {
  "base": "usix",
  "creator": "'$SUPER_ADMIN_ADDRESS'",
  "maxSupply": { "amount": "0", "denom": "usix" },
  "mintee": "'$SUPER_ADMIN_ADDRESS'",
  "name": "usix"
}'

update_genesis '.app_state.tokenmngr.tokenList[1] |= . + {
  "base": "asix",
  "creator": "'$SUPER_ADMIN_ADDRESS'",
  "maxSupply": { "amount": "0", "denom": "asix" },
  "mintee": "'$SUPER_ADMIN_ADDRESS'",
  "name": "asix"
}'

# Governance configuration - Match working genesis
update_genesis '.app_state.gov.deposit_params.max_deposit_period = "172800s"'
update_genesis '.app_state.gov.voting_params.voting_period = "300s"'

# Feemarket configuration - Match working genesis exactly
update_genesis '.app_state.feemarket.params = {
  "base_fee": "5000000000000",
  "base_fee_change_denominator": 8,
  "elasticity_multiplier": 4,
  "enable_height": "0",
  "min_gas_multiplier": "0.5",
  "min_gas_price": "5000000000000.0",
  "no_base_fee": false,
  "legacy_base_fee": "20000000000",
  "legacy_min_gas_price": "20000000000.0"
}'

# =====================================================
# PLATFORM SPECIFIC CONFIGURATIONS
# =====================================================
if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "Configuring for macOS..."
    sed -i '' 's/stake/'$STAKING_TOKEN'/g' ${SIX_HOME}/config/genesis.json
else
    echo "Configuring for Linux..."
    sed -i 's/stake/'$STAKING_TOKEN'/g' ${SIX_HOME}/config/genesis.json
fi

# =====================================================
# CONFIG.TOML CONFIGURATION - Match ignite settings
# =====================================================
echo "Configuring config.toml to match ignite settings..."

if [[ "$OSTYPE" == "darwin"* ]]; then
    # RPC settings
    sed -i '' 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*",\]/g' ${SIX_HOME}/config/config.toml
    # Consensus settings
    sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "10s"/g' ${SIX_HOME}/config/config.toml
    sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' ${SIX_HOME}/config/config.toml
    # Critical: match the timeout settings from ignite
    sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "2m30s"/g' ${SIX_HOME}/config/config.toml
else
    # RPC settings
    sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["*",\]/g' ${SIX_HOME}/config/config.toml
    # Consensus settings
    sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "500ms"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "10s"/g' ${SIX_HOME}/config/config.toml
    sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' ${SIX_HOME}/config/config.toml
    # Critical: match the timeout settings from ignite
    sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "2m30s"/g' ${SIX_HOME}/config/config.toml
fi

# =====================================================
# APP.TOML CONFIGURATION - Match ignite settings
# =====================================================
echo "Configuring app.toml to match ignite settings..."

if [[ "$OSTYPE" == "darwin"* ]]; then
    # Minimum gas prices - CRITICAL
    sed -i '' 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "0usix"/g' ${SIX_HOME}/config/app.toml

    # API configuration
    sed -i '' 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/swagger = false/swagger = true/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/address = "tcp:\/\/localhost:1317"/address = "tcp:\/\/0.0.0.0:1317"/g' ${SIX_HOME}/config/app.toml
    
    # gRPC configuration
    sed -i '' 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/address = "0.0.0.0:9090"/address = "0.0.0.0:9090"/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/address = "0.0.0.0:9091"/address = "0.0.0.0:9091"/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    
    # JSON-RPC configuration - CRITICAL for immediate startup
    sed -i '' 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3"/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/address = "127.0.0.1:8545"/address = "0.0.0.0:8545"/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/ws-address = "127.0.0.1:8546"/ws-address = "0.0.0.0:8546"/g' ${SIX_HOME}/config/app.toml
    sed -i '' 's/allow-unprotected-txs = false/allow-unprotected-txs = true/g' ${SIX_HOME}/config/app.toml
else
    # Minimum gas prices - CRITICAL
    sed -i 's/minimum-gas-prices = "0stake"/minimum-gas-prices = "0usix"/g' ${SIX_HOME}/config/app.toml
    
    # API configuration
    sed -i 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i 's/swagger = false/swagger = true/g' ${SIX_HOME}/config/app.toml
    sed -i 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' ${SIX_HOME}/config/app.toml
    sed -i 's/address = "tcp:\/\/localhost:1317"/address = "tcp:\/\/0.0.0.0:1317"/g' ${SIX_HOME}/config/app.toml
    
    # gRPC configuration
    sed -i 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i 's/address = "0.0.0.0:9090"/address = "0.0.0.0:9090"/g' ${SIX_HOME}/config/app.toml
    sed -i 's/address = "0.0.0.0:9091"/address = "0.0.0.0:9091"/g' ${SIX_HOME}/config/app.toml
    sed -i 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    
    # JSON-RPC configuration - CRITICAL for immediate startup
    sed -i 's/enable = false/enable = true/g' ${SIX_HOME}/config/app.toml
    sed -i 's/api = "eth,net,web3"/api = "eth,txpool,personal,net,debug,web3"/g' ${SIX_HOME}/config/app.toml
    sed -i 's/address = "127.0.0.1:8545"/address = "0.0.0.0:8545"/g' ${SIX_HOME}/config/app.toml
    sed -i 's/ws-address = "127.0.0.1:8546"/ws-address = "0.0.0.0:8546"/g' ${SIX_HOME}/config/app.toml
    sed -i 's/allow-unprotected-txs = false/allow-unprotected-txs = true/g' ${SIX_HOME}/config/app.toml
fi

# Add EVM-RPC section if it doesn't exist (this is in the ignite config)
if ! grep -q "\[evm-rpc\]" ${SIX_HOME}/config/app.toml; then
    echo -e "\n[evm-rpc]\naddress = \"0.0.0.0:8545\"\nws-address = \"0.0.0.0:8546\"" >> ${SIX_HOME}/config/app.toml
fi

# =====================================================
# ACCOUNT ALLOCATION SECTION
# =====================================================
echo "Allocating genesis accounts..."

# Function to add genesis accounts
add_genesis_account() {
    local address=$1
    local amount=$2
    sixd add-genesis-account $address $amount --home ${SIX_HOME}
}

add_genesis_account "$ALICE_ADDRESS" "100000000000000${STAKING_TOKEN}"
add_genesis_account "$BOB_ADDRESS" "200000000000000${STAKING_TOKEN}"
add_genesis_account "$SUPER_ADMIN_ADDRESS" "300000000000000${STAKING_TOKEN}"
add_genesis_account "$SPECIAL_EVM_ADDRESS" "10000000000000000000${EVM_TOKEN}"

# =====================================================
# GENTX SECTION
# =====================================================
echo "Creating and collecting gentxs with bob as validator..."

sixd gentx bob 1000000000000${STAKING_TOKEN} --chain-id ${CHAINID} --home ${SIX_HOME} --keyring-backend ${KEYRING}

# Collect genesis txs
sixd collect-gentxs --home ${SIX_HOME}

# =====================================================
# VALIDATION AND START SECTION
# =====================================================
echo "Validating genesis..."
sixd validate-genesis --home ${SIX_HOME}

echo "====================================================="
echo "Initialization complete!"
echo "Chain ID: $CHAINID"
echo "Moniker: $MONIKER"
echo "Home: $SIX_HOME"
echo ""
echo "To start your node, run:"
echo "sixd start --json-rpc.enable --json-rpc.api eth,txpool,personal,net,debug,web3 --json-rpc.address 0.0.0.0:8545 --json-rpc.ws-address 0.0.0.0:8546 --api.enable --rpc.laddr tcp://0.0.0.0:26657 --home ${SIX_HOME}"
echo "====================================================="

# Start the node if requested
if [ "$2" == "start" ]; then
    echo "Starting node..."
    sixd start --minimum-gas-prices=1.25usix,1250000000000asix --json-rpc.api eth,txpool,personal,net,debug,web3 --rpc.laddr tcp://0.0.0.0:26657 --api.enable true --log_level info --json-rpc.allow-unprotected-txs true --home ${SIX_HOME}
fi