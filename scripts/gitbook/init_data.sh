RPC_ENDPOINT=https://rpc1.fivenet.sixprotocol.net:443
CHAIN_ID=fivenet
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} \
    --node ${RPC_ENDPOINT} ${BASE64_SCHEMA}