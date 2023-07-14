token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
uuid=$(uuidgen)
sixd tx nftmngr perform-multi-token-action techsauce.mocking3 ${token_id} \
    attend_zone,attend_zone,attend_zone,attend_zone multi_${uuid} \
    '[[{"name":"zone","value":"bussiness_zone"}],[{"name":"zone","value":"relax_zone"}],[{"name":"zone","value":"exp_zone"}],[{"name":"zone","value":"cvc_zone"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet --node http://localhost:26657
