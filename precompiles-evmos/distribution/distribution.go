package distribution

import (
	"embed"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/evmos/evmos/v20/precompiles/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"
	evmtypes "github.com/evmos/evmos/v20/x/evm/types"
	stakingkeeper "github.com/evmos/evmos/v20/x/staking/keeper"
)

var (
	_ vm.PrecompiledContract = &Precompile{}
	_ cmn.Executor           = &DistributionExecutor{}
)

type Precompile struct {
	*cmn.Precompile
}

type DistributionExecutor struct {
	distributionKeeper distributionkeeper.Keeper
	stakingKeeper      stakingkeeper.Keeper
	authzKeeper        authzkeeper.Keeper

	precompile *Precompile
	address    common.Address
}

//go:embed abi.json
var f embed.FS

func GetABI() (abi.ABI, error) {
	return cmn.LoadABI(f, "abi.json")
}

func NewPrecompile(
	dk distributionkeeper.Keeper,
	sk stakingkeeper.Keeper,
	ak authzkeeper.Keeper,
) (*Precompile, error) {
	abi, err := GetABI()
	if err != nil {
		return nil, fmt.Errorf("error loading distribution ABI: %w", err)
	}
	precompile := &Precompile{}
	executor := &DistributionExecutor{
		distributionKeeper: dk,
		stakingKeeper:      sk,
		authzKeeper:        ak,
		address:            common.HexToAddress(evmtypes.DistributionPrecompileAddress),
		precompile:         precompile,
	}
	precompile.Precompile = cmn.NewPrecompile(abi, executor, executor.address, "dist")
	return precompile, nil
}

func NewDistributionExecutor(
	dk distributionkeeper.Keeper,
	sk stakingkeeper.Keeper,
	ak authzkeeper.Keeper,
) *DistributionExecutor {
	return &DistributionExecutor{
		distributionKeeper: dk,
		stakingKeeper:      sk,
		authzKeeper:        ak,
		address:            common.HexToAddress(evmtypes.DistributionPrecompileAddress),
	}
}

func (e *DistributionExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return cmn.DefaultGasCost(input, e.IsTransaction(method.Name))
}

func (e *DistributionExecutor) IsTransaction(methodName string) bool {
	switch methodName {
	case ClaimRewardsMethod,
		SetWithdrawAddressMethod,
		WithdrawDelegatorRewardsMethod,
		WithdrawValidatorCommissionMethod,
		FundCommunityPoolMethod:
		return true
	default:
		return false
	}
}

func (e *DistributionExecutor) Address() common.Address {
	// All methods are queries for this precompile
	return e.address
}

func (e *DistributionExecutor) GetABI() abi.ABI {
	// All methods are queries for this precompile
	abi, err := GetABI()
	if err != nil {
		panic(err)
	}
	return abi
}

func (e *DistributionExecutor) Execute(
	ctx sdk.Context,
	method *abi.Method,
	caller common.Address,
	callingContract common.Address,
	args []interface{},
	value *big.Int,
	readOnly bool,
	evm *vm.EVM,
) ([]byte, error) {
	switch method.Name {
	// Custom transactions
	case ClaimRewardsMethod:
		return e.ClaimRewards(ctx, evm.StateDB, evm.Origin, caller, method, args, value, readOnly)
	// Distribution transactions
	case SetWithdrawAddressMethod:
		return e.SetWithdrawAddress(ctx, evm.StateDB, evm.Origin, caller, method, args, value, readOnly)
	case WithdrawDelegatorRewardsMethod:
		return e.WithdrawDelegatorRewards(ctx, evm.StateDB, evm.Origin, caller, method, args, value, readOnly)
	case WithdrawValidatorCommissionMethod:
		return e.WithdrawValidatorCommission(ctx, evm.StateDB, evm.Origin, caller, method, args, value, readOnly)
	case FundCommunityPoolMethod:
		return e.FundCommunityPool(ctx, evm.StateDB, evm.Origin, caller, method, args, value, readOnly)
	// Distribution queries
	case ValidatorDistributionInfoMethod:
		return e.ValidatorDistributionInfo(ctx, callingContract, method, args)
	case ValidatorOutstandingRewardsMethod:
		return e.ValidatorOutstandingRewards(ctx, callingContract, method, args)
	case ValidatorCommissionMethod:
		return e.ValidatorCommission(ctx, callingContract, method, args)
	case ValidatorSlashesMethod:
		return e.ValidatorSlashes(ctx, callingContract, method, args)
	case DelegationRewardsMethod:
		return e.DelegationRewards(ctx, callingContract, method, args)
	case DelegationTotalRewardsMethod:
		return e.DelegationTotalRewards(ctx, callingContract, method, args)
	case DelegatorValidatorsMethod:
		return e.DelegatorValidators(ctx, callingContract, method, args)
	case DelegatorWithdrawAddressMethod:
		return e.DelegatorWithdrawAddress(ctx, callingContract, method, args)
	default:
		return nil, fmt.Errorf("distribution precompile: unknown method: %s", method.Name)
	}
}
