token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
action=$2
# if action is not provided, ask for it
if [ -z "$action" ]; then
    read -p "Enter Action: " action
fi
uuid=$(uuidgen)
sixd tx nftmngr perform-action-by-nftadmin techsauce.mocking3 ${token_id} \
    ${action} ${uuid} '[]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --node http://localhost:26657 --chain-id testnet
