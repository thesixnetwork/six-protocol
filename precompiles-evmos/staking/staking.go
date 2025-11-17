// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/evmos/evmos/blob/main/LICENSE)

package staking

import (
	"embed"
	"fmt"
	"math/big"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v20/precompiles/authorization"
	cmn "github.com/evmos/evmos/v20/precompiles/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"
	"github.com/evmos/evmos/v20/x/evm/statedb"
	evmtypes "github.com/evmos/evmos/v20/x/evm/types"
	stakingkeeper "github.com/evmos/evmos/v20/x/staking/keeper"
)

var (
	_ vm.PrecompiledContract = &Precompile{}
	_ cmn.Executor           = &StakingExecutor{}
)

// Precompile defines the precompiled contract for staking.
type Precompile struct {
	*cmn.Precompile
}

// StakingExecutor is the implementation of the staking executor contract logic.
type StakingExecutor struct {
	stakingKeeper    stakingkeeper.Keeper
	authzKeeper      authzkeeper.Keeper
	expiration       time.Duration
	kvGasConfig      storetypes.GasConfig
	transientGasConf storetypes.GasConfig

	precompile *Precompile
	address    common.Address
}

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

// LoadABI loads the staking ABI from the embedded abi.json file
// for the staking precompile.
func GetABI() (abi.ABI, error) {
	return cmn.LoadABI(f, "abi.json")
}

// NewPrecompile creates a new staking Precompile instance as a
// PrecompiledContract interface.
func NewPrecompile(
	stakingKeeper stakingkeeper.Keeper,
	authzKeeper authzkeeper.Keeper,
) (*Precompile, error) {
	abi, err := GetABI()
	if err != nil {
		return nil, err
	}

	precompile := &Precompile{}
	executor := &StakingExecutor{
		stakingKeeper:    stakingKeeper,
		authzKeeper:      authzKeeper,
		address:          common.HexToAddress(evmtypes.StakingPrecompileAddress),
		expiration:       cmn.DefaultExpirationDuration,
		kvGasConfig:      storetypes.KVGasConfig(),
		transientGasConf: storetypes.TransientGasConfig(),
		precompile:       precompile,
	}

	precompile.Precompile = cmn.NewPrecompile(abi, executor, executor.address, "staking")
	return precompile, nil
}

// NewStakingExecutor creates a new instance of the StakingExecutor.
func NewStakingExecutor(
	stakingKeeper stakingkeeper.Keeper,
	authzKeeper authzkeeper.Keeper,
) *StakingExecutor {
	return &StakingExecutor{
		stakingKeeper:    stakingKeeper,
		authzKeeper:      authzKeeper,
		address:          common.HexToAddress(evmtypes.StakingPrecompileAddress),
		expiration:       cmn.DefaultExpirationDuration,
		kvGasConfig:      storetypes.KVGasConfig(),
		transientGasConf: storetypes.TransientGasConfig(),
	}
}

// RequiredGas returns the required gas for contract execution
func (e *StakingExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	if e.IsTransaction(method.Name) {
		return e.kvGasConfig.WriteCostFlat + (e.kvGasConfig.WriteCostPerByte * uint64(len(input)))
	}

	return e.kvGasConfig.ReadCostFlat + (e.kvGasConfig.ReadCostPerByte * uint64(len(input)))
}

// Execute implements the Executor interface
func (e *StakingExecutor) Execute(
	ctx sdk.Context,
	method *abi.Method,
	caller common.Address,
	callingContract common.Address,
	args []interface{},
	value *big.Int,
	readOnly bool,
	evm *vm.EVM,
) ([]byte, error) {
	stateDB, ok := evm.StateDB.(*statedb.StateDB)
	if !ok {
		return nil, fmt.Errorf("invalid StateDB type")
	}

	if readOnly && e.IsTransaction(method.Name) {
		return nil, fmt.Errorf("cannot call non-view method in read-only mode")
	}

	switch method.Name {
	// Authorization methods
	case authorization.ApproveMethod:
		return e.Approve(ctx, evm.Origin, stateDB, method, args)
	case authorization.RevokeMethod:
		return e.Revoke(ctx, evm.Origin, stateDB, method, args)
	case authorization.IncreaseAllowanceMethod:
		return e.IncreaseAllowance(ctx, evm.Origin, stateDB, method, args)
	case authorization.DecreaseAllowanceMethod:
		return e.DecreaseAllowance(ctx, evm.Origin, stateDB, method, args)
		// Staking transactions
	case CreateValidatorMethod:
		return e.CreateValidator(ctx, evm.Origin, caller, stateDB, method, args)
	case EditValidatorMethod:
		return e.EditValidator(ctx, evm.Origin, caller, stateDB, method, args)
	// Transactions
	case DelegateMethod:
		return e.Delegate(ctx, evm.Origin, caller, stateDB, method, args)
	case UndelegateMethod:
		return e.Undelegate(ctx, evm.Origin, caller, stateDB, method, args)
	case RedelegateMethod:
		return e.Redelegate(ctx, evm.Origin, caller, stateDB, method, args)
	case CancelUnbondingDelegationMethod:
		return e.CancelUnbondingDelegation(ctx, evm.Origin, caller, stateDB, method, args)
	case DelegationMethod:
		return e.Delegation(ctx, evm.Origin, method, args)
	case UnbondingDelegationMethod:
		return e.UnbondingDelegation(ctx, caller, method, args)
	case ValidatorMethod:
		return e.Validator(ctx, method, caller, args)
	case ValidatorsMethod:
		return e.Validators(ctx, method, caller, args)
	case RedelegationMethod:
		return e.Redelegation(ctx, method, caller, args)
	case RedelegationsMethod:
		return e.Redelegations(ctx, method, caller, args)
	case authorization.AllowanceMethod:
		return e.Allowance(ctx, method, caller, args)
	default:
		return nil, fmt.Errorf(cmn.ErrUnknownMethod, method.Name)
	}
}

// IsTransaction checks if the method is a transaction or not, depending on its name
func (e *StakingExecutor) IsTransaction(methodName string) bool {
	switch methodName {
	case CreateValidatorMethod,
		EditValidatorMethod,
		DelegateMethod,
		UndelegateMethod,
		RedelegateMethod,
		CancelUnbondingDelegationMethod,
		authorization.ApproveMethod,
		authorization.RevokeMethod,
		authorization.IncreaseAllowanceMethod,
		authorization.DecreaseAllowanceMethod:
		return true
	default:
		return false
	}
}

// Address implements the Executor interface.
func (e *StakingExecutor) Address() common.Address {
	return e.address
}

func (e *StakingExecutor) GetABI() abi.ABI {
	// All methods are queries for this precompile
	abi, err := GetABI()
	if err != nil {
		panic(err)
	}
	return abi
}

// Logger returns a precompile-specific logger.
func (p StakingExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("evm extension", "staking")
}
