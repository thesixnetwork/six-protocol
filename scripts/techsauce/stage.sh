token_id=$1
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
uuid=$(uuidgen)
sixd tx nftmngr perform-multi-token-action techsauce.mocking3 ${token_id} \
    attend_stage,attend_stage,attend_stage,attend_stage,attend_stage attend_stage_${uuid} \
    '[[{"name":"stage","value":"stage_1"}],[{"name":"stage","value":"stage_2"}],[{"name":"stage","value":"stage_3"}],[{"name":"stage","value":"stage_4"}],[{"name":"stage","value":"stage_5"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet
