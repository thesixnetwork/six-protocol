package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// SetNFTFeeBalance set nFTFeeBalance in the store
func (k Keeper) SetNFTFeeBalance(ctx sdk.Context, nFTFeeBalance types.NFTFeeBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))
	b := k.cdc.MustMarshal(&nFTFeeBalance)
	store.Set([]byte{0}, b)
}

// GetNFTFeeBalance returns nFTFeeBalance
func (k Keeper) GetNFTFeeBalance(ctx sdk.Context) (val types.NFTFeeBalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetNFTFeeConfig set nFTFeeConfig in the store
func (k Keeper) SetNFTFeeConfig(ctx sdk.Context, nFTFeeConfig types.NFTFeeConfig) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))
	b := k.cdc.MustMarshal(&nFTFeeConfig)
	store.Set([]byte{0}, b)
}

// GetNFTFeeConfig returns nFTFeeConfig
func (k Keeper) GetNFTFeeConfig(ctx sdk.Context) (val types.NFTFeeConfig, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNFTFeeConfig removes nFTFeeConfig from the store
func (k Keeper) RemoveNFTFeeConfig(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeConfigKey))
	store.Delete([]byte{0})
}

// RemoveNFTFeeBalance removes nFTFeeBalance from the store
func (k Keeper) RemoveNFTFeeBalance(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NFTFeeBalanceKey))
	store.Delete([]byte{0})
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
