export DAEMON_NAME=/usr/bin/sixd         
export DAEMON_HOME=/opt/build/six_home
export DAEMON_RESTART_AFTER_UPGRADE=true
cosmovisor run start --rpc.laddr tcp://0.0.0.0:26657 --log_level=info