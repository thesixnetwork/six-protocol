sixd tx nftmngr perform-multi-token-action techsauce.eventname 1 \
    attend_stage,attend_stage,attend_stage,attend_stage,attend_stage attend_stage_token_1 \
    '[[{"name":"stage","value":"stage_1"}],[{"name":"stage","value":"stage_2"}],[{"name":"stage","value":"stage_3"}],[{"name":"stage","value":"stage_4"}],[{"name":"stage","value":"stage_5"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet