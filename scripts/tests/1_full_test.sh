PLATFORM=$1
key=$2
if [ -z "$PLATFORM" ]; then
    read -p "Enter test platform: [local(defaule), docker, fivenet, sixnet] " _PLATFORM
    PLATFORM=$(echo "$_PLATFORM" | tr '[:upper:]' '[:lower:]')
    # if platform is not set, set it to local
    if [ -z "$PLATFORM" ]; then
        PLATFORM="local"
    fi
fi

if [ -z "$key" ]; then
    read -p "Enter key: [alice] " key
fi

if [ -z "$key" ] && [ "$PLATFORM" == "local" ] || [ "$PLATFORM" == "docker" ]; then
    key="alice"
else
    while [ -z "$key" ]; do
        read -p "Cannot use defult key for this platform please specify your key >> " key
    done
fi

# switch case
case $PLATFORM in
"local")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID="testnet"
    ;;
"docker")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID="sixnet"
    ;;
"fivenet")
    RPC_ENDPOINT="https://rpc1.fivenet.sixprotocol.net:443"
    CHAIN_ID="fivenet"
    ;;
"sixnet")
    RPC_ENDPOINT="https://sixnet-rpc.sixprotocol.net:443"
    CHAIN_ID="sixnet"
    ;;
*)
    echo "Error: unsupported PLATFORM '$PLATFORM'" >&2
    exit 1
    ;;
esac


# array
modules=(
    protocoladmin
    tokenmngr
    nftadmin
    nftmngr
)

for mod in ${modules[@]}
do
    echo "#######################################"
    echo "Starting ${mod} tests on ${RPC_ENDPOINT} and chainId ${CHAIN_ID}..."
    echo "#######################################"
    bash ${mod}.sh ${RPC_ENDPOINT} ${CHAIN_ID} ${key} || exit 1
done