#!/bin/sh
#set -o errexit -o nounset -o pipefail

PASSWORD=${PASSWORD:-1234567890}
STAKE=${STAKE_TOKEN:-usix}
FEE=${FEE_TOKEN:-usix}
CHAIN_ID=${CHAIN_ID:-testnet}
MONIKER=${MONIKER:-node001}
KEYRING="--keyring-backend test"
BLOCK_GAS_LIMIT=${GAS_LIMIT:-10000000} # should mirror mainnet

echo "Configured Block Gas Limit: $BLOCK_GAS_LIMIT"

# check the genesis file
GENESIS_FILE="$HOME"/.six/config/genesis.json
if [ -f "$GENESIS_FILE" ]; then
  echo "$GENESIS_FILE exists..."
else
  echo "$GENESIS_FILE does not exist. Generating..."

  sixd init --chain-id "$CHAIN_ID" "$MONIKER"
  sixd add-ica-config
  # staking/governance token is hardcoded in config, change this
  sed -i "s/\"stake\"/\"$STAKE\"/" "$GENESIS_FILE"
  # this is essential for sub-1s block times (or header times go crazy)
  sed -i 's/"time_iota_ms": "1000"/"time_iota_ms": "10"/' "$GENESIS_FILE"
  # change gas limit to mainnet value
  sed -i 's/"max_gas": "-1"/"max_gas": "'"$BLOCK_GAS_LIMIT"'"/' "$GENESIS_FILE"
  # change default keyring-backend to test
  sed -i 's/keyring-backend = "os"/keyring-backend = "test"/' "$HOME"/.six/config/client.toml
fi

APP_TOML_CONFIG="$HOME"/.six/config/app.toml
APP_TOML_CONFIG_NEW="$HOME"/.six/config/app_new.toml
CONFIG_TOML_CONFIG="$HOME"/.six/config/config.toml
if [ -n $UNSAFE_CORS ]; then
  echo "Unsafe CORS set... updating app.toml and config.toml"
  # sorry about this bit, but toml is rubbish for structural editing
  sed -n '1h;1!H;${g;s/# Enable defines if the API server should be enabled.\nenable = false/enable = true/;p;}' "$APP_TOML_CONFIG" > "$APP_TOML_CONFIG_NEW"
  mv "$APP_TOML_CONFIG_NEW" "$APP_TOML_CONFIG"
  # ...and breathe
  sed -i "s/enabled-unsafe-cors = false/enabled-unsafe-cors = true/" "$APP_TOML_CONFIG"
  sed -i "s/cors_allowed_origins = \[\]/cors_allowed_origins = \[\"\*\"\]/" "$CONFIG_TOML_CONFIG"
fi

# are we running for the first time?
if ! sixd keys show validator $KEYRING; then
  (echo "$PASSWORD"; echo "$PASSWORD") | sixd keys add validator $KEYRING

  # hardcode the validator account for this instance
  echo "$PASSWORD" | sixd add-genesis-account validator "1000000000$STAKE,1000000000$FEE" $KEYRING

  # (optionally) add a few more genesis accounts
  for addr in "$@"; do
    echo $addr
    sixd add-genesis-account "$addr" "1000000000$STAKE,1000000000$FEE"
  done

  # submit a genesis validator tx
  ## Workraround for https://github.com/cosmos/cosmos-sdk/issues/8251
  (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | sixd gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID" --amount="250000000$STAKE" $KEYRING
  ## should be:
  # (echo "$PASSWORD"; echo "$PASSWORD"; echo "$PASSWORD") | sixd gentx validator "250000000$STAKE" --chain-id="$CHAIN_ID"
  sixd collect-gentxs
fi
