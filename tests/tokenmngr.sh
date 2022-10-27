TOTAL=0
PASSED=0

# create-token
if sixd tx tokenmngr create-token umango 21000000000000 $(sixd keys show -a alice) "{\"description\":\"Mango\",\"denom_units\":[{\"denom\":\"umango\",\"exponent\":0,\"aliases\":[\"micromango\"]},{\"denom\":\"mmango\",\"exponent\":3,\"aliases\":[\"millimango\"]},{\"denom\":\"mango\",\"exponent\":6,\"aliases\":[]}],\"base\":\"umango\",\"display\":\"umango\",\"name\":\"MangoToken\",\"symbol\":\"umango\"}" --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
  echo "✅ create-token success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
  if sixd q tokenmngr list-token --chain-id testnet | grep -q 'umango'; then
    echo "✅ query list-token success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
  else
    echo "🛑 query list-token failed"
  fi
else
    echo "🛑 create-token failed"
fi

# create-mintperm alice
if sixd tx tokenmngr create-mintperm umango $(sixd keys show -a alice) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ create-mintperm alice success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q tokenmngr list-mintperm --chain-id testnet | grep -q $(sixd keys show -a alice); then
        echo "✅ query list-mintperm success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-mintperm failed"
    fi
else
    echo "🛑 create-mintperm alice failed"
fi

# create-mintperm bob
if sixd tx tokenmngr create-mintperm umango $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ create-mintperm bob success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q tokenmngr list-mintperm --chain-id testnet | grep -q $(sixd keys show -a bob); then
        echo "✅ query list-mintperm success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-mintperm failed"
    fi
else
    echo "🛑 create-mintperm bob failed"
fi

# mint
if sixd tx tokenmngr mint 1000000000000 umango --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "✅ mint success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q bank total --chain-id testnet | grep -q 'umango'; then
        echo "✅ query bank total umango success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query bank total umango failed"
    fi
else
    echo "🛑 mint failed"
fi

# burn
if sixd tx tokenmngr burn 500000000000 umango --chain-id testnet --from alice -y | grep -q 'msg_index: 0'; then
    echo "✅ burn success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    # 500000000000 umango exists
    if sixd q tokenmngr show-burn umango --chain-id testnet | grep -q 'umango'; then
        echo "✅ query bank total umango success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query bank total umango failed"
    fi
else
    echo "🛑 burn failed"
fi

# create-token banana
if sixd tx tokenmngr create-token banana 21000000000000 $(sixd keys show -a bob) "{\"description\":\"banana\",\"denom_units\":[{\"denom\":\"ubanana\",\"exponent\":0,\"aliases\":[\"microbanana\"]},{\"denom\":\"mbanana\",\"exponent\":3,\"aliases\":[\"millibanana\"]},{\"denom\":\"banana\",\"exponent\":6,\"aliases\":[]}],\"base\":\"ubanana\",\"display\":\"banana\",\"name\":\"bananaToken\",\"symbol\":\"banana\"}" --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ create-token banana success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q tokenmngr list-token --chain-id testnet | grep -q 'banana'; then
        echo "✅ query list-token success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    else
        echo "🛑 query list-token failed"
    fi
    else 
        echo "🛑 create-token banana failed"
    fi

# delete-mintperm
if sixd tx tokenmngr delete-mintperm umango $(sixd keys show -a bob) --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ delete-mintperm success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q tokenmngr list-mintperm --chain-id testnet | grep -q $(sixd keys show -a bob); then
        echo "🛑 query list-mintperm failed"
    else
        echo "✅ query list-mintperm success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    fi
else
    echo "🛑 delete-mintperm failed"
fi

# delete-token
if sixd tx tokenmngr delete-token banana --chain-id testnet --from super-admin -y | grep -q 'msg_index: 0'; then
    echo "✅ delete-token success"
    TOTAL=$((TOTAL+1))
    PASSED=$((PASSED+1))
    if sixd q tokenmngr list-token --chain-id testnet | grep -q 'banana'; then
        echo "🛑 query list-token failed"
    else
        echo "✅ query list-token success"
        TOTAL=$((TOTAL+1))
        PASSED=$((PASSED+1))
    fi
else
    echo "🛑 delete-token failed"
fi
echo "========================================"
echo "tokenmngr: Passed $PASSED out of $TOTAL tests"