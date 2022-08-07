export SIX_HOME=~/.six_test
export CHAIN_ID=six
export MONIKER=deenode
export VALKEY=validator1
export ORCKEY=orch1
export VAL_ADDRESS="6x1fdts53zq5xtnmmap3a8enffjxzcuvv2tddldds"
export ORC_ADDRESS="6x14kee3xxg6v88akhyu3ha3dwhctqm6ze4kkys9m"
export ETH_ADDRESS="0xD224824bBE868095132ee2d3A50aE770D0DFbb8c"

rm -Rf ${SIX_HOME}
cp -r ${SIX_HOME}_backup ${SIX_HOME}

sixd start --home ${SIX_HOME}