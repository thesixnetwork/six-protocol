package bank

import (
	"embed"
	"fmt"
	"math/big"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cmn "github.com/evmos/evmos/v20/precompiles/common"
	erc20keeper "github.com/evmos/evmos/v20/x/erc20/keeper"
	"github.com/evmos/evmos/v20/x/evm/core/vm"
	evmtypes "github.com/evmos/evmos/v20/x/evm/types"
)

const (
	GasBalanceOf   = 2_851
	GasTotalSupply = 2_477
	GasSupplyOf    = 2_477

	BalancesMethod    = "balances"
	TotalSupplyMethod = "totalSupply"
	SupplyOfMethod    = "supplyOf"
)

var (
	_ vm.PrecompiledContract = &Precompile{}
	_ cmn.Executor           = &BankExecutor{}
)

type Precompile struct {
	*cmn.Precompile
}

// BankExecutor implements business logic for precompile methods.
type BankExecutor struct {
	bankKeeper  bankkeeper.Keeper
	erc20Keeper erc20keeper.Keeper

	precompile *Precompile
	address    common.Address
}

//go:embed abi.json
var f embed.FS

func GetABI() (abi.ABI, error) {
	return cmn.LoadABI(f, "abi.json")
}

// Constructor for the precompile contract.
func NewPrecompile(bankKeeper bankkeeper.Keeper, erc20Keeper erc20keeper.Keeper) (*Precompile, error) {
	abi, err := GetABI()
	if err != nil {
		return nil, fmt.Errorf("error loading distribution ABI: %w", err)
	}
	precompile := &Precompile{}
	executor := &BankExecutor{
		bankKeeper:  bankKeeper,
		erc20Keeper: erc20Keeper,
		address:     common.HexToAddress(evmtypes.BankPrecompileAddress),
		precompile:  precompile,
	}
	precompile.Precompile = cmn.NewPrecompile(abi, executor, executor.address, "bank")
	return precompile, nil
}

func NewBankExecutor(bankKeeper bankkeeper.Keeper, erc20Keeper erc20keeper.Keeper) *BankExecutor {
	return &BankExecutor{
		bankKeeper:  bankKeeper,
		erc20Keeper: erc20Keeper,
		address:     common.HexToAddress(evmtypes.BankPrecompileAddress),
	}
}

// cmn.Executor interface implementations:
func (e *BankExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	switch method.Name {
	case BalancesMethod:
		return GasBalanceOf
	case TotalSupplyMethod:
		return GasTotalSupply
	case SupplyOfMethod:
		return GasSupplyOf
	default:
		return cmn.DefaultGasCost(input, false)
	}
}

func (e *BankExecutor) Execute(
	ctx sdk.Context,
	method *abi.Method,
	caller, callingContract common.Address,
	args []interface{},
	value *big.Int,
	readOnly bool,
	evm *vm.EVM,
) ([]byte, error) {
	switch method.Name {
	case BalancesMethod:
		return e.Balances(ctx, caller, method, args, value, readOnly)
	case TotalSupplyMethod:
		return e.TotalSupply(ctx, caller, method, args, value, readOnly)
	case SupplyOfMethod:
		return e.SupplyOf(ctx, caller, method, args, value, readOnly)
	default:
		return nil, fmt.Errorf("bank precompile: unknown method: %s", method.Name)
	}
}

func (e *BankExecutor) IsTransaction(method string) bool {
	return false
}

func (e *BankExecutor) Address() common.Address {
	return e.address
}

func (e *BankExecutor) GetABI() abi.ABI {
	// All methods are queries for this precompile
	abi, err := GetABI()
	if err != nil {
		panic(err)
	}
	return abi
}

// Logger returns a precompile-specific logger.
func (p BankExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("evm extension", "staking")
}
