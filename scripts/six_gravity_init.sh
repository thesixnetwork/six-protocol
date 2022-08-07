export CHAIN_ID=six
export MONIKER=deenode
export VALKEY=validator1
export ORCKEY=orch1
export SIX_HOME=~/.six_test

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

sixd keys add ${VALKEY} --keyring-backend test --home ${SIX_HOME}
sixd keys add ${ORCKEY} --keyring-backend test --home ${SIX_HOME}
export VAL_ADDRESS="6x1fdts53zq5xtnmmap3a8enffjxzcuvv2tddldds"
export ORC_ADDRESS="6x14kee3xxg6v88akhyu3ha3dwhctqm6ze4kkys9m"

sixd eth_keys add --keyring-backend test --home ${SIX_HOME}
export ETH_ADDRESS="0xD224824bBE868095132ee2d3A50aE770D0DFbb8c"

sixd add-genesis-account ${VALKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account ${ORCKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}

# modify nativeHRP
code ${SIX_HOME}/config/genesis.json
# jq '.nativeHRP = "six"' ${SIX_HOME}/config/genesis.json > ${SIX_HOME}/config/genesis.json.tmp && mv ${SIX_HOME}/config/genesis.json.tmp ${SIX_HOME}/config/genesis.json
sixd gengate --moniker=${MONIKER} ${VALKEY} 1000000000stake \
    ${ETH_ADDRESS} ${VAL_ADDRESS} --chain-id=${CHAIN_ID} \
    --keyring-backend test --home ${SIX_HOME}
sixd collect-gengate --home ${SIX_HOME}


 # init protocoladmin
sixd keys add alice --keyring-backend test --home ${SIX_HOME}
sixd keys add bob --keyring-backend test --home ${SIX_HOME}
export ALICE_ADDRESS="6x1dhldndym6g543k980xp7wrpjk008z7j6p6nffr"
export BOB_ADDRESS="6x1h29m0g35kjvgkar3fz7nsdq4g5656nlh4hhjlj"

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