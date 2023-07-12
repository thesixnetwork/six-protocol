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

BASE64_SCHEMA=$(cat ./mock-data/nft-schema.json | base64 | tr -d '\n')

sixd tx nftmngr create-nft-schema --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}

schema_code=$2
if [ -z "$schema_code" ]; then
    read -p "Enter schema code: " schema_code
fi

token_id=1,2,3,4,5,6,7,8,9,10
BASE64_META=$(cat ./mock-data/nft-data_vip.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=11,12,13,14,15,16,17,18,19,20
BASE64_META=$(cat ./mock-data/nft-data_speaker.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=21,22,23,24,25,26,27,28,29,30
BASE64_META=$(cat ./mock-data/nft-data_investor.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=31,32,33,34,35,36,37,38,39,40
BASE64_META=$(cat ./mock-data/nft-data_partner.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=41,42,43,44,45,46,47,48,49,50
BASE64_META=$(cat ./mock-data/nft-data_media.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=51,52,53,54,55,56,57,58,59,60
BASE64_META=$(cat ./mock-data/nft-data_exhibitor.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

token_id=61,62,63,64,65,66,67,68,69,70
BASE64_META=$(cat ./mock-data/nft-data_general.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} ${BASE64_META} --from $KEY_NAME --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
