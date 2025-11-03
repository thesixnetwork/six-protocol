package keeper

import (
	"context"
	"encoding/json"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value already exists
	_, isFound := k.GetToken(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var denomMetaData banktypes.Metadata

	err := json.Unmarshal([]byte(msg.DenomMetaData), &denomMetaData)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrJSONUnmarshal, "error unmarshal denom metadata")
	}

	if denomMetaData.Display != msg.Name {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token name does not correspond to the denom")
	}

	// Set denom metadata to bank module
	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)

	token := types.Token{
		Creator:   msg.Creator,
		Name:      msg.Name,
		Base:      denomMetaData.Base,
		MaxSupply: msg.MaxSupply,
		Mintee:    msg.Mintee,
	}

	k.SetToken(
		ctx,
		token,
	)
	return &types.MsgCreateTokenResponse{}, nil
}

func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {
	return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "operation not available")
	// ctx := sdk.UnwrapSDKContext(goCtx)

	// // Check if the value exists
	// valFound, isFound := k.GetToken(
	// 	ctx,
	// 	msg.Name,
	// )
	// if !isFound {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	// }

	// // Checks if the msg creator is the same as the current owner
	// if msg.Creator != valFound.Creator {
	// 	return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }

	// var token = types.Token{
	// 	Creator:   msg.Creator,
	// 	Name:      msg.Name,
	// 	Base:      msg.Base,
	// 	MaxSupply: msg.MaxSupply,
	// 	Mintee:    msg.Mintee,
	// }

	// k.SetToken(ctx, token)

	// return &types.MsgUpdateTokenResponse{}, nil
}

func (k msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetToken(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveToken(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteTokenResponse{}, nil
}
