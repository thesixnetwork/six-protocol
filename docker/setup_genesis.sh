default_six_home=six_home
SIX_HOME=$1
if [ -z "$SIX_HOME" ]; then
    SIX_HOME=$default_six_home
fi

## funtion setup_genesis
function setupGenesis() {
    NODE_PEER=$(jq '.app_state.genutil.gen_txs[0].body.memo' ./build/sixnode0/config/genesis.json)

    if [[ "$OSTYPE" == "darwin"* ]]; then
        ## replace NODE_PEER in config.toml to persistent_peers
        sed -i '' "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
        ## replace mininum gas price
        sed -i '' "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix\"/g" ./build/${SIX_HOME}/config/app.toml
        ## replace to enalbe api
        sed -i '' "108s/.*/enable = true/" ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i '' "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml
        ## replace consensus params
        sed -i '' "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        sed -i '' "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
           ## from stake to usix
        sed -i '' "s/stake/usix/g" ./build/${SIX_HOME}/config/genesis.json
    else 
        ## replace NODE_PEER in config.toml to persistent_peers
        sed -i "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
        ## replace mininum gas price
        sed -i "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix\"/g" ./build/${SIX_HOME}/config/app.toml
        ## replace to enalbe api
        sed -i "108s/.*/enable = true/" ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml
        ## replace consensus params
        sed -i "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        sed -i "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
    fi 

    ## config genesis.json
    jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## demom metadata
    jq '.app_state.bank.denom_metadata[0] =  {"description": "The native staking token of the SIX Protocol.","denom_units": [{"denom": "usix","exponent": 0,"aliases": ["microsix"]},{"denom": "msix","exponent": 3,"aliases": ["millisix"]},{"denom": "six","exponent": 6,"aliases": []}],"base": "usix","display": "six","name": "Six token","symbol": "six"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json


    ## nftadmin
    jq '.app_state.nftadmin.authorization = {"root_admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## nftmngr
    jq '.app_state.nftmngr.nft_fee_config = {"schema_fee": {"fee_amount": "200000000usix","fee_distributions": [{"method": "BURN","portion": 0.5},{"method": "REWARD_POOL","portion": 0.5}]}}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## nftoracle
    jq '.app_state.nftoracle.params = {"action_request_active_duration": "120s","mint_request_active_duration": "120s","verify_request_active_duration": "120s", "action_signer_active_duration": "2592000s","sync_action_signer_active_duration": "300s"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.nftoracle.oracle_config = {"minimum_confirmation": 4}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## protocoladmin
    jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "super.admin"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.adminList[1] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "token.admin"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.groupList[1] |= . + {"name": "token.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## staking 
    jq '.app_state.staking.validator_approval.approver_address = "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## tokenmngr
    jq '.app_state.tokenmngr.mintpermList[0] |= . + {"address": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","token": "usix"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.tokenmngr.options = {"defaultMintee": "6x1cws3ex5yqwlu4my49htq06nsnhuxw3v7rt20g6"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.tokenmngr.tokenList[0] |= . +  {"base": "usix","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","maxSupply": { "amount": "0", "denom": "usix" },"mintee": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","name": "usix"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json 
    
    ## gov
    jq '.app_state.gov.deposit_params.max_deposit_period = "300s"' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.gov.voting_params.voting_period = "300s"' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    echo "Setup Genesis Success 🟢"

}

if [[ ! -e ./build/sixnode0/config/genesis.json ]]; then
    echo "File does not exist 🖕"
else
    setupGenesis
fi