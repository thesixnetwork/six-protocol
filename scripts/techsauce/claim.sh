token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
action=$2
# if action is not provided, ask for it
if [ -z "$action" ]; then
    read -p "Enter Action: " action
fi

PLATFORM=$3
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

schema=$4

# if platform is not local or docker then input key
if [ "$PLATFORM" != "local" ] && [ "$PLATFORM" != "docker" ]; then
    read -p "Enter your key name: " KEY_NAME
fi

uuid=$(uuidgen)
sixd tx nftmngr perform-action-by-nftadmin ${schema} ${token_id} \
    ${action} ${uuid} '[]' \
    --from ${KEY_NAME} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
