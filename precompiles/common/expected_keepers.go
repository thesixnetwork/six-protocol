package common

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	nftmngrtype "github.com/thesixnetwork/sixnft/x/nftmngr/types"
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
	GetCodec() codec.BinaryCodec
	// ####################
	// #                  #
	// #     SETTER       #
	// #                  #
	// ####################
	ActionByAdmin(ctx sdk.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*nftmngrtype.ActionParameter) (changelist []byte, err error)
	AddAttributeKeeper(ctx sdk.Context, creator string, nftSchemaName string, new_add_attribute nftmngrtype.AttributeDefinition, location nftmngrtype.AttributeLocation) error
	UpdateAttributeKeeper(ctx sdk.Context, creator, nftSchemaName string, update_attribute nftmngrtype.AttributeDefinition) error
	ResyncAttibutesKeeper(ctx sdk.Context, creator, nftSchemaName, tokenId string) error
	SetAttributeOveridingKeeper(ctx sdk.Context, creator, nftSchemaName string, newOveridingType int32) error
	ShowAttributeKeeper(ctx sdk.Context, creator, nftSchemaName string, status bool, attributesName []string) error
	ToggleActionKeeper(ctx sdk.Context, creator, nftSchemaName, actionName string, status bool) error
	UpdateActionKeeper(ctx sdk.Context, creator, nftSchemaName string, updateAction nftmngrtype.Action) error
	AddActionKeeper(ctx sdk.Context, creator string, nftSchemaName string, newAction nftmngrtype.Action) error
	CreateNewMetadataKeeper(ctx sdk.Context, creator, nftSchemaName, tokenId string, metadata nftmngrtype.NftData) error
	CreateNftSchemaKeeper(ctx sdk.Context, creator string, schemaInput nftmngrtype.NFTSchemaINPUT) error
	SetBaseURIKeeper(ctx sdk.Context, creator, nftSchemaName, baseURI string) error
	SetMetadataFormatKeeper(ctx sdk.Context, creator, nftSchemaName, format string) error
	SetMintAuthKeeper(ctx sdk.Context, creator, nftSchemaName string, authTo nftmngrtype.AuthorizeTo) error
	SetOriginChainKeeper(ctx sdk.Context, creator, nftSchemaName, originChain string) error
	SetOriginContractKeeper(ctx sdk.Context, creator, nftSchemaName, contract string) error
	SetURIRetrievalKeeper(ctx sdk.Context, creator, nftSchemaName string, method int32) error
	ChangeOrgOwner(ctx sdk.Context, creator, newOwner, orgName string) error
	ChangeSchemaOwner(ctx sdk.Context, creator, newOwner, nftSchemaName string) error
	AddActionExecutor(ctx sdk.Context, creator string, nftSchemaName string, executorAddress string) error
	DelActionExecutor(ctx sdk.Context, creator, nftSchemaName, executorAddress string) error
}
