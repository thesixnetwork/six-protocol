
export SIX_HOME=~/.sixtest
export CHAIN_ID=six
export MONIKER=deenode
export VALKEY=validator1
export ORCKEY=orch1

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}
sixd keys add ${VALKEY} --keyring-backend test --home ${SIX_HOME}
sixd keys add ${ORCKEY} --keyring-backend test --home ${SIX_HOME}
export VAL_ADDRESS="6x1xj3muz3jt5f5ze7px3t5a3ly9cwpz68j4er9vp"
export ORC_ADDRESS="6x1uavet0axk833y8wl67qmpamdmqe0tmrxrxzd8m"

sixd eth_keys add --keyring-backend test --home ${SIX_HOME}
export ETH_ADDRESS="0x645479015C528b33D657375B2198E69A66121c57"

sixd add-genesis-account ${VALKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account ${ORCKEY} 1000000000000stake --keyring-backend test --home ${SIX_HOME}

# modify nativeHRP
code ${SIX_HOME}/config/genesis.json

sixd gengate --moniker=${MONIKER} ${VALKEY} 1000000000stake \
    ${ETH_ADDRESS} ${VAL_ADDRESS} --chain-id=${CHAIN_ID} \
    --keyring-backend test --home ${SIX_HOME}
sixd collect-gengate --home ${SIX_HOME}

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