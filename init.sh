export CHAIN_ID=six
export MONIKER=sixnode4 ## should be: export as docker env var
export VALKEY=val1 # should be: export as docker env var
export SIX_HOME=./build/six_validator5
VAL0_MNEMONIC="note base stone list envelope tail start forget alarm acoustic cook occur divert giant bike curtain chase shuffle fade glow capital slot file provide"
VAL1_MNEMONIC="strike tower consider despair bridge diesel clay celery violin base hello ride they weather tunnel elite truth oblige spot hen wise flag pet battle"
VAL2_MNEMONIC="canvas human require month loan oak december blame grit palm slice error absorb total spice autumn trouble soda repeat shove quit bid forward organ"
VAL3_MNEMONIC="grant raw marine drink text dove flat waste wish buzz output hand merge cluster civil clog stay alert silent reunion idea cake village almost"
SUPER_ADMIN_MNEMONIC="expect peace defense conduct virtual flight flip unit equip solve broccoli protect shed group else useless tree such tornado minimum decade tower warfare galaxy"

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

# mint to validator
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend test
echo $VAL0_MNEMONIC | sixd keys add val1 --recover --home ${SIX_HOME} --keyring-backend test
echo $VAL1_MNEMONIC | sixd keys add val2 --recover --home ${SIX_HOME} --keyring-backend test
echo $VAL2_MNEMONIC | sixd keys add val3 --recover --home ${SIX_HOME} --keyring-backend test
echo $VAL3_MNEMONIC | sixd keys add val4 --recover --home ${SIX_HOME} --keyring-backend test

sixd add-genesis-account $(sixd keys show -a val1 --keyring-backend=test --home ${SIX_HOME}) 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val2 --keyring-backend=test --home ${SIX_HOME}) 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val3 --keyring-backend=test --home ${SIX_HOME}) 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val4 --keyring-backend=test --home ${SIX_HOME}) 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a super-admin --keyring-backend=test --home ${SIX_HOME}) 1000000000000stake --keyring-backend test --home ${SIX_HOME}

# add super.admin to grouplist
jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin", "owner": "'`$SUPERADMIN_ADDRESS`'"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json
jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "'"$SUPERADMIN_ADDRESS"'", "group": "super.admin"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json

sixd gentx val1 100000000stake --chain-id=six --keyring-backend=test --home ${SIX_HOME}
# sixd collect-gentxs --home ${SIX_HOME}