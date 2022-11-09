RPC_ENDPOINT=$1
# NFT Schema
BASE64_SCHEMA=`cat ./mock-data/nft-schema.json | base64 | tr -d '\n'`
# NFT Data
BASE64_DATA=`cat ./mock-data/nft-data.json | base64 | tr -d '\n'`
# New Action
BASE64_ACTION=`cat ./mock-data/new-action.json | base64 | tr -d '\n'`
# New Attribute
BASE64_ATTRIBUTE=`cat ./mock-data/new-attribute.json | base64 | tr -d '\n'`

TOTAL=0
PASSED=0

# create-nft-schema
if sixd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from tester --node ${RPC_ENDPOINT} --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ create-nft-schema success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q $(cat ./mock-data/nft-schema.json | jq -r '.code'); then
        echo "✅ query create-nft-schema success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query create-nft-schema failed"
    fi
else
    echo "🛑 create-nft-schema failed"
fi

# add-action
if sixd tx nftmngr add-action sixnetwork.nftexpo ${BASE64_ACTION} --from tester --node ${RPC_ENDPOINT} --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ add-action success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1)) 
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q $(cat ./mock-data/new-action.json | jq -r '.name'); then
        echo "✅ query add-action success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query add-action failed"
    fi
else
    echo "🛑 add-action failed"
fi

# add-attribute
if sixd tx nftmngr add-attribute sixnetwork.nftexpo 1 ${BASE64_ATTRIBUTE} --from tester --node ${RPC_ENDPOINT} --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ add-attribute success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q $(cat ./mock-data/new-attribute.json | jq -r '.name'); then
        echo "✅ query add-attribute success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query add-attribute failed"
    fi
else
    echo "🛑 add-attribute failed"
fi

# add-system-actioner
if sixd tx nftmngr add-system-actioner sixnetwork.nftexpo $(sixd keys show -a tester2 --keyring-backend=os) --from tester --node ${RPC_ENDPOINT} --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ add-system-actioner success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 add-system-actioner failed"
fi

# change-schema-owner
if sixd tx nftmngr change-schema-owner sixnetwork.nftexpo $(sixd keys show -a tester2 --keyring-backend=os) --from tester --node ${RPC_ENDPOINT} --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ change-schema-owner success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 change-schema-owner failed"
fi

# create-metadata
if sixd tx nftmngr create-metadata sixnetwork.nftexpo 0 ${BASE64_DATA} --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ create-metadata success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 create-metadata failed"
fi

# perform-action-by-nftadmin
if sixd tx nftmngr perform-action-by-nftadmin sixnetwork.nftexpo 0 check_in --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ perform-action-by-nft-admin success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 perform-action-by-nft-admin failed"
fi

# remove-system-actioner
if sixd tx nftmngr remove-system-actioner sixnetwork.nftexpo $(sixd keys show -a tester2 --keyring-backend=os) --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ remove-system-actioner success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 remove-system-actioner failed"
fi

# resync-attributes
if sixd tx nftmngr resync-attributes sixnetwork.nftexpo 0 --from tester2 --keyring-backend=os --node ${RPC_ENDPOINT} --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ resync-attributes success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 resync-attributes failed"
fi

# set-base-uri
if sixd tx nftmngr set-base-uri sixnetwork.nftexpo https://nftexpo.six.network/ --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ set-base-uri success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q 'https://nftexpo.six.network/'; then
        echo "✅ query set-base-uri success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query set-base-uri failed"
    fi
else
    echo "🛑 set-base-uri failed"
fi

# set-fee-config
# if sixd tx nftmngr set-fee-config sixnetwork.nftexpo ${BASE64_FEE_CONFIG} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
#     echo "set-fee-config success"
# else
#     echo "set-fee-config failed"
# fi

# set-mintauth 0
if sixd tx nftmngr set-mintauth sixnetwork.nftexpo 0 --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ set-mintauth success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q 'mint_authorization: system'; then
        echo "✅ query set-mintauth success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query set-mintauth failed"
    fi
else
    echo "🛑 set-mintauth failed"
fi

# set-mintauth 1
if sixd tx nftmngr set-mintauth sixnetwork.nftexpo 1 --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ set-mintauth success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
    if sixd q nftmngr show-nft-schema --node ${RPC_ENDPOINT} $(cat ./mock-data/nft-schema.json | jq -r '.code') | grep -q 'mint_authorization: all'; then
        echo "✅ query set-mintauth success"
        TOTAL=$(($TOTAL+1))
        PASSED=$(($PASSED+1))
    else
        echo "🛑 query set-mintauth failed"
    fi
else
    echo "🛑 set-mintauth failed"
fi

# set-nft-attribute
# if sixd tx nftmngr set-nft-attribute sixnetwork.nftexpo ${BASE64_DATA} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
#     echo "set-nft-attribute success"
# else
#     echo "set-nft-attribute failed"
# fi

# show-attributes
if sixd tx nftmngr show-attributes sixnetwork.nftexpo true points --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ show-attributes success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 show-attributes failed"
fi

# toggle-action
if sixd tx nftmngr toggle-action sixnetwork.nftexpo check_in --node ${RPC_ENDPOINT} --from tester2 --keyring-backend=os --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y --chain-id fivenet | grep -q 'msg_index: 0'; then
    echo "✅ toggle-action success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 toggle-action failed"
fi
echo "========================================"
echo "nftmngr: Passed $PASSED out of $TOTAL tests"