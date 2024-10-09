package app

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	evmenc "github.com/thesixnetwork/six-protocol/encoding"
)

// MakeEncodingConfig creates the EncodingConfig for the application.
func MakeEncodingConfig() params.EncodingConfig {
	return evmenc.MakeConfig(ModuleBasics)
}
