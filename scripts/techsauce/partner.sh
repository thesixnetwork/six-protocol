token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]
then
    read -p "Enter Token ID: " token_id
fi
sixd tx nftmngr perform-action-by-nftadmin techsauce.eventname ${token_id} \
    check_in check_intechsauce.eventname_${token_id} '[]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --node http://localhost:26657 --chain-id testnet

sixd tx nftmngr perform-multi-token-action techsauce.eventname ${token_id} \
    attend_partner,attend_partner,attend_partner,attend_partner attend_parter_${token_id} \
    '[[{"name":"partner","value":"partner_1"}],[{"name":"partner","value":"partner_2"}],[{"name":"partner","value":"partner_3"}],[{"name":"partner","value":"partner_4"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet --node http://localhost:26657