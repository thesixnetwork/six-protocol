default_six_home=six_home
default_docker_tag="4.0.3"
node_homes=(
    sixnode0
    sixnode1
    sixnode2
    sixnode3
)
validator_keys=(
    val1
    val2
    val3
    val4
)

SUPER_ADMIN_ADDRESS="6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"

# ---------------------------------------------------------------------------
# Helper: get the local chain's P2P peer string suitable for Docker containers.
# Replaces IP addresses (0.0.0.0 or LAN IPs) with host.docker.internal so containers can reach the host.
# ---------------------------------------------------------------------------
get_local_peer_for_docker() {
    local raw_peer
    raw_peer=$(jq -r '.app_state.genutil.gen_txs[0].body.memo' ~/.six/config/genesis.json 2>/dev/null || echo "")
    if [ -z "$raw_peer" ]; then
        echo ""
        return
    fi
    # Extract node ID and port, replace IP with host.docker.internal
    # Format: <node-id>@<ip>:<port>
    local node_id=$(echo "$raw_peer" | cut -d'@' -f1)
    local port=$(echo "$raw_peer" | cut -d':' -f2)
    echo "${node_id}@host.docker.internal:${port}"
}

# ---------------------------------------------------------------------------
# Helper: get sixnode0's peer string for hub-and-spoke topology
# Uses docker to get the node ID from sixnode0
# ---------------------------------------------------------------------------
get_sixnode0_peer() {
    if [ ! -f ./build/sixnode0/config/node_key.json ]; then
        echo ""
        return
    fi

    # Try to get node ID using sixd command in a temporary container
    local node_id=$(docker run --rm \
        -v "$(pwd)/build/sixnode0:/opt/build/six_home" \
        asia-southeast1-docker.pkg.dev/six-protocol/six-node-docker-repo/sixnode:${default_docker_tag} \
        sixd tendermint show-node-id --home /opt/build/six_home 2>/dev/null || echo "")

    if [ -z "$node_id" ]; then
        echo ""
        return
    fi

    # sixnode0 is accessible at 10.10.0.2:26656 within Docker network
    echo "${node_id}@10.10.0.2:26656"
}

# ---------------------------------------------------------------------------
function setUpGenesis() {
    ## config genesis.json
    jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## demom metadata
    ## bank
    jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.bank.denom_metadata[0] =  {"description": "The native staking token of the SIX Protocol.","denom_units": [{"denom": "usix","exponent": 0,"aliases": ["microsix"]},{"denom": "msix","exponent": 3,"aliases": ["millisix"]},{"denom": "six","exponent": 6,"aliases": []}],"base": "usix","display": "six","name": "Six token","symbol": "six"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.bank.denom_metadata[1] =  {"description": "The native evm token of the SIX Protocol.","denom_units": [{"denom": "asix","exponent": 0,"aliases": ["attosix"]},{"denom": "usix","exponent": 12,"aliases": ["microsix"]},{"denom": "msix","exponent": 15,"aliases": ["millisix"]},{"denom": "six","exponent": 18,"aliases": []}],"base": "asix","display": "asix","name": "aSIX token","symbol": "asix"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## from stake to usix
    sed -i '' "s/stake/usix/g" ./build/sixnode0/config/genesis.json

    ## evm
    jq '.app_state.evm.params.evm_denom="asix"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## feemarket
    jq '.app_state.feemarket.params.base_fee = "5000000000000"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.feemarket.params.elasticity_multiplier = 4' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.feemarket.params.min_gas_price = "5000000000000.000000000000000000"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## nftadmin
    jq '.app_state.nftadmin.authorization = {"root_admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## nftmngr
    jq '.app_state.nftmngr.nft_fee_config = {"schema_fee": {"fee_amount": "200000000usix","fee_distributions": [{"method": "BURN","portion": 0.5},{"method": "REWARD_POOL","portion": 0.5}]}}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## nftoracle
    jq '.app_state.nftoracle.params = {"action_request_active_duration": "120s","mint_request_active_duration": "120s","verify_request_active_duration": "120s", "action_signer_active_duration": "2592000s","sync_action_signer_active_duration": "300s"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.nftoracle.oracle_config = {"minimum_confirmation": 4}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## protocoladmin
    jq '.app_state.protocoladmin.adminList[0] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "super.admin"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.protocoladmin.adminList[1] |= . + {"admin": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","group": "token.admin"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.protocoladmin.groupList[0] |= . + {"name": "super.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.protocoladmin.groupList[1] |= . + {"name": "token.admin","owner": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## staking
    jq '.app_state.staking.validator_approval.approver_address = "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.staking.params.unbonding_time = "300s"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## tokenmngr
    jq '.app_state.tokenmngr.mintpermList[0] |= . + {"address": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","token": "usix"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.tokenmngr.options = {"defaultMintee": "6x1cws3ex5yqwlu4my49htq06nsnhuxw3v7rt20g6"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.tokenmngr.tokenList[0] |= . +  {"base": "usix","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","maxSupply": {"amount": "0","denom": "usix"},"mintee": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","name": "usix"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.tokenmngr.tokenList[1] |= . +  {"base": "asix","creator": "6x1t3p2vzd7w036ahxf4kefsc9sn24pvlqphcuauv","maxSupply": {"amount": "0","denom": "asix"},"mintee": "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq","name": "asix"}' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json

    ## gov
    jq '.app_state.gov.deposit_params.max_deposit_period = "300s"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
    jq '.app_state.gov.voting_params.voting_period = "300s"' ./build/sixnode0/config/genesis.json | sponge ./build/sixnode0/config/genesis.json
}

function setUpConfig() {
    echo "#######################################"
    echo "Setup ${SIX_HOME} genesis..."

    if [[ ${SIX_HOME} == "sixnode0" ]]; then
        echo "sixnode0"
        setUpGenesis
    else
        NODE_PEER=$(jq '.app_state.genutil.gen_txs[0].body.memo' ./build/sixnode0/config/genesis.json)
        if [[ "$OSTYPE" == "darwin"* ]]; then
            ## replace NODE_PEER in config.toml to persistent_peers
            sed -i '' "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
        else
            sed -i "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml
        fi
        ## replace genesis of node0 to all node
        cp ./build/sixnode0/config/genesis.json ./build/${SIX_HOME}/config/genesis.json
    fi

    # if $TYPE = 0 then ignore this step
    if [[ ${TYPE} == "1" ]]; then
        echo "Running Fast Node"
        ## replace consensus params
        if [[ "$OSTYPE" == "darwin"* ]]; then
            sed -i '' "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
            sed -i '' "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        else
            sed -i "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
            sed -i "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        fi
    else
        echo "Running Default Node"
    fi

    if [[ "$OSTYPE" == "darwin"* ]]; then
        ## replace to enalbe api
        sed -i '' '/^\[api\]$/,/^\[/ s/enable = false/enable = true/' ./build/${SIX_HOME}/config/app.toml
        sed -i '' '/^\[api\]$/,/^[^[]/ s/^swagger = false$/swagger = true/' ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i '' "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

        ## replace mininum gas price
        sed -i '' "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix,1250000000000asix\"/g" ./build/${SIX_HOME}/config/app.toml
    else
        sed -i '/^\[api\]$/,/^\[/ s/enable = false/enable = true/' ./build/${SIX_HOME}/config/app.toml
        sed -i '/^\[api\]$/,/^[^[]/ s/^swagger = false$/swagger = true/' ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

        ## replace mininum gas price
        sed -i "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix,1250000000000asix\"/g" ./build/${SIX_HOME}/config/app.toml
    fi

    echo "Setup Genesis Success 🟢"

}

function setupLocalSyncConfig() {
    echo "#######################################"
    echo "Setup ${SIX_HOME} for local chain sync..."

    # Get local chain peer string
    LOCAL_PEER=$(get_local_peer_for_docker)
    if [ -z "$LOCAL_PEER" ]; then
        echo "ERROR: Could not read peer from ~/.six/config/genesis.json 🖕"
        exit 1
    fi

    # Determine peers based on node role:
    # - sixnode0: connects only to local chain (acts as gateway/hub)
    # - sixnode1-3: connect to both local chain AND sixnode0 (hub-and-spoke)
    if [[ "${SIX_HOME}" == "sixnode0" ]]; then
        echo "Configuring sixnode0 as gateway (connects to local chain only)"
        PEERS="${LOCAL_PEER}"
    else
        echo "Configuring ${SIX_HOME} with hub-and-spoke topology (local chain + sixnode0)"
        SIXNODE0_PEER=$(get_sixnode0_peer)
        if [ -z "$SIXNODE0_PEER" ]; then
            echo "WARNING: Could not get sixnode0 peer ID. Using local chain peer only."
            PEERS="${LOCAL_PEER}"
        else
            # Connect to both local chain and sixnode0
            PEERS="${LOCAL_PEER},${SIXNODE0_PEER}"
            echo "Peers: local chain + sixnode0 (${SIXNODE0_PEER})"
        fi
    fi

    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/g" ./build/${SIX_HOME}/config/config.toml
        # Enable PEX (peer exchange) to discover other peers
        sed -i '' "s/pex = false/pex = true/g" ./build/${SIX_HOME}/config/config.toml
        # Allow private peer IDs
        sed -i '' "s/addr_book_strict = true/addr_book_strict = false/g" ./build/${SIX_HOME}/config/config.toml
    else
        sed -i "s/persistent_peers = \"\"/persistent_peers = \"${PEERS}\"/g" ./build/${SIX_HOME}/config/config.toml
        # Enable PEX (peer exchange) to discover other peers
        sed -i "s/pex = false/pex = true/g" ./build/${SIX_HOME}/config/config.toml
        # Allow private peer IDs
        sed -i "s/addr_book_strict = true/addr_book_strict = false/g" ./build/${SIX_HOME}/config/config.toml
    fi

    ## Copy genesis from local chain
    cp ~/.six/config/genesis.json ./build/${SIX_HOME}/config/genesis.json

    # if $TYPE = 0 then ignore this step
    if [[ ${TYPE} == "1" ]]; then
        echo "Running Fast Node"
        ## replace consensus params
        if [[ "$OSTYPE" == "darwin"* ]]; then
            sed -i '' "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
            sed -i '' "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        else
            sed -i "s/timeout_propose = \"3s\"/timeout_propose = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
            sed -i "s/timeout_commit = \"5s\"/timeout_commit = \"1s\"/g" ./build/${SIX_HOME}/config/config.toml
        fi
    else
        echo "Running Default Node"
    fi

    if [[ "$OSTYPE" == "darwin"* ]]; then
        ## replace to enalbe api
        sed -i '' '/^\[api\]$/,/^\[/ s/enable = false/enable = true/' ./build/${SIX_HOME}/config/app.toml
        sed -i '' '/^\[api\]$/,/^[^[]/ s/^swagger = false$/swagger = true/' ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i '' "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

        ## replace mininum gas price
        sed -i '' "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix,1250000000000asix\"/g" ./build/${SIX_HOME}/config/app.toml
    else
        sed -i '/^\[api\]$/,/^\[/ s/enable = false/enable = true/' ./build/${SIX_HOME}/config/app.toml
        sed -i '/^\[api\]$/,/^[^[]/ s/^swagger = false$/swagger = true/' ./build/${SIX_HOME}/config/app.toml
        ## replace to from 127.0.0.1 to 0.0.0.0
        sed -i "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

        ## replace mininum gas price
        sed -i "s/minimum-gas-prices = \"0stake\"/minimum-gas-prices = \"1.25usix,1250000000000asix\"/g" ./build/${SIX_HOME}/config/app.toml
    fi

    echo "Setup Genesis Success 🟢"

}

echo "#############################################"
echo "## 1.  Build Docker Image                  ##"
echo "## 2.  Docker Compose init chain           ##"
echo "## 3.  Start chain validator               ##"
echo "## 4.  Stop chain validator                ##"
echo "## 5.  Config Genesis (docker cluster)     ##"
echo "## 5.2 Config Genesis for Local Sync       ##"
echo "## 6.  Reset chain validator               ##"
echo "## 7.  Staking validator (docker cluster)  ##"
echo "## 7.2 Staking validator (local sync)      ##"
echo "## 8.  Query Validator set                 ##"
echo "## 9.  Setup Cosmovisor                    ##"
echo "## 10. Start Cosmovisor                    ##"
echo "#############################################"
read -p "Enter your choice: " choice
case $choice in
1)
    echo "Building Docker Image"
    read -p "Enter Docker Tag: " docker_tag
    if [ -z "$docker_tag" ]; then
        docker_tag=$default_docker_tag
    fi
    docker build . -t asia-southeast1-docker.pkg.dev/six-protocol/six-node-docker-repo/sixnode:${docker_tag}
    ;;
2)
    echo "Run init Chain validator"
    export COMMAND="init"
    docker compose -f ./docker-compose.yml up -d
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
    read -p "Enter Node Type [0:Default, 1:Fast] : " TYPE
    if [ -z "$TYPE" ]; then
        TYPE=0
    fi
    for home in ${node_homes[@]}; do
        (
            export SIX_HOME=${home}
            if [[ -e !./build/sixnode0/config/genesis.json ]]; then
                echo "File does not exist 🖕"
            else
                setUpConfig
            fi
        ) || exit 1
    done
    ;;
"5.2")
    echo "Config Genesis for Local Sync"
    read -p "Enter Node Type [0:Default, 1:Fast] : " TYPE
    if [ -z "$TYPE" ]; then
        TYPE=0
    fi
    if ! [[ -e ~/.six/config/genesis.json ]]; then
        echo "~/.six/config/genesis.json does not exist. Run bash init_testnet.sh first. 🖕"
        exit 1
    fi
    for home in ${node_homes[@]}; do
        (
            export SIX_HOME=${home}
            if ! [[ -e ./build/${SIX_HOME}/config/config.toml ]]; then
                echo "ERROR: ./build/${SIX_HOME}/config/config.toml not found. Initialize the nodes first. 🖕"
                exit 1
            fi
            setupLocalSyncConfig
        ) || exit 1
    done
    ;;
6)
    echo "Reset Docker Container"
    for home in ${node_homes[@]}; do
        echo "#######################################"
        echo "Starting ${home} reset..."

        (
            export DAEMON_HOME=./build/${home}
            rm -rf $DAEMON_HOME/data
            rm -rf $DAEMON_HOME/wasm
            rm $DAEMON_HOME/config/addrbook.json
            mkdir $DAEMON_HOME/data/
            touch $DAEMON_HOME/data/priv_validator_state.json
            echo '{"height": "0", "round": 0,"step": 0}' >$DAEMON_HOME/data/priv_validator_state.json

            echo "Reset ${home} Success 🟢"
        ) || exit 1
    done
    ;;
7)
    echo "Staking Docker Container"
    read -p "Chain ID [testnet] : " CHAIN_ID
    if [ -z "$CHAIN_ID" ]; then
        CHAIN_ID="testnet"
    fi
    i=1
    amount=100000000000
    # i=0
    # for val in ${validator_keys[@]}
    for val in ${validator_keys[@]:1:3}; do
        echo "#######################################"
            (
                echo "Creating validators ${val}"
                echo ${node_homes[i]}
                export DAEMON_HOME=./build/${node_homes[i]}
                sixd tx staking create-validator-legacy --amount="${amount}usix" --moniker ${node_homes[i]} --pubkey $(sixd tendermint show-validator --home ./build/${node_homes[i]}) \
                    --validator-mode="${i-1}" --max-license=100 --min-delegation 10000000000 --delegation-increment 10000000000 --enable-redelegation=false --min-self-delegation 10000000000 \
                    --commission-rate "0.1" --commission-max-rate "0.1" --commission-max-change-rate "0.1" \
                    --details "node_test_${i}" --security-contact "node_test_${i}" --website "www.idk_${i}.com" --identity "idk_${i}" \
                    --sign-mode amino-json --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
                    --keyring-backend test --chain-id $CHAIN_ID --from=${val} --home build/${node_homes[i]} -y --node http://0.0.0.0:26662
                echo "Config Genesis at ${home} Success 🟢"
        ) || exit 1
        i=$((i + 1))
    done
    ;;
"7.2")
    echo "Staking validators for local chain sync"
    read -p "Chain ID [testnet] : " CHAIN_ID
    if [ -z "$CHAIN_ID" ]; then
        CHAIN_ID="testnet"
    fi
    # Local chain RPC is at 0.0.0.0:26657 (init_testnet.sh sets it up this way)
    LOCAL_RPC="http://0.0.0.0:26657"
    amount=100000000000
    i=0
    for val in ${validator_keys[@]:0:4}; do
        echo "#######################################"
        (
            NODE_HOME=${node_homes[i]}
            echo "Creating validator ${val} on ${NODE_HOME}"
            sixd tx staking create-validator-legacy --amount="${amount}usix" --moniker ${NODE_HOME} \
                --pubkey $(sixd tendermint show-validator --home ./build/${NODE_HOME}) \
                --validator-mode=0 --max-license=100 --min-delegation 10000000000 --delegation-increment 10000000000 --enable-redelegation=false --min-self-delegation 10000000000 \
                --commission-rate "1" --commission-max-rate "1" --commission-max-change-rate "1" \
                --details "local_sync_${i}" --security-contact "local_sync_${i}" --website "www.six_${i}.com" --identity "six_${i}" \
                --sign-mode amino-json --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
                --approver $SUPER_ADMIN_ADDRESS \
                --keyring-backend test --chain-id $CHAIN_ID --from=${val} --home build/${NODE_HOME} -y --node $LOCAL_RPC
            echo "Staking for ${NODE_HOME} Success 🟢"
        ) || exit 1
        i=$((i + 1))
    done
    ;;
8)
    echo "Query Validator set"
    sixd q tendermint-validator-set --home ./build/sixnode0
    ;;
9)
    echo "Set up Cosmovisor"
    export COMMAND="cosmovisor_setup"
    docker compose -f ./docker-compose.yml up -d
    ;;
10)
    echo "Cosmovisor start"
    export COMMAND="cosmovisor_start"
    docker compose -f ./docker-compose.yml up -d
    ;;
*)
    echo "Invalid Choice"
    ;;
esac
