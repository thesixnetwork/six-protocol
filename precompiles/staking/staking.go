package staking

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/evmos/ethermint/utils"
	pcommon "github.com/thesixnetwork/six-protocol/precompiles/common"
)

const (
	DelegateMethod   = "delegate"
	RedelegateMethod = "redelegate"
	UndelegateMethod = "undelegate"
	DelegationMethod = "delegation"
)

const (
	StakingAddress         = "0x0000000000000000000000000000000000001005"
	bridgeDiffTreshold     = 1
	defaultAttoToMicroDiff = 1_000_000_000_000
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
	pcommon.Precompile

	stakingKeeper   pcommon.StakingKeeper
	stakingQuerier  pcommon.StakingQuerier
	bankKeeper      pcommon.BankKeeper
	tokenmngrKeeper pcommon.TokenmngrKeeper
	address         common.Address

	/*
	   #################
	   #### GETTER #####
	   #################
	*/
	DelegationID []byte

	/*
	   #################
	   #### SETTER #####
	   #################
	*/
	DelegateID   []byte
	RedelegateID []byte
	UndelegateID []byte
}

func NewExecutor(stakingKeeper pcommon.StakingKeeper, stakingQuerier pcommon.StakingQuerier, bankKeeper pcommon.BankKeeper, tokenmngrKeeper pcommon.TokenmngrKeeper) (*PrecompileExecutor, error) {

	p := &PrecompileExecutor{
		stakingKeeper:   stakingKeeper,
		stakingQuerier:  stakingQuerier,
		bankKeeper:      bankKeeper,
		tokenmngrKeeper: tokenmngrKeeper,
		address:         common.HexToAddress(StakingAddress),
	}

	return p, nil

}

func NewPrecompile(stakingKeeper pcommon.StakingKeeper, stakingQuerier pcommon.StakingQuerier, bankKeeper pcommon.BankKeeper, tokenmngrKeeper pcommon.TokenmngrKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()

	p := &PrecompileExecutor{
		stakingKeeper:   stakingKeeper,
		stakingQuerier:  stakingQuerier,
		bankKeeper:      bankKeeper,
		tokenmngrKeeper: tokenmngrKeeper,
		address:         common.HexToAddress(StakingAddress),
	}

	for name, m := range newAbi.Methods {
		switch name {
		case DelegateMethod:
			p.DelegateID = m.ID
		case RedelegateMethod:
			p.RedelegateID = m.ID
		case UndelegateMethod:
			p.UndelegateID = m.ID
		case DelegationMethod:
			p.DelegationID = m.ID
		}
	}

	return pcommon.NewPrecompile(newAbi, p, p.address, "staking"), nil
}

func (p *PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) ([]byte, error) {
	switch method.Name {
	case DelegateMethod:
		return p.delegate(ctx, caller, method, args, value, readOnly)
	case RedelegateMethod:
		return p.redelegate(ctx, caller, method, args, value, readOnly)
	case UndelegateMethod:
		return p.undelegate(ctx, caller, method, args, value, readOnly)
	case DelegationMethod:
		return p.delegation(ctx, method, args, value)
	}
	return nil, nil
}

func (p *PrecompileExecutor) delegate(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
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

	if err = pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	validatorBech32 := args[0].(string)

	amount := args[1].(*big.Int)

	fmt.Println("amount:", amount)
	delegateAmount, err := p.convertCoinFromArg(amount)
	if err != nil {
		return nil, err
	}

	fmt.Println("delegateAmount:", delegateAmount)
	ctx.Logger().Info("delegation request", "amount", delegateAmount.String())

	// conver wei to staking coin
	err = p.convertWeiToStakingCoin(ctx, amount, senderCosmoAddr)
	if err != nil {
		return nil, err
	}

	_, err = p.stakingKeeper.Delegate(sdk.WrapSDKContext(ctx), &stakingtypes.MsgDelegate{
		DelegatorAddress: senderCosmoAddr.String(),
		ValidatorAddress: validatorBech32,
		Amount:           delegateAmount,
	})
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p *PrecompileExecutor) undelegate(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
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

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	validatorBech32 := args[0].(string)

	amount := args[1].(*big.Int)
	delegateAmount, err := p.convertCoinFromArg(amount)
	if err != nil {
		return nil, err
	}

	_, err = p.stakingKeeper.Undelegate(sdk.WrapSDKContext(ctx), &stakingtypes.MsgUndelegate{
		DelegatorAddress: senderCosmoAddr.String(),
		ValidatorAddress: validatorBech32,
		Amount:           delegateAmount,
	})
	if err != nil {
		return nil, err
	}

	// TODO: need to convert usix to asix after redelegation is done

	return method.Outputs.Pack(true)
}

func (p *PrecompileExecutor) redelegate(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
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

	if err = pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	srcValidatorBech32 := args[0].(string)
	dstValidatorBech32 := args[1].(string)

	// TODO: need to convert from asix to usix
	amount := args[2].(*big.Int)
	delegateAmount, err := p.convertCoinFromArg(amount)
	if err != nil {
		return nil, err
	}

	_, err = p.stakingKeeper.BeginRedelegate(sdk.WrapSDKContext(ctx), &stakingtypes.MsgBeginRedelegate{
		DelegatorAddress:    senderCosmoAddr.String(),
		ValidatorSrcAddress: srcValidatorBech32,
		ValidatorDstAddress: dstValidatorBech32,
		Amount:              delegateAmount,
	})
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

type Delegation struct {
	Balance    Balance
	Delegation DelegationDetails
}

type Balance struct {
	Amount *big.Int
	Denom  string
}

type DelegationDetails struct {
	DelegatorAddress string
	Shares           *big.Int
	Decimals         *big.Int
	ValidatorAddress string
}

func (p PrecompileExecutor) delegation(ctx sdk.Context, method *abi.Method, args []interface{}, value *big.Int) ([]byte, error) {
	var err error

	if err = pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err = pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	var senderCosmoAddr sdk.AccAddress
	senderCosmoAddr, err = p.accAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	validatorBech32 := args[1].(string)

	delegationRequest := &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: senderCosmoAddr.String(),
		ValidatorAddr: validatorBech32,
	}

	var delegationResponse *stakingtypes.QueryDelegationResponse
	delegationResponse, err = p.stakingQuerier.Delegation(sdk.WrapSDKContext(ctx), delegationRequest)
	if err != nil {
		return nil, err
	}

	delegation := Delegation{
		Balance: Balance{
			Amount: delegationResponse.GetDelegationResponse().GetBalance().Amount.BigInt(),
			Denom:  delegationResponse.GetDelegationResponse().GetBalance().Denom,
		},
		Delegation: DelegationDetails{
			DelegatorAddress: delegationResponse.GetDelegationResponse().GetDelegation().DelegatorAddress,
			Shares:           delegationResponse.GetDelegationResponse().GetDelegation().Shares.BigInt(),
			Decimals:         big.NewInt(sdk.Precision),
			ValidatorAddress: delegationResponse.GetDelegationResponse().GetDelegation().ValidatorAddress,
		},
	}

	return method.Outputs.Pack(delegation)
}

func (p *PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (p *PrecompileExecutor) accAddressFromArg(arg interface{}) (sdk.AccAddress, error) {
	addr := arg.(common.Address)
	if addr == (common.Address{}) {
		return nil, errors.New("invalid addr")
	}
	bec32Addr := utils.EthToCosmosAddr(addr)
	return bec32Addr, nil
}

func (p PrecompileExecutor) accAddressFromBech32(arg interface{}) (bec32Addr sdk.AccAddress, err error) {
	addr := arg.(string)
	bec32Addr, err = sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("invalid addr")
	}
	return bec32Addr, nil
}

func (p *PrecompileExecutor) convertCoinFromArg(amount *big.Int) (sdk.Coin, error) {
	if amount.Cmp(utils.Big0) == 0 {
		// short circuit
		return sdk.Coin{}, errors.New("invalid amount value")
	}

	intAmount := sdk.NewIntFromBigInt(amount)
	if intAmount.IsZero() {
		return sdk.Coin{}, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	microSix := sdk.NewCoin("usix", intAmount.QuoRaw(int64(defaultAttoToMicroDiff)))

	return microSix, nil
}

func (p *PrecompileExecutor) convertWeiToStakingCoin(ctx sdk.Context, weiAmount *big.Int, bech32Address sdk.AccAddress) error {
	// check if amount is valid
	intAmount := sdk.NewIntFromBigInt(weiAmount)
	if intAmount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// check if balance and input are valid
	if balance := p.bankKeeper.GetBalance(ctx, bech32Address, "asix"); balance.Amount.LT(intAmount) {
		// if current_balance + 1 >= inputAmount then convert all token of the account

		tresshold_balance := balance.Amount.Add(sdk.NewInt(bridgeDiffTreshold))
		if tresshold_balance.LT(intAmount) {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
		}
		intAmount = balance.Amount
	}

	// check total supply of evm denom
	supply := p.bankKeeper.GetSupply(ctx, "asix")
	if supply.Amount.LT(intAmount) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	// send convert coin to itself
	err := p.tokenmngrKeeper.AttoCoinConverter(ctx, bech32Address, bech32Address, intAmount)
	if err != nil {
		fmt.Println("err AttoCoinConverter : ", err)
		return err
	}

	return nil
}

func (PrecompileExecutor) IsTransaction(method string) bool {
	switch method {
	case DelegateMethod:
		return true
	case RedelegateMethod:
		return true
	case UndelegateMethod:
		return true
	case DelegationMethod:
		return false
	default:
		return false
	}
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "staking")
}
