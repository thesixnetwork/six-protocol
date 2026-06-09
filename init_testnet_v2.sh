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
VAL_MODE=$2
TRACE="" # Set to "--trace" for tracing
VALIDATOR_APPROVAL_ENABLED="${3:-false}" # New parameter for validator approval

if [ -z "$VAL_MODE" ]; then
  VAL_MODE=0
fi

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

# =====================================================
# MNEMONICS SECTION - From config.yml
# =====================================================

ALICE_MNEMONIC="history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
BOB_MNEMONIC="limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"
VAL1_MNEMONIC="note base stone list envelope tail start forget alarm acoustic cook occur divert giant bike curtain chase shuffle fade glow capital slot file provide"
VAL2_MNEMONIC="strike tower consider despair bridge diesel clay celery violin base hello ride they weather tunnel elite truth oblige spot hen wise flag pet battle"
VAL3_MNEMONIC="canvas human require month loan oak december blame grit palm slice error absorb total spice autumn trouble soda repeat shove quit bid forward organ"
VAL4_MNEMONIC="grant raw marine drink text dove flat waste wish buzz output hand merge cluster civil clog stay alert silent reunion idea cake village almost"

# =====================================================
# INITIAL BALANCES - From config.yml
# =====================================================

ALICE_BALANCE="1000000000000${STAKING_TOKEN}"
BOB_BALANCE="11000000000000${STAKING_TOKEN}"
SUPER_ADMIN_BALANCE="11000000000000${STAKING_TOKEN}"
VAL1_BALANCE="100000000000000${STAKING_TOKEN}"
VAL2_BALANCE="11000000000000${STAKING_TOKEN}"
VAL3_BALANCE="11000000000000${STAKING_TOKEN}"
VAL4_BALANCE="11000000000000${STAKING_TOKEN}"
FAUCET_BALANCE="1000000000000${STAKING_TOKEN}"

# Validator staking amount
VALIDATOR_STAKE="1000000000000${STAKING_TOKEN}"

# =====================================================
# FUNCTIONS
# =====================================================

print_section() {
  echo ""
  echo "=================================================="
  echo "$1"
  echo "=================================================="
}

# =====================================================
# MAIN SCRIPT
# =====================================================

print_section "Starting Six Protocol Testnet Initialization"
echo "Chain ID: $CHAINID"
echo "Moniker: $MONIKER"
echo "Validator Mode: $VAL_MODE"
echo "Validator Approval Enabled: $VALIDATOR_APPROVAL_ENABLED"

# Clean up old data
print_section "Cleaning up old data"
rm -rf $SIX_HOME
sixd tendermint unsafe-reset-all --home $SIX_HOME

# Initialize chain
print_section "Initializing chain"
sixd config set client chain-id $CHAINID --home $SIX_HOME
sixd config set client keyring-backend $KEYRING --home $SIX_HOME
sixd init $MONIKER --chain-id $CHAINID --home $SIX_HOME $TRACE

# =====================================================
# CREATE KEYS
# =====================================================

print_section "Creating keys from mnemonics"

echo "$ALICE_MNEMONIC" | sixd keys add alice --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$BOB_MNEMONIC" | sixd keys add bob --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$SUPER_ADMIN_MNEMONIC" | sixd keys add super-admin --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$VAL1_MNEMONIC" | sixd keys add val1 --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$VAL2_MNEMONIC" | sixd keys add val2 --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$VAL3_MNEMONIC" | sixd keys add val3 --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
echo "$VAL4_MNEMONIC" | sixd keys add val4 --recover --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME
sixd keys add faucet --keyring-backend $KEYRING --algo $KEYALGO --home $SIX_HOME

echo ""
echo "Key addresses:"
echo "Alice: $(sixd keys show alice -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo "Bob: $(sixd keys show bob -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo "Super Admin: $(sixd keys show super-admin -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo "Val1: $(sixd keys show val1 -a --keyring-backend $KEYRING --home $SIX_HOME)"

# =====================================================
# ADD GENESIS ACCOUNTS
# =====================================================

print_section "Adding genesis accounts"

sixd genesis add-genesis-account alice $ALICE_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account bob $BOB_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account super-admin $SUPER_ADMIN_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account val1 $VAL1_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account val2 $VAL2_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account val3 $VAL3_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account val4 $VAL4_BALANCE --keyring-backend $KEYRING --home $SIX_HOME
sixd genesis add-genesis-account faucet $FAUCET_BALANCE --keyring-backend $KEYRING --home $SIX_HOME

# =====================================================
# GENESIS CONFIGURATION
# =====================================================

print_section "Configuring genesis file"

GENESIS_FILE="$SIX_HOME/config/genesis.json"

# Update staking params
jq '.app_state.staking.params.bond_denom = "usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.staking.params.max_validators = 3' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.staking.params.unbonding_time = "300s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update validator approval configuration
jq --arg approver "$SUPER_ADMIN_ADDRESS" --arg enabled "$VALIDATOR_APPROVAL_ENABLED" \
  '.app_state.staking.validator_approval.approver_address = $approver | .app_state.staking.validator_approval.enabled = ($enabled == "true")' \
  $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update mint params
jq '.app_state.mint.params.mint_denom = "usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update crisis params
jq '.app_state.crisis.constant_fee.denom = "usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update gov params
jq '.app_state.gov.params.voting_period = "300s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.gov.params.expedited_voting_period = "180s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.gov.params.min_deposit[0].denom = "usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.gov.params.min_deposit[0].amount = "1000000"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.gov.params.expedited_min_deposit[0].denom = "usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.gov.params.expedited_min_deposit[0].amount = "50000000"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update EVM params
jq '.app_state.evm.params.evm_denom = "asix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.evm.params.allow_unprotected_txs = true' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update feemarket params
jq '.app_state.feemarket.params.base_fee = "100000000000"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.feemarket.params.min_gas_price = "100000000000.0"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.feemarket.params.min_gas_multiplier = "0.5"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update bank denom metadata
jq --arg super_admin "$SUPER_ADMIN_ADDRESS" \
  '.app_state.bank.denom_metadata = [
    {
      "description": "The native staking token of the SIX Protocol.",
      "denom_units": [
        {"denom": "usix", "exponent": 0, "aliases": ["microsix"]},
        {"denom": "msix", "exponent": 3, "aliases": ["millisix"]},
        {"denom": "six", "exponent": 6, "aliases": []}
      ],
      "base": "usix",
      "display": "six",
      "name": "Six token",
      "symbol": "six",
      "uri": "",
      "uri_hash": ""
    },
    {
      "description": "The native evm token of the SIX Protocol.",
      "denom_units": [
        {"denom": "asix", "exponent": 0, "aliases": ["attosix"]},
        {"denom": "usix", "exponent": 12, "aliases": ["microsix"]},
        {"denom": "msix", "exponent": 15, "aliases": ["millisix"]},
        {"denom": "six", "exponent": 18, "aliases": []}
      ],
      "base": "asix",
      "display": "six",
      "name": "eSix token",
      "symbol": "asix",
      "uri": "",
      "uri_hash": ""
    }
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update circuit permissions
jq --arg super_admin "$SUPER_ADMIN_ADDRESS" --arg alice "$ALICE_ADDRESS" \
  '.app_state.circuit.account_permissions = [
    {
      "address": $super_admin,
      "permissions": {
        "level": "LEVEL_SUPER_ADMIN",
        "limit_type_urls": []
      }
    },
    {
      "address": $alice,
      "permissions": {
        "level": "LEVEL_SUPER_ADMIN",
        "limit_type_urls": []
      }
    }
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update protocol admin
jq --arg super_admin "$SUPER_ADMIN_ADDRESS" \
  '.app_state.protocoladmin.adminList = [
    {"admin": $super_admin, "group": "super.admin"},
    {"admin": $super_admin, "group": "token.admin"}
  ] |
  .app_state.protocoladmin.groupList = [
    {"name": "super.admin", "owner": $super_admin},
    {"name": "token.admin", "owner": $super_admin}
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update nftadmin
jq --arg super_admin "$SUPER_ADMIN_ADDRESS" \
  '.app_state.nftadmin.authorization.root_admin = $super_admin |
  .app_state.nftadmin.authorization.permissions = [
    {
      "name": "nft_fee_admin",
      "addresses": [$super_admin]
    }
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update tokenmngr
jq --arg super_admin "$SUPER_ADMIN_ADDRESS" --arg alice "$ALICE_ADDRESS" \
  '.app_state.tokenmngr.options.defaultMintee = $super_admin |
  .app_state.tokenmngr.mintpermList = [
    {"address": $alice, "creator": $super_admin, "token": "usix"},
    {"address": $alice, "creator": $super_admin, "token": "asix"}
  ] |
  .app_state.tokenmngr.tokenList = [
    {
      "base": "usix",
      "creator": $super_admin,
      "maxSupply": {"amount": "0", "denom": "usix"},
      "mintee": $super_admin,
      "name": "usix"
    },
    {
      "base": "asix",
      "creator": $super_admin,
      "maxSupply": {"amount": "0", "denom": "asix"},
      "mintee": $super_admin,
      "name": "asix"
    }
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update nftoracle params
jq '.app_state.nftoracle.oracle_config.minimum_confirmation = 4' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.nftoracle.params.action_request_active_duration = "120s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.nftoracle.params.mint_request_active_duration = "120s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.nftoracle.params.verify_request_active_duration = "120s"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# Update nftmngr fee config
jq '.app_state.nftmngr.nft_fee_config.schema_fee.fee_amount = "5000000usix"' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
jq '.app_state.nftmngr.nft_fee_config.schema_fee.fee_distributions = [
    {"method": "BURN", "portion": 0.5},
    {"method": "REWARD_POOL", "portion": 0.5}
  ]' $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE

# =====================================================
# CREATE GENTX
# =====================================================

print_section "Creating genesis transaction for validator"

if [ "$VAL_MODE" -eq 1 ]; then
  VALIDATOR_KEY="bob"
else
  VALIDATOR_KEY="bob"
fi

# Create gentx with approver flag (when available)
if [ "$VALIDATOR_APPROVAL_ENABLED" = "true" ]; then
  echo "Creating gentx with approver address: $SUPER_ADMIN_ADDRESS"
  sixd genesis gentx $VALIDATOR_KEY $VALIDATOR_STAKE \
    --chain-id $CHAINID \
    --keyring-backend $KEYRING \
    --home $SIX_HOME \
    --moniker $MONIKER \
    --commission-rate 0.1 \
    --commission-max-rate 0.2 \
    --commission-max-change-rate 0.01 \
    --min-self-delegation 1 \
    --approver $SUPER_ADMIN_ADDRESS
else
  echo "Creating gentx without approver (approval disabled)"
  sixd genesis gentx $VALIDATOR_KEY $VALIDATOR_STAKE \
    --chain-id $CHAINID \
    --keyring-backend $KEYRING \
    --home $SIX_HOME \
    --moniker $MONIKER \
    --commission-rate 0.1 \
    --commission-max-rate 0.2 \
    --commission-max-change-rate 0.01 \
    --min-self-delegation 1
fi

# =====================================================
# VERIFY/FIX GENTX APPROVER ADDRESS (if needed)
# =====================================================

if [ "$VALIDATOR_APPROVAL_ENABLED" = "true" ]; then
  print_section "Verifying gentx approver_address"
  
  # The gentx file is created in $SIX_HOME/config/gentx/
  GENTX_FILE=$(ls $SIX_HOME/config/gentx/*.json | head -n 1)
  
  if [ -f "$GENTX_FILE" ]; then
    echo "Found gentx file: $GENTX_FILE"
    
    # Check if approver_address is already set
    CURRENT_APPROVER=$(jq -r '.body.messages[0].approver_address' $GENTX_FILE)
    
    if [ "$CURRENT_APPROVER" = "$SUPER_ADMIN_ADDRESS" ]; then
      echo "✅ Approver address already correct: $SUPER_ADMIN_ADDRESS"
    elif [ -z "$CURRENT_APPROVER" ] || [ "$CURRENT_APPROVER" = "null" ] || [ "$CURRENT_APPROVER" = "" ]; then
      echo "⚠️  Approver address empty, patching gentx file..."
      jq --arg approver "$SUPER_ADMIN_ADDRESS" \
        '(.body.messages[0].approver_address) = $approver' \
        $GENTX_FILE > tmp.json && mv tmp.json $GENTX_FILE
      echo "✅ Updated approver_address in gentx to: $SUPER_ADMIN_ADDRESS"
    else
      echo "⚠️  Found different approver: $CURRENT_APPROVER, updating to: $SUPER_ADMIN_ADDRESS"
      jq --arg approver "$SUPER_ADMIN_ADDRESS" \
        '(.body.messages[0].approver_address) = $approver' \
        $GENTX_FILE > tmp.json && mv tmp.json $GENTX_FILE
      echo "✅ Updated approver_address in gentx"
    fi
  fi
else
  echo "Validator approval disabled, skipping approver verification"
fi

# Collect genesis transactions
sixd genesis collect-gentxs --home $SIX_HOME

# =====================================================
# VERIFY/FIX GENTX IN GENESIS.JSON (if needed)
# =====================================================

if [ "$VALIDATOR_APPROVAL_ENABLED" = "true" ]; then
  print_section "Verifying gentx approver_address in final genesis.json"
  
  # After collect-gentxs, the gentx is now in genesis.json
  # Check if approver_address is correct
  GENESIS_APPROVER=$(jq -r '.app_state.genutil.gen_txs[0].body.messages[0].approver_address' $GENESIS_FILE)
  
  if [ "$GENESIS_APPROVER" = "$SUPER_ADMIN_ADDRESS" ]; then
    echo "✅ Genesis approver address already correct: $SUPER_ADMIN_ADDRESS"
  elif [ -z "$GENESIS_APPROVER" ] || [ "$GENESIS_APPROVER" = "null" ] || [ "$GENESIS_APPROVER" = "" ]; then
    echo "⚠️  Genesis approver address empty, patching genesis.json..."
    jq --arg approver "$SUPER_ADMIN_ADDRESS" \
      '(.app_state.genutil.gen_txs[0].body.messages[0].approver_address) = $approver' \
      $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
    echo "✅ Updated approver_address in genesis.json gentx to: $SUPER_ADMIN_ADDRESS"
  else
    echo "⚠️  Found different approver in genesis: $GENESIS_APPROVER, updating to: $SUPER_ADMIN_ADDRESS"
    jq --arg approver "$SUPER_ADMIN_ADDRESS" \
      '(.app_state.genutil.gen_txs[0].body.messages[0].approver_address) = $approver' \
      $GENESIS_FILE > tmp.json && mv tmp.json $GENESIS_FILE
    echo "✅ Updated approver_address in genesis.json"
  fi
else
  echo "Validator approval disabled, skipping genesis approver verification"
fi

# =====================================================
# CONFIGURATION FILES
# =====================================================

print_section "Configuring node settings"

# Update config.toml
CONFIG_FILE="$SIX_HOME/config/config.toml"
sed -i.bak "s/^moniker = .*/moniker = \"$MONIKER\"/" $CONFIG_FILE
sed -i.bak "s/^log_level = .*/log_level = \"$LOGLEVEL\"/" $CONFIG_FILE

# Update app.toml
APP_FILE="$SIX_HOME/config/app.toml"
sed -i.bak "s/^minimum-gas-prices = .*/minimum-gas-prices = \"1.25usix\"/" $APP_FILE

# Enable API
sed -i.bak '/\[api\]/,/\[/ s/^enable = .*/enable = true/' $APP_FILE
sed -i.bak '/\[api\]/,/\[/ s/^swagger = .*/swagger = true/' $APP_FILE

# Enable EVM RPC
sed -i.bak '/\[json-rpc\]/,/\[/ s/^enable = .*/enable = true/' $APP_FILE
sed -i.bak '/\[json-rpc\]/,/\[/ s/^address = .*/address = "0.0.0.0:8545"/' $APP_FILE
sed -i.bak '/\[json-rpc\]/,/\[/ s/^ws-address = .*/ws-address = "0.0.0.0:8546"/' $APP_FILE

# =====================================================
# VALIDATION
# =====================================================

print_section "Validating genesis file"

sixd genesis validate-genesis --home $SIX_HOME

# =====================================================
# COMPLETION
# =====================================================

print_section "Initialization Complete!"

echo ""
echo "Chain initialized with the following configuration:"
echo "  Chain ID: $CHAINID"
echo "  Moniker: $MONIKER"
echo "  Home: $SIX_HOME"
echo "  Validator Approval Enabled: $VALIDATOR_APPROVAL_ENABLED"
echo "  Approver Address: $SUPER_ADMIN_ADDRESS"
echo ""
echo "Key accounts:"
echo "  Alice: $(sixd keys show alice -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo "  Bob: $(sixd keys show bob -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo "  Super Admin: $(sixd keys show super-admin -a --keyring-backend $KEYRING --home $SIX_HOME)"
echo ""
echo "To start the chain, run:"
echo "  sixd start --home $SIX_HOME"
echo ""
echo "Or with trace enabled:"
echo "  sixd start --home $SIX_HOME --trace"
