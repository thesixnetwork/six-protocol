module github.com/thesixnetwork/six-protocol

go 1.16

require (
	github.com/CosmWasm/wasmd v0.24.0
	github.com/confio/ics23/go v0.7.0 // indirect
	github.com/cosmos/cosmos-sdk v0.45.4
	github.com/cosmos/ibc-go/v2 v2.0.3
	github.com/gin-gonic/gin v1.7.0 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/starport v0.19.2
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.7
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	google.golang.org/genproto v0.0.0-20220718134204-073382fd740c
	google.golang.org/grpc v1.48.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/cosmos/cosmos-sdk => /Users/hamdeeduere/Documents/Repositories/six-dev/cosmos-sdk
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
