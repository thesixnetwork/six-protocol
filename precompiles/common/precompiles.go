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

// Operation is a type that defines if the precompile call
// produced an addition or subtraction of an account's balance
type Operation int8

const (
	Sub Operation = iota
	Add
)

type BalanceChangeEntry struct {
	Account common.Address
	Amount  *big.Int
	Op      Operation
}

// snapshot captures a MultiStore and Events for revert logic.
type snapshot struct {
	MultiStore storetypes.CacheMultiStore
	Events     sdk.Events
}

func NewBalanceChangeEntry(acc common.Address, amt *big.Int, op Operation) BalanceChangeEntry {
	return BalanceChangeEntry{acc, amt, op}
}

const UnknownMethodCallGas uint64 = 3000

type Executor interface {
	vm.ContractRef
	RequiredGas([]byte, *abi.Method) uint64
	Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) ([]byte, error)
}

type Precompile struct {
	executor Executor
	name     string
	abi.ABI
	address        common.Address
	balanceChanges []BalanceChangeEntry
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
		// This should never happen since this method is going to fail during Run
		return UnknownMethodCallGas
	}
	requiredGas := p.executor.RequiredGas(input[4:], method)
	return requiredGas
}

func (p *Precompile) Run(evm *vm.EVM, caller common.Address, callingContract common.Address, input []byte, value *big.Int, readOnly bool) (bz []byte, err error) {
	ctx, method, args, snap, err := p.Prepare(evm, input)
	if err != nil {
		return nil, err
	}
	em := ctx.EventManager()
	ctx = ctx.WithEventManager(sdk.NewEventManager())

	// Handle gas errors gracefully
	initialGas := ctx.GasMeter().GasConsumed()
	defer HandleGasError(ctx, evm, initialGas, &err)()

	// Clear any previous balance changes
	p.balanceChanges = nil

	bz, err = p.executor.Execute(ctx, method, caller, callingContract, args, value, readOnly, evm)
	if err != nil {
		return bz, err
	}
	events := ctx.EventManager().Events()
	if len(events) > 0 {
		em.EmitEvents(ctx.EventManager().Events())
	}

	// Apply balance changes to stateDB if any were recorded
	if len(p.balanceChanges) > 0 {
		stateDB, ok := evm.StateDB.(*statedb.StateDB)
		if ok {
			p.applyBalanceChanges(stateDB, snap)
		}
	}

	return bz, err
}

func (p Precompile) Prepare(evm *vm.EVM, input []byte) (sdk.Context, *abi.Method, []interface{}, snapshot, error) {
	var snap snapshot
	stateDB, ok := evm.StateDB.(*statedb.StateDB)

	if !ok {
		return sdk.Context{}, nil, nil, snap, errors.New("not run in EVM")
	}

	ctx, err := stateDB.GetCacheContext()
	if err != nil {
		return sdk.Context{}, nil, nil, snap, err
	}

	methodID, err := ExtractMethodID(input)
	if err != nil {
		return sdk.Context{}, nil, nil, snap, err
	}
	method, err := p.MethodById(methodID)
	if err != nil {
		return sdk.Context{}, nil, nil, snap, err
	}

	argsBz := input[4:]
	args, err := method.Inputs.Unpack(argsBz)
	if err != nil {
		return sdk.Context{}, nil, nil, snap, err
	}

	return ctx, method, args, snap, nil
}

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
		return errors.New("sending funds to a non-payable function")
	}

	return nil
}

func ExtractMethodID(input []byte) ([]byte, error) {
	// Check if the input has at least the length needed for methodID
	if len(input) < 4 {
		return nil, errors.New("input too short to extract method ID")
	}
	return input[:4], nil
}

func DefaultGasCost(input []byte, isTransaction bool) uint64 {
	if isTransaction {
		defaultGast := storetypes.KVGasConfig().WriteCostFlat + (storetypes.KVGasConfig().WriteCostPerByte * uint64(len(input)))
		return defaultGast
	}

	return storetypes.KVGasConfig().ReadCostFlat + (storetypes.KVGasConfig().ReadCostPerByte * uint64(len(input)))
}

// SetBalanceChangeEntries records balance changes that need to be applied to the EVM stateDB
// This prevents the stateDB from overwriting the changed balance in the bank keeper when committing the EVM state
func (p *Precompile) SetBalanceChangeEntries(entries ...BalanceChangeEntry) {
	p.balanceChanges = entries
}

// applyBalanceChanges applies the recorded balance changes to the EVM stateDB
func (p *Precompile) applyBalanceChanges(stateDB *statedb.StateDB, s snapshot) error {
	for _, entry := range p.balanceChanges {
		switch entry.Op {
		case Add:
			stateDB.AddBalance(entry.Account, entry.Amount)
		case Sub:
			stateDB.SubBalance(entry.Account, entry.Amount)
		}
	}
	return stateDB.AddPrecompileFn(p.Address(), s.MultiStore, s.Events)
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
