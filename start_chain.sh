VERBOSE=$1
rm ./docs/static/openapi.yml
rm go.sum && touch go.sum
# export GOPRIVATE=github.com/thesixnetwork/six-protocol
go mod tidy
ignite chain serve --config ./config.yml -r -f $VERBOSE