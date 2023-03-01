echo "Deploy Schema"
read -p "Enter test platform: [local(defaule), docker, fivenet, sixnet] " _PLATFORM
read -p "Enter key: [alice] " key
PLATFORM=$(echo "$_PLATFORM" | tr '[:upper:]' '[:lower:]')
# if platform is not set, set it to local
if [ -z "$PLATFORM" ]; then
    PLATFORM="local"
fi

if [ -z "$key" ] && [ "$PLATFORM" == "local" ] || [ "$PLATFORM" == "docker" ]; then
    key="alice"
fi

# switch case
case $PLATFORM in
"local")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID="testnet"
    ;;
"docker")
    RPC_ENDPOINT="http://localhost:26657"
    CHAIN_ID="sixnet"
    ;;
"fivenet")
    RPC_ENDPOINT="https://rpc1.fivenet.sixprotocol.net:443"
    CHAIN_ID="fivenet"
    ;;
"sixnet")
    RPC_ENDPOINT="https://sixnet-rpc.sixprotocol.net:443"
    CHAIN_ID="sixnet"
    ;;
*)
    echo "Error: unsupported PLATFORM '$PLATFORM'" >&2
    exit 1
    ;;
esac

grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'
}

## if plat from local
if [ "$PLATFORM" == "local" ]; then
    echo "INIT LOCAL ENV"
    sixd tx nftadmin grant-permission oracle_admin $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}  --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix | grep -q 'msg_index: 0'
    sixd tx nftoracle set-minimum-confirmation 4 --from super-admin  -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}  --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix | grep -q 'msg_index: 0'
    sixd tx nftadmin grant-permission admin_signer_config $(sixd keys show super-admin -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix | grep -q 'msg_index: 0'
    sixd tx nftoracle create-action-signer-config baobab 0x45AaF440FbA71E52cCb096D66230A7FaAd9b31ac --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix | grep -q 'msg_index: 0'
    grantOracle $(sixd keys show oracle1 -a --keyring-backend test)
    grantOracle $(sixd keys show oracle2 -a --keyring-backend test)
    grantOracle $(sixd keys show oracle3 -a --keyring-backend test)
    grantOracle $(sixd keys show oracle4 -a --keyring-backend test)
fi

TOTAL=0
PASSED=0

echo ""
echo "############### TESTING NFT MANAGER MODULE ###############" 
echo ""

#read -p "Press enter to continue"
BASE64_SCHEMA=$(cat ./mock-data/nft-schema.json | base64 | tr -d '\n')
if sixd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from ${key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ create-nft-schema"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ create-nft-schema 😭"
    TOTAL=$((TOTAL+1))
fi

echo "Create Multi NFT"
#read -p "Press enter to continue"
BASE64_META=$(cat ./mock-data/nft-data.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/six-protocol.develop_v220/g" | base64 | tr -d '\n')
if sixd tx nftmngr create-multi-metadata six-protocol.develop_v220 0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30 --from ${key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
    ${BASE64_META} --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ create-multi-metadata"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ create-multi-metadata 😭" 
    TOTAL=$((TOTAL+1))
fi


action_list=("start_mission" "test_read_nft" "test_split" "test_lowercase" "test_uppercase" "test_hidden" "transfer")
params_list=("[]" "[]" "[]" "[]" "[]" "[{\"name\":\"attribute_name\",\"value\":\"hide_pass\"},{\"name\":\"show\",\"value\":\"true\"}]" "[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"0\"}]")
for i in "${!action_list[@]}"; do
    action=${action_list[$i]}
    params=${params_list[$i]}
    echo "Perform Action: $action"
    #read -p "Press enter to continue"
    if sixd tx nftmngr perform-action-by-nftadmin six-protocol.develop_v220 1 $action six-protocol.develop_v220_TK1_${action} $params --from ${key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
        echo "✅ $action"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "❌ $action 😭"
        TOTAL=$((TOTAL+1))
    fi
done

# echo "Test Hide Fail"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-action-by-nftadmin six-protocol.develop_v220 1 test_hidden hidden_fail "[{\"name\":\"attribute_name\",\"value\":\"hide_fail\"},{\"name\":\"show\",\"value\":\"false\"}]" \
    --from ${key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}  2>&1 | grep -q 'Attribute overriding is not allowed'; then
    echo "✅ test_hidden"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ test_hidden 😭"
    TOTAL=$((TOTAL+1))
fi
echo "______________________________________________________________________________________________"

# echo "Test Multi Action(one token one action) - Case param required = 0"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 10 start_mission oneonezero "[[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token one action) - Case param required = 0"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token one action) - Case param required = 0 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token one action) - Case param required > 0"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 10 transfer oneonenonzero "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token one action) - Case param required > 0"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token one action) - Case param required > 0 😭"
    TOTAL=$((TOTAL+1))
fi

echo "______________________________________________________________________________________________"

# echo "Test Multi Action(multi token one action) - Case param required = 0"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 3,4,5 start_mission manyonezero "[[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token one action) - Case param required = 0"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token one action) - Case param required = 0 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token one action) - Case param required > 0"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 3,4,5 transfer manyonenonzero "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token one action) - Case param required > 0"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token one action) - Case param required > 0 😭"
    TOTAL=$((TOTAL+1))
fi

echo "______________________________________________________________________________________________"

# echo "Test Multi Action(one token multi action) - Case Fail"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 11 start_mission,test_read_nft,test_hidden onemanyfail "[[],[],[],[],[],[{\"name\":\"attribute_name\",\"value\":\"hide_pass\"},{\"name\":\"show\",\"value\":\"true\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} 2>&1 | grep -q 'failed to execute message'; then
    echo "✅ Test Multi Action(one token multi action) - Case Fail"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case Fail 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case Pass"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 11 start_mission,test_read_nft,test_split,test_lowercase,test_uppercase,test_hidden onemanypass "[[],[],[],[],[],[{\"name\":\"attribute_name\",\"value\":\"hide_pass\"},{\"name\":\"show\",\"value\":\"true\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case Pass"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case Pass 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 202"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 11 transfer,transfer onemany202 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 202"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 202 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 211"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 12 start_mission,transfer onemany211 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 211"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 211 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 220"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 13 start_mission,test_read_nft onemany220 "[[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 220"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 220 😭"
    TOTAL=$((TOTAL+1))
fi

echo "Test Multi Action(one token multi action) - Case 303"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 13 transfer,transfer,transfer onemany303 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 303"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 303 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 312"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 14 start_mission,transfer,transfer onemany312 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 312"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 312 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 321"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 15 start_mission,test_read_nft,transfer onemany321 "[[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 321"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 321 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 330"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 16 start_mission,test_read_nft,test_split onemany330 "[[],[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 330"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 330 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 404"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 16 transfer,transfer,transfer,transfer onemany404 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 404"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 404 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 413"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 17 start_mission,transfer,transfer,transfer onemany413 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}  | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 413"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 413 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 422"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 18 start_mission,test_read_nft,transfer,transfer onemany422 "[[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 422"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 422 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 431"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 19 start_mission,test_read_nft,test_split,transfer onemany431 "[[],[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 431"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 431 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(one token multi action) - Case 440"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 20 start_mission,test_read_nft,test_split,test_lowercase onemany440 "[[],[],[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(one token multi action) - Case 440"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(one token multi action) - Case 440 😭"
    TOTAL=$((TOTAL+1))
fi

echo "______________________________________________________________________________________________"

# echo "Test Multi Action(multi token multi action) - Case 202"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 20,20 transfer,transfer manymany202 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 202"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 202 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 211"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 21,21 start_mission,transfer manymany211 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 211"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 211 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 220"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 22,22 start_mission,test_read_nft manymany220 "[[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT}  --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 220"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 220 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 303"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 22,22,22 transfer,transfer,transfer manymany303 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 303"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 303 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 312"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 23,23,23 start_mission,transfer,transfer manymany312 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 312"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 312 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 321"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 24,24,24 start_mission,test_read_nft,transfer manymany321 "[[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 321"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 321 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 330"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 25,25,25 start_mission,test_read_nft,test_split manymany330 "[[],[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then 
    echo "✅ Test Multi Action(multi token multi action) - Case 330"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 330 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 404"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 25,25,25,25 transfer,transfer,transfer,transfer manymany404 "[[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 404"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 404 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 413"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 26,26,26,26 start_mission,transfer,transfer,transfer manymany413 "[[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 413"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 413 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 422"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 27,27,27,27 start_mission,test_read_nft,transfer,transfer manymany422 "[[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 422"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 422 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 431"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 28,28,28,28 start_mission,test_read_nft,test_split,transfer manymany431 "[[],[],[],[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"1\"}]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 431"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 431 😭"
    TOTAL=$((TOTAL+1))
fi

# echo "Test Multi Action(multi token multi action) - Case 440"
#read -p "Press enter to continue"
if sixd tx nftmngr perform-multi-token-action six-protocol.develop_v220 29 start_mission,test_read_nft,test_split,test_lowercase manymany440 "[[],[],[],[]]" --from ${key} \
    --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Test Multi Action(multi token multi action) - Case 440"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
else
    echo "❌ Test Multi Action(multi token multi action) - Case 440 😭"
    TOTAL=$((TOTAL+1))
fi

echo ""
echo "############### TESTING NFT ORACLE MODULE ###############" 
echo ""

array_owner_request=()

action_list=("start_mission" "test_read_nft" "test_split" "test_lowercase" "test_uppercase" "test_hidden" "transfer")
params_list=("[]" "[]" "[]" "[]" "[]" "[{\"name\":\"attribute_name\",\"value\":\"hide_fail\"},{\"name\":\"show\",\"value\":\"false\"}]" "[{\"name\":\"points\",\"value\":\"20\"},{\"name\":\"token_id\",\"value\":\"0\"}]")
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
        echo "🛑 create-action-request failed 😭😭😭😭😭😭😭😭😭😭"
    fi
done

#read -p "Press enter to continue"
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

#read -p "Press enter to continue"
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
        echo "🛑 create-action-request failed 😭😭😭😭😭😭😭😭😭😭"
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

echo "______________________________________________________________________________________________"
echo "Test: Passed $PASSED out of $TOTAL tests"