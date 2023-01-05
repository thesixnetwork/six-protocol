sixd tx gov submit-proposal software-upgrade v2.1.0 \
    --title="idk" \
    --description="idk" \
    --upgrade-height=350 \
    --from=alice \
    --home ./build/sixnode0 \
    --gas=auto \
    --gas-adjustment=1.5 \
    --gas-prices=1.25usix \
    --keyring-backend test \
    --chain-id six \
    --node http://localhost:26657

sixd tx gov deposit 1 100000000usix --from alice --home ./build/sixnode0 --keyring-backend test \
    --gas=auto \
    --gas-adjustment=1.5 \
    --gas-prices=1.25usix \
    --chain-id six \
    --node http://localhost:26657


sixd tx gov vote 1 yes --from val1 --home ./build/sixnode0 --keyring-backend test \
    --chain-id six --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
    --node http://localhost:26657 -y

sixd tx gov vote 1 yes --from val2 --home ./build/sixnode1 --keyring-backend test \
    --chain-id six --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
    --node http://localhost:26657 -y

sixd tx gov vote 1 yes --from val3 --home ./build/sixnode2 --keyring-backend test \
    --chain-id six --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
    --node http://localhost:26657 -y

sixd tx gov vote 1 yes --from val4 --home ./build/sixnode3 --keyring-backend test \
    --chain-id six --gas auto --gas-adjustment 1.5 --gas-prices 1.25usix \
    --node http://localhost:26657 -y



sixnftd tx gov submit-proposal software-upgrade v0.7.0 --title="idk" \
    --description="idk" \
    --upgrade-height=700 \
    --from=alice \
    --home ~/.sixnft --gas=auto \
    --gas-adjustment=1.5 \
    --gas-prices=1.25stake \
    --keyring-backend test \
    --chain-id sixnft \
    --node http://localhost:26657


sixnftd tx gov deposit 1 100000000stake --from alice --home ~/.sixnft --keyring-backend test \
    --gas=auto \
    --gas-adjustment=1.5 \
    --gas-prices=1.25stake \
    --chain-id sixnft \
    --node http://localhost:26657

sixnftd tx gov vote 1 yes --from bob --home ~/.sixnft  --keyring-backend test \
    --chain-id sixnft --gas auto --gas-adjustment 1.5 --gas-prices 1.25stake \
    --node http://localhost:26657 -y


