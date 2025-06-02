package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k msgServer) SetFeeConfig(goCtx context.Context, msg *types.MsgSetFeeConfig) (*types.MsgSetFeeConfigResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate if creator has permission to set fee config
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionNftFeeAdmin, creator)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoNftFeeAdminPermission, msg.Creator)
	}

	feeConfig := types.NFTFeeConfig{}
	feeConfig.SchemaFee = msg.FeeConfig
	err = k.ValidateFeeConfig(&feeConfig)
	if err != nil {
		return nil, err
	}

	// Set fee config
	k.SetNFTFeeConfig(ctx, feeConfig)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSetFeeConfig,
			sdk.NewAttribute(types.AttributeKeyFeeConfig, string(feeConfig.SchemaFee.String())),
		),
	)
	return &types.MsgSetFeeConfigResponse{}, nil
}
