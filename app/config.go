package app

import (
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func init() {
	// Set prefixes
	accountPubKeyPrefix := AccountAddressPrefix + "pub"
	validatorAddressPrefix := AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.SetCoinType(sdk.CoinType)
	config.SetPurpose(sdk.Purpose)
	config.SetFullFundraiserPath(sdk.FullFundraiserPath)
	config.Seal()

	RegisterDenoms()
}

const (
	// DisplayDenom defines the denomination displayed to users in client applications.
	DisplayDenom = "six"
	EVMDenom     = "asix"
)

// RegisterDenoms registers the base and display denominations to the SDK.
func RegisterDenoms() {
	if err := sdk.RegisterDenom(DisplayDenom, sdkmath.LegacyOneDec()); err != nil {
		panic(err)
	}

	if err := sdk.RegisterDenom(BaseDenom, sdkmath.LegacyNewDecWithPrec(1, BaseDenomUnit)); err != nil {
		panic(err)
	}
}