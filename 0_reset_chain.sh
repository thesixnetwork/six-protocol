export DAEMON_HOME=./build/sixnode1
rm -rf $DAEMON_HOME/data
rm -rf $DAEMON_HOME/wasm
rm $DAEMON_HOME/config/addrbook.json
mkdir $DAEMON_HOME/data/
touch $DAEMON_HOME/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > $DAEMON_HOME/data/priv_validator_state.json

export DAEMON_HOME=./build/sixnode2
rm -rf $DAEMON_HOME/data
rm -rf $DAEMON_HOME/wasm
rm $DAEMON_HOME/config/addrbook.json
mkdir $DAEMON_HOME/data/
touch $DAEMON_HOME/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > $DAEMON_HOME/data/priv_validator_state.json

export DAEMON_HOME=./build/sixnode3
rm -rf $DAEMON_HOME/data
rm -rf $DAEMON_HOME/wasm
rm $DAEMON_HOME/config/addrbook.json
mkdir $DAEMON_HOME/data/
touch $DAEMON_HOME/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > $DAEMON_HOME/data/priv_validator_state.json

export DAEMON_HOME=./build/sixnode0
rm -rf $DAEMON_HOME/data
rm -rf $DAEMON_HOME/wasm
rm $DAEMON_HOME/config/addrbook.json
mkdir $DAEMON_HOME/data/
touch $DAEMON_HOME/data/priv_validator_state.json
echo '{"height": "0", "round": 0,"step": 0}' > $DAEMON_HOME/data/priv_validator_state.json