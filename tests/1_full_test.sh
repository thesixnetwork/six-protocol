RPC_ENDPOINT=$1
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
    echo "Starting ${mod} tests on ${RPC_ENDPOINT}..."
    echo "#######################################"
    bash ${mod}.sh || exit 1
done