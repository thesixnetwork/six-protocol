export CHAIN_ID=six
export MONIKER=mynode
export VALKEY=validator1
export ORCKEY=orch1
export SIX_HOME=~/.six_test

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

sixd keys add ${VALKEY} --keyring-backend test --home ${SIX_HOME}
sixd keys add ${ORCKEY} --keyring-backend test --home ${SIX_HOME}
sixd keys add super-admin --keyring-backend test --home ${SIX_HOME}
# replace address of val and orc here
export VAL_ADDRESS="6x1fjy5cfjp2pqqt430lexalwtm872jjlfjy9qgzt" 
export ORC_ADDRESS="6x1gvdc9zgc9m9ap5hgs2w7g4mcdsun93qzt84a2z"
export SUPERADMIN_ADDRESS="6x1l0ceauyrkuhte463halxz8tawrlsv5vxc3jxer"

sixd eth_keys add --keyring-backend test --home ${SIX_HOME}
export ETH_ADDRESS="0xb1561B494fC1E99cd2353f773Baa7e7907a93d44"

sixd add-genesis-account ${VALKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account ${ORCKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account super-admin 1000000000000stake --keyring-backend test --home ${SIX_HOME}

# modify nativeHRP
# code ${SIX_HOME}/config/genesis.json
jq '.app_state.bech32ibc.nativeHRP = "6x"' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json

# replave nativeHRP with 6x
# add grouplist to genesis.json
# add super.admin to grouplist
jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin", "owner": "genesis"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json
jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "'"$SUPERADMIN_ADDRESS"'", "group": "super.admin"}' ${SIX_HOME}/config/genesis.json | sponge ${SIX_HOME}/config/genesis.json

sixd gengate --moniker=${MONIKER} ${VALKEY} 1000000000stake \
    ${ETH_ADDRESS} ${VAL_ADDRESS} --chain-id=${CHAIN_ID} \
    --keyring-backend test --home ${SIX_HOME}
sixd collect-gengate --home ${SIX_HOME}


 # init protocoladmin
sixd tx protocoladmin add-admin-to-group token.admin ${ADDRESS} --from ${VAL_ADDRESS} --home ${SIX_HOME}--chain-id ${CHAIN_ID}


# backup
cp -r ${SIX_HOME} ${SIX_HOME}_backup

# start
sixd start --home ${SIX_HOME}


# other cmds
sixd keys list --keyring-backend test --home ${SIX_HOME}
# restore
rm -Rf ${SIX_HOME}
cp -r ${SIX_HOME}_backup ${SIX_HOME}

#  notes

# private: 0xec87c71a921ee242a0a95cc0d5eb34806c63cb8b095d4912e5fa740d7a8c61aa 
# public: 0x04717ad160cccbfabb16b438d0bfa075b75bc99a8ea2e43bfe1ad74bc9106d541ed2f37f518b5badf9f632da17a7913faccdf550cd9f55be79b580e315ca7ee426 
# address: 0x837f7E4DcCB7AEd3807e51f0535E7Fa2718bc0E7