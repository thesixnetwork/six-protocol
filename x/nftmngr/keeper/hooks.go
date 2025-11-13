package keeper

import (
	"context"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	distribtype "github.com/cosmos/cosmos-sdk/x/distribution/types"
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

	if currentCreateSchemaFeeBalance.Amount.GT(sdkmath.NewInt(0)) {
		// Loop over feeConfig.SchemaFee.FeeDistributions
		for _, feeDistribution := range feeConfig.SchemaFee.FeeDistributions {
			switch feeDistribution.Method {
			case types.FeeDistributionMethod_BURN:
				burnBalance := currentCreateSchemaFeeBalance.Amount.ToLegacyDec().Mul(sdkmath.LegacyNewDecWithPrec(int64(feeDistribution.Portion*100), 2)).TruncateInt()
				err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(currentCreateSchemaFeeBalance.Denom, burnBalance)))
				if err != nil {
					return err
				}
			case types.FeeDistributionMethod_REWARD_POOL:

				// totalPreviousPower = sum of all previous validators' power
				totalPreviousPower := int64(0)
				// Loop over votes
				for _, vote := range bondedVotes {
					totalPreviousPower += vote.Validator.Power
				}

				rewardBalance := currentCreateSchemaFeeBalance.Amount.ToLegacyDec().Mul(sdkmath.LegacyNewDecWithPrec(int64(feeDistribution.Portion*100), 2)).TruncateInt()
				reward := sdk.NewDecCoins(sdk.NewDecCoin(currentCreateSchemaFeeBalance.Denom, rewardBalance))

				err := k.bankKeeper.SendCoinsFromModuleToModule(
					ctx, types.ModuleName, distribtype.ModuleName,
					sdk.NewCoins(sdk.NewCoin(currentCreateSchemaFeeBalance.Denom, rewardBalance)))
				if err != nil {
					panic(err)
				}

				remaining := reward
				for i, vote := range bondedVotes {
					validator, err := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)
					if err != nil {
						panic(err)
					}

					powerFraction := sdkmath.LegacyNewDec(vote.Validator.Power).QuoTruncate(sdkmath.LegacyNewDec(totalPreviousPower))
					toAllocate := reward.MulDecTruncate(powerFraction)
					if i == len(bondedVotes)-1 {
						// last validator, allocate the remaining coins
						toAllocate = remaining
					} else {
						remaining = remaining.Sub(toAllocate)
					}
					k.distrKeeper.AllocateTokensToValidator(ctx, validator, reward.MulDecTruncate(powerFraction))

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
	if virtualSchemaProposal.VirtualSchema.VirtualNftSchemaCode == "" {
		k.Logger().Error("empty virtual schema code")
		return false
	}

	if len(virtualSchemaProposal.VirtualSchema.Registry) == 0 {
		k.Logger().Error("empty registry")
		return false
	}

	// Count votes
	acceptCount, totalVotes := countProposalVotes(virtualSchemaProposal.VirtualSchema.Registry)
	voteThreshold := len(virtualSchemaProposal.VirtualSchema.Registry)

	k.Logger().Info("Processing virtual schema proposal",
		"id", virtualSchemaProposal.Id,
		"type", virtualSchemaProposal.ProposalType,
		"accept_count", acceptCount,
		"total_votes", totalVotes,
		"threshold", voteThreshold,
	)

	if totalVotes != voteThreshold {
		k.Logger().Info("Not all votes received yet")
		return false
	}

	if acceptCount < voteThreshold {
		k.Logger().Info("Proposal rejected")
		k.RemoveActiveVirtualSchemaProposal(ctx, virtualSchemaProposal.Id)
		k.SetInactiveVirtualSchemaProposal(ctx, types.InactiveVirtualSchemaProposal{Id: virtualSchemaProposal.Id})
		// **** SCHEMA FEE ****
		if err := k.processSchemaFee(ctx, virtualSchemaProposal, false); err != nil {
			k.Logger().Error("failed to process schema fee", "error", err)
			return false
		}
		return false
	} else {
		// Process proposal
		virtualSchema := types.VirtualSchema{
			VirtualNftSchemaCode: virtualSchemaProposal.VirtualSchema.VirtualNftSchemaCode,
			Registry:             virtualSchemaProposal.VirtualSchema.Registry,
			Enable:               virtualSchemaProposal.VirtualSchema.Enable,
		}

		k.Logger().Info("Updating virtual schema", "code", virtualSchema.VirtualNftSchemaCode, "enabled", virtualSchema.Enable)
		k.SetVirtualSchema(ctx, virtualSchema)

		if virtualSchemaProposal.ProposalType == types.ProposalType_EDIT {
			for _, action := range virtualSchemaProposal.Actions {
				_, found := k.GetVirtualAction(ctx, virtualSchema.VirtualNftSchemaCode, action.Name)
				if found {
					k.UpdateVirtualActionKeeper(ctx, virtualSchema.VirtualNftSchemaCode, *action)
				} else {
					k.AddVirtualActionKeeper(ctx, virtualSchema.VirtualNftSchemaCode, *action)
				}
			}
		} else {
			for _, action := range virtualSchemaProposal.Actions {
				k.AddVirtualActionKeeper(ctx, virtualSchema.VirtualNftSchemaCode, *action)
			}
		}

		if err := k.processSchemaFee(ctx, virtualSchemaProposal, true); err != nil {
			k.Logger().Error("failed to process schema fee", "error", err)
			return false
		}

		// compare current executor and updated executor to add new one or remove some
		currentExecutors, found := k.GetExecutorOfSchema(ctx, virtualSchema.VirtualNftSchemaCode)
		toAddExecutor := []string{}
		toRmExecutor := []string{}

		if found {
			currentExecutorMap := make(map[string]bool)
			for _, executor := range currentExecutors.ExecutorAddress {
				currentExecutorMap[executor] = true
			}

			for _, executor := range virtualSchemaProposal.Executors {
				if !currentExecutorMap[executor] {
					toAddExecutor = append(toAddExecutor, executor)
				}
				delete(currentExecutorMap, executor)
			}

			for executor := range currentExecutorMap {
				toRmExecutor = append(toRmExecutor, executor)
			}
		} else {
			toAddExecutor = virtualSchemaProposal.Executors
		}

		for _, executor := range toAddExecutor {
			k.AddActionExecutor(ctx, k.GetModuleAddress().String(), virtualSchema.VirtualNftSchemaCode, executor)
		}

		for _, executor := range toRmExecutor {
			k.DelActionExecutor(ctx, k.GetModuleAddress().String(), virtualSchema.VirtualNftSchemaCode, executor)
		}

		// Update proposal status
		k.RemoveActiveVirtualSchemaProposal(ctx, virtualSchemaProposal.Id)
		k.SetInactiveVirtualSchemaProposal(ctx, types.InactiveVirtualSchemaProposal{Id: virtualSchemaProposal.Id})

		k.Logger().Info("Virtual schema proposal processed successfully", "id", virtualSchemaProposal.Id)
		return true
	}
}

func countProposalVotes(registry []*types.VirtualSchemaRegistry) (acceptCount, totalVotes int) {
	for _, reg := range registry {
		if reg.Decision != types.RegistryStatus_PENDING {
			if reg.Decision == types.RegistryStatus_ACCEPT {
				acceptCount++
			}
		}
	}

	for _, reg := range registry {
		if reg.Decision != types.RegistryStatus_PENDING {
			totalVotes++
		}
	}
	return
}

func (k Keeper) processSchemaFee(ctx context.Context, virtualSchemaProposal types.VirtualSchemaProposal, isAccepted bool) error {
	if virtualSchemaProposal.ProposalType == types.ProposalType_EDIT {
		return nil
	}

	feeConfig, found := k.GetNFTFeeConfig(ctx)
	if !found {
		return nil
	}

	amount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	if err != nil {
		k.Logger().Error("failed to parse fee amount", "error", err)
		return fmt.Errorf("failed to parse fee amount: %w", err)
	}

	feeBalances, found := k.GetNFTFeeBalance(ctx)
	if !found {
		feeBalances = types.NFTFeeBalance{
			FeeBalances: []string{
				"0" + amount.Denom,
			},
		}
	}

	if len(feeBalances.FeeBalances) > 0 {
		feeBalances.FeeBalances[types.FeeSubject_CREATE_NFT_SCHEMA] = "0" + amount.Denom
	}

	err = k.VirtualSchemaProcessFee(ctx, &feeConfig, &feeBalances, types.FeeSubject_CREATE_NFT_SCHEMA, isAccepted, virtualSchemaProposal.Id)
	if err != nil {
		k.Logger().Error("failed to process fee", "error", err)
		return fmt.Errorf("failed to process fee: %w", err)
	}

	k.SetNFTFeeBalance(ctx, feeBalances)

	return nil
}
