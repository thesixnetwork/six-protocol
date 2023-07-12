token_id=$1
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
stock=$2
if [ -z "$stock" ]; then
    read -p "Enter Stock: " stock
fi
uuid=$(uuidgen)
params='[{"name":"stock","value":"'$stock'"}]'
sixd tx nftmngr perform-action-by-nftadmin techsauce.eventname ${token_id} \
    redeem_prize ${uuid} ${params} \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet
