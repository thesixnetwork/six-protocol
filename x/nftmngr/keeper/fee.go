package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

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
