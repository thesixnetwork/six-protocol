# NFT Schema
BASE64_SCHEMA=`cat ./mock-data/nft-schema.json | base64 | tr -d '\n'`
# NFT Data
BASE64_DATA=`cat ./mock-data/nft-data.json | base64 | tr -d '\n'`
# New Action
BASE64_ACTION=`cat ./mock-data/new-action.json | base64 | tr -d '\n'`
# New Attribute
BASE64_ATTRIBUTE=`cat ./mock-data/new-attribute.json | base64 | tr -d '\n'`

# create-nft-schema
if sixd tx nftmngr create-nft-schema ${BASE64_SCHEMA} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "create-nft-schema success"
else
    echo "create-nft-schema failed"
fi

# add-action
if sixd tx nftmngr add-action sixnetwork.nftexpo ${BASE64_ACTION} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "add-action success" 
else
    echo "add-action failed"
fi

# add-attribute
if sixd tx nftmngr add-attribute sixnetwork.nftexpo 1 ${BASE64_ATTRIBUTE} --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "add-attribute success"
else
    echo "add-attribute failed"
fi

# add-system-actioner
if sixd tx nftmngr add-system-actioner sixnetwork.nftexpo $(sixd keys show -a bob) --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "add-system-actioner success"
else
    echo "add-system-actioner failed"
fi

# change-schema-owner
if sixd tx nftmngr change-schema-owner sixnetwork.nftexpo $(sixd keys show -a bob) --from alice --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "change-schema-owner success"
else
    echo "change-schema-owner failed"
fi

# create-metadata
if sixd tx nftmngr create-metadata sixnetwork.nftexpo 0 ${BASE64_DATA} --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "create-metadata success"
else
    echo "create-metadata failed"
fi

# perform-action-by-nftadmin
if sixd tx nftmngr perform-action-by-nftadmin sixnetwork.nftexpo 0 check_in --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "perform-action-by-nft-admin success"
else
    echo "perform-action-by-nft-admin failed"
fi

# remove-system-actioner
if sixd tx nftmngr remove-system-actioner sixnetwork.nftexpo $(sixd keys show -a bob) --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "remove-system-actioner success"
else
    echo "remove-system-actioner failed"
fi

# resync-attributes
if sixd tx nftmngr resync-attributes sixnetwork.nftexpo 0 --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "resync-attributes success"
else
    echo "resync-attributes failed"
fi

# set-base-uri
if sixd tx nftmngr set-base-uri sixnetwork.nftexpo https://nftexpo.six.network/ --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "set-base-uri success"
else
    echo "set-base-uri failed"
fi

# set-fee-config
# if sixd tx nftmngr set-fee-config sixnetwork.nftexpo ${BASE64_FEE_CONFIG} --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
#     echo "set-fee-config success"
# else
#     echo "set-fee-config failed"
# fi

# set-mintauth 0
if sixd tx nftmngr set-mintauth sixnetwork.nftexpo 0 --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "set-mintauth success"
else
    echo "set-mintauth failed"
fi

# set-mintauth 0
if sixd tx nftmngr set-mintauth sixnetwork.nftexpo 1 --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "set-mintauth success"
else
    echo "set-mintauth failed"
fi

# set-nft-attribute
# if sixd tx nftmngr set-nft-attribute sixnetwork.nftexpo ${BASE64_DATA} --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
#     echo "set-nft-attribute success"
# else
#     echo "set-nft-attribute failed"
# fi

# show-attributes
if sixd tx nftmngr show-attributes sixnetwork.nftexpo true points --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "show-attributes success"
else
    echo "show-attributes failed"
fi

# toggle-action
if sixd tx nftmngr toggle-action sixnetwork.nftexpo check_in --from bob --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "toggle-action success"
else
    echo "toggle-action failed"
fi