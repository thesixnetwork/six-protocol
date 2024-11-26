package common

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	// "github.com/ethereum/go-ethereum/common"
)

type BankKeeper interface {
	SendCoins(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetBalance(sdk.Context, sdk.AccAddress, string) sdk.Coin
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetAllAccounts(ctx sdk.Context) (accounts []authtypes.AccountI)
	IterateAccounts(ctx sdk.Context, cb func(account authtypes.AccountI) bool)
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type TokenmngrKeeper interface {
	AttoCoinConverter(sdk.Context, sdk.AccAddress, sdk.AccAddress, sdk.Int) error
}

type NftmngrKeeper interface {
	GetCodec() codec.BinaryCodec
	GetNFTFeeConfig(ctx sdk.Context) (val nftmngrtypes.NFTFeeConfig, found bool)
	GetActionExecutor(ctx sdk.Context, nftSchemaCode string, executorAddress string) (val nftmngrtypes.ActionExecutor, found bool)
	GetSchemaOwner(ctx sdk.Context, nftSchemaName string) (string, error)
	IsSchemaOwner(ctx sdk.Context, nftSchemaName, inputAddress string) (bool, error)
	// ####################
	// #                  #
	// #     SETTER       #
	// #                  #
	// ####################
	ActionByAdmin(ctx sdk.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*nftmngrtypes.ActionParameter) (changelist []byte, err error)
	AddAttributeKeeper(ctx sdk.Context, creator string, nftSchemaName string, new_add_attribute nftmngrtypes.AttributeDefinition, location nftmngrtypes.AttributeLocation) error
	UpdateAttributeKeeper(ctx sdk.Context, creator, nftSchemaName string, update_attribute nftmngrtypes.AttributeDefinition) error
	ResyncAttibutesKeeper(ctx sdk.Context, creator, nftSchemaName, tokenId string) error
	SetAttributeOveridingKeeper(ctx sdk.Context, creator, nftSchemaName string, newOveridingType int32) error
	ShowAttributeKeeper(ctx sdk.Context, creator, nftSchemaName string, status bool, attributesName []string) error
	ToggleActionKeeper(ctx sdk.Context, creator, nftSchemaName, actionName string, status bool) error
	UpdateActionKeeper(ctx sdk.Context, creator, nftSchemaName string, updateAction nftmngrtypes.Action) error
	AddActionKeeper(ctx sdk.Context, creator string, nftSchemaName string, newAction nftmngrtypes.Action) error
	CreateNewMetadataKeeper(ctx sdk.Context, creator, nftSchemaName, tokenId string, metadata nftmngrtypes.NftData) error
	CreateNftSchemaKeeper(ctx sdk.Context, creator string, schemaInput nftmngrtypes.NFTSchemaINPUT) error
	SetBaseURIKeeper(ctx sdk.Context, creator, nftSchemaName, baseURI string) error
	SetMetadataFormatKeeper(ctx sdk.Context, creator, nftSchemaName, format string) error
	SetMintAuthKeeper(ctx sdk.Context, creator, nftSchemaName string, authTo nftmngrtypes.AuthorizeTo) error
	SetOriginChainKeeper(ctx sdk.Context, creator, nftSchemaName, originChain string) error
	SetOriginContractKeeper(ctx sdk.Context, creator, nftSchemaName, contract string) error
	SetURIRetrievalKeeper(ctx sdk.Context, creator, nftSchemaName string, method int32) error
	ChangeOrgOwner(ctx sdk.Context, creator, newOwner, orgName string) error
	ChangeSchemaOwner(ctx sdk.Context, creator, newOwner, nftSchemaName string) error
	AddActionExecutor(ctx sdk.Context, creator string, nftSchemaName string, executorAddress string) error
	DelActionExecutor(ctx sdk.Context, creator, nftSchemaName, executorAddress string) error
}
