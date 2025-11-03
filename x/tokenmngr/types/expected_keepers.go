package types

import (
	"context"
	"time"

	addresscodec "cosmossdk.io/core/address"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"

	evmtypes "github.com/evmos/evmos/v20/x/evm/types"

	protocoladmintypes "github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

type ProtocoladminKeeper interface {
	GetAdmin(ctx context.Context, group string, admin string) (val protocoladmintypes.Admin, found bool)
	GetAllAdmin(ctx context.Context) (list []protocoladmintypes.Admin)
	Authenticate(ctx context.Context, group string, address string) bool
	GetGroup(ctx context.Context, group string) (val protocoladmintypes.Group, found bool)
}

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	GetModuleAddress(moduleName string) sdk.AccAddress
	AddressCodec() addresscodec.Codec
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	GetSupply(ctx context.Context, denom string) sdk.Coin
	GetDenomMetaData(ctx context.Context, denom string) (banktypes.Metadata, bool)
	SetDenomMetaData(ctx context.Context, denomMetaData banktypes.Metadata)
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoins(ctx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

type EVMKeeper interface {
	GetParams(ctx sdk.Context) evmtypes.Params
	SetParams(ctx sdk.Context, params evmtypes.Params) error
	GetStaticPrecompileInstance(params *evmtypes.Params, address common.Address) (vm.PrecompiledContract, bool, error)
	GetPrecompiles() map[common.Address]vm.PrecompiledContract
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// StakingKeeper defines the expected staking keeper methods
type StakingKeeper interface {
	ValidatorByConsAddr(context.Context, sdk.ConsAddress) (stakingtypes.ValidatorI, error)
	GetValidator(ctx context.Context, addr sdk.ValAddress) (stakingtypes.Validator, error)

	// Delegation methods
	GetDelegatorDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.Delegation, error)
	SetDelegation(ctx context.Context, delegation stakingtypes.Delegation) error
	RemoveDelegation(ctx context.Context, delegation stakingtypes.Delegation) error
	GetDelegatorBonded(ctx context.Context, delegator sdk.AccAddress) (math.Int, error)

	// Unbonding delegation methods
	GetUnbondingDelegations(ctx context.Context, delegator sdk.AccAddress, maxRetrieve uint16) ([]stakingtypes.UnbondingDelegation, error)
	SetUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error
	RemoveUnbondingDelegation(ctx context.Context, ubd stakingtypes.UnbondingDelegation) error
	InsertUBDQueue(ctx context.Context, ubd stakingtypes.UnbondingDelegation, completionTime time.Time) error
	GetDelegatorUnbonding(ctx context.Context, delegator sdk.AccAddress) (math.Int, error)
}

// DistributionKeeper defines the expected distribution keeper methods
type DistributionKeeper interface {
	// Delegator reward methods
	GetDelegatorStartingInfo(ctx context.Context, val sdk.ValAddress, del sdk.AccAddress) (distrtypes.DelegatorStartingInfo, error)
	SetDelegatorStartingInfo(ctx context.Context, val sdk.ValAddress, del sdk.AccAddress, period distrtypes.DelegatorStartingInfo) error
	DeleteDelegatorStartingInfo(ctx context.Context, val sdk.ValAddress, del sdk.AccAddress) error

	// Reward withdrawal
	WithdrawDelegationRewards(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)

	// Withdraw address management - key for address migration
	SetWithdrawAddr(ctx context.Context, delegatorAddr sdk.AccAddress, withdrawAddr sdk.AccAddress) error

	// Validator period methods needed for proper delegation initialization
	GetValidatorCurrentRewards(ctx context.Context, val sdk.ValAddress) (distrtypes.ValidatorCurrentRewards, error)
	GetValidatorHistoricalRewards(ctx context.Context, val sdk.ValAddress, period uint64) (distrtypes.ValidatorHistoricalRewards, error)
	SetValidatorHistoricalRewards(ctx context.Context, val sdk.ValAddress, period uint64, rewards distrtypes.ValidatorHistoricalRewards) error
}
