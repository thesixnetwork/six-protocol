sixd tx nftmngr perform-multi-token-action techsauce.eventname 1 \
    attend_partner,attend_partner,attend_partner,attend_partner attend_parter_token_1 \
    '[[{"name":"partner","value":"partner_1"}],[{"name":"partner","value":"partner_2"}],[{"name":"partner","value":"partner_3"}],[{"name":"partner","value":"partner_4"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet --node http://localhost:26657