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

// TODO:: TEST(VirtualSchema)
// ON POC WE WILL JUST CREATE SCHEMA
func (k Keeper) AfterProposalCreateVirtualSchemaSuccess(ctx sdk.Context, virtualSchemaProposal types.VirtualSchemaProposal) (pass bool) {
	if len(virtualSchemaProposal.Registry) == 0 {
		return false
	}

	acceptCount, totalVotes := countProposalVotes(virtualSchemaProposal.Registry)
	voteThreshold := len(virtualSchemaProposal.Registry)

	if totalVotes != voteThreshold {
		return false
	}

	virtualSchema := types.VirtualSchema{
		VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchemaCode,
		Registry:             virtualSchemaProposal.Registry,
		Enable:               acceptCount == totalVotes,
	}

	k.SetVirtualSchema(ctx, virtualSchema)
	k.RemoveActiveVirtualSchemaProposal(ctx, virtualSchemaProposal.Id)
	k.SetInactiveVirtualSchemaProposal(ctx, types.InactiveVirtualSchemaProposal{Id: virtualSchemaProposal.Id})

	return true
}

func (k Keeper) AfterProposalDisableVirtualSchemaSuccess(ctx sdk.Context, disableVirtualSchemaProposal types.DisableVirtualSchemaProposal) (pass bool) {
	if len(disableVirtualSchemaProposal.Registry) == 0 {
		return false
	}

	acceptCount, totalVotes := countProposalVotes(disableVirtualSchemaProposal.Registry)
	voteThreshold := len(disableVirtualSchemaProposal.Registry)

	if totalVotes != voteThreshold {
		return false
	}

	virtualSchema := types.VirtualSchema{
		VirtualNftSchemaCode: disableVirtualSchemaProposal.VirtualSchemaCode,
		Registry:             disableVirtualSchemaProposal.Registry,
		Enable:               !(acceptCount == totalVotes),
	}

	k.SetVirtualSchema(ctx, virtualSchema)
	k.RemoveActiveDisableVirtualSchemaProposal(ctx, disableVirtualSchemaProposal.Id)
	k.SetInactiveDisableVirtualSchemaProposal(ctx, types.InactiveDisableVirtualSchemaProposal{Id: disableVirtualSchemaProposal.Id})

	return true
}

func (k Keeper) AfterProposalEnableVirtualSchemaSuccess(ctx sdk.Context, enableVirtualSchemaProposal types.EnableVirtualSchemaProposal) (pass bool) {
	if len(enableVirtualSchemaProposal.Registry) == 0 {
		return false
	}

	acceptCount, totalVotes := countProposalVotes(enableVirtualSchemaProposal.Registry)
	voteThreshold := len(enableVirtualSchemaProposal.Registry)

	if totalVotes != voteThreshold {
		return false
	}

	virtualSchema := types.VirtualSchema{
		VirtualNftSchemaCode: enableVirtualSchemaProposal.VirtualSchemaCode,
		Registry:             enableVirtualSchemaProposal.Registry,
		Enable:               acceptCount == totalVotes,
	}

	k.SetVirtualSchema(ctx, virtualSchema)
	k.RemoveActiveEnableVirtualSchemaProposal(ctx, enableVirtualSchemaProposal.Id)
	k.SetInactiveEnableVirtualSchemaProposal(ctx, types.InactiveEnableVirtualSchemaProposal{Id: enableVirtualSchemaProposal.Id})

	return true
}


func countProposalVotes(registry []*types.VirtualSchemaRegistry) (acceptCount, totalVotes int) {
	for _, reg := range registry {
		if reg.Status != types.RegistryStatus_PENDING {
			if reg.Status == types.RegistryStatus_ACCEPT {
				acceptCount++
			}
		}
	}

	for _, reg := range registry {
		if reg.Status != types.RegistryStatus_PENDING {
			totalVotes++
		}
	}
	return
}
