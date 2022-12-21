RPC_ENDPOINT=http://localhost:26657
CHAIN_ID=six
read -p "Enter Schema Code: " schema_code 
if [ -z "$schema_code" ]; then
    schema_code=six-protocol.test_v071_2
fi

for i in {0..6}
do
    echo "Mockup Token ${i}"
    BASE64_META=`cat nft-data.json | sed "s/TOKENID/${i}/g"  | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n'`
    sixd tx nftmngr create-metadata "${schema_code}" ${i} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        ${BASE64_META} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
done

# Test case 1: Action with non-required param
echo "Action with non-required param (Test case 1)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 check_in ${schema_code}_tk0_ci "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 2: Action that use utils function
echo "perform-action-by-nftadmin token 0 - start_mission (Test case 2)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 start_mission ${schema_code}_tk0_sm "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}  | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 3: Action with required param
echo "perform-action-by-nftadmin token 0 - test_param (Test case 3)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_param ${schema_code}_tk0_burn "[{\"name\":\"points\",\"value\":\"7\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}  | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 4: Using meta.{utils} as condition 
echo "perform-action-by-nftadmin token 0 - Using meta.{utils} as condition to Pass (Test case 4)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 claim_dec ${schema_code}_tk0_dec "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test cast 5: Using meta.{utils} as condition
echo "perform-action-by-nftadmin token 0 - Using meta.{utils} as condition to Fail (Test case 5)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 claim_jan ${schema_code}_tk0_jan "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "🛑 failed"
else
    echo "✅ success"
fi

# Test case 6: To disable action
echo "perform-action-by-nftadmin token 0 - disable action (Test case 6)"
if sixd tx nftmngr toggle-action ${schema_code} test_disable true --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 7: Action to disable function
echo "perform-action-by-nftadmin token 0 - Action to disable function (Test case 7)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_disable ${schema_code}_tk0_test_disable_action "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "🛑 failed"
else
    echo "✅ success"
fi

# Test case 8: Perform action that locate after disabled action
echo "perform-action-by-nftadmin token 0 - Perform action that locate after disabled action (Test case 8)"
for i in {1..4}
do
    if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 disable_consequent${i} ${schema_code}_tk0_test_disable_bug${i} "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
        echo "✅ success"
    else
        echo "🛑 failed"
    fi
done

# Test case 9: Action with invalid param
echo "perform-action-by-nftadmin token 0 - invalid param (Test case 9)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_param ${schema_code}_tk0_fail_param "[{\"name\":\"boints\",\"value\":\"7\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}  | grep -q 'msg_index: 0'; then
    echo "🛑 failed"
else
    echo "✅ success"
fi


# Test case 10: Transfer point to another token_id
echo "perform-action-by-nftadmin token 0 - test_transfer (Test case 10)"
for i in {1..6}
do
    if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_transfer ${schema_code}_tk0_tf0-${i} "[{\"name\":\"points\",\"value\":\"10\"},{\"name\":\"token_id\",\"value\":\"${i}\"}]" \
        --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
        echo "✅ success"
    else
        echo "🛑 failed"
    fi
done

for i in {0..6}
do
_POINT=$(sixd query nftmngr show-nft-data ${schema_code} ${i} --output json | jq '.nftData.onchain_attributes[] | select(.name=="points").number_attribute_value.value')
POINT=$(echo $_POINT | sed 's/\"//g')
# if id = 0 point is 333 return success else 10 return success
if [ $i -eq 0 ]; then
  if [ "$POINT" == "333" ]; then
    echo "Success"
  else
    echo "Fail $POINT"
  fi
else
  if [ "$POINT" == "10" ]; then
    echo "Success"
  else
    echo "Fail $POINT"
  fi
fi

done

# Test case 11: meta.SetDisplayArribute
echo "perform-action-by-nftadmin token 0 - meta.SetDisplayArribut (Test case 11)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_hidden ${schema_code}_tk0_th "[{\"name\":\"attribute_name\",\"value\":\"hidden_tested\"},{\"name\":\"show\",\"value\":\"true\"}]" \
    --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 12: Default value of param when input of value unmet
echo "perform-action-by-nftadmin token 0 - test_param (Test case 12)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_param ${schema_code}_tk0_burn_fail "[{\"name\":\"points\",\"value\":\"\"}]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}  | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

# Test case 13: Perform start_mission once more (after checked_in = true)
echo "Perform start_mission once more (after checked_in = true)"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 check_in ${schema_code}_tk0_ci_willfail "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "🛑 failed"
else
    echo "✅ success"
fi

# Test All Utils
echo "Test All Utils"
echo "Test GetBlockHeigh"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_GetBlockHeight ${schema_code}_tk0_GetBlockHeight "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

echo "Test test_utils_GetUTCBlockTimestamp"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_GetUTCBlockTimestamp ${schema_code}test_utils_GetUTCBlockTimestamp "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

echo "Test test_utils_GetBlockTimestampByZone"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_GetBlockTimestampByZone ${schema_code}test_utils_GetBlockTimestampByZone "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

echo "Test test_utils_GetLocalBlockTimestamp"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_GetLocalBlockTimestamp ${schema_code}test_utils_GetLocalBlockTimestamp "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

echo "Test test_utils_BlockTimeBeforeByZone"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_BlockTimeBeforeByZone ${schema_code}test_utils_BlockTimeBeforeByZone "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi

echo "Test test_utils_BlockTimeAfterByZone"
if sixd tx nftmngr perform-action-by-nftadmin ${schema_code} 0 test_utils_BlockTimeAfterByZone ${schema_code}test_utils_BlockTimeAfterByZone "[]" --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ success"
else
    echo "🛑 failed"
fi