RPC_ENDPOINT=$1
CHAIN_ID=$2
key=$3


echo ""
echo "############### TESTING NFT ORACLE MODULE (OWNER REQUEST) ###############" 
echo ""

TOTAL=0
PASSED=0
array_owner_request=()


action_list=("start_mission" "test_read_nft" "test_split" "test_lowercase" "test_uppercase" "test_hidden" "transfer")
params_list=("[]" "[]" "[]" "[]" "[]" "[{\"name\":\"attribute_name\",\"value\":\"hide_fail\"},{\"name\":\"show\",\"value\":\"false\"}]" "[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"0\"}]")
#read -p "Press enter to continue"
for i in "${!action_list[@]}"; do
    action=${action_list[$i]}
    params=${params_list[$i]}
    echo "Perform Action: $action"
    BASE64JSON=`cat ./mock-data/action-param.json | sed "s/ACTION/${action}/g" | sed "s/TOKEN_ID/2/g" | sed "s/SCHEMA_CODE/six-protocol.develop_v220/g" | sed "s/REFID/six-protocol.develop_v220_TK2_${action}/g" | sed "s/\"PARAMS\"/${params}/g" | sed "s/ONBEHALFOF/""/g"`
    BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
    MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | ./evmsign ./.secret`
    BASE64_ACTION_SIG=`cat ./mock-data/action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`
    #read -p "Press enter to continue"
    id=$(sixd tx nftoracle create-action-request ethereum ${BASE64_ACTION_SIG} 4 --from alice --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y | grep 'raw_log' | sed -n 's/.*"action_request_id","value":"\([0-9]*\)".*/\1/p')
    if [ -n "$id" ]; then
        echo "✅ create-action-request success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
        array_owner_request+=($id)
    else
        echo "🛑 create-action-request failed 🖕🖕🖕🖕🖕🖕🖕🖕🖕🖕🖕"
    fi
done

oracles=(oracle1 oracle2 oracle3 oracle4)
BASE64_ORIGINDATA=`cat ./mock-data/nft-origin-data.json | base64 | tr -d '\n'`
for i in ${array_owner_request[@]}
do
    for j in ${oracles[@]}
    do
        echo "Confirm action request $i"
        if sixd tx nftoracle submit-action-response $i ${BASE64_ORIGINDATA} --from $j --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y | grep -q 'msg_index: 0'; then
            echo "✅ confirm-action-request $i completed"
            TOTAL=$(($TOTAL+1))
            PASSED=$(($PASSED+1))
        else
            echo "🛑 confirm-action-request $i failed "
        fi
    done
done

echo "______________________________________________________________________________________________"
echo ""
echo "############### TESTING SET ACTION SIGNER ###############" 
echo ""
BASE64JSON=`cat ./mock-data/set-signer.json`
BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | ./evmsign ./.secret`
BASE64_VERIFY_SIG=`cat ./mock-data/verify-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`
sixd tx nftoracle create-action-signer ${BASE64_VERIFY_SIG} --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}

echo "______________________________________________________________________________________________"
echo ""
echo "############### TESTING NFT ORACLE MODULE (ACTOR REQUEST) ###############" 
echo ""

array_actor_request=()
action_list=("start_mission" "test_read_nft" "test_split" "test_lowercase" "test_uppercase" "test_hidden" "transfer")
params_list=("[]" "[]" "[]" "[]" "[]" "[{\"name\":\"attribute_name\",\"value\":\"hide_fail\"},{\"name\":\"show\",\"value\":\"false\"}]" "[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"0\"}]")
#read -p "Press enter to continue"
for i in "${!action_list[@]}"; do
    action=${action_list[$i]}
    params=${params_list[$i]}
    echo "Perform Action: $action"
    BASE64JSON=`cat ./mock-data/action-param.json | sed "s/ACTION/${action}/g" | sed "s/TOKEN_ID/3/g" | sed "s/SCHEMA_CODE/six-protocol.develop_v220/g" | sed "s/REFID/six-protocol.develop_v220_TK3_${action}/g" | sed "s/\"PARAMS\"/${params}/g" | sed "s/ONBEHALFOF/0xb7c2468b9481CbDfD029998d6bA98c55072d932e/g"`
    BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
    MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | ./evmsign ./.secret2`
    BASE64_ACTION_SIG=`cat ./mock-data/action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`
    #read -p "Press enter to continue"
    id=$(sixd tx nftoracle create-action-request ethereum ${BASE64_ACTION_SIG} 4 --from alice --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y | grep 'raw_log' | sed -n 's/.*"action_request_id","value":"\([0-9]*\)".*/\1/p')
    if [ -n "$id" ]; then
        echo "✅ create-action-request success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
        array_actor_request+=($id)
    else
        echo "🛑 create-action-request failed 🖕🖕🖕🖕🖕🖕🖕🖕🖕🖕🖕"
    fi
done

BASE64_ORIGINDATA=`cat ./mock-data/nft-origin-data.json | base64 | tr -d '\n'`
for i in ${array_actor_request[@]}
do
    for j in ${oracles[@]}
    do
        echo "Confirm action request $i"
        if sixd tx nftoracle submit-action-response $i ${BASE64_ORIGINDATA} --from $j --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y | grep -q 'msg_index: 0'; then
            echo "✅ confirm-action-request $i completed"
            TOTAL=$(($TOTAL+1))
            PASSED=$(($PASSED+1))
        else
            echo "🛑 confirm-action-request $i failed "
        fi
    done
done



echo "Test: Passed $PASSED out of $TOTAL tests"