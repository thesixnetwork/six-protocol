token_id=$1
# if token_id is not provided, ask for it
if [ -z "$token_id" ]; then
    read -p "Enter Token ID: " token_id
fi
uuid=$(uuidgen)
sixd tx nftmngr perform-action-by-nftadmin techsauce.mocking3 ${token_id} \
    check_in single_${uuid} '[]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --node http://localhost:26657 --chain-id testnet

uuid=$(uuidgen)
sixd tx nftmngr perform-multi-token-action techsauce.mocking3 ${token_id} \
    attend_exhibition,attend_exhibition,attend_exhibition multi_${uuid} \
    '[[{"name":"exhibition","value":"exhibition_a"}],[{"name":"exhibition","value":"exhibition_b"}],[{"name":"exhibition","value":"exhibition_c"}]]' \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id testnet --node http://localhost:26657
