package keeper

import (
	"context"
	"strconv"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	errormod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetNFTFeeBalance set nFTFeeBalance in the store
func (k Keeper) SetNFTFeeBalance(ctx context.Context, nFTFeeBalance types.NFTFeeBalance) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeBalanceKey))
	b := k.cdc.MustMarshal(&nFTFeeBalance)
	store.Set([]byte{0}, b)
}

// GetNFTFeeBalance returns nFTFeeBalance
func (k Keeper) GetNFTFeeBalance(ctx context.Context) (val types.NFTFeeBalance, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeBalanceKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetNFTFeeConfig set nFTFeeConfig in the store
func (k Keeper) SetNFTFeeConfig(ctx context.Context, nFTFeeConfig types.NFTFeeConfig) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeConfigKey))

	b := k.cdc.MustMarshal(&nFTFeeConfig)
	store.Set([]byte{0}, b)
}

// GetNFTFeeConfig returns nFTFeeConfig
func (k Keeper) GetNFTFeeConfig(ctx context.Context) (val types.NFTFeeConfig, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeConfigKey))
	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTFeeConfig removes nFTFeeConfig from the store
func (k Keeper) RemoveNFTFeeConfig(ctx sdk.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeConfigKey))
	store.Delete([]byte{0})
}

// RemoveNFTFeeBalance removes nFTFeeBalance from the store
func (k Keeper) RemoveNFTFeeBalance(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.NFTFeeBalanceKey))
	store.Delete([]byte{0})
}

func (k Keeper) ValidateFeeConfig(feeConfig *types.NFTFeeConfig) error {
	// Validate Amount
	feeAmount, err := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
	if err != nil {
		return errormod.Wrap(types.ErrInvalidFeeAmount, err.Error())
	}
	// check if feeAmount.Amount > 0
	if feeAmount.Amount.LTE(sdkmath.NewInt(0)) {
		return errormod.Wrap(types.ErrInvalidFeeAmount, "Fee amount must be greater than 0")
	}
	// loop over feeConfig.SchemaFee.FeeDistributions
	totalPortion := float32(0)
	for _, feeDistribution := range feeConfig.SchemaFee.FeeDistributions {
		// validate portion
		if feeDistribution.Portion <= 0 || feeDistribution.Portion > 1 {
			return errormod.Wrap(types.ErrInvalidFeePortion, "Fee portion must be between 0 and 1")
		}
		totalPortion += feeDistribution.Portion
	}
	if totalPortion != 1 {
		return errormod.Wrap(types.ErrInvalidFeePortion, "Total fee portion must be equal to 1")
	}
	return nil
}

func (k Keeper) VirtualSchemaProcessFee(ctx context.Context, feeConfig *types.NFTFeeConfig, feeBalances *types.NFTFeeBalance, feeSubject types.FeeSubject, pass bool, proposalId string) error {
	lockedAsset, found := k.GetLockSchemaFee(ctx, proposalId)
	if !found {
		return errormod.Wrap(types.ErrLockedAssetNotFound, proposalId)
	}
	creatorAddress, err := sdk.AccAddressFromBech32(lockedAsset.Proposer)
	if err != nil {
		return err
	}

	if !pass {
		err := k.bankKeeper.SendCoinsFromModuleToAccount(
			ctx, types.ModuleName, creatorAddress, sdk.NewCoins(lockedAsset.Amount),
		)
		if err != nil {
			return err
		}
	} else {
		currentFeeBalance, _ := sdk.ParseCoinNormalized(feeBalances.FeeBalances[int32(feeSubject)])
		feeAmount, _ := sdk.ParseCoinNormalized(feeConfig.SchemaFee.FeeAmount)
		// make sure fee lockamount enough for feeAmount
		if feeAmount.Amount.LT(lockedAsset.Amount.Amount) {
			return errormod.Wrap(sdkerrors.ErrInsufficientFunds, "insufficient fee balance")
		}
		// Plus fee amount to fee balance
		currentFeeBalance = currentFeeBalance.Add(feeAmount)
		feeBalances.FeeBalances[int32(feeSubject)] = strconv.FormatInt(currentFeeBalance.Amount.Int64(), 10) + feeAmount.Denom
	}

	k.RemoveLockSchemaFee(ctx, proposalId)
	return nil
}
