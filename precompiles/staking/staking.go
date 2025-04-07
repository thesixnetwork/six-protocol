package staking

import (
	"bytes"
	"embed"
	"errors"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	StakingAddress = "0x0000000000000000000000000000000000001005"
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
	stakingKeeper  pcommon.StakingKeeper
	stakingQuerier pcommon.StakingQuerier
	bankKeeper     pcommon.BankKeeper
	address        common.Address

	/*
	   #################
	   #### GETTER #####
	   #################
	*/

	/*
	   #################
	   #### SETTER #####
	   #################
	*/
	DelegateID   []byte
	RedelegateID []byte
	UndelegateID []byte
	DelegationID []byte
}

func NewPrecompile(stakingKeeper pcommon.StakingKeeper, stakingQuerier pcommon.StakingQuerier, bankKeeper pcommon.BankKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()

	p := &PrecompileExecutor{
		stakingKeeper:  stakingKeeper,
		stakingQuerier: stakingQuerier,
		bankKeeper:     bankKeeper,
		address:        common.HexToAddress(StakingAddress),
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
	if value == nil || value.Sign() == 0 {
		return nil, errors.New("set `value` field to non-zero to send delegate fund")
	}

	// TODO: need to convert from asix to usix
	amount := args[1].(*big.Int)
	delegateAmount, err := p.convertCoinFromArg(amount)
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
	if value == nil || value.Sign() == 0 {
		return nil, errors.New("set `value` field to non-zero to send delegate fund")
	}

	// TODO: need to convert from asix to usix
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

	return method.Outputs.Pack(true)
}

func (p *PrecompileExecutor) redelegate(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	if readOnly {
		return nil, errors.New("cannot call send from staticcall")
	}

	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 3); err != nil {
		return nil, err
	}

	if value == nil || value.Sign() == 0 {
		return nil, errors.New("set `value` field to non-zero to send delegate fund")
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
	if err := pcommon.ValidateNonPayable(value); err != nil {
		return nil, err
	}

	if err := pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	senderCosmoAddr, err := p.accAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	validatorBech32 := args[1].(string)

	delegationRequest := &stakingtypes.QueryDelegationRequest{
		DelegatorAddr: senderCosmoAddr.String(),
		ValidatorAddr: validatorBech32,
	}

	delegationResponse, err := p.stakingQuerier.Delegation(sdk.WrapSDKContext(ctx), delegationRequest)
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

	convAmount := sdk.NewCoin("usix", sdk.NewIntFromBigInt(amount))

	return convAmount, nil
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
