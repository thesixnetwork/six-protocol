token_id=$1
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
uuid=$(uuidgen)
sixd tx nftmngr perform-multi-token-action techsauce.eventname ${token_id} \
    attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage,attend_stage attend_stage_${uuid} \
    '[[{"name":"stage","value":"stage_1"}],[{"name":"stage","value":"stage_2"}],[{"name":"stage","value":"stage_3"}],[{"name":"stage","value":"stage_4"}],[{"name":"stage","value":"stage_5"}],[{"name":"stage","value":"stage_6"}],[{"name":"stage","value":"stage_7"}],[{"name":"stage","value":"stage_8"}],[{"name":"stage","value":"stage_9"}],[{"name":"stage","value":"stage_10"}],[{"name":"stage","value":"stage_11"}],[{"name":"stage","value":"stage_12"}],[{"name":"stage","value":"stage_13"}],[{"name":"stage","value":"stage_14"}],[{"name":"stage","value":"stage_15"}],[{"name":"stage","value":"stage_16"}],[{"name":"stage","value":"stage_17"}],[{"name":"stage","value":"stage_18"}],[{"name":"stage","value":"stage_19"}],[{"name":"stage","value":"stage_20"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet
