token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]
then
    read -p "Enter Token ID: " token_id
fi
sixd tx nftmngr perform-action-by-nftadmin techsauce.eventname ${token_id} \
    transform transform-check_intechsauce.eventname_${token_id} '[]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --node http://localhost:26657 --chain-id testnet