grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y \
        --node ${RPC_ENDPOINT}
}

RPC_ENDPOINT=http://localhost:26657
BASE64_SCHEMA=`cat nft-schema.json | base64 | tr -d '\n'`

sixd tx nftadmin grant-permission oracle_admin $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id testnet
sixd tx nftoracle set-minimum-confirmation 1 --from super-admin --chain-id sixnft -y --node ${RPC_ENDPOINT} --chain-id testnet
sixd tx nftmngr create-nft-schema --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet \
    --node ${RPC_ENDPOINT} \
    ${BASE64_SCHEMA}

# grantOracle 6nft1clup6q6ucfdp4tg2r6zv82eu2m8xdegsjg3sk7
# grantOracle 6nft1mvw6jk5wuhxmyd0edgj9d3d0nr5amrmxsejcc7
# grantOracle 6nft1d5auhw4hfg49fmzfngpaz98vn4f59yazu6j72s
# grantOracle 6nft1kl7qw9jw0zatph3jc7mdaf3xd5p7aw4edy8svq

# sixd q nftadmin show-authorization \
#     --node ${RPC_ENDPOINT}