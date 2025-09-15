package common

import (
	"context"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	disttypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	nftmngrtypes "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	tokenmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

type BankKeeper interface {
	SendCoins(context.Context, sdk.AccAddress, sdk.AccAddress, sdk.Coins) error
	GetBalance(context.Context, sdk.AccAddress, string) sdk.Coin
	GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
	GetSupply(ctx context.Context, denom string) sdk.Coin
}

type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetAllAccounts(ctx context.Context) (accounts []sdk.AccountI)
	IterateAccounts(ctx context.Context, cb func(account sdk.AccountI) bool)
	GetSequence(context.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
}

type TokenmngrKeeper interface {
	AttoCoinConverter(context.Context, sdk.AccAddress, sdk.AccAddress, sdkmath.Int) error
	ChangeDelegatorAddress(ctx context.Context, oldDelAddr sdk.AccAddress, newDelAddr sdk.AccAddress) error
}

type TokenmngrMsgServer interface {
	UnwrapToken(goCtx context.Context, msg *tokenmngrtypes.MsgUnwrapToken) (*tokenmngrtypes.MsgUnwrapTokenResponse, error)
	WrapToken(goCtx context.Context, msg *tokenmngrtypes.MsgWrapToken) (*tokenmngrtypes.MsgWrapTokenResponse, error)
}

type NftmngrKeeper interface {
	GetCodec() codec.BinaryCodec
	GetNFTFeeConfig(ctx context.Context) (val nftmngrtypes.NFTFeeConfig, found bool)
	GetActionExecutor(ctx context.Context, nftSchemaCode string, executorAddress string) (val nftmngrtypes.ActionExecutor, found bool)
	GetSchemaOwner(ctx context.Context, nftSchemaName string) (string, error)
	IsSchemaOwner(ctx context.Context, nftSchemaName, inputAddress string) (bool, error)
	GetAttributeValue(ctx context.Context, nftSchemaName, tokenId, attributeName string) (string, error)
	// ####################
	// #                  #
	// #     SETTER       #
	// #                  #
	// ####################
	ActionByAdmin(ctx context.Context, creator, nftSchemaName, tokenId, actionName, refId string, parameters []*nftmngrtypes.ActionParameter) (nftmngrtypes.ActionChangeList, error)
	AddAttributeKeeper(ctx context.Context, creator string, nftSchemaName string, new_add_attribute nftmngrtypes.AttributeDefinition, location nftmngrtypes.AttributeLocation) error
	UpdateAttributeKeeper(ctx context.Context, creator, nftSchemaName string, update_attribute nftmngrtypes.AttributeDefinition) error
	ResyncAttibutesKeeper(ctx context.Context, creator, nftSchemaName, tokenId string) error
	SetAttributeOveridingKeeper(ctx context.Context, creator, nftSchemaName string, newOveridingType int32) error
	ShowAttributeKeeper(ctx context.Context, creator, nftSchemaName string, status bool, attributesName []string) error
	ToggleActionKeeper(ctx context.Context, creator, nftSchemaName, actionName string, status bool) error
	UpdateActionKeeper(ctx context.Context, creator, nftSchemaName string, updateAction nftmngrtypes.Action) error
	AddActionKeeper(ctx context.Context, creator string, nftSchemaName string, newAction nftmngrtypes.Action) error
	CreateNewMetadataKeeper(ctx context.Context, creator, nftSchemaName, tokenId string, metadata nftmngrtypes.NftData) error
	CreateNftSchemaKeeper(ctx context.Context, creator string, schemaInput nftmngrtypes.NFTSchemaINPUT) error
	SetBaseURIKeeper(ctx context.Context, creator, nftSchemaName, baseURI string) error
	SetMetadataFormatKeeper(ctx context.Context, creator, nftSchemaName, format string) error
	SetMintAuthKeeper(ctx context.Context, creator, nftSchemaName string, authTo nftmngrtypes.AuthorizeTo) error
	SetOriginChainKeeper(ctx context.Context, creator, nftSchemaName, originChain string) error
	SetOriginContractKeeper(ctx context.Context, creator, nftSchemaName, contract string) error
	SetURIRetrievalKeeper(ctx context.Context, creator, nftSchemaName string, method int32) error
	ChangeOrgOwner(ctx context.Context, creator, newOwner, orgName string) error
	ChangeSchemaOwner(ctx context.Context, creator, newOwner, nftSchemaName string) error
	AddActionExecutor(ctx context.Context, creator string, nftSchemaName string, executorAddress string) error
	DelActionExecutor(ctx context.Context, creator, nftSchemaName, executorAddress string) error
	PerformVirtualActionKeeper(ctx context.Context, creator, vitualSchemaName string, tokenIdMap []*nftmngrtypes.TokenIdMap, actionName, refId string, parameters []*nftmngrtypes.ActionParameter) (changeList nftmngrtypes.ActionChangeList, err error)
	VoteVirtualSchemaProposalKeeper(ctx context.Context, creator, proposalId, srcNftSchemaCode string, option nftmngrtypes.RegistryStatus) error
	ProposalVirtualSchemaKeeper(ctx context.Context, creator, virtualNftSchemaCode string, proposalType nftmngrtypes.ProposalType, registryReq []*nftmngrtypes.VirtualSchemaRegistryRequest, actions []*nftmngrtypes.Action, executors []string, enable bool) (string, error)
}

type StakingKeeper interface {
	Delegate(goCtx context.Context, msg *stakingtypes.MsgDelegate) (*stakingtypes.MsgDelegateResponse, error)
	BeginRedelegate(goCtx context.Context, msg *stakingtypes.MsgBeginRedelegate) (*stakingtypes.MsgBeginRedelegateResponse, error)
	Undelegate(goCtx context.Context, msg *stakingtypes.MsgUndelegate) (*stakingtypes.MsgUndelegateResponse, error)
}

type StakingQuerier interface {
	Delegation(c context.Context, req *stakingtypes.QueryDelegationRequest) (*stakingtypes.QueryDelegationResponse, error)
}

type DistributionKeeper interface {
	DelegationRewards(c context.Context, req *disttypes.QueryDelegationRewardsRequest) (*disttypes.QueryDelegationRewardsResponse, error)
	DelegationTotalRewards(c context.Context, req *disttypes.QueryDelegationTotalRewardsRequest) (*disttypes.QueryDelegationTotalRewardsResponse, error)
	SetWithdrawAddr(ctx context.Context, delegatorAddr sdk.AccAddress, withdrawAddr sdk.AccAddress) error
	WithdrawDelegationRewards(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
}
