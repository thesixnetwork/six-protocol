# create-group
if sixd tx protocoladmin create-group artist --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "create-group success"
else
    echo "create-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist 6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "add-admin-to-group success"
else
    echo "add-admin-to-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist 6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "add-admin-to-group success"
else
    echo "add-admin-to-group failed"
fi

# remove-admin-from-group
if sixd tx protocoladmin remove-admin-from-group artist 6x13g50hqdqsjk85fmgqz2h5xdxq49lsmjdwlemsp --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "remove-admin-from-group success"
else
    echo "remove-admin-from-group failed"
fi

# delete-group
if sixd tx protocoladmin delete-group artist --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "delete-group success"
else
    echo "delete-group failed"
fi