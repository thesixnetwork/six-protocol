export CHAIN_ID=testnet
export MONIKER=sixnode1 ## should be: export as docker env var
export VALKEY=validator
export SIX_HOME=~/.six_docker
VALIDATOR_MNEMONIC="${MONIKER} perfect across group seek acoustic delay captain sauce audit carpet tattoo exhaust green there giant cluster want pond bulk close screen scissors remind"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

sixd keys add ${VALKEY} --keyring-backend test --home ${SIX_HOME}
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --algo secp256k1 --recover --home ${SIX_HOME} --keyring-backend test
echo $VALIDATOR_MNEMONIC | sixd keys add validator --algo secp256k1 --recover --home ${SIX_HOME} --keyring-backend test

sixd add-genesis-account ${sixd keys show -a validator --keyring-backend=test --home ${SIX_HOME}} 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account ${sixd keys show -a super-admin --keyring-backend=test --home ${SIX_HOME}} 1000000000000stake --keyring-backend test --home ${SIX_HOME}

# add super.admin to grouplist
jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin", "owner": "genesis"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json
jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "'"$SUPERADMIN_ADDRESS"'", "group": "super.admin"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json

sixd gentx validator 100000000stake --chain-id=six --keyring-backend=test --home ${SIX_HOME}
sixd collect-gentxs --home ${SIX_HOME}

# start
sixd start --home ${SIX_HOME}

# address: 0x837f7E4DcCB7AEd3807e51f0535E7Fa2718bc0E7