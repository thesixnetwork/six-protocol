package keeper

import (
	"context"
	"math"
	"strconv"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SetMinimumConfirmation(goCtx context.Context, msg *types.MsgSetMinimumConfirmation) (*types.MsgSetMinimumConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracleAdmin, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracleAdmin, oracleAdmin)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	// convert msg.NewConfirmation into number
	newConfirmation, err := strconv.ParseUint(msg.NewConfirmation, 10, 64)
	if err != nil {
		return nil, err
	}

	// A minimum confirmation of 0 collapses the oracle quorum: a mint/action
	// request created with RequiredConfirm=1 would reach consensus on the FIRST
	// oracle submission, so a single (or compromised) oracle key could forge
	// arbitrary cross-chain mints. Require at least one confirmation, and reject
	// values that would not fit the int32 storage field (silent truncation could
	// otherwise wrap a huge value back into range).
	if newConfirmation < 1 {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "minimum confirmation must be at least 1")
	}
	if newConfirmation > math.MaxInt32 {
		return nil, errormod.Wrapf(sdkerrors.ErrInvalidRequest, "minimum confirmation too large; max %d", math.MaxInt32)
	}

	// Retrieve the current oracle configuration
	var oracleConfig types.OracleConfig
	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		oracleConfig = types.OracleConfig{
			MinimumConfirmation: 0,
		}
	}
	oracleConfig.MinimumConfirmation = int32(newConfirmation)
	// Store the new oracle configuration
	k.SetOracleConfig(ctx, oracleConfig)

	// Emit events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSetMinimumConfirmation,
			sdk.NewAttribute(types.AttributeKeyMinimumConfirmation, msg.NewConfirmation),
		),
	)

	return &types.MsgSetMinimumConfirmationResponse{}, nil
}
