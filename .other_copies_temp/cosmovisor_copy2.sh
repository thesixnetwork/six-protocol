export DAEMON_NAME=sixd
export DAEMON_HOME=$HOME/.six
export DAEMON_RESTART_AFTER_UPGRADE=true
export DAEMON_BACKUP_DIR=$HOME/.six
export DAEMON_SHUTDOWN_GRACE=10s
cosmovisor run start --minimum-gas-prices=1.25usix,1250000000000asix --api.enable true --json-rpc.api eth,txpool,personal,net,debug,web3 --rpc.laddr tcp://0.0.0.0:26657  --log_level info --json-rpc.allow-unprotected-txs true --home $DAEMON_HOME