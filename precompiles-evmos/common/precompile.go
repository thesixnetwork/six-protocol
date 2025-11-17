package common

import (
	"errors"
	"fmt"
	"math/big"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"
	"github.com/evmos/evmos/v20/x/evm/statedb"
)

const UnknownMethodCallGas uint64 = 3000

// snapshot captures a MultiStore and Events for revert logic.
type snapshot struct {
	MultiStore storetypes.CacheMultiStore
	Events     sdk.Events
}

// Executor defines the contract logic interface.
type Executor interface {
	vm.ContractRef
	RequiredGas([]byte, *abi.Method) uint64
	Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) ([]byte, error)
}

// Precompile provides EVM precompile logic with journaling.
type Precompile struct {
	executor Executor
	name     string
	abi.ABI
	address        common.Address
	journalEntries []BalanceChangeEntry
}

var _ vm.PrecompiledContract = &Precompile{}

func NewPrecompile(a abi.ABI, executor Executor, address common.Address, name string) *Precompile {
	return &Precompile{ABI: a, executor: executor, address: address, name: name}
}

func (p Precompile) RequiredGas(input []byte) uint64 {
	methodID, err := ExtractMethodID(input)
	if err != nil {
		return UnknownMethodCallGas
	}
	method, err := p.MethodById(methodID)
	if err != nil {
		return UnknownMethodCallGas
	}
	return p.executor.RequiredGas(input[4:], method)
}

// Run executes the precompile, with gas error handling and journaling.
func (p *Precompile) Run(
	evm *vm.EVM,
	caller common.Address,
	callingContract common.Address,
	input []byte,
	value *big.Int,
	readOnly bool,
) ([]byte, error) {
	err := ValidateNonPayable(value)
	if err != nil {
		return nil, err
	}

	// NOTE: This is a special case where the calling transaction does not specify a function name.
	// In this case we default to a `fallback` or `receive` function on the contract.

	// Simplify the calldata checks
	isEmptyCallData := len(input) == 0
	isShortCallData := len(input) > 0 && len(input) < 4
	isStandardCallData := len(input) >= 4

	var method *abi.Method

	switch {
	// Case 1: Calldata is empty
	case isEmptyCallData:
		method, err = p.emptyCallData(value)

	// Case 2: calldata is non-empty but less than 4 bytes needed for a method
	case isShortCallData:
		method, err = p.methodIDCallData()

	// Case 3: calldata is non-empty and contains the minimum 4 bytes needed for a method
	case isStandardCallData:
		method, err = p.standardCallData(input)
	}
	if err != nil {
		return nil, err
	}

	ctx, args, snap, err := p.Prepare(evm, input, method)
	if err != nil {
		return nil, err
	}
	em := ctx.EventManager()
	ctx = ctx.WithEventManager(sdk.NewEventManager())

	// Handle gas errors gracefully
	initialGas := ctx.GasMeter().GasConsumed()
	defer HandleGasError(ctx, evm, initialGas, &err)()

	// Execute logic
	bz, execErr := p.executor.Execute(ctx, method, caller, callingContract, args, value, readOnly, evm)
	if execErr != nil {
		return bz, execErr
	}

	events := ctx.EventManager().Events()
	if len(events) > 0 {
		em.EmitEvents(events)
	}

	// Add journal entries if present
	if len(p.journalEntries) > 0 {
		if e := p.AddJournalEntries(evm.StateDB.(*statedb.StateDB), snap); e != nil {
			return bz, e
		}
	}
	return bz, nil
}

// Prepare returns context, method, arguments, and a snapshot for journaling.
func (p *Precompile) Prepare(evm *vm.EVM, input []byte, method *abi.Method) (sdk.Context, []interface{}, snapshot, error) {
	var snap snapshot
	stateDB, ok := evm.StateDB.(*statedb.StateDB)
	if !ok {
		return sdk.Context{}, nil, snap, errors.New("not run in EVM")
	}

	// get the stateDB cache ctx
	ctx, err := stateDB.GetCacheContext()
	if err != nil {
		return sdk.Context{}, nil, snap, err
	}
	// ctx := stateDB.GetContext()
	snap.MultiStore = stateDB.MultiStoreSnapshot()
	snap.Events = ctx.EventManager().Events()

	// commit the current changes in the cache ctx
	// to get the updated state for the precompile call
	if err := stateDB.CommitWithCacheCtx(); err != nil {
		return sdk.Context{}, nil, snap, err
	}

	argsBz := input[4:]
	args, err := method.Inputs.Unpack(argsBz)
	if err != nil {
		return sdk.Context{}, nil, snap, err
	}
	return ctx, args, snap, nil
}

// AddJournalEntries records balance changes and state snapshot for potential revert.
func (p *Precompile) AddJournalEntries(stateDB *statedb.StateDB, s snapshot) error {
	for _, entry := range p.journalEntries {
		switch entry.Op {
		case Sub:
			stateDB.SubBalance(entry.Account, entry.Amount)
		case Add:
			stateDB.AddBalance(entry.Account, entry.Amount)
		}
	}
	return stateDB.AddPrecompileFn(p.Address(), s.MultiStore, s.Events)
}

// SetBalanceChangeEntries sets entries for journaling.
func (p *Precompile) SetBalanceChangeEntries(entries ...BalanceChangeEntry) {
	p.journalEntries = entries
}

// HandleGasError resets the gas meter and returns an error if out of gas (use in defer).
func HandleGasError(ctx sdk.Context, evm *vm.EVM, initialGas storetypes.Gas, err *error) func() {
	return func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case storetypes.ErrorOutOfGas:
				// usedGas := ctx.GasMeter().GasConsumed() - initialGas
				// _ = evm.UseGas(usedGas)
				*err = vm.ErrOutOfGas
				ctx = ctx.WithKVGasConfig(storetypes.GasConfig{}).
					WithTransientKVGasConfig(storetypes.GasConfig{})
			default:
				panic(r)
			}
		}
	}
}

// --- Utility functions ---

func (p Precompile) GetABI() abi.ABI {
	return p.ABI
}

func (p Precompile) Address() common.Address {
	return p.address
}

func (p Precompile) GetName() string {
	return p.name
}

func (p Precompile) GetExecutor() Executor {
	return p.executor
}

func ValidateArgsLength(args []interface{}, length int) error {
	if len(args) != length {
		return fmt.Errorf("expected %d arguments but got %d", length, len(args))
	}
	return nil
}

func ValidateNonPayable(value *big.Int) error {
	if value != nil && value.Sign() != 0 {
		return vm.ErrExecutionReverted
	}
	return nil
}

func ExtractMethodID(input []byte) ([]byte, error) {
	if len(input) < 4 {
		return nil, errors.New("input too short to extract method ID")
	}
	return input[:4], nil
}

func DefaultGasCost(input []byte, isTransaction bool) uint64 {
	if isTransaction {
		return storetypes.KVGasConfig().WriteCostFlat + (storetypes.KVGasConfig().WriteCostPerByte * uint64(len(input)))
	}
	return storetypes.KVGasConfig().ReadCostFlat + (storetypes.KVGasConfig().ReadCostPerByte * uint64(len(input)))
}

// emptyCallData is a helper function that returns the method to be called when the calldata is empty.
func (p Precompile) emptyCallData(value *big.Int) (method *abi.Method, err error) {
	switch {
	// Case 1.1: Send call or transfer tx - 'receive' is called if present and value is transferred
	case value.Sign() > 0 && p.HasReceive():
		return &p.Receive, nil
	// Case 1.2: Either 'receive' is not present, or no value is transferred - call 'fallback' if present
	case p.HasFallback():
		return &p.Fallback, nil
	// Case 1.3: Neither 'receive' nor 'fallback' are present - return error
	default:
		return nil, vm.ErrExecutionReverted
	}
}

// methodIDCallData is a helper function that returns the method to be called when the calldata is less than 4 bytes.
func (p Precompile) methodIDCallData() (method *abi.Method, err error) {
	// Case 2.2: calldata contains less than 4 bytes needed for a method and 'fallback' is not present - return error
	if !p.HasFallback() {
		return nil, vm.ErrExecutionReverted
	}
	// Case 2.1: calldata contains less than 4 bytes needed for a method - 'fallback' is called if present
	return &p.Fallback, nil
}

// standardCallData is a helper function that returns the method to be called when the calldata is 4 bytes or more.
func (p Precompile) standardCallData(input []byte) (method *abi.Method, err error) {
	methodID := input[:4]
	// NOTE: this function iterates over the method map and returns
	// the method with the given ID
	method, err = p.MethodById(methodID)

	// Case 3.1 calldata contains a non-existing method ID, and `fallback` is not present - return error
	if err != nil && !p.HasFallback() {
		return nil, err
	}

	// Case 3.2: calldata contains a non-existing method ID - 'fallback' is called if present
	if err != nil && p.HasFallback() {
		return &p.Fallback, nil
	}

	return method, nil
}
