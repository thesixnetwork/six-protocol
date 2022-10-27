RPC_ENDPOINT=http://localhost:26657
grantOracle()
{
    echo "Grant 'oracle' to $1"
    sixd tx nftadmin grant-permission oracle $1 --from super-admin --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix -y \
        --node ${RPC_ENDPOINT} --chain-id testnet
}

if sixd tx nftadmin grant-permission oracle_admin $(sixd keys show alice -a) --from super-admin -y --node ${RPC_ENDPOINT} --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "grant-permission oracle_admin success"
else
    echo "grant-permission oracle_admin failed"
fi
# set oracle admin

if grantOracle $(sixd keys show oracle1 -a) | grep -q 'msg_index: 0'; then
    echo "grantOracle1 success"
else
    echo "grantOracle1 failed"
fi

if grantOracle $(sixd keys show oracle2 -a) | grep -q 'msg_index: 0'; then
    echo "grantOracle2 success"
else
    echo "grantOracle2 failed"
fi

if grantOracle $(sixd keys show oracle3 -a) | grep -q 'msg_index: 0'; then
    echo "grantOracle3 success"
else
    echo "grantOracle3 failed"
fi

if grantOracle $(sixd keys show oracle4 -a) | grep -q 'msg_index: 0'; then
    echo "grantOracle4 success"
else
    echo "grantOracle4 failed"
fi

# set minter
if sixd tx nftadmin grant-permission minter $(sixd keys show alice -a) --gas auto --gas-adjustment 1.5 \
    --gas-prices 0.1usix --from super-admin -y --node ${RPC_ENDPOINT} --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "Set minter success"
else
    echo "Set minter failed"
fi

# Mint usix
if sixd tx nftadmin mint 5 usix --from alice --chain-id testnet -y --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix \
    --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "Mint success"
else
    echo "Mint failed"
fi

# set burner
if sixd tx nftadmin grant-permission burner $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix \
    --node ${RPC_ENDPOINT} --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "Set burner success"
else
    echo "Set burner failed"
fi

# Burn usix
if sixd tx nftadmin burn 5 usix --from alice --chain-id testnet -y --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix \
    --node ${RPC_ENDPOINT} | grep -q 'msg_index: 0'; then
    echo "Burn success"
else
    echo "Burn failed"
fi

# set test perm
if sixd tx nftadmin grant-permission test-perm $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix \
    --chain-id testnet | grep -q 'msg_index: 0'; then
    echo " "
else
    echo " "
fi

if sixd tx nftadmin revoke-permission test-perm $(sixd keys show alice -a) --from super-admin -y --gas auto --gas-adjustment 1.5 --gas-prices 0.1usix \
    --chain-id testnet | grep -q 'msg_index: 0'; then
    echo "Revoke test-perm success"
else
    echo "Revoke test-perm faild"
fi