MONIKER=$1
# if empty, default to "testnet"
if [ -z "$MONIKER" ]; then
  MONIKER="mynode"
fi
CHAIN_ID=six_666-1
VALKEY=val1 # should be: export as docker env var
SIX_HOME=~/.six
ALICE_MNEMONIC="history perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
BOB_MNEMONIC="limb sister humor wisdom elephant weasel beyond must any desert glance stem reform soccer include chest chef clerk call popular display nerve priority venture"
VAL1_MNEMONIC="note base stone list envelope tail start forget alarm acoustic cook occur divert giant bike curtain chase shuffle fade glow capital slot file provide"
VAL2_MNEMONIC="strike tower consider despair bridge diesel clay celery violin base hello ride they weather tunnel elite truth oblige spot hen wise flag pet battle"
VAL3_MNEMONIC="canvas human require month loan oak december blame grit palm slice error absorb total spice autumn trouble soda repeat shove quit bid forward organ"
VAL4_MNEMONIC="grant raw marine drink text dove flat waste wish buzz output hand merge cluster civil clog stay alert silent reunion idea cake village almost"
ORACLE1_MNEMONIC="list split future remain scene cheap pledge forum siren purse bright ivory split morning swing dumb fabric rapid remove worth diary task island donkey"
ORACLE2_MNEMONIC="achieve rice anger junk delay glove slam find poem feed emerge next core twice kitchen road proof remain notice slice walk super piece father"
ORACLE3_MNEMONIC="hint expose mix lemon leave genuine host fiction peasant daughter enable region mixture bean soda auction armed turtle iron become bracket wasp drama front"
ORACLE4_MNEMONIC="clown cabbage clean design mosquito surround citizen virus kite castle sponsor wife lesson coffee alien panel hand together good crazy fabric mouse hat town"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"
KEY="mykey"
KEY2="mykey2"
CHAINID="six_666-1"
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Reinstall daemon
rm -rf ~/.six*
make install

# Set client config
sixd config keyring-backend $KEYRING
sixd config chain-id $CHAINID

# if $KEY exists it should be deleted
# mint to validator
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ALICE_MNEMONIC | sixd keys add alice --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $BOB_MNEMONIC | sixd keys add bob --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL1_MNEMONIC | sixd keys add val1 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL2_MNEMONIC | sixd keys add val2 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL3_MNEMONIC | sixd keys add val3 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $VAL4_MNEMONIC | sixd keys add val4 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ORACLE1_MNEMONIC | sixd keys add oracle1 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ORACLE2_MNEMONIC | sixd keys add oracle2 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ORACLE3_MNEMONIC | sixd keys add oracle3 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}
echo $ORACLE4_MNEMONIC | sixd keys add oracle4 --recover --home ${SIX_HOME} --keyring-backend ${KEYRING} --algo ${KEYALGO}

# Set moniker and chain-id for six (Moniker can be anything, chain-id must be an integer)
sixd init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to usix
cat $HOME/.six/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="usix"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json
cat $HOME/.six/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="usix"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json
cat $HOME/.six/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="usix"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json
cat $HOME/.six/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="usix"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json
# cat $HOME/.six/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="usix"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json

# Change voting params so that submitted proposals pass immediately for testing
cat $HOME/.six/config/genesis.json| jq '.app_state.gov.voting_params.voting_period="600s"' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json


# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.six/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.six/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.six/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.six/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.six/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.six/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses)
sixd add-genesis-account $(sixd keys show -a val1 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val2 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val3 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val4 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle1 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle2 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle3 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle4 --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a alice --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a bob --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a super-admin --keyring-backend ${KEYRING} --home ${SIX_HOME}) 100000000000000000000000000usix --keyring-backend ${KEYRING} --home ${SIX_HOME}
                                 
# Update total supply with claim values
#validators_supply=$(cat $HOME/.six/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=1100000000000000000000000000
cat $HOME/.six/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.six/config/tmp_genesis.json && mv $HOME/.six/config/tmp_genesis.json $HOME/.six/config/genesis.json

echo $KEYRING
echo $KEY
# Sign genesis transaction
sixd gentx val1 100000000000000000000000usix --keyring-backend $KEYRING --chain-id $CHAINID
#sixd gentx $KEY2 1000000000000000000000usix --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
sixd collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
sixd validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
sixd start --pruning=nothing --trace --log_level info --minimum-gas-prices=0.0001usix --json-rpc.api eth,txpool,personal,net,debug,web3 --rpc.laddr "tcp://0.0.0.0:26657" --api.enable true

