sixd tx staking create-validator --amount="100000000stake" --from=val1 --moniker sixnode1 \
    --pubkey $(sixd tendermint show-validator --home ./build/sixnode1) --home build/sixnode1 \
    --keyring-backend test --commission-rate 0.1 --commission-max-rate 0.5 --commission-max-change-rate 0.1 \
    --min-self-delegation 1000000 --node http://0.0.0.0:26662 -y --min-delegation 1000000 --delegation-increment 1000000

sixd tx staking create-validator --amount="100000000stake" --from=val2 --moniker sixnode2 \
    --pubkey $(sixd tendermint show-validator --home ./build/sixnode2) --home build/sixnode2 \
    --keyring-backend test --commission-rate 0.1 --commission-max-rate 0.5 --commission-max-change-rate 0.1 \
    --min-self-delegation 1000000 --node http://0.0.0.0:26662 -y --min-delegation 1000000 --delegation-increment 1000000

sixd tx staking create-validator --amount="100000000stake" --from=val3 --moniker sixnode3 \
    --pubkey $(sixd tendermint show-validator --home ./build/sixnode3) --home build/sixnode3 \
    --keyring-backend test --commission-rate 0.1 --commission-max-rate 0.5 --commission-max-change-rate 0.1 \
    --min-self-delegation 1000000 --node http://0.0.0.0:26662 -y --min-delegation 1000000 --delegation-increment 1000000
