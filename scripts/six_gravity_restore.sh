export SIX_HOME=~/.sixtest
export CHAIN_ID=six
export MONIKER=deenode
export VALKEY=validator1
export ORCKEY=orch1
export ETH_ADDRESS="0x645479015C528b33D657375B2198E69A66121c57"
export VAL_ADDRESS="6x1xj3muz3jt5f5ze7px3t5a3ly9cwpz68j4er9vp"
export ORC_ADDRESS="6x1uavet0axk833y8wl67qmpamdmqe0tmrxrxzd8m"

rm -Rf ${SIX_HOME}
cp -r ${SIX_HOME}_backup ${SIX_HOME}

sixd start --home ${SIX_HOME}