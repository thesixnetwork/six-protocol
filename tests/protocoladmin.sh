TOTAL=0
PASSED=0

# create-group
if sixd tx protocoladmin create-group artist --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ create-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-group --chain-id testnet | grep -q 'artist'; then
        echo "✅ query list-group success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-group failed"
    fi
else
    echo "🛑 create-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist $(sixd keys show -a alice) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ add-admin-to-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-admin --chain-id testnet | grep -q $(sixd keys show -a alice); then
        echo "✅ query list-admin success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-admin failed"
    fi
else
    echo "🛑 add-admin-to-group failed"
fi

# add-admin-to-group
if sixd tx protocoladmin add-admin-to-group artist $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ add-admin-to-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-admin --chain-id testnet | grep -q $(sixd keys show -a bob); then
        echo "✅ query list-admin success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-admin failed"
    fi
else
    echo "🛑 add-admin-to-group failed"
fi

# remove-admin-from-group
if sixd tx protocoladmin remove-admin-from-group artist $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ remove-admin-from-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-admin --chain-id testnet | grep -q $(sixd keys show -a bob); then
        echo "🛑 query list-admin failed"
    else
        echo "✅ query list-admin success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    fi
else
    echo "🛑 remove-admin-from-group failed"
fi

#create for delete
if sixd tx protocoladmin create-group artist_test --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ create-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-group --chain-id testnet | grep -q 'artist_test'; then
        echo "✅ query list-group success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-group failed"
    fi
else
    echo "🛑 create-group failed"
fi
# delete-group
if sixd tx protocoladmin delete-group artist_test --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ delete-group success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q protocoladmin list-group --chain-id testnet | grep -q 'artist_test'; then
        echo "🛑 query list-group failed"
    else
        echo "✅ query list-group success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    fi
else
    echo "🛑 delete-group failed"
fi
echo "========================================"
echo "protocoladmin: Passed $PASSED out of $TOTAL tests"