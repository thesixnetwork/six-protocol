export DAEMON_NAME=sixd         
export DAEMON_HOME=/opt/build/six_home
export DAEMON_RESTART_AFTER_UPGRADE=true
cosmovisor run start --home $DAEMON_HOME 
