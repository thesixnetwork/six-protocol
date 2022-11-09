default_github_token=$1
default_docker_tag = "latest"
node_homes=(
    sixnode0
    sixnode1
    sixnode2
    sixnode3
);
validator_keys=(
    val0
    val1
    val2
    val3
);
echo "#############################################"
echo "## 1. Build Docker Image                   ##"
echo "## 2. Docker Compose init chain            ##"
echo "## 3. Run Start chain validator            ##"
echo "## 4. Run Stop chain validator             ##"
echo "## 5. Reset chain validator                ##"
echo "## 6. Staking validator                    ##"
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
        docker compose -f ./docker-compose-init.yml up
        ;;
    3)
        echo "Running Docker Container in Interactive Mode"
        docker compose -f ./docker-compose-start.yml up -d
        ;;
    4)
        echo "Stop Docker Container"
        docker compose -f ./docker-compose-start.yml down
        ;;
    5)
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
    6)
        echo "Staking Docker Container"
        i=1
        for val in ${validator_keys[@]:1:3}
        do
            echo "#######################################"
            ( 
            echo "Creating validators ${val}"
            export DAEMON_HOME=./build/${node_homes[i]}
            sixd tx staking create-validator --amount="100000000stake" --from=${val} --moniker ${node_homes[i]} \
                --pubkey $(sixd tendermint show-validator --home ./build/${node_homes[i]}) --home build/${node_homes[i]} \
                --keyring-backend test --commission-rate 0.1 --commission-max-rate 0.5 --commission-max-change-rate 0.1 \
                --min-self-delegation 1000000 --node http://0.0.0.0:26662 -y --min-delegation 1000000 --delegation-increment 1000000
            ) || exit 1
            i=$((i+1))
        done
        ;;
    *)
        echo "Invalid Choice"
        ;;
esac