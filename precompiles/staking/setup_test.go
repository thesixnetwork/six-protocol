package staking_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/precompiles/staking"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/thesixnetwork/six-protocol/precompiles/common"

	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	testkeeper "github.com/thesixnetwork/six-protocol/testutil/keeper"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	tokenmngrkeeper "github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	tokenmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	testutil "github.com/thesixnetwork/six-protocol/testutil"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// StakingPrecompileTestSuite defines the test suite for staking precompile
type StakingPrecompileTestSuite struct {
	suite.Suite
	ctx               sdk.Context
	stakingQuerier    common.StakingQuerier
	tokenmngrKeeper   *tokenmngrkeeper.Keeper
	stakingKeeper     *stakingkeeper.Keeper
	bankKeeper        bankkeeper.Keeper
	stakingPrecompile *staking.PrecompileExecutor

	// Additional fields for testing
	user         testutil.TestAccount
	validator    testutil.TestValidator
	srcValidator testutil.TestValidator
	dstValidator testutil.TestValidator
}

func TestStakingPrecompile(t *testing.T) {
	suite.Run(t, new(StakingPrecompileTestSuite))
}

func (s *StakingPrecompileTestSuite) SetupTest() {

	// Initialize context and keepers
	ctx, bankKeeper, stakingKeeper, tokenmngrKeeper, accountKeeper := testkeeper.NewAllKeepers(s.T())

	stakingModuleAcct := authtypes.NewEmptyModuleAccount(
		stakingtypes.ModuleName,
		authtypes.Burner, authtypes.Staking, authtypes.Minter,
	)

	accountKeeper.SetModuleAccount(ctx, stakingModuleAcct)

	mintModuleAcct := authtypes.NewEmptyModuleAccount("mint", authtypes.Minter)

	accountKeeper.SetModuleAccount(ctx, mintModuleAcct)

	acc := accountKeeper.GetModuleAccount(ctx, stakingtypes.ModuleName)
	if acc == nil {
		accountKeeper.SetModuleAccount(ctx, stakingModuleAcct)
	}

	s.Require().NotNil(acc, "staking module account is not set")

	accI := accountKeeper.GetAccount(ctx, acc.GetAddress())
	if accI == nil {
		panic("module account still missing")
	}

	s.ctx = ctx
	s.bankKeeper = bankKeeper
	s.stakingKeeper = stakingKeeper
	s.tokenmngrKeeper = tokenmngrKeeper

	// Setup staking querier
	s.stakingQuerier = testkeeper.NewStakingQuerier(stakingKeeper)

	// Create a delegator account (EVM user)
	user := testutil.CreateRandomAccount()
	s.user = user
	// fmt.Println("user:",user)

	mintCoin := sdk.NewCoin("asix", sdk.NewInt(1_000_000_000_000_000_000)) // 1,000 SIX
	mintCoins := sdk.NewCoins(mintCoin)

	// Mint coins to the 'mint' module account
	err := s.bankKeeper.MintCoins(s.ctx, "mint", mintCoins)
	s.Require().NoError(err)

	// Send coins from 'mint' to the user
	err = s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, "mint", s.user.Address, mintCoins)
	s.Require().NoError(err)

	balance := s.bankKeeper.GetBalance(s.ctx, s.user.Address, "asix")
	fmt.Println("User balance:", balance)

	// Fund user with enough tokens
	// coins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1_000_000_000_000)))
	coins := sdk.NewCoins(sdk.NewCoin("asix", sdk.NewInt(1_000_000_000_000_000_000)))

	err = testutil.FundAccount(bankKeeper, ctx, user.Address, coins)
	s.Require().NoError(err)

	// Set staking params before using BondDenom()
	stakingParams := stakingtypes.DefaultParams()
	stakingParams.BondDenom = "usix"
	s.stakingKeeper.SetParams(ctx, stakingParams)

	params := s.stakingKeeper.GetParams(s.ctx)
	params.MaxValidators = 2
	s.stakingKeeper.SetParams(s.ctx, params)

	// ----- Generate private keys for validators -----
	srcPriv, _ := testutil.GeneratePrivKeyAndValAddress("srcVal")
	dstPriv, _ := testutil.GeneratePrivKeyAndValAddress("dstVal")
	defaultPriv, _ := testutil.GeneratePrivKeyAndValAddress("val")

	// srcPriv, dstPriv, and defaultPriv are of type crypto.PrivKey
	srcAddr := sdk.AccAddress(srcPriv.PubKey().Address())
	dstAddr := sdk.AccAddress(dstPriv.PubKey().Address())
	defaultAddr := sdk.AccAddress(defaultPriv.PubKey().Address())

	// ✅ Approve and create srcValidator
	s.stakingKeeper.SetNewValidatorApprovalState(ctx, stakingtypes.ValidatorApproval{
		ApproverAddress: srcAddr.String(),
		Enabled:         true,
	})
	srcValidator := testutil.CreateValidatorWithPrivKey(ctx, stakingKeeper, bankKeeper, "srcVal", sdk.NewInt(10_000_000_000), srcPriv)

	// ✅ Approve and create dstValidator
	s.stakingKeeper.SetNewValidatorApprovalState(ctx, stakingtypes.ValidatorApproval{
		ApproverAddress: dstAddr.String(),
		Enabled:         true,
	})
	dstValidator := testutil.CreateValidatorWithPrivKey(ctx, stakingKeeper, bankKeeper, "dstVal", sdk.NewInt(10_000_000_000), dstPriv)

	// ✅ Approve and create defaultValidator
	s.stakingKeeper.SetNewValidatorApprovalState(ctx, stakingtypes.ValidatorApproval{
		ApproverAddress: defaultAddr.String(),
		Enabled:         true,
	})
	defaultValidator := testutil.CreateValidatorWithPrivKey(ctx, stakingKeeper, bankKeeper, "val", sdk.NewInt(10_000_000_000), defaultPriv)

	s.srcValidator = srcValidator
	s.dstValidator = dstValidator
	s.validator = defaultValidator

	srcVal, _ := s.stakingKeeper.GetValidator(s.ctx, s.srcValidator.Address)
	srcVal = srcVal.UpdateStatus(stakingtypes.Bonded)
	s.stakingKeeper.SetValidator(s.ctx, srcVal)

	dstVal, _ := s.stakingKeeper.GetValidator(s.ctx, s.dstValidator.Address)
	dstVal = dstVal.UpdateStatus(stakingtypes.Bonded)
	s.stakingKeeper.SetValidator(s.ctx, dstVal)

	// Wrap the staking keeper into MsgServer
	stakingWrapper := testkeeper.StakingKeeperWrapper{
		MsgServer: stakingkeeper.NewMsgServerImpl(*stakingKeeper),
		Keeper:    stakingKeeper,
	}

	RegisterAsixToken(s.ctx, s.tokenmngrKeeper, s.user.Address)

	token, found := s.tokenmngrKeeper.GetToken(s.ctx, "asix")
	fmt.Println("Token found:", found, "Token:", token)

	// Instantiate precompile executor
	precompile, err := staking.NewExecutor(stakingWrapper, s.stakingQuerier, bankKeeper, tokenmngrKeeper)
	s.Require().NoError(err)

	// ✅ Parse ABI and assign to precompile.Methods
	parsedABI := staking.GetABI() // ✅ already parsed ABI

	s.Require().NoError(err, "failed to parse staking precompile ABI")

	precompile.Methods = parsedABI.Methods

	s.stakingPrecompile = precompile

	srcValfound, found := s.stakingKeeper.GetValidator(s.ctx, s.srcValidator.Address)
	fmt.Println("srcValidator found:", found, "status:", srcValfound.Status) // Bonded = 2

	dstValfound, found := s.stakingKeeper.GetValidator(s.ctx, s.dstValidator.Address)
	fmt.Println("dstValidator found:", found, "status:", dstValfound.Status)



}

func RegisterAsixToken(ctx sdk.Context, k *tokenmngrkeeper.Keeper, creator sdk.AccAddress) {
	token := tokenmngrtypes.Token{
		Name:      "asix",
		Base:      "asix",                                                     // MUST match DefaultAttoDenom
		MaxSupply: sdk.NewCoin("asix", sdk.NewInt(1_000_000_000_000_000_000)), // arbitrary high supply
		Mintee:    creator.String(),
		Creator:   creator.String(),
	}

	k.SetToken(ctx, token)
}
