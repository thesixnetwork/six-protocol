token_id=$1
if [ -z "$token_id" ]
then
    read -p "Enter Token ID: " token_id
fi
sixd tx nftmngr perform-multi-token-action techsauce.eventname ${token_id} \
    attend_session,attend_session,attend_session,attend_session,attend_session,attend_session,attend_session,attend_session,attend_session,attend_session \
    session_attend_${token_id} \
    '[[{"name":"session","value":"session_1"}],[{"name":"session","value":"session_2"}],[{"name":"session","value":"session_3"}],[{"name":"session","value":"session_4"}],[{"name":"session","value":"session_5"}],[{"name":"session","value":"session_6"}],[{"name":"session","value":"session_7"}],[{"name":"session","value":"session_8"}],[{"name":"session","value":"session_9"}],[{"name":"session","value":"session_10"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet --node http://localhost:26657 -o json