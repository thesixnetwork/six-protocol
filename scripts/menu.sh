EVMSIGN=./evmsign
default_schema_code=$1
PLATFORM=$2
if [ -z "$PLATFORM" ]; then
    read -p "Enter test platform: [local(defaule), docker, fivenet, sixnet] " _PLATFORM
    PLATFORM=$(echo "$_PLATFORM" | tr '[:upper:]' '[:lower:]')
    # if platform is not set, set it to local
    if [ -z "$PLATFORM" ]; then
        PLATFORM="local"
    fi
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

timestamp=$(date -u +"%Y-%m-%dT%H:%M:%S.000z")
echo "#############################################"
echo "##                                         ##"
echo "##  Welcome to the menu script             ##"
echo "##                                         ##"
echo "##  Please select an option                ##"
echo "##                                         ##"
echo "##  1. Show Schema                         ##"
echo "##  2. Show NFTs                           ##"
echo "##  3. Mockup Token                        ##"
echo "##  4. Mockup Multi Token                  ##"
echo "##  5. Do Action                           ##"
echo "##  6. Do Action Multi Tokens              ##"
echo "##  7. Set NFT Attribute                   ##"
echo "##  8. Oracle - Create Mint Request        ##"
echo "##  9. Oracle - Get Mint Request           ##"
echo "##  10. Oracle - Submit Mint Response      ##"
echo "##  11. Oracle - Create Action Request     ##"
echo "##  12. Oracle - Get Action Request        ##"
echo "##  13. Oracle - Submit Action Response    ##"
echo "##  14. Oracle - Create Verfify Request    ##"
echo "##  15. Oracle - Get Verify Request        ##"
echo "##  16. Oracle - Submit Verify Response    ##"
echo "##  17. Add Attribute                      ##"
echo "##  18. Add Action                         ##"
echo "##  19. Oracle - Set Signer                ##"
echo "##  20. Show ActionSigner By Address       ##"
echo "##  21. Oracle - Action Request By Signer  ##"
echo "##  22. Oracle - Request Sync Signer       ##"
echo "##  23. Oracle - Submit Sync Signer        ##"
echo "##  Your choice:                           ##"
echo "##                                         ##"
echo "#############################################"
read -p "Your choice: " choice
case $choice in
    1) echo "Showing Schema"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixd q nftmngr show-nft-schema ${schema_code} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    2) echo "Showing NFT"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixd q nftmngr show-nft-data ${schema_code} ${token_id} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    3) echo "Mockup Token"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_META=`cat nft-data.json | sed "s/TOKENID/${token_id}/g"  | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n'`
        sixd tx nftmngr create-metadata "${schema_code}" ${token_id} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            ${BASE64_META} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    4) echo "Mockup Multi Token"
        read -p "Enter Schema Code: " schema_code
        read -p "Enter Token IDs: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_META=$(cat nft-data.json | sed "s/TOKENID/MULTIMINT/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n')
        sixd tx nftmngr create-multi-metadata ${schema_code} ${token_id} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            ${BASE64_META} --chain-id ${CHAIN_ID}
        ;;
    5) echo "Do Action"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Enter Action: " action
        read -p "Enter Ref ID: " ref_id
        read -p "Enter Required Params: " num_params
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        # check if required_params is empty
        if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
            required_params="[]"
        else
            for ((i=1; i<=num_params; i++)); do
                read -p "Enter name of param $i: " param_name
                read -p "Enter value of >> $param_name << : " param_value
                required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
            done
            required_params=$(echo ${required_params[@]} | tr ' ' ',')
            required_params="["$required_params"]"
            echo $required_params
        fi

        sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${token_id} ${action} ${ref_id} ${required_params} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -o json | grep -q 'Error:'
        ;;
    6) echo "Do Action Multi token"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token IDs: " token_id
        read -p "Enter Action: " action
        read -p "Enter Ref ID: " ref_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        # array from action
        arrAction=(${action//,/ })
        all_required_params=()
        # iterate through array using for loop
        for i in "${arrAction[@]}"
        do
            echo "$i"
            read -p "Enter number of Required Params of $i: " num_params
            required_params=()
            # check if required_params is empty
            if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
                required_params="[]"
            else
                for ((j=1; j<=num_params; j++)); do
                    read -p "Enter name of param $j: " param_name
                    read -p "Enter value of >> $param_name << : " param_value
                    required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
                done
                required_params="["$(echo ${required_params[@]} | tr ' ' ',')"]"
                echo $required_params
            fi
            all_required_params+=($required_params)
        done
        all_required_params="["$(echo ${all_required_params[@]} | tr ' ' ',')"]"
        sixd tx nftmngr perform-multi-token-action ${schema_code} ${token_id} ${action} ${ref_id} ${all_required_params} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID}
        ;;
    7) echo "Set NFT Attribute"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Value (attribute_name=N[value]): " value
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi

        ATTRIBUTE_NAME=`echo $value | cut -d'=' -f1`
        ATTRIBUTE_VALUE_STRING=`echo $value | cut -d'=' -f2`
        # get one character from ATTRIBUTE_VALUE
        ATTRIBUTE_VALUE_CHAR=`echo $ATTRIBUTE_VALUE_STRING | cut -c1`
        # get characters between [] from ATTRIBUTE_VALUE_CHAR
        ATTRIBUTE_VALUE_VALUE=`echo $ATTRIBUTE_VALUE_STRING | cut -d'[' -f2 | cut -d']' -f1`

        if [ "$ATTRIBUTE_VALUE_CHAR" = "N" ]; then
            ATTRIBUTE_VALUE_TYPE="number"
            ATTRIBUTE_VALUE_TYPE_VALUE=${ATTRIBUTE_VALUE_VALUE}
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "S" ]; then
            ATTRIBUTE_VALUE_TYPE="string"
            ATTRIBUTE_VALUE_TYPE_VALUE="\"${ATTRIBUTE_VALUE_VALUE}\""
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "B" ]; then
            ATTRIBUTE_VALUE_TYPE="boolean"
            # check if ATTRIBUTE_VALUE_VALUE is true or false
            if [ "$ATTRIBUTE_VALUE_VALUE" = "true" ]; then
                ATTRIBUTE_VALUE_TYPE_VALUE="true"
            elif [ "$ATTRIBUTE_VALUE_VALUE" = "false" ]; then
                ATTRIBUTE_VALUE_TYPE_VALUE="false"
            else
                echo "Invalid boolean value"
                exit 1
            fi
        elif [ "$ATTRIBUTE_VALUE_CHAR" = "F" ]; then
            ATTRIBUTE_VALUE_TYPE="float"
            ATTRIBUTE_VALUE_TYPE_VALUE=${ATTRIBUTE_VALUE_VALUE}
        fi

        BASE64_ATTR=`cat nft-data-test071.json \
            | sed "s/#ATTRIBUTE_NAME#/${ATTRIBUTE_NAME}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE#/${ATTRIBUTE_VALUE_TYPE}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE_VALUE#/${ATTRIBUTE_VALUE_TYPE_VALUE}/g" \
            | base64 | tr -d '\n'`

        echo "BASE64_ATTR: ${BASE64_ATTR}"

        sixd tx nftmngr set-nft-attribute ${schema_code} ${BASE64_ATTR} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    8) echo "Oracle - Create Mint Request"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Require confirmations: " require_confirmations
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixd tx nftoracle create-mint-request ${schema_code} ${token_id} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    9) echo "Oracle - Get Mint Request"
        read -p "Mint Request ID: " mint_request_id 
        sixd q nftoracle show-mint-request ${mint_request_id} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    10) echo "Oracle - Submit Mint Response"
        read -p "Mint Request ID: " mint_request_id
        read -p "Oracle : " oracle_key_name
        BASE64_ORIGINDATA=`cat nft-origin-data.json | base64 | tr -d '\n'`

        sixd tx nftoracle submit-mint-response ${mint_request_id} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    11) echo "Oracle - Create Action Request"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Enter Action: " action
        read -p "Require confirmations: " require_confirmations
        read -p "Reference ID: " reference_id
        read -p "Enter Required Params: " num_params
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        # check if required_params is empty
        if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
            required_params="[]"
        else
            for ((i=1; i<=num_params; i++)); do
                read -p "Enter name of param $i: " param_name
                read -p "Enter value of >> $param_name << : " param_value
                required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
            done
            required_params=$(echo ${required_params[@]} | tr ' ' ',')
            required_params=[$required_params]
            echo $required_params
        fi

        BASE64JSON=`cat action-param.json | sed "s/ACTION/${action}/g" | sed "s/TOKEN_ID/${token_id}/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | sed "s/REFID/${reference_id}/g" | sed "s/\"PARAMS\"/${required_params}/g" | sed "s/ONBEHALFOF/""/g"`

        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_ACTION_SIG=`cat action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        # echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret 1
        # echo  ${BASE64_ACTION_SIG} 
        sixd tx nftoracle create-action-request ethereum ${BASE64_ACTION_SIG} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    12) echo "Oracle - Get Action Request"
        read -p "Action Request ID: " action_request_id 
        sixd q nftoracle show-action-request ${action_request_id} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    13) echo "Oracle - Submit Action Response"
        read -p "Action Request ID: " action_request_id
        read -p "Oracle : " oracle_key_name
        BASE64_ORIGINDATA=`cat nft-origin-data.json | base64 | tr -d '\n'`

        sixd tx nftoracle submit-action-response ${action_request_id} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    14) echo "Oracle - Create Verify Schema Request"
        read -p "Enter Schema Code: " schema_code
        read -p "Require confirmations: " require_confirmations
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi

        BASE64JSON=`cat verify-collection-owner.json`
        # echo "BASE64JSON: ${BASE64JSON}"
        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_VERIFY_SIG=`cat verify-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        sixd tx nftoracle create-verify-collection-owner-request ${schema_code} ${BASE64_VERIFY_SIG} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    15) echo "Oracle - Get Verify Request"
        read -p "Verify Request ID: " verfiry_request_id 
        sixd q nftoracle show-collection-owner-request ${verfiry_request_id} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    16) echo "Oracle - Submit Verify Response"
        read -p "Enter Schema Code: " schema_code
        read -p "Verify Request ID: " verfiry_request_id
        read -p "Oracle : " oracle_key_name
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_ORIGINDATA=`cat verify-collection-owner.json | base64 | tr -d '\n'`

        sixd tx nftoracle submit-verify-collection-owner ${verfiry_request_id} ${schema_code} ${BASE64_ORIGINDATA} --from ${oracle_key_name} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
     17) echo "Add Attribute"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        read -p "Location of attribute (0 or 1): " location
        BASE64_ATTRIBUTE=`cat new-attribute.json | base64 | tr -d '\n'`
        sixd tx nftmngr add-attribute ${schema_code} ${location} ${BASE64_ATTRIBUTE} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
     18) echo "Add Action"
        read -p "Enter Schema Code: " schema_code 
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_ACTION=`cat new-action.json | base64 | tr -d '\n'`
        sixd tx nftmngr add-action ${schema_code} ${BASE64_ACTION} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
     19) echo "Set Signer"
        BASE64JSON=`cat set-signer.json`
        # echo "BASE64JSON: ${BASE64JSON}"
        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_VERIFY_SIG=`cat verify-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        sixd tx nftoracle create-action-signer ${BASE64_VERIFY_SIG} --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    20) echo "Show Action Signer"
        read -p "Enter Signer Address (ETH): " signer_address
        read -p "Enter Owner Address (ETH): " owner_address 
        sixd q nftoracle show-action-signer ${signer_address} ${owner_address} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -o json | jq .
        ;;
    21) echo "Oracle - ActionSigner Action Request"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "Enter Action: " action
        read -p "Enter OnBehalfOf: " on_behalf_of
        read -p "Require confirmations: " require_confirmations
        read -p "Reference ID: " reference_id
        read -p "Enter Required Params: " num_params
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        # check if required_params is empty
        if [[ -z "$num_params" || "$num_params" -eq 0 ]]; then
            required_params="[]"
        else
            for ((i=1; i<=num_params; i++)); do
                read -p "Enter name of param $i: " param_name
                read -p "Enter value of >> $param_name << : " param_value
                required_params+=( "{\"name\":\"$param_name\",\"value\":\"$param_value\"}" )
            done
            required_params=$(echo ${required_params[@]} | tr ' ' ',')
            required_params=[$required_params]
            echo $required_params
        fi

        BASE64JSON=`cat action-param.json | sed "s/ACTION/${action}/g" | sed "s/TOKEN_ID/${token_id}/g" | sed "s/SCHEMA_CODE/${schema_code}/g" | sed "s/REFID/${reference_id}/g" | sed "s/\"PARAMS\"/${required_params}/g" | sed "s/ONBEHALFOF/${on_behalf_of}/g"`
        BASE64_MESSAGE=`echo -n $BASE64JSON | base64 | tr -d '\n'`
        # echo "BASE64_MESSAGE: ${BASE64_MESSAGE}"
        MESSAGE_SIG=`echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret2`
        # echo "MESSAGE_SIG: ${MESSAGE_SIG}"

        BASE64_ACTION_SIG=`cat action-signature.json | sed "s/SIGNATURE/${MESSAGE_SIG}/g" | sed "s/MESSAGE/${BASE64_MESSAGE}/g" | base64 | tr -d '\n'`

        # echo -n ${BASE64_MESSAGE} | $EVMSIGN ./.secret 1
        # echo  ${BASE64_ACTION_SIG} 
        sixd tx nftoracle create-action-request ethereum ${BASE64_ACTION_SIG} ${require_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y 
        ;;
    22) echo "Oracle - Request Sync Signer"
        read -p "Enter Signer Address (ETH): " signer_address
        read -p "Enter Owner Address (ETH): " owner_address 
        read -p "Enter Chain: " chain
        read -p "Enter Required Confirmations: " required_confirmations
        sixd tx nftoracle create-sync-action-signer ${chain} ${signer_address} ${owner_address} ${required_confirmations} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -y
        ;;
    23) echo "Oracle - Submit Sync Signer"
        read -p "Enter Request ID: " request_id
        read -p "Enter Chain: " chain
        read -p "Enter Signer Address (ETH): " signer_address
        read -p "Enter Owner Address (ETH): " owner_address 
        read -p "Enter Expire Epoch (default end of day): " expire_epoch
        read -p "Enter Required Confirmations: " required_confirmations
        if [ -z "$expire_epoch" ]; then
            now=$(date +%s)
            end_of_day=$(( now - now%86400 + 86399))
            expire_epoch=$end_of_day
        fi
        sixd tx nftoracle submit-sync-action-signer ${request_id} ${chain} ${signer_address} ${owner_address} ${expire_epoch} --from oracle4 --chain-id testnet  --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y
        ;;
    *) echo "Invalid choice"
       ;;
esac