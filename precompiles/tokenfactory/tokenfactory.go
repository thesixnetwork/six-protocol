package tokenfactory

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
	"github.com/tendermint/tendermint/libs/log"

	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
	tokenmngr "github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	tokenmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

const (
	SendToCosmos     = "transferToCosmos"
	SendToCrossChain = "transferToCrossChain"
	UnwrapStakeToken = "unwrapStakeToken"
)

const (
	BridgeAddress = "0x0000000000000000000000000000000000001069"
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
	bankKeeper         pcommon.BankKeeper
	accountKeeper      pcommon.AccountKeeper
	tokenmngrKeeper    pcommon.TokenmngrKeeper
	tokenmngrMsgServer pcommon.TokenmngrMsgServer
	SendToCosmosID     []byte
	address            common.Address
}

func NewPrecompile(bankKeeper pcommon.BankKeeper, accountKeeper pcommon.AccountKeeper, tokenmngrKeeper pcommon.TokenmngrKeeper, tokennmngrMsgServer pcommon.TokenmngrMsgServer) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := &PrecompileExecutor{
		bankKeeper:         bankKeeper,
		accountKeeper:      accountKeeper,
		tokenmngrKeeper:    tokenmngrKeeper,
		tokenmngrMsgServer: tokennmngrMsgServer,
		address:            common.HexToAddress(BridgeAddress),
	}

	for name, m := range newAbi.Methods {
		switch name {
		case SendToCosmos:
			p.SendToCosmosID = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "tokenfactory"), nil
}

// RequiredGas returns the required bare minimum gas to execute the precompile.
func (p PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (p PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) (bz []byte, err error) {
	switch method.Name {
	case SendToCosmos:
		return p.sendToCosmos(ctx, caller, method, args, value, readOnly)
	case SendToCrossChain:
		return p.sendToCrossChain(ctx, caller, method, args, value, readOnly)
	case UnwrapStakeToken:
		return p.unwrapStakeToken(ctx, caller, method, args, value, readOnly)
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
		return method.Outputs.Pack(false)
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
	balance := p.bankKeeper.GetBalance(ctx, senderCosmoAddr, "asix")
	if balance.Amount.LT(intAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	// check total supply of evm denom
	supply := p.bankKeeper.GetSupply(ctx, tokenmngr.DefaultAttoDenom)
	if supply.Amount.LT(intAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	err = p.tokenmngrKeeper.AttoCoinConverter(ctx, senderCosmoAddr, receiverCosmoAddr, intAmount)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"precompile",
			sdk.NewAttribute(sdk.AttributeKeyModule, tokenmoduletypes.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, tokenmoduletypes.EventTypesConvertCoinToMicro),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyEvmSender, caller.Hex()),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyDestAddress, receiverCosmoAddr.String()),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyAmount, amount.String()),
		),
	})

	return method.Outputs.Pack(true)
}

func (p PrecompileExecutor) sendToCrossChain(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 4); err != nil {
		return nil, err
	}

	amount := args[1].(*big.Int)
	if amount.Cmp(utils.Big0) == 0 {
		// short circuit
		return method.Outputs.Pack(true)
	}

	memo, err := pcommon.StringFromArg(args[2])
	if err != nil {
		return nil, err
	}

	chain, err := pcommon.StringFromArg(args[3])
	if err != nil {
		return nil, err
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
	balance := p.bankKeeper.GetBalance(ctx, senderCosmoAddr, "asix")
	if balance.Amount.LT(intAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	// check total supply of evm denom
	supply := p.bankKeeper.GetSupply(ctx, "asix")
	if supply.Amount.LT(intAmount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	err = p.tokenmngrKeeper.AttoCoinConverter(ctx, senderCosmoAddr, receiverCosmoAddr, intAmount)
	if err != nil {
		return nil, err
	}

	// emit events
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"precompile",
			sdk.NewAttribute(sdk.AttributeKeyModule, tokenmoduletypes.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, tokenmoduletypes.EventTypesSentToCrossChain),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyEvmSender, caller.Hex()),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyDestChain, chain),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyMemo, memo),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyDestAddress, receiverCosmoAddr.String()),
			sdk.NewAttribute(tokenmoduletypes.AttributeKeyAmount, amount.String()),
		),
	})

	return method.Outputs.Pack(true)
}

// from evm-end usix consider as wrapedToken
// unwraped of evm-end = wrap from  cosmos-end
func (p PrecompileExecutor) unwrapStakeToken(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	var err error

	defer func() {
		if err != nil {
			ctx.Logger().Error("delegation precompile execution failed",
				"error", err.Error(),
			)
		}
	}()
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err = pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err = pcommon.ValidateArgsLength(args, 1); err != nil {
		return nil, err
	}

	amount := args[0].(*big.Int)
	if amount.Cmp(utils.Big0) == 0 {
		// short circuit
		return method.Outputs.Pack(true)
	}

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	msg := &tokenmoduletypes.MsgWrapToken{
		Creator:  senderCosmoAddr.String(),
		Receiver: senderCosmoAddr.String(),
		Amount:   sdk.NewCoin(tokenmngr.DefaultMicroDenom, sdk.NewIntFromBigInt(amount)),
	}

	_, err = p.tokenmngrMsgServer.WrapToken(sdk.WrapSDKContext(ctx), msg)
	if err != nil {
		return nil, err
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
