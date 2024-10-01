package bridge

import (
	"bytes"
	"embed"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/evmos/ethermint/utils"
	tkmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"github.com/tendermint/tendermint/libs/log"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
)

const (
	SendToCosmos = "transferToCosmos"
)

const (
	BridgeAddress = "0x0000000000000000000000000000000000001069"

	tokenmngrModuleName = "tokenmngr"
	bridgeDiffTreshold  = 1
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
	bankKeeper     pcommon.BankKeeper
	accountKeeper  pcommon.AccountKeeper
	SendToCosmosID []byte
	address        common.Address
}

func NewPrecompile(bankKeeper pcommon.BankKeeper, accountKeeper pcommon.AccountKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := &PrecompileExecutor{
		bankKeeper:    bankKeeper,
		accountKeeper: accountKeeper,
		address:       common.HexToAddress(BridgeAddress),
	}

	for name, m := range newAbi.Methods {
		switch name {
		case SendToCosmos:
			p.SendToCosmosID = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "bridge"), nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) (bz []byte, err error) {
	switch method.Name {
	case SendToCosmos:
		return p.sendToCosmos(ctx, caller, method, args, value, readOnly)
	}
	return
}

func (p PrecompileExecutor) sendToCosmos(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	amount := args[1].(*big.Int)
	if amount.Cmp(utils.Big0) == 0 {
		// short circuit
		return method.Outputs.Pack(true)
	}

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}
	receiverCosmoAddr, err := p.accAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	// check if amount is valid
	intAmount := sdk.NewIntFromBigInt(amount)
	if intAmount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// ------------------------------------
	// |                                  |
	// |          CORE CONVERTOR          |
	// |                                  |
	// ------------------------------------

	// check if balance and input are valid
	if balance := p.bankKeeper.GetBalance(ctx, senderCosmoAddr, "asix"); balance.Amount.LT(intAmount) {
		// if current_balance + 1 >= inputAmount then convert all token of the account

		tresshold_balance := balance.Amount.Add(sdk.NewInt(bridgeDiffTreshold))
		if tresshold_balance.LT(intAmount) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
		}
		intAmount = balance.Amount
	}

	// check total supply of evm denom
	supply := p.bankKeeper.GetSupply(ctx, "asix")
	if supply.Amount.LT(intAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	//send evm coin to module account
	convertAmount := sdk.NewCoins(sdk.NewCoin("asix", intAmount))
	if err := p.bankKeeper.SendCoinsFromAccountToModule(ctx, senderCosmoAddr, tokenmngrModuleName, convertAmount); err != nil {
		return nil, sdkerrors.Wrap(tkmngrtypes.ErrSendCoinsFromAccountToModule, "Amount of token is too high than current balance due")
	}

	if err := p.bankKeeper.BurnCoins(ctx, tokenmngrModuleName, convertAmount); err != nil {
		return nil, sdkerrors.Wrap(tkmngrtypes.ErrBurnCoinsFromModuleAccount, "invali amount to burn")
	}

	// convert amount for burn to usix
	microSix := sdk.NewCoin("usix", intAmount.QuoRaw(1_000_000_000_000))

	// get the module account balance
	tokenmngrModuleAccount := p.accountKeeper.GetModuleAddress(tokenmngrModuleName)
	moduleBalance := p.bankKeeper.GetBalance(ctx, tokenmngrModuleAccount, "usix")

	// check if module account balance is enough to send
	if moduleBalance.Amount.LT(microSix.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "module account balance is not enough to send")
	}

	// send to receiver
	if err := p.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, tokenmngrModuleName, receiverCosmoAddr, sdk.NewCoins(microSix),
	); err != nil {
		return nil, sdkerrors.Wrap(tkmngrtypes.ErrSendCoinsFromAccountToModule, "unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account")
	}

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) accAddressFromBech32(arg interface{}) (bec32Addr sdk.AccAddress, err error) {
	addr := arg.(string)
	bec32Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("invalid addr")
	}
	return bec32Addr, nil
}

func (p PrecompileExecutor) accAddressFromArg(arg interface{}) (sdk.AccAddress, error) {
	addr := arg.(common.Address)
	if addr == (common.Address{}) {
		return nil, errors.New("invalid addr")
	}
	bec32Addr := utils.EthToCosmosAddr(addr)
	return bec32Addr, nil
}

func (PrecompileExecutor) IsTransaction(method string) bool {
	switch method {
	case SendToCosmos:
		return true
	default:
		return false
	}
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "bridge")
}
