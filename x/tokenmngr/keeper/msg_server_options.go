package keeper

import (
	"context"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateOptions(goCtx context.Context, msg *types.MsgCreateOptions) (*types.MsgCreateOptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value already exists
	_, isFound := k.GetOptions(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	options := types.Options{
		DefaultMintee: msg.DefaultMintee,
	}

	k.SetOptions(
		ctx,
		options,
	)
	return &types.MsgCreateOptionsResponse{}, nil
}

func (k msgServer) UpdateOptions(goCtx context.Context, msg *types.MsgUpdateOptions) (*types.MsgUpdateOptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value exists
	_, isFound := k.GetOptions(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	options := types.Options{
		DefaultMintee: msg.DefaultMintee,
	}

	k.SetOptions(ctx, options)

	return &types.MsgUpdateOptionsResponse{}, nil
}

func (k msgServer) DeleteOptions(goCtx context.Context, msg *types.MsgDeleteOptions) (*types.MsgDeleteOptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value exists
	_, isFound := k.GetOptions(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	k.RemoveOptions(ctx)

	return &types.MsgDeleteOptionsResponse{}, nil
}
