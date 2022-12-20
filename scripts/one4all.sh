default_schema_code=$1
read -p "Enter Schema Code: " schema_code 
if [ -z "$schema_code" ]; then
    schema_code=$default_schema_code
fi

for i in {0..9}
do
    echo "Mockup Token ${i}"
    BASE64_META=`cat nft-data.json | sed "s/TOKENID/${i}/g"  | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n'`
    sixd tx nftmngr create-metadata "${schema_code}" ${i} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        ${BASE64_META} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
    
    # echo "perform-action-by-nftadmin ${i} - start_mission"
    # sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${i} start_mission tk_${i}_sm "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    #     --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
    
    # echo "perform-action-by-nftadmin ${i} - claim_dec"
    # sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${i} claim_dec tk_${i}_dec "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    #     --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

    # echo "perform-action-by-nftadmin ${i} - test_param"
    # sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${i} test_param tk_${i}_tp "[{\"name\":\"points\",\"value\":\"10\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    #     --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

    # echo "perform-action-by-nftadmin ${i} - test_hidden"
    # sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${i} test_param tk_${i}_th "[{\"name\":\"attribute_name\",\"value\":\"missions_completed\"},{\"name\":\"show\",\"value\":\"true\"}]" \
    #     --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    #     --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

    # echo "perform-action-by-nftadmin ${i} - test_transfer"
    # sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${i} test_transfer tk_${i}_tf0-1 "[{\"name\":\"recipient\",\"value\":\"$(sixd keys show bob -a)\"}]" \
    #     --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    #     --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
done

echo "perform-action-by-nftadmin token 0 - start_mission"
sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 start_mission tk0_sm "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

echo "perform-action-by-nftadmin token 0 - claim_dec"
sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 claim_dec tk0_dec "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

 echo "perform-action-by-nftadmin token 0 - test_param"
sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_param tk0_tp "[{\"name\":\"points\",\"value\":\"10\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

echo "perform-action-by-nftadmin token 0 - test_hidden"
sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_param tk0_th "[{\"name\":\"attribute_name\",\"value\":\"missions_completed\"},{\"name\":\"show\",\"value\":\"true\"}]" \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

for i in {1..9}
do
    echo "perform-action-by-nftadmin token 0 - test_transfer"
    sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_transfer tk0_tf0-${i} "[{\"name\":\"points\",\"value\":\"10\"},{\"name\":\"token_id\",\"value\":\"${i}\"}]" \
        --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
done

