RPC_ENDPOINT=$1
CHAIN_ID=$2
key=$3

grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix -y \
        --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID}
}
TOTAL=0
PASSED=0

if sixd tx nftadmin grant-permission oracle_admin $(sixd keys show alice -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ grant-permission oracle_admin success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 grant-permission oracle_admin failed"
fi
# set oracle admin

if grantOracle $(sixd keys show oracle1 -a) | grep -q 'msg_index: 0'; then
    echo "✅ grantOracle1 success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 grantOracle1 failed"
fi

if grantOracle $(sixd keys show oracle2 -a) | grep -q 'msg_index: 0'; then
    echo "✅ grantOracle2 success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 grantOracle2 failed"
fi

if grantOracle $(sixd keys show oracle3 -a) | grep -q 'msg_index: 0'; then
    echo "✅ grantOracle3 success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))

else
    echo "🛑 grantOracle3 failed"
fi

if grantOracle $(sixd keys show oracle4 -a) | grep -q 'msg_index: 0'; then
    echo "✅ grantOracle4 success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 grantOracle4 failed"
fi

# set minter
if sixd tx nftadmin grant-permission minter $(sixd keys show alice -a) --gas auto --gas-adjustment 1.5 \
    --gas-prices 1.25usix--from super-admin -y --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Set minter success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 Set minter failed"
fi

# Mint usix
if sixd tx nftadmin mint 5 usix --from alice -y --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix\
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Mint success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 Mint failed"
fi

# set burner
if sixd tx nftadmin grant-permission burner $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix\
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Set burner success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 Set burner failed"
fi

# Burn usix
if sixd tx nftadmin burn 5 usix --from alice --chain-id ${CHAIN_ID} -y --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix\
    --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "✅ Burn success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 Burn failed"
fi

# set test perm
if sixd tx nftadmin grant-permission test-perm $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix\
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo
else
    echo
fi

if sixd tx nftadmin revoke-permission test-perm $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix\
    --node ${RPC_ENDPOINT} --chain-id ${CHAIN_ID} | grep -q 'msg_index: 0'; then
    echo "✅ Revoke test-perm success"
    TOTAL=$(($TOTAL+1))
    PASSED=$(($PASSED+1))
else
    echo "🛑 Revoke test-perm faild"
fi
echo "========================================"
echo "nftadmin: Passed $PASSED out of $TOTAL tests"