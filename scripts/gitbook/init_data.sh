RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=six
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}