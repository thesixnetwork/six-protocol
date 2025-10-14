grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 2usix -y \
        --node ${RPC_ENDPOINT} --chain-id testnet
}

RPC_ENDPOINT=http://localhost:26657
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

sixd tx nftadmin grant-permission oracle_admin $(sixd keys show alice -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id testnet --gas-prices 2usix
sixd tx nftoracle set-minimum-confirmation 1 --from super-admin -y --node ${RPC_ENDPOINT} --chain-id testnet --gas-prices 2usix
sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 2usix -y --chain-id testnet \
    --node ${RPC_ENDPOINT} \
    ${BASE64_SCHEMA}

grantOracle $(sixd keys show oracle1 -a)
grantOracle $(sixd keys show oracle2 -a)
grantOracle $(sixd keys show oracle3 -a)
grantOracle $(sixd keys show oracle4 -a)

# sixd q nftadmin show-authorization \
#     --node ${RPC_ENDPOINT}