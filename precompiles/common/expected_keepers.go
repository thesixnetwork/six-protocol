package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// "github.com/ethereum/go-ethereum/common"
)

type BankKeeper interface {
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetBalance(sdk.Context, sdk.AccAddress, string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

type EVMKeeper interface {
	// GetCodeHash(sdk.Context, common.Address) common.Hash
	// GetBaseDenom(ctx sdk.Context) string
}

