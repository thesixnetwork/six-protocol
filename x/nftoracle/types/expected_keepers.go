package types

import (
	"context"

	nftadmintypes "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// StakingKeeper defines the expected interface for the Staking module.
type NftmngrKeeper interface {
	GetNFTSchema(
		context context.Context,
		code string,
	) (val nftmngrtypes.NFTSchema, found bool)
	GetAllSchemaAttribute(context context.Context) (list []nftmngrtypes.SchemaAttribute)

	SetNFTSchema(ctx context.Context, nFTSchema nftmngrtypes.NFTSchema)

	GetNftData(
		ctx context.Context,
		nftSchemaCode string,
		tokenId string,
	) (val nftmngrtypes.NftData, found bool)

	// ValidateNFTData(data *nftmngrtypes.NftData, schema *nftmngrtypes.NFTSchema) (bool, error)
	SetNftData(ctx context.Context, nftData nftmngrtypes.NftData)
	GetActionByRefId(
		ctx context.Context,
		refId string,
	) (val nftmngrtypes.ActionByRefId, found bool)
	SetActionByRefId(ctx context.Context, actionByRefId nftmngrtypes.ActionByRefId)

	GetSchemaAttribute(ctx context.Context, nftSchemaCode string, name string) (val nftmngrtypes.SchemaAttribute, found bool)
	SetSchemaAttribute(ctx context.Context, schemaAttribute nftmngrtypes.SchemaAttribute)

	GetActionOfSchema(tx context.Context, nftSchemaCode string, name string) (val nftmngrtypes.ActionOfSchema, found bool)
}

type NftadminKeeper interface {
	GetAuthorization(ctx context.Context) (val nftadmintypes.Authorization, found bool)
	HasPermission(ctx context.Context, name string, addr sdk.AccAddress) bool
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
