MONIKER=$1
if [ -z "$MONIKER" ]; then
  MONIKER="mynode"
fi
export CHAIN_ID=sixnet
export VALKEY=val1 # should be: export as docker env var
export SIX_HOME=./build/six_home
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

rm -Rf ${SIX_HOME}

sixd init ${MONIKER} --chain-id=${CHAIN_ID} --home ${SIX_HOME}

# # Change parameter token denominations to usix
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="usix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="usix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="usix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="asix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="usix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state["mint"]["params"]["mint_denom"]="usix"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.bank.denom_metadata[0] =  {"description": "The native staking token of the SIX Protocol.","denom_units": [{"denom": "usix","exponent": 0,"aliases": ["microsix"]},{"denom": "msix","exponent": 3,"aliases": ["millisix"]},{"denom": "six","exponent": 6,"aliases": []}],"base": "usix","display": "six","name": "Six token","symbol": "six"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.bank.denom_metadata[1] =  {"description": "The native evm token of the SIX Protocol.","denom_units": [{"denom": "asix","exponent": 0,"aliases": ["attosix"]},{"denom": "usix","exponent": 12,"aliases": ["microsix"]},{"denom": "msix","exponent": 15,"aliases": ["millisix"]},{"denom": "six","exponent": 18,"aliases": []}],"base": "asix","display": "asix","name": "aSIX token","symbol": "asix"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.nftadmin.authorization = {"root_admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.nftmngr.nft_fee_config = {"schema_fee": {"fee_amount": "200000000usix","fee_distributions": [{"method": "BURN","portion": 0.5},{"method": "REWARD_POOL","portion": 0.5}]}}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.nftoracle.params = {"action_request_active_duration": "120s","mint_request_active_duration": "120s","verify_request_active_duration": "120s", "action_signer_active_duration": "2592000s","sync_action_signer_active_duration": "300s"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.nftoracle.oracle_config = {"minimum_confirmation": 4}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "super.admin"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.protocoladmin.adminList[1] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "token.admin"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.protocoladmin.groupList[1] |= . + {"name": "token.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.staking.validator_approval.approver_address = "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.tokenmngr.mintpermList[0] |= . + {"address": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","token": "usix"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.tokenmngr.options = {"defaultMintee": "6x1cws3ex5yqwlu4my49htq06nsnhuxw3v7rt20g6"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.tokenmngr.tokenList[0] |= . +  {"base": "usix","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","maxSupply": { "amount": "0", "denom": "usix" },"mintee": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","name": "usix"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.tokenmngr.tokenList[1] |= . +  {"base": "asix","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","maxSupply": { "amount": "0", "denom": "asix" },"mintee": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","name": "asix"}' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.gov.deposit_params.max_deposit_period = "300s"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json
# cat ${SIX_HOME}/config/genesis.json | jq '.app_state.gov.voting_params.voting_period = "300s"' > ${SIX_HOME}/config/tmp_genesis.json && mv ${SIX_HOME}/config/tmp_genesis.json ${SIX_HOME}/config/genesis.json

# mint to validator
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ALICE_MNEMONIC | sixd keys add alice --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $BOB_MNEMONIC | sixd keys add bob --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL1_MNEMONIC | sixd keys add val1 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL2_MNEMONIC | sixd keys add val2 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL3_MNEMONIC | sixd keys add val3 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $VAL4_MNEMONIC | sixd keys add val4 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE1_MNEMONIC | sixd keys add oracle1 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE2_MNEMONIC | sixd keys add oracle2 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE3_MNEMONIC | sixd keys add oracle3 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118
echo $ORACLE4_MNEMONIC | sixd keys add oracle4 --recover --home ${SIX_HOME} --keyring-backend test --algo secp256k1 --coin-type 118

sixd add-genesis-account $(sixd keys show -a val1 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val2 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val3 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a val4 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle1 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle2 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle3 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a oracle4 --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a alice --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a bob --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}
sixd add-genesis-account $(sixd keys show -a super-admin --keyring-backend=test --home ${SIX_HOME}) 1000000000000usix --keyring-backend test --home ${SIX_HOME}

# eth_secp256k1 address
echo $SUPER_ADMIN_MNEMONIC | sixd keys add super-admin_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $ALICE_MNEMONIC | sixd keys add alice_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $BOB_MNEMONIC | sixd keys add bob_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $VAL1_MNEMONIC | sixd keys add val1_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $VAL2_MNEMONIC | sixd keys add val2_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $VAL3_MNEMONIC | sixd keys add val3_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $VAL4_MNEMONIC | sixd keys add val4_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $ORACLE1_MNEMONIC | sixd keys add oracle1_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $ORACLE2_MNEMONIC | sixd keys add oracle2_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $ORACLE3_MNEMONIC | sixd keys add oracle3_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1
echo $ORACLE4_MNEMONIC | sixd keys add oracle4_eth --recover --home ${SIX_HOME} --keyring-backend test --algo eth_secp256k1

sixd gentx ${VALKEY} 100000000usix --chain-id=${CHAIN_ID} --keyring-backend=test --home ${SIX_HOME}
sixd collect-gentxs --home ${SIX_HOME}