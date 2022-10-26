# array
modules=(
    protocoladmin
    tokenmngr
    nftadmin
    nftmngr
)

for mod in ${modules[@]}
do
    echo "#######################################"
    echo "Starting ${mod} tests..."
    echo "#######################################"
    sh ${mod}.sh || exit 1
done