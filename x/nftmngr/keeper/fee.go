package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	distribtype "github.com/cosmos/cosmos-sdk/x/distribution/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k Keeper) ProcessFeeAmount(ctx sdk.Context, bondedVotes []abci.VoteInfo) error {
	feeConfig, found := k.GetNFTFeeConfig(ctx)
	if !found {
		return nil
	}
	feeBalances, found := k.GetNFTFeeBalance(ctx)
	if !found {
		return nil
	}
	currentCreateSchemaFeeBalance, err := sdk.ParseCoinNormalized(feeBalances.FeeBalances[int32(types.FeeSubject_CREATE_NFT_SCHEMA)])
	if err != nil {
		return err
	}
	if currentCreateSchemaFeeBalance.Amount.GT(sdk.NewInt(0)) {
		// Loop over feeConfig.SchemaFee.FeeDistributions
		for _, feeDistribution := range feeConfig.SchemaFee.FeeDistributions {
			if feeDistribution.Method == types.FeeDistributionMethod_BURN {
				burnBalance := currentCreateSchemaFeeBalance.Amount.ToDec().Mul(sdk.NewDecWithPrec(int64(feeDistribution.Portion*100), 2)).TruncateInt()
				err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(currentCreateSchemaFeeBalance.Denom, burnBalance)))
				if err != nil {
					return err
				}
			} else if feeDistribution.Method == types.FeeDistributionMethod_REWARD_POOL {

				// totalPreviousPower = sum of all previous validators' power
				totalPreviousPower := int64(0)
				// Loop over votes
				for _, vote := range bondedVotes {
					totalPreviousPower += vote.Validator.Power
				}

				rewardBalance := currentCreateSchemaFeeBalance.Amount.ToDec().Mul(sdk.NewDecWithPrec(int64(feeDistribution.Portion*100), 2)).TruncateInt()
				reward := sdk.NewDecCoins(sdk.NewDecCoin(currentCreateSchemaFeeBalance.Denom, rewardBalance))

				err := k.bankKeeper.SendCoinsFromModuleToModule(
					ctx, types.ModuleName, distribtype.ModuleName,
					sdk.NewCoins(sdk.NewCoin(currentCreateSchemaFeeBalance.Denom, rewardBalance)))
				if err != nil {
					panic(err)
				}

				remaining := reward
				for i, vote := range bondedVotes {
					validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

					powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
					toAllocate := reward.MulDecTruncate(powerFraction)
					if i == len(bondedVotes)-1 {
						// last validator, allocate the remaining coins
						toAllocate = remaining
					} else {
						remaining = remaining.Sub(toAllocate)
					}
					k.distributionKeeper.AllocateTokensToValidator(ctx, validator, reward.MulDecTruncate(powerFraction))

				}
			}
		}

		// Set FeeBlance to 0
		feeBalances.FeeBalances[int32(types.FeeSubject_CREATE_NFT_SCHEMA)] = "0" + currentCreateSchemaFeeBalance.Denom
		k.SetNFTFeeBalance(ctx, feeBalances)
	}
	return nil
}

func (k Keeper) ValidateFeeConfig(feeConfig *types.NFTFeeConfig) error {
	// Validate Amount
	feeAmount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidFeeAmount, err.Error())
	}
	// check if feeAmount.Amount > 0
	if feeAmount.Amount.LTE(sdk.NewInt(0)) {
		return sdkerrors.Wrap(types.ErrInvalidFeeAmount, "Fee amount must be greater than 0")
	}
	// loop over feeConfig.SchemaFee.FeeDistributions
	totalPortion := float32(0)
	for _, feeDistribution := range feeConfig.SchemaFee.FeeDistributions {
		// validate portion
		if feeDistribution.Portion <= 0 || feeDistribution.Portion > 1 {
			return sdkerrors.Wrap(types.ErrInvalidFeePortion, "Fee portion must be between 0 and 1")
		}
		totalPortion += feeDistribution.Portion
	}
	if totalPortion != 1 {
		return sdkerrors.Wrap(types.ErrInvalidFeePortion, "Total fee portion must be equal to 1")
	}
	return nil
}
