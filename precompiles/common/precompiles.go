// Package common contains shared types and utilities for EVM precompiled contracts.
package common

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

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
	address common.Address

	// mu serializes Run: precompile instances are process-global singletons
	// (registered once into the vm.PrecompiledContracts maps), so consensus
	// DeliverTx and concurrent JSON-RPC queries (eth_call/estimateGas) share
	// this instance — without the lock they race on balanceChanges. Executors
	// never re-enter the EVM, so holding the lock across Execute cannot
	// deadlock.
	mu             sync.Mutex
	balanceChanges []BalanceChangeEntry
}

var _ vm.PrecompiledContract = &Precompile{}

func NewPrecompile(a abi.ABI, executor Executor, address common.Address, name string) *Precompile {
	return &Precompile{ABI: a, executor: executor, address: address, name: name}
}

func (p *Precompile) RequiredGas(input []byte) uint64 {
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

	p.mu.Lock()
	defer p.mu.Unlock()

	// Clear any previous balance changes
	p.balanceChanges = nil

	if !readOnly {
		if stateDB, ok := evm.StateDB.(*statedb.StateDB); ok {
			if err := stateDB.AddPrecompileFn(p.address, snap.MultiStore, snap.Events); err != nil {
				return nil, err
			}
		}
	}

	bz, err = p.executor.Execute(ctx, method, caller, callingContract, args, value, readOnly, evm)
	if err != nil {
		return bz, err
	}
	events := ctx.EventManager().Events()
	if len(events) > 0 {
		em.EmitEvents(events)
	}

	// Apply any recorded balance changes so the EVM stateDB balance view stays
	// consistent with the bank state the executor just modified.
	if !readOnly {
		if stateDB, ok := evm.StateDB.(*statedb.StateDB); ok {
			p.applyBalanceChanges(stateDB)
		}
	}

	return bz, err
}

func (p *Precompile) Prepare(evm *vm.EVM, input []byte) (sdk.Context, *abi.Method, []interface{}, snapshot, error) {
	var snap snapshot
	stateDB, ok := evm.StateDB.(*statedb.StateDB)

	if !ok {
		return sdk.Context{}, nil, nil, snap, errors.New("not run in EVM")
	}

	ctx, err := stateDB.GetCacheContext()
	if err != nil {
		return sdk.Context{}, nil, nil, snap, err
	}

	// Capture a snapshot of the cache multi-store and events BEFORE this
	// precompile call mutates any cosmos state. This snapshot is registered in
	// the stateDB journal (see Run) so that an EVM revert of the calling frame
	// rolls the cosmos-state changes back to exactly this point. Without it a
	// reverted precompile call's writes would survive the unconditional
	// cacheCtx flush at Commit (state-through-revert, cf. Cosmos EVM
	// ASA-2026-002).
	snap.MultiStore = stateDB.MultiStoreSnapshot()
	snap.Events = ctx.EventManager().Events()

	// Flush the stateDB's pending EVM journal changes into the cache context so
	// the executor observes up-to-date balances/state for this call.
	if err := stateDB.CommitWithCacheCtx(); err != nil {
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

func (p *Precompile) GetABI() abi.ABI {
	return p.ABI
}

func (p *Precompile) Address() common.Address {
	return p.address
}

func (p *Precompile) GetName() string {
	return p.name
}

func (p *Precompile) GetExecutor() Executor {
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
// This prevents the stateDB from overwriting the changed balance in the bank keeper when committing the EVM state.
// Entries accumulate across calls within one Run (the slice is cleared at the start of each Run),
// so an executor performing several tracked operations does not clobber earlier entries.
func (p *Precompile) SetBalanceChangeEntries(entries ...BalanceChangeEntry) {
	p.balanceChanges = append(p.balanceChanges, entries...)
}

// applyBalanceChanges applies the recorded balance changes to the EVM stateDB.
// The revert-protection journal entry is registered in Run BEFORE Execute, not
// here, so that error/out-of-gas exits are covered too.
func (p *Precompile) applyBalanceChanges(stateDB *statedb.StateDB) {
	for _, entry := range p.balanceChanges {
		switch entry.Op {
		case Add:
			stateDB.AddBalance(entry.Account, entry.Amount)
		case Sub:
			stateDB.SubBalance(entry.Account, entry.Amount)
		}
	}
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
