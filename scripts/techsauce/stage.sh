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

# if platform is not local or docker then input key
if [ "$PLATFORM" != "local" ] && [ "$PLATFORM" != "docker" ]; then
    read -p "Enter your key name: " KEY_NAME
fi

schema=$3

uuid=$(uuidgen)
sixd tx nftmngr perform-multi-token-action ${schema} ${token_id} \
    attend_stage,attend_stage,attend_stage,attend_stage,attend_stage attend_stage_${uuid} \
    '[[{"name":"stage","value":"stage_1"}],[{"name":"stage","value":"stage_2"}],[{"name":"stage","value":"stage_3"}],[{"name":"stage","value":"stage_4"}],[{"name":"stage","value":"stage_5"}]]' \
    --from ${KEY_NAME} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
