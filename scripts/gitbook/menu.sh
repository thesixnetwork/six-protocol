default_schema_code=$1
RPC_ENDPOINT=https://rpc1.fivenet.sixprotocol.net:443
CHAIN_ID=fivenet
echo "#############################################"
echo "##                                         ##"
echo "##  Welcome to the menu script             ##"
echo "##                                         ##"
echo "##  Please select an option                ##"
echo "##                                         ##"
echo "##  1. Show Schema                         ##"
echo "##  2. Create NFT Metadata (mint)          ##"
echo "##  3. Show NFTs                           ##"
echo "##  4. Do Action                           ##"
echo "##  5. Set NFT Attribute                   ##"
echo "##  6. Add Attribute                       ##"
echo "##  7. Add Action                          ##"
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
    2) echo "Mockup Token"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "From (address or key) : " address_key
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_META=`cat nft-data.json | sed "s/TOKENID/${token_id}/g"  | sed "s/SCHEMA_CODE/${schema_code}/g" | base64 | tr -d '\n'`
        sixd tx nftmngr create-metadata "${schema_code}" ${token_id} --from ${address_key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            ${BASE64_META} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    3) echo "Showing NFT"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        sixd q nftmngr show-nft-data ${schema_code} ${token_id} --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} --output json | jq .
        ;;
    4) echo "Do Action"
        read -p "Enter Schema Code: " schema_code 
        read -p "Enter Token ID: " token_id
        read -p "From (address or key) : " address_key
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

        sixd tx nftmngr perform-action-by-nftadmin ${schema_code} ${token_id} ${action} ${ref_id} ${required_params} --from ${address_key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT} -o json | grep -q 'Error:'
        ;;
    5) echo "Set NFT Attribute"
        read -p "Enter Schema Code: " schema_code 
        read -p "From (address or key) : " address_key
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

        BASE64_ATTR=`cat nft-data.json \
            | sed "s/#ATTRIBUTE_NAME#/${ATTRIBUTE_NAME}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE#/${ATTRIBUTE_VALUE_TYPE}/g" \
            | sed "s/#ATTRIBUTE_VALUE_TYPE_VALUE#/${ATTRIBUTE_VALUE_TYPE_VALUE}/g" \
            | base64 | tr -d '\n'`

        echo "BASE64_ATTR: ${BASE64_ATTR}"

        sixd tx nftmngr set-nft-attribute ${schema_code} ${BASE64_ATTR} --from ${address_key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
     6) echo "Add Attribute"
        read -p "Enter Schema Code: " schema_code 
        read -p "From (address or key) : " address_key
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        read -p "Location of attribute (0 or 1): " location
        BASE64_ATTRIBUTE=`cat new-attribute.json | base64 | tr -d '\n'`
        sixd tx nftmngr add-attribute ${schema_code} ${location} ${BASE64_ATTRIBUTE} --from ${address_key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
     7) echo "Add Action"
        read -p "Enter Schema Code: " schema_code 
        read -p "From (address or key) : " address_key
        if [ -z "$schema_code" ]; then
            schema_code=$default_schema_code
        fi
        BASE64_ACTION=`cat new-action.json | base64 | tr -d '\n'`
        sixd tx nftmngr add-action ${schema_code} ${BASE64_ACTION} --from ${address_key} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
            --chain-id ${CHAIN_ID} --node ${RPC_ENDPOINT}
        ;;
    *) echo "Invalid choice"
       ;;
esac