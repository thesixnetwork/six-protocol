grantOracle() {
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
}

EVMSIGN=./evmsign
PLATFORM=$1
if [ -z "$PLATFORM" ]; then
    read -p "Enter test platform: [local(defaule), docker, fivenet, sixnet] " _PLATFORM
    PLATFORM=$(echo "$_PLATFORM" | tr '[:upper:]' '[:lower:]')
    # if platform is not set, set it to local
    if [ -z "$PLATFORM" ]; then
        PLATFORM="local"
    fi
fi

# switch case
case $PLATFORM in
"local")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID=testnet
    ;;
"docker")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID=testnet
    ;;
"fivenet")
    RPC_ENDPOINT="https://rpc1.fivenet.sixprotocol.net:443"
    CHAIN_ID=fivenet
    ;;
"sixnet")
    RPC_ENDPOINT="https://sixnet-rpc.sixprotocol.net:443"
    CHAIN_ID=sixnet
    ;;
*)
    echo "Error: unsupported PLATFORM '$PLATFORM'" >&2
    exit 1
    ;;
esac


BASE64_SCHEMA=`cat ./mock-data/nft-schema.json | base64 | tr -d '\n'`

sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}