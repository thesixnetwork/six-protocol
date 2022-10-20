# create-token

if sixd tx tokenmngr create-token umango 21000000000000 6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq "{\"description\":\"Mango\",\"denom_units\":[{\"denom\":\"umango\",\"exponent\":0,\"aliases\":[\"micromango\"]},{\"denom\":\"mmango\",\"exponent\":3,\"aliases\":[\"millimango\"]},{\"denom\":\"mango\",\"exponent\":6,\"aliases\":[]}],\"base\":\"umango\",\"display\":\"umango\",\"name\":\"MangoToken\",\"symbol\":\"umango\"}" --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
  echo "create-token success"
else
    echo "create-token failed"
fi

# create-mintperm
if sixd tx tokenmngr create-mintperm umango 6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "create-mintperm success"
else
    echo "create-mintperm failed"
fi

# mint
if sixd tx tokenmngr mint 1000000000000 umango --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "mint success"
else
    echo "mint failed"
fi

# burn
if sixd tx tokenmngr burn 500000000000 umango --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "burn success"
else
    echo "burn failed"
fi

# create-token banana
if sixd tx tokenmngr create-token banana 21000000000000 6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq "{\"description\":\"banana\",\"denom_units\":[{\"denom\":\"ubanana\",\"exponent\":0,\"aliases\":[\"microbanana\"]},{\"denom\":\"mbanana\",\"exponent\":3,\"aliases\":[\"millibanana\"]},{\"denom\":\"banana\",\"exponent\":6,\"aliases\":[]}],\"base\":\"ubanana\",\"display\":\"banana\",\"name\":\"bananaToken\",\"symbol\":\"banana\"}" --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "create-token banana success"
    else 
        echo "create-token banana failed"
    fi

# delete-mintperm
if sixd tx tokenmngr delete-mintperm umango 6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "delete-mintperm success"
else
    echo "delete-mintperm failed"
fi

# delete-token
if sixd tx tokenmngr delete-token banana --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "delete-token success"
else
    echo "delete-token failed"
fi
