package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tkmtype "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	evmbindtype "github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

type TokenmngrKeeper interface {
	GetTokenBurn(ctx sdk.Context, token string) (val tkmtype.TokenBurn, found bool)
	GetToken(ctx sdk.Context, name string) (val tkmtype.Token, found bool)
	SetTokenBurn(ctx sdk.Context, tokenBurn tkmtype.TokenBurn)
	UpdateBurn(ctx sdk.Context, burn tkmtype.Burn) uint64
	// Methods imported from tokenmngr should be defined here
}

type EvmbindKeeper interface {
	GetBinding(ctx sdk.Context,ethAddress string,) (val evmbindtype.Binding, found bool) 
	// Methods imported from evmbind should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}
