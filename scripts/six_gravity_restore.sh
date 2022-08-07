export CHAIN_ID=six
export MONIKER=mynode
export VALKEY=validator1
export ORCKEY=orch1
export SIX_HOME=~/.six_test
export VAL_ADDRESS="6x1fjy5cfjp2pqqt430lexalwtm872jjlfjy9qgzt"
export SUPERADMIN_ADDRESS="6x1l0ceauyrkuhte463halxz8tawrlsv5vxc3jxer"
export ORC_ADDRESS="6x1gvdc9zgc9m9ap5hgs2w7g4mcdsun93qzt84a2z"

rm -Rf ${SIX_HOME}
cp -r ${SIX_HOME}_backup ${SIX_HOME}

sixd start --home ${SIX_HOME}