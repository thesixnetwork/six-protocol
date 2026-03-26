package precisebank

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(
	ctx context.Context,
	keeper keeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	gs types.GenesisState,
) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Initialize module account
	if moduleAcc := ak.GetModuleAccount(sdkCtx, types.ModuleName); moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	// Check module balance matches sum of fractional balances + remainder
	totalAmt := gs.TotalAmountWithRemainder()
	moduleAddr := ak.GetModuleAddress(types.ModuleName)
	moduleBal := bk.GetBalance(sdkCtx, moduleAddr, types.IntegerCoinDenom)
	moduleBalExtended := moduleBal.Amount.Mul(types.ConversionFactor())

	if !totalAmt.Equal(moduleBalExtended) {
		panic(fmt.Sprintf(
			"module account balance does not match sum of fractional balances and remainder, balance is %s but expected %v%s (%v%s)",
			moduleBal,
			totalAmt, types.ExtendedCoinDenom,
			totalAmt.Quo(types.ConversionFactor()), types.IntegerCoinDenom,
		))
	}

	// Set FractionalBalances in state
	for _, bal := range gs.Balances {
		addr := sdk.MustAccAddressFromBech32(bal.Address)
		keeper.SetFractionalBalance(ctx, addr, bal.Amount)
	}

	// Set remainder amount in state
	keeper.SetRemainderAmount(ctx, gs.Remainder)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx context.Context, keeper keeper.Keeper) *types.GenesisState {
	balances := types.FractionalBalances{}

	keeper.IterateFractionalBalances(ctx, func(addr sdk.AccAddress, amount sdkmath.Int) bool {
		balances = append(balances, types.NewFractionalBalance(addr.String(), amount))
		return false
	})

	remainder := keeper.GetRemainderAmount(ctx)

	return types.NewGenesisState(balances, remainder)
}
