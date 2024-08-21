package common

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	// nftmngrtype "github.com/thesixnetwork/sixnft/x/nftmngr/types"
	// "github.com/ethereum/go-ethereum/common"
)

type BankKeeper interface {
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetBalance(sdk.Context, sdk.AccAddress, string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin

	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	// SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetAllAccounts(ctx sdk.Context) (accounts []authtypes.AccountI)
	IterateAccounts(ctx sdk.Context, cb func(account authtypes.AccountI) bool)
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type NftmngrKeeper interface {
	// ActionByAdmin(ctx sdk.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*nftmngrtype.ActionParameter) (changelist []byte, err error)
}
