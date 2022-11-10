default_github_token=$1
default_six_home=six_home
default_docker_tag="latest"
node_homes=(
    sixnode0
    sixnode1
    sixnode2
    sixnode3
);
validator_keys=(
    val1
    val2
    val3
    val4
);

function setupGenesis() {
    echo "#######################################"
    echo "Setup ${SIX_HOME} genesis..."
    echo "#######################################"
    NODE_PEER=$(jq '.app_state.genutil.gen_txs[0].body.memo' ./build/sixnode0/config/genesis.json)

    if [[ ${SIX_HOME} == "sixnode0" ]]; then
        echo "sixnode0"
        NODE_PEER=$(jq '.app_state.genutil.gen_txs[0].body.memo' ./build/sixnode1/config/genesis.json)
        sed -i '' "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
    else
        ## replace NODE_PEER in config.toml to persistent_peers
        sed -i '' "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
    fi

    ## replace to enalbe api
    sed -i '' "108s/.*/enable = true/" ./build/${SIX_HOME}/config/app.toml

    ## replace to from 127.0.0.1 to 0.0.0.0
    sed -i '' "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

    ## config genesis.json
    jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## demom metadata
    jq '.app_state.bank.denom_metadata[0] =  {"description": "The native staking token of the SIX Protocol.","denom_units": [{"denom": "usix","exponent": 0,"aliases": ["microsix"]},{"denom": "msix","exponent": 3,"aliases": ["millisix"]},{"denom": "six","exponent": 6,"aliases": []}],"base": "usix","display": "six","name": "Six token","symbol": "six"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## from stake to usix
    sed -i '' "s/stake/usix/g" ./build/${SIX_HOME}/config/genesis.json

    ## nftadmin
    jq '.app_state.nftadmin.authorization = {"root_admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## nftmngr
    jq '.app_state.nftmngr.nft_fee_config = {"schema_fee": {"fee_amount": "200000000usix","fee_distributions": [{"method": "BURN","portion": 0.5},{"method": "REWARD_POOL","portion": 0.5}]}}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## nftoracle
    jq '.app_state.nftoracle.params = {"action_request_active_duration": "120s","mint_request_active_duration": "120s","verify_request_active_duration": "120s"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.nftoracle.oracle_config = {"minimum_confirmation": 4}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## protocoladmin
    jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "super.admin"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.adminList[1] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "token.admin"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin","owner": "genesis"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.protocoladmin.groupList[1] |= . + {"name": "token.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## staking 
    jq '.app_state.staking.validator_approval.approver_address = "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json

    ## tokenmngr
    jq '.app_state.tokenmngr.mintpermList[0] |= . + {"address": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","token": "usix"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.tokenmngr.options = {"defaultMintee": "6x1cws3ex5yqwlu4my49htq06nsnhuxw3v7rt20g6"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
    jq '.app_state.tokenmngr.tokenList[0] |= . +  {"base": "usix","creator": "6x1eau6xz2kdv6wy7rhj2nxv0xrgnjy79hcm2tr9t","maxSupply": 0,"mintee": "6x1cws3ex5yqwlu4my49htq06nsnhuxw3v7rt20g6","name": "usix"}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json 
    
    echo "Setup Genesis Success 🟢"

}

echo "#############################################"
echo "## 1. Build Docker Image                   ##"
echo "## 2. Docker Compose init chain            ##"
echo "## 3. Start chain validator                ##"
echo "## 4. Stop chain validator                 ##"
echo "## 5. Config Genesis                       ##"
echo "## 6. Reset chain validator                ##"
echo "## 7. Staking validator                    ##"
echo "#############################################"

read -p "Enter your choice: " choice
case $choice in
    1)
        echo "Building Docker Image"
        read -p "Enter Github Token: " github_token 
        read -p "Enter Docker Tag: " docker_tag
        if [ -z "$github_token" ]; then
            github_token=$default_github_token
        fi
        if [ -z "$docker_tag" ]; then
            docker_tag=$default_docker_tag
        fi
        docker build . -t six/node:${docker_tag} --build-arg GITHUB_TOKEN=${github_token}
        ;;
    2)
        echo "Run init Chain validator"
        export COMMAND="init"
        docker compose -f ./docker-compose.yml up
        ;;
    3)
        echo "Running Docker Container in Interactive Mode"
        export COMMAND="start_chain"
        docker compose -f ./docker-compose.yml up -d
        ;;
    4)
        echo "Stop Docker Container"
        export COMMAND="start_chain"
        docker compose -f ./docker-compose.yml down
        ;;
    5) 
        echo "Config Genesis"
        for home in ${node_homes[@]}
        do  
            (
            export SIX_HOME=${home}
            if [[ -e !./build/sixnode0/config/genesis.json ]]; then
                echo "File does not exist 🖕"
            else
                setupGenesis
            fi 
            )|| exit 1
        done
        ;;
    6) 
        echo "Reset Docker Container"
        for home in ${node_homes[@]}
        do
            echo "#######################################"
            echo "Starting ${home} reset..."
            echo "#######################################"
            ( export DAEMON_HOME=./build/${home}
            rm -rf $DAEMON_HOME/data
            rm -rf $DAEMON_HOME/wasm
            rm $DAEMON_HOME/config/addrbook.json
            mkdir $DAEMON_HOME/data/
            touch $DAEMON_HOME/data/priv_validator_state.json
            echo '{"height": "0", "round": 0,"step": 0}' > $DAEMON_HOME/data/priv_validator_state.json
            )|| exit 1
        done
        ;;
    7)
        echo "Staking Docker Container"
        i=0
        amount=100000001
        for val in ${validator_keys[@]:1:3}
        # for val in ${validator_keys[@]}
        do
            echo "#######################################"
            ( 
            echo "Creating validators ${val}"
            echo ${node_homes[i]}
            export DAEMON_HOME=./build/${node_homes[i]}
            sixd tx staking create-validator --amount="${amount}usix" --from=${val} --moniker ${node_homes[i]} \
                --pubkey $(sixd tendermint show-validator --home ./build/${node_homes[i]}) --home build/${node_homes[i]} \
                --keyring-backend test --commission-rate 0.1 --commission-max-rate 0.5 --commission-max-change-rate 0.1 \
                --min-self-delegation 1000000 --node http://0.0.0.0:26662 -y --min-delegation 1000000 --delegation-increment 1000000
            ) || exit 1
            i=$((i+1))
            amount=$((amount+1))
        done
        ;;
    *)
        echo "Invalid Choice"
        ;;
esac