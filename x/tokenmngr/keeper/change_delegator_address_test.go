package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/thesixnetwork/six-protocol/testutil/keeper"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	tokenmngrkeeper "github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
)

type ChangeDelegatorAddressTestSuite struct {
	suite.Suite

	ctx           sdk.Context
	keeper        *tokenmngrkeeper.Keeper
	stakingKeeper *stakingkeeper.Keeper
	// bankKeeper    bankkeeper.Keeper
	// accountKeeper authkeeper.AccountKeeper

	oldAddress sdk.AccAddress
	newAddress sdk.AccAddress
	validator  stakingtypes.Validator
}

func TestChangeDelegatorAddressTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeDelegatorAddressTestSuite))
}

func (suite *ChangeDelegatorAddressTestSuite) SetupTest() {

	// Create individual keepers using the test helpers with dependencies
	sk, ak, ctx := keeper.StakingKeeperWithDeps(suite.T())

	tk, _ := keeper.TokenmngrKeeperWithDeps(suite.T(), ak, sk)

	// Create a validator address and public key
	privKey := secp256k1.GenPrivKey()
	pubKey := privKey.PubKey()

	// Create validator address from sample address
	accAddr := sample.AccAddress() // This returns string
	valAddrBytes, _ := sdk.AccAddressFromBech32(accAddr)
	valAddr := sdk.ValAddress(valAddrBytes)

	validator, err := stakingtypes.NewValidator(valAddr.String(), pubKey, stakingtypes.Description{})
	suite.Require().NoError(err)
	validator.Status = stakingtypes.Bonded
	validator.Tokens = math.NewInt(1000000)
	validator.DelegatorShares = math.LegacyNewDec(1000000)
	suite.Require().NoError(sk.SetValidator(ctx, validator))
	suite.Require().NoError(sk.SetValidatorByConsAddr(ctx, validator))

	// Setup test addresses
	oldAddrStr := sample.AccAddress()
	newAddrStr := sample.AccAddress()
	oldAddr, _ := sdk.AccAddressFromBech32(oldAddrStr)
	newAddr, _ := sdk.AccAddressFromBech32(newAddrStr)

	// Setup test addresses
	suite.ctx = ctx
	suite.keeper = &tk
	suite.stakingKeeper = sk
	suite.validator = validator
	suite.oldAddress = oldAddr
	suite.newAddress = newAddr

	// Note: For this test, we'll create delegations directly in the staking keeper
	// to bypass the bank/funding requirements, since we're testing delegation address changes,
	// not the funding mechanics
}

func (suite *ChangeDelegatorAddressTestSuite) TestChangeDelegatorAddress_SuccessfulMigration() {
	// Instead of using the Delegate function (which requires bank operations),
	// we'll directly create a delegation record for testing purposes
	delegateAmount := math.NewInt(100)
	shares := math.LegacyNewDecFromInt(delegateAmount)

	// Get validator address
	valAddr, _ := sdk.ValAddressFromBech32(suite.validator.GetOperator())

	// Create delegation directly using NewDelegation
	delegation := stakingtypes.NewDelegation(suite.oldAddress.String(), valAddr.String(), shares)

	// Set the delegation directly in the staking keeper
	err := suite.stakingKeeper.SetDelegation(suite.ctx, delegation)
	suite.Require().NoError(err)

	// Sanity check: Verify the initial delegation exists
	retrievedDelegation, err := suite.stakingKeeper.GetDelegation(suite.ctx, suite.oldAddress, valAddr)
	suite.Require().NoError(err)
	suite.Require().Equal(shares, retrievedDelegation.Shares)

	// 2. Execute the address change
	err = suite.keeper.ChangeDelegatorAddress(suite.ctx, suite.oldAddress, suite.newAddress)
	suite.Require().NoError(err)

	// 3. Verify the results
	// Check that the old address has no delegations
	delegations, err := suite.stakingKeeper.GetDelegatorDelegations(suite.ctx, suite.oldAddress, 1)
	suite.Require().NoError(err)
	suite.Require().Empty(delegations, "old address should have no delegations after migration")

	// Check that the new address now has the delegation
	newDelegation, err := suite.stakingKeeper.GetDelegation(suite.ctx, suite.newAddress, valAddr)
	suite.Require().NoError(err)
	suite.Require().NotNil(newDelegation, "new address should have a delegation")

	// Check that the shares were correctly transferred
	suite.Require().Equal(shares, newDelegation.Shares, "shares should be equal")
}

func (suite *ChangeDelegatorAddressTestSuite) TestChangeDelegatorAddress_WithUnbondingDelegations() {
	// Setup: Create a regular delegation first
	delegateAmount := math.NewInt(200)
	shares := math.LegacyNewDecFromInt(delegateAmount)
	valAddr, _ := sdk.ValAddressFromBech32(suite.validator.GetOperator())

	// Create initial delegation
	delegation := stakingtypes.NewDelegation(suite.oldAddress.String(), valAddr.String(), shares)
	err := suite.stakingKeeper.SetDelegation(suite.ctx, delegation)
	suite.Require().NoError(err)

	// Create an unbonding delegation (simulating a user starting to undelegate)
	unbondingAmount := math.NewInt(50)
	completionTime := suite.ctx.BlockTime().AddDate(0, 0, 21) // 21 days from now

	// Create unbonding delegation entry
	unbondingEntry := stakingtypes.NewUnbondingDelegationEntry(
		suite.ctx.BlockHeight(),
		completionTime,
		unbondingAmount,
		uint64(1), // validator share index
	)

	// Create unbonding delegation
	unbondingDelegation := stakingtypes.UnbondingDelegation{
		DelegatorAddress: suite.oldAddress.String(),
		ValidatorAddress: valAddr.String(),
		Entries:          []stakingtypes.UnbondingDelegationEntry{unbondingEntry},
	}

	// Set unbonding delegation in staking keeper
	err = suite.stakingKeeper.SetUnbondingDelegation(suite.ctx, unbondingDelegation)
	suite.Require().NoError(err)

	// Verify initial state: old address has both delegation and unbonding delegation
	oldDelegations, err := suite.stakingKeeper.GetDelegatorDelegations(suite.ctx, suite.oldAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Len(oldDelegations, 1)

	oldUnbondingDelegations, err := suite.stakingKeeper.GetUnbondingDelegations(suite.ctx, suite.oldAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Len(oldUnbondingDelegations, 1)

	// Execute the address change
	err = suite.keeper.ChangeDelegatorAddress(suite.ctx, suite.oldAddress, suite.newAddress)
	suite.Require().NoError(err)

	// Verify results: old address should have no delegations or unbonding delegations
	finalOldDelegations, err := suite.stakingKeeper.GetDelegatorDelegations(suite.ctx, suite.oldAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Empty(finalOldDelegations, "old address should have no delegations")

	finalOldUnbonding, err := suite.stakingKeeper.GetUnbondingDelegations(suite.ctx, suite.oldAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Empty(finalOldUnbonding, "old address should have no unbonding delegations")

	// Verify results: new address should have both delegations and unbonding delegations
	newDelegations, err := suite.stakingKeeper.GetDelegatorDelegations(suite.ctx, suite.newAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Len(newDelegations, 1)
	suite.Require().Equal(shares, newDelegations[0].Shares)

	newUnbondingDelegations, err := suite.stakingKeeper.GetUnbondingDelegations(suite.ctx, suite.newAddress, 10)
	suite.Require().NoError(err)
	suite.Require().Len(newUnbondingDelegations, 1)
	suite.Require().Equal(suite.newAddress.String(), newUnbondingDelegations[0].DelegatorAddress)
	suite.Require().Equal(valAddr.String(), newUnbondingDelegations[0].ValidatorAddress)
	suite.Require().Len(newUnbondingDelegations[0].Entries, 1)
	suite.Require().Equal(unbondingAmount, newUnbondingDelegations[0].Entries[0].Balance)
}
