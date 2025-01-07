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

// VirtualSchemaHook processes virtual schema proposals and updates the state accordingly.
// It checks voting results and performs the requested operation (create/enable/disable/delete).
// Returns true if the proposal was successfully processed, false otherwise.
func (k Keeper) VirtualSchemaHook(ctx sdk.Context, virtualSchemaProposal types.VirtualSchemaProposal) bool {
	// Validate input
	if virtualSchemaProposal.VirtualSchemaCode == "" {
		k.Logger(ctx).Error("empty virtual schema code")
		return false
	}

	if len(virtualSchemaProposal.Registry) == 0 {
		k.Logger(ctx).Error("empty registry")
		return false
	}

	// Count votes
	acceptCount, totalVotes := countProposalVotes(virtualSchemaProposal.Registry)
	voteThreshold := len(virtualSchemaProposal.Registry)

	k.Logger(ctx).Info("Processing virtual schema proposal",
		"id", virtualSchemaProposal.Id,
		"type", virtualSchemaProposal.ProposalType,
		"accept_count", acceptCount,
		"total_votes", totalVotes,
		"threshold", voteThreshold)

	if totalVotes != voteThreshold {
		k.Logger(ctx).Info("Not all votes received yet")
		return false
	}

	// Process proposal
	var virtualSchema types.VirtualSchema
	var deleteProcess bool

	switch virtualSchemaProposal.ProposalType {
	case types.ProposalType_CREATE:
		virtualSchema = types.VirtualSchema{
			VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchemaCode,
			Registry:             virtualSchemaProposal.Registry,
			Enable:               acceptCount == totalVotes,
		}
	case types.ProposalType_ENABLE:
		virtualSchema = types.VirtualSchema{
			VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchemaCode,
			Registry:             virtualSchemaProposal.Registry,
			Enable:               acceptCount == totalVotes,
		}
	case types.ProposalType_DISABLE:
		virtualSchema = types.VirtualSchema{
			VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchemaCode,
			Registry:             virtualSchemaProposal.Registry,
			Enable:               !(acceptCount == totalVotes),
		}
	case types.ProposalType_DELETE:
		deleteProcess = true
	}

	// Update state
	if deleteProcess {
		k.Logger(ctx).Info("Deleting virtual schema",
			"code", virtualSchema.VirtualNftSchemaCode)
		k.RemoveVirtualSchema(ctx, virtualSchema.VirtualNftSchemaCode)
	} else {
		k.Logger(ctx).Info("Updating virtual schema",
			"code", virtualSchema.VirtualNftSchemaCode,
			"enabled", virtualSchema.Enable)
		k.SetVirtualSchema(ctx, virtualSchema)
	}

	// Update proposal status
	k.RemoveActiveVirtualSchemaProposal(ctx, virtualSchemaProposal.Id)
	k.SetInactiveVirtualSchemaProposal(ctx, types.InactiveVirtualSchemaProposal{Id: virtualSchemaProposal.Id})

	k.Logger(ctx).Info("Virtual schema proposal processed successfully", "id", virtualSchemaProposal.Id)
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
