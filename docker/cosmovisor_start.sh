export DAEMON_NAME=sixd         
export DAEMON_HOME=$HOME/.six
export DAEMON_RESTART_AFTER_UPGRADE=true
cosmovisor run start --home $DAEMON_HOME 