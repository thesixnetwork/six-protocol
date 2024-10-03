export DAEMON_NAME=sixd         
export DAEMON_HOME=/opt/build/six_home
export DAEMON_DATA_BACKUP_DIR=/opt/build/six_home
export UNSAFE_SKIP_BACKUP=false
export DAEMON_RESTART_AFTER_UPGRADE=true
cosmovisor start --home $DAEMON_HOME  --minimum-gas-prices 1.25usix