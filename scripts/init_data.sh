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


BASE64_SCHEMA=`cat ./mock-data/nft-schema.json | base64 | tr -d '\n'`

sixd tx nftadmin grant-permission oracle_admin $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftadmin grant-permission admin_signer_config $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftoracle set-minimum-confirmation 1 --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}

sixd tx nftoracle create-action-signer-config baobab 0x45AaF440FbA71E52cCb096D66230A7FaAd9b31ac --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftoracle create-action-signer-config goerli 0x4aEe985A876Deb8413c8ee509a8803bF634A247f --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
grantOracle $(sixd keys show oracle1 -a --keyring-backend test)
grantOracle $(sixd keys show oracle2 -a --keyring-backend test)
grantOracle $(sixd keys show oracle3 -a --keyring-backend test)
grantOracle $(sixd keys show oracle4 -a --keyring-backend test)

sixd q nftadmin show-authorization \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}