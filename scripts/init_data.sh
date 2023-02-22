grantOracle() {
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
}

RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=$1
BASE64_SCHEMA=$(cat nft-schema.json | base64 | tr -d '\n')
if [ -z "$CHAIN_ID" ]; then
    CHAIN_ID=testnet
fi

sixd tx nftadmin grant-permission oracle_admin $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftoracle set-minimum-confirmation 1 --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix
sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}

grantOracle $(sixd keys show oracle1 -a --keyring-backend test)
grantOracle $(sixd keys show oracle2 -a --keyring-backend test)
grantOracle $(sixd keys show oracle3 -a --keyring-backend test)
grantOracle $(sixd keys show oracle4 -a --keyring-backend test)

sixd q nftadmin show-authorization \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}