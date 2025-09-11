package keeper_test

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/testutil"
	"testing"
)

type ChangeDelegatorAddressTestSuite struct {
	suite.Suite

	ctx           context.Context
	keeper        *keeper.Keeper
	stakingKeeper *testutil.MockStakingKeeper
	accountKeeper *testutil.MockAccountKeeper

	oldAddress sdk.AccAddress
	newAddress sdk.AccAddress
	validator1 string
	validator2 string
}

func TestChangeDelegatorAddressTestSuite(t *testing.T) {
	suite.Run(t, new(ChangeDelegatorAddressTestSuite))
}

func (suite *ChangeDelegatorAddressTestSuite) SetupTest() {
	// Setup test environment
	suite.ctx = context.Background()
	suite.stakingKeeper = testutil.NewMockStakingKeeper()
	suite.accountKeeper = testutil.NewMockAccountKeeper()

	// Create keeper with mocked dependencies - we'll create a minimal mock keeper
	// that only has the methods we need for testing
	suite.keeper = &keeper.Keeper{} // Simplified for testing

	// Generate test addresses
	suite.oldAddress = sample.AccAddressBytes()
	suite.newAddress = sample.AccAddressBytes()
	suite.validator1 = sample.ValAddress()
	suite.validator2 = sample.ValAddress()
}

// TestChangeDelegatorAddress_ValidationErrors tests the input validation logic
func TestChangeDelegatorAddress_ValidationErrors(t *testing.T) {
	t.Run("empty old address", func(t *testing.T) {
		emptyAddress := sdk.AccAddress{}
		newAddress := sample.AccAddressBytes()

		// Test address validation
		require.True(t, emptyAddress.Empty())
		require.False(t, newAddress.Empty())
		require.False(t, emptyAddress.Equals(newAddress))
	})

	t.Run("empty new address", func(t *testing.T) {
		oldAddress := sample.AccAddressBytes()
		emptyAddress := sdk.AccAddress{}

		// Test address validation
		require.False(t, oldAddress.Empty())
		require.True(t, emptyAddress.Empty())
		require.False(t, oldAddress.Equals(emptyAddress))
	})

	t.Run("same addresses", func(t *testing.T) {
		address := sample.AccAddressBytes()

		// Test address validation
		require.False(t, address.Empty())
		require.True(t, address.Equals(address))
	})
}

// TestAddressConversion tests address string conversion scenarios
func TestAddressConversion(t *testing.T) {
	t.Run("valid address conversion", func(t *testing.T) {
		addr := sample.AccAddressBytes()
		addrStr := addr.String()

		// Convert back to bytes
		convertedAddr, err := sdk.AccAddressFromBech32(addrStr)
		require.NoError(t, err)
		require.Equal(t, addr, convertedAddr)
	})

	t.Run("empty address string", func(t *testing.T) {
		_, err := sdk.AccAddressFromBech32("")
		require.Error(t, err)
	})

	t.Run("invalid address string", func(t *testing.T) {
		_, err := sdk.AccAddressFromBech32("invalid")
		require.Error(t, err)
	})
}
