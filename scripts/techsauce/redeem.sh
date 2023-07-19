token_id=$1
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi

PLATFORM=$2
case $PLATFORM in
"local")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID=testnet
    KEY_NAME=alice
    ;;
"docker")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID=testnet
    KEY_NAME=alice
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

schema=$3

stock=$4
if [ -z "$stock" ]; then
    read -p "Enter Stock: " stock
fi

# if platform is not local or docker then input key
if [ "$PLATFORM" != "local" ] && [ "$PLATFORM" != "docker" ]; then
    read -p "Enter your key name: " KEY_NAME
fi

uuid=$(uuidgen)
params='[{"name":"stock","value":"'$stock'"}]'
sixd tx nftmngr perform-action-by-nftadmin ${schema} ${token_id} \
    redeem_prize ${uuid} ${params} \
    --from ${KEY_NAME} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
