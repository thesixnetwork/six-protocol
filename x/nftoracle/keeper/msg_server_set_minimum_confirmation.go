package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func (k msgServer) SetMinimumConfirmation(goCtx context.Context, msg *types.MsgSetMinimumConfirmation) (*types.MsgSetMinimumConfirmationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracleAdmin, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracleAdmin, oracleAdmin)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	// convert msg.NewConfirmation into number
	newConfirmation, err := strconv.ParseUint(msg.NewConfirmation, 10, 64)
	if err != nil {
		return nil, err
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
