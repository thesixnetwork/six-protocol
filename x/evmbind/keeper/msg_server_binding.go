package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
	handler "github.com/thesixnetwork/six-protocol/handler"
)

func (k msgServer) CreateBinding(goCtx context.Context, msg *types.MsgCreateBinding) (*types.MsgCreateBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	//chaeck creator is valid
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, err.Error())
	}

	_ ,err = handler.ValidateEVM(msg.SignMessage, msg.EthSignature, msg.EthAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	timestamp := time.Now().Unix()
	timestamp_epoch := time.Unix(timestamp,0)
	timestamp_string := timestamp_epoch.Format("2006-01-02T15:04:05Z")

	var binding = types.Binding{
		Creator:      msg.Creator,
		EthAddress:   msg.EthAddress,
		Timestamp:   timestamp_string,
	}

	k.SetBinding(
		ctx,
		binding,
	)
	return &types.MsgCreateBindingResponse{}, nil
}

func (k msgServer) UpdateBinding(goCtx context.Context, msg *types.MsgUpdateBinding) (*types.MsgUpdateBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	_ ,err := handler.ValidateEVM(msg.SignMessage, msg.EthSignature, msg.EthAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	timestamp := time.Now().Unix()
	timestamp_epoch := time.Unix(timestamp,0)
	timestamp_string := timestamp_epoch.Format("2006-01-02T15:04:05Z")


	var binding = types.Binding{
		Creator:      msg.Creator,
		EthAddress:   msg.EthAddress,
		Timestamp:   timestamp_string,
	}

	k.SetBinding(ctx, binding)

	return &types.MsgUpdateBindingResponse{}, nil
}

func (k msgServer) DeleteBinding(goCtx context.Context, msg *types.MsgDeleteBinding) (*types.MsgDeleteBindingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBinding(
		ctx,
		msg.EthAddress,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBinding(
		ctx,
		msg.EthAddress,
	)

	return &types.MsgDeleteBindingResponse{}, nil
}
