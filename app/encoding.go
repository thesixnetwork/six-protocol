package app

import (
	"github.com/cosmos/cosmos-sdk/std"

	ethermintcryptocodec "github.com/evmos/ethermint/crypto/codec"
	ethermintcodec "github.com/evmos/ethermint/encoding/codec"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"

)

// MakeEncodingConfig creates an EncodingConfig for gravity.
func MakeEncodingConfig() cosmoscmd.EncodingConfig {
	encodingConfig := cosmoscmd.MakeEncodingConfig()
	ethermintcodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ethermintcryptocodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	ModuleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	ModuleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}
