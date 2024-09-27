VERBOSE=$1
rm ./docs/static/openapi.yml
rm go.sum && touch go.sum
# export GOPRIVATE=github.com/thesixnetwork/sixnft
go mod tidy -e -go=1.18 && go mod tidy -e -go=1.19
ignite chain serve --config ./config.yml -r -f $VERBOSE