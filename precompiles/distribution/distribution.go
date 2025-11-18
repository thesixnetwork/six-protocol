package distribution

import (
	"bytes"
	"embed"
	"errors"
	"math/big"

	"cosmossdk.io/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/evmos/v20/x/evm/core/vm"

	"github.com/thesixnetwork/six-protocol/v4/utils"

	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	pcommon "github.com/thesixnetwork/six-protocol/v4/precompiles/common"
)

const (
	SetWithdrawAddressMethod = "setWithdrawAddress"
	WithdrawRewardsMethod    = "withdrawRewards"
	RewardsMethod            = "rewards"
	AllRewardMethod          = "allRewards"
)

const (
	DistrAddress = "0x0000000000000000000000000000000000001007"
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
	distrKeeper     pcommon.DistributionKeeper
	distrQuerier    pcommon.DistributionQuerier
	tokenmngrKeeper pcommon.TokenmngrKeeper
	address         common.Address
	precompile      *pcommon.Precompile

	/*
	   #################
	   #### GETTER #####
	   #################
	*/
	RewardsMethodID    []byte
	AllRewardsMethodID []byte
	/*
	   #################
	   #### SETTER #####
	   #################
	*/
	SetWithdrawAddressMethodID []byte
	WithdrawRewardsMethodID    []byte
}

func NewExecutor(distKeeper pcommon.DistributionKeeper, distQuerier pcommon.DistributionQuerier, tokenmngrKeeper pcommon.TokenmngrKeeper) *PrecompileExecutor {
	return &PrecompileExecutor{
		distrKeeper:     distKeeper,
		distrQuerier:    distQuerier,
		tokenmngrKeeper: tokenmngrKeeper,
		address:         common.HexToAddress(DistrAddress),
	}
}

func NewPrecompile(distKeeper pcommon.DistributionKeeper, distQuerier pcommon.DistributionQuerier, tokenmngrKeeper pcommon.TokenmngrKeeper) (*pcommon.Precompile, error) {
	newAbi := GetABI()
	p := NewExecutor(distKeeper, distQuerier, tokenmngrKeeper)
	for name, m := range newAbi.Methods {
		switch name {
		case SetWithdrawAddressMethod:
			p.SetWithdrawAddressMethodID = m.ID
		case WithdrawRewardsMethod:
			p.WithdrawRewardsMethodID = m.ID
		case RewardsMethod:
			p.RewardsMethodID = m.ID
		case AllRewardMethod:
			p.AllRewardsMethodID = m.ID
		}
	}

	precompile := pcommon.NewPrecompile(newAbi, p, p.address, "distribution")
	p.precompile = precompile
	return precompile, nil
}

// Address implements common.PrecompileExecutor.
func (p *PrecompileExecutor) Address() common.Address {
	return p.address
}

func (p *PrecompileExecutor) Execute(ctx sdk.Context, method *abi.Method, caller common.Address, callingContract common.Address, args []interface{}, value *big.Int, readOnly bool, evm *vm.EVM) ([]byte, error) {
	switch method.Name {
	/*
		TODO: (@ddeedev): add balance state tracking
		NOTE: disable function relate with bank module on v4.0.0
		case WithdrawRewardsMethod:
			return p.withdrawRewards(ctx, caller, method, args, value, readOnly)
	*/
	case SetWithdrawAddressMethod:
		return p.setWithdrawAddressctx(ctx, caller, method, args, value, readOnly)
	case RewardsMethod:
		return p.rewards(ctx, method, args)
	case AllRewardMethod:
		return p.allRewards(ctx, method, args)
	}
	return nil, nil
}

func (p *PrecompileExecutor) setWithdrawAddressctx(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	var err error

	defer func() {
		if err != nil {
			ctx.Logger().Error("distribution precompile execution failed",
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

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	withdrawerAddressBech32 := args[0].(string)
	withdrawerAccAddress, err := p.accAddressFromBech32(withdrawerAddressBech32)
	if err != nil {
		return nil, err
	}

	err = p.distrKeeper.SetWithdrawAddr(ctx, senderCosmoAddr, withdrawerAccAddress)
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

func (p *PrecompileExecutor) withdrawRewards(ctx sdk.Context, caller common.Address, method *abi.Method, args []interface{}, value *big.Int, readOnly bool) ([]byte, error) {
	var err error

	defer func() {
		if err != nil {
			ctx.Logger().Error("distribution precompile execution failed",
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

	senderCosmoAddr, err := p.accAddressFromArg(caller)
	if err != nil {
		return nil, err
	}

	validatorAddressBech32 := args[0].(string)
	valAddress, err := p.valAddressFromBech32(validatorAddressBech32)
	if err != nil {
		return nil, err
	}

	coins, err := p.distrKeeper.WithdrawDelegationRewards(ctx, senderCosmoAddr, valAddress)
	if err != nil {
		return nil, err
	}

	// NOTE: This ensures that the changes in the bank keeper are correctly mirrored to the EVM stateDB.
	// This prevents the stateDB from overwriting the changed balance in the bank keeper when committing the EVM state.
	// This happens when the precompile is called from a smart contract
	if pcommon.ShouldTrackFromContract(caller, senderCosmoAddr) {
		tracker := pcommon.NewBalanceTracker(p.precompile)

		withdrawerHexAddr, err := p.getWithdrawerHexAddr(ctx, senderCosmoAddr)
		if err != nil {
			return nil, err
		}

		baseDenomAmount := pcommon.GetBaseDenomAmount(coins)
		tracker.TrackRewardWithdrawal(withdrawerHexAddr, baseDenomAmount, pcommon.BaseDenom)
	}

	return method.Outputs.Pack(true)
}

type Coin struct {
	Amount *big.Int
	Denom  string
}

type Reward struct {
	Coins            []Coin
	ValidatorAddress string
}

type Rewards struct {
	Reward []Reward
	Total  []Coin
}

func (p *PrecompileExecutor) rewards(ctx sdk.Context, method *abi.Method, args []interface{}) ([]byte, error) {
	var err error

	defer func() {
		if err != nil {
			ctx.Logger().Error("distribution precompile execution failed",
				"error", err.Error(),
			)
		}
	}()

	if err = pcommon.ValidateArgsLength(args, 2); err != nil {
		return nil, err
	}

	validatorAddressBech32 := args[0].(string)
	_, err = p.valAddressFromBech32(validatorAddressBech32)
	if err != nil {
		return nil, err
	}

	delegatorAddressBech32 := args[1].(string)
	_, err = p.accAddressFromBech32(delegatorAddressBech32)
	if err != nil {
		return nil, err
	}

	req := &distrtypes.QueryDelegationRewardsRequest{
		DelegatorAddress: delegatorAddressBech32,
		ValidatorAddress: validatorAddressBech32,
	}

	res, err := p.distrQuerier.DelegationRewards(ctx, req)
	if err != nil {
		return nil, err
	}

	coins := make([]Coin, 0, res.Rewards.Len())
	for _, rewardCoin := range res.Rewards {
		coins = append(coins, Coin{
			Amount: rewardCoin.Amount.BigInt(),
			Denom:  rewardCoin.Denom,
		})
	}

	reward := Reward{
		Coins:            coins,
		ValidatorAddress: validatorAddressBech32,
	}

	return method.Outputs.Pack(reward)
}

func (p PrecompileExecutor) allRewards(ctx sdk.Context, method *abi.Method, args []interface{}) (ret []byte, rerr error) {
	var err error

	defer func() {
		if err != nil {
			ctx.Logger().Error("distribution precompile execution failed",
				"error", err.Error(),
			)
		}
	}()

	err = pcommon.ValidateArgsLength(args, 1)
	if err != nil {
		return nil, err
	}

	delegatorAddress, err := p.accAddressFromBech32(args[0])
	if err != nil {
		return nil, err
	}

	req := &distrtypes.QueryDelegationTotalRewardsRequest{
		DelegatorAddress: delegatorAddress.String(),
	}

	response, err := p.distrQuerier.DelegationTotalRewards(ctx, req)
	if err != nil {
		return nil, err
	}

	rewardsOutput := getResponseOutput(response)
	return method.Outputs.Pack(rewardsOutput)
}

func getResponseOutput(response *distrtypes.QueryDelegationTotalRewardsResponse) Rewards {
	rewards := make([]Reward, 0, len(response.Rewards))
	for _, rewardInfo := range response.Rewards {
		coins := make([]Coin, 0, len(rewardInfo.Reward))
		for _, coin := range rewardInfo.Reward {
			coins = append(coins, Coin{
				Amount: coin.Amount.BigInt(),
				Denom:  coin.Denom,
			})
		}
		rewards = append(rewards, Reward{
			ValidatorAddress: rewardInfo.ValidatorAddress,
			Coins:            coins,
		})
	}

	totalCoins := make([]Coin, 0, len(response.Total))
	for _, coin := range response.Total {
		totalCoins = append(totalCoins, Coin{
			Amount: coin.Amount.BigInt(),
			Denom:  coin.Denom,
		})
	}

	return Rewards{
		Reward: rewards,
		Total:  totalCoins,
	}
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

func (p PrecompileExecutor) valAddressFromBech32(arg interface{}) (valAddress sdk.ValAddress, err error) {
	addr := arg.(string)
	valAddress, err = sdk.ValAddressFromBech32(addr)
	if err != nil {
		return nil, errors.New("invalid addr")
	}
	return valAddress, nil
}

func (p *PrecompileExecutor) RequiredGas(input []byte, method *abi.Method) uint64 {
	return pcommon.DefaultGasCost(input, p.IsTransaction(method.Name))
}

func (PrecompileExecutor) IsTransaction(method string) bool {
	switch method {
	case WithdrawRewardsMethod:
		return true
	case SetWithdrawAddressMethod:
		return true
	case RewardsMethod:
		return false
	default:
		return false
	}
}

func (p PrecompileExecutor) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("precompile", "distribution")
}

// getWithdrawerHexAddr is a helper function to get the hex address
// of the withdrawer for the specified account address
func (p PrecompileExecutor) getWithdrawerHexAddr(ctx sdk.Context, delegatorAddr sdk.AccAddress) (common.Address, error) {
	// Try to get the delegator's withdraw address
	withdrawerAccAddr, err := p.distrKeeper.GetDelegatorWithdrawAddr(ctx, delegatorAddr)
	if err != nil {
		// If GetDelegatorWithdrawAddr is not implemented or fails,
		// return the delegator address as the default withdrawer
		return utils.CosmosToEthAddr(delegatorAddr), nil
	}
	return utils.CosmosToEthAddr(withdrawerAccAddr), nil
}
