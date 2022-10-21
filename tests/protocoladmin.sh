# create-group
if sixd tx protocoladmin create-group artist --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "create-group success"
else
    echo "create-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist $(sixd keys show -a alice) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "add-admin-to-group success"
else
    echo "add-admin-to-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "add-admin-to-group success"
else
    echo "add-admin-to-group failed"
fi

# remove-admin-from-group
if sixd tx protocoladmin remove-admin-from-group artist $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "remove-admin-from-group success"
else
    echo "remove-admin-from-group failed"
fi

#create for delete
if sixd tx protocoladmin create-group artist_test --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo " "
else
    echo " "
fi
# delete-group
if sixd tx protocoladmin delete-group artist_test --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "delete-group success"
else
    echo "delete-group failed"
fi