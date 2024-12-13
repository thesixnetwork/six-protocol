package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
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

func (k Keeper) AfterProposalSubmission(ctx sdk.Context, proposalId string) {
}

func (k Keeper) AfterVoteProposal(ctx sdk.Context, proposalId string) {
}

// TODO:: Feat(VirtualSchema)
// ON POC WE WILL JUST CREATE SCHEMA
func (k Keeper) AfterProposalSuccess(ctx sdk.Context, proposalId string) {
	// validate if proposal complete

	// get virtual info from proposal
	virtualSchemaProposal, _ := k.GetVirtualSchemaProposal(ctx, proposalId)

	// Count votes
	var (
		acceptCount  int
		totalVotes   int
		voteTreshold = len(virtualSchemaProposal.Registry)
	)

	for _, registry := range virtualSchemaProposal.Registry {
		if registry.Status == types.RegistryStatus_ACCEPT {
			acceptCount++
		}
		if registry.Status != types.RegistryStatus_PENDING {
			totalVotes++
		}
	}

	var virSchema = types.VirtualSchema{
		VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchemaCode,
		Registry:             virtualSchemaProposal.Registry,
		Enable:               true,
		ExpiredAtBlock:       "0",
	}

	if acceptCount == voteTreshold {
		virSchema.Enable = true
	} else {
		virSchema.Enable = false
	}

	k.SetVirtualSchema(
		ctx,
		virSchema,
	)
}
