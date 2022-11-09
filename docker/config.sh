NODE_PEER=$(jq '.app_state.genutil.gen_txs[0].body.memo' ./build/sixnode0/config/genesis.json)

## replace NODE_PEER in config.toml to persistent_peers
sed -i '' "s/persistent_peers = \"\"/persistent_peers = ${NODE_PEER}/g" ./build/${SIX_HOME}/config/config.toml

## replace to enalbe api
sed -i '' "108s/.*/enable = true/" ./build/${SIX_HOME}/config/app.toml

## replace to from 127.0.0.1 to 0.0.0.0
sed -i '' "s/127.0.0.1/0.0.0.0/g" ./build/${SIX_HOME}/config/config.toml

## config genesis.json
jq '.app_state.bank.params.send_enabled[0] = {"denom": "usix","enabled": true}' ./build/${SIX_HOME}/config/genesis.json | sponge ./build/${SIX_HOME}/config/genesis.json
{
            "denom": "usix",
            "enabled": true
}

## copy genesis.json from build/sixnode0/config/genesis.json to build/${SIX_HOME}/config/genesis.json
cp ./build/sixnode0/config/genesis.json ./build/${SIX_HOME}/config/genesis.json