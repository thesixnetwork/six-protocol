package bank

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"math/big"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v20/x/evm/core/vm"

	"github.com/thesixnetwork/six-protocol/v4/utils"

	pcommon "github.com/thesixnetwork/six-protocol/v4/precompiles/common"
)

const (
	SendMethod        = "send"
	BalanceMethod     = "balance"
	AllBalancesMethod = "all_balances"
	NameMethod        = "name"
	SymbolMethod      = "symbol"
	DecimalsMethod    = "decimals"
	SupplyMethod      = "supply"
)

const (
	BankAddress = "0x0000000000000000000000000000000000001001"
)

// Embed abi json file to the executable binary. Needed when importing as dependency.
//
//go:embed abi.json
var f embed.FS

func GetABI() abi.ABI {
	abiBz, err := f.ReadFile("abi.json")
	if err != nil {
		panic(err)
	}

	newAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		panic(err)
	}
	return newAbi
}

type PrecompileExecutor struct {
	bankKeeper    pcommon.BankKeeper
	SendID        []byte
	BalanceID     []byte
	AllBalancesID []byte
	NameID        []byte
	SymbolID      []byte
	DecimalsID    []byte
	SupplyID      []byte
	address       common.Address
}

// Address implements common.PrecompileExecutor.
func (p *PrecompileExecutor) Address() common.Address {
	return p.address
}

type CoinBalance struct {
	Amount *big.Int
	Denom  string
}

func NewExecutor(bankKeeper pcommon.BankKeeper) *PrecompileExecutor {
	return &PrecompileExecutor{
		bankKeeper: bankKeeper,
		address:    common.HexToAddress(BankAddress),
	}
}

func NewPrecompile(bankKeeper pcommon.BankKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := NewExecutor(bankKeeper)

	for name, m := range newAbi.Methods {
		switch name {
		case SendMethod:
			p.SendID = m.ID
		case BalanceMethod:
			p.BalanceID = m.ID
		case AllBalancesMethod:
			p.AllBalancesID = m.ID
		case NameMethod:
			p.NameID = m.ID
		case SymbolMethod:
			p.SymbolID = m.ID
		case DecimalsMethod:
			p.DecimalsID = m.ID
		case SupplyMethod:
			p.SupplyID = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "bank"), nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) (bz []byte, err error) {
	switch method.Name {
	case SendMethod:
		return p.send(ctx, caller, method, args, value, readOnly)
	case BalanceMethod:
		return p.balance(ctx, method, args, value)
	case AllBalancesMethod:
		return p.all_balances(ctx, method, args, value)
	case NameMethod:
		return p.name(ctx, method, args, value)
	case SymbolMethod:
		return p.symbol(ctx, method, args, value)
	case DecimalsMethod:
		return p.decimals(ctx, method, args, value)
	case SupplyMethod:
		return p.totalSupply(ctx, method, args, value)
	}
	return
}

func (p PrecompileExecutor) send(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 4); err != nil {
		return nil, err
	}
	denom := args[2].(string)
	if denom == "" {
		return nil, errors.New("invalid denom")
	}
	amount := args[3].(*big.Int)
	if amount.Cmp(utils.Big0) == 0 {
		// short circuit
		return method.Outputs.Pack(true)
	}
	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}
	receiverCosmoAddr, err := p.accAddressFromArg(args[1])
	if err != nil {
		return nil, err
	}

	if err := p.bankKeeper.SendCoins(ctx, senderCosmoAddr, receiverCosmoAddr, sdk.NewCoins(sdk.NewCoin(denom, sdkmath.NewIntFromBigInt(amount)))); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) balance(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	addr, err := p.accAddressFromArg(args[0])
	if err != nil {
		return nil, err
	}
	denom := args[1].(string)
	if denom == "" {
		return nil, errors.New("invalid denom")
	}

	return method.Outputs.Pack(p.bankKeeper.GetBalance(ctx, addr, denom).Amount.BigInt())
}

func (p PrecompileExecutor) all_balances(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	addr, err := p.accAddressFromArg(args[0])
	if err != nil {
		return nil, err
	}

	coins := p.bankKeeper.GetAllBalances(ctx, addr)

	// convert to coin balance structs
	coinBalances := make([]CoinBalance, 0, len(coins))

	for _, coin := range coins {
		coinBalances = append(coinBalances, CoinBalance{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		})
	}

	return method.Outputs.Pack(coinBalances)
}

func (p PrecompileExecutor) name(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	denom := args[0].(string)
	metadata, found := p.bankKeeper.GetDenomMetaData(ctx, denom)
	if !found {
		return nil, fmt.Errorf("denom %s not found", denom)
	}
	return method.Outputs.Pack(metadata.Name)
}

func (p PrecompileExecutor) symbol(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	denom := args[0].(string)
	metadata, found := p.bankKeeper.GetDenomMetaData(ctx, denom)
	if !found {
		return nil, fmt.Errorf("denom %s not found", denom)
	}
	return method.Outputs.Pack(metadata.Symbol)
}

func (p PrecompileExecutor) decimals(_ sdk.Context, method *abi.Method, _ []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	// all native tokens are integer-based, returns decimals for microdenom (ucosmo)
	return method.Outputs.Pack(uint8(0))
}

func (p PrecompileExecutor) totalSupply(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	denom := args[0].(string)
	coin := p.bankKeeper.GetSupply(ctx, denom)
	return method.Outputs.Pack(coin.Amount.BigInt())
}

func (p PrecompileExecutor) accAddressFromArg(arg interface{}) (sdk.AccAddress, error) {
	addr := arg.(common.Address)
	if addr == (common.Address{}) {
		return nil, errors.New("invalid addr")
	}
	bec32Addr := utils.EthToCosmosAddr(addr)
	return bec32Addr, nil
}

func (p PrecompileExecutor) IsTransaction(method string) bool {
	switch method {
	case SendMethod:
		return true
	default:
		return false
	}
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "bank")
}
