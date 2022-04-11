package keeper

import (
	"context"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) CreateToken(goCtx context.Context, msg *types.MsgCreateToken) (*types.MsgCreateTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value already exists
	_, isFound := k.GetToken(
		ctx,
		msg.Name,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index already set")
	}

	var denomMetaData banktypes.Metadata

	err := json.Unmarshal([]byte(msg.DenomMetaData), &denomMetaData)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, "error unmarshal denom metadata")
	}

	if denomMetaData.Display != msg.Name {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token name does not correspond to the denom")
	}

	// Set denom metadata to bank module
	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)

	var token = types.Token{
		Creator:   msg.Creator,
		Name:      msg.Name,
		Base:      denomMetaData.Base,
		Mintee:    msg.Mintee,
		MaxSupply: msg.MaxSupply,
	}

	k.SetToken(
		ctx,
		token,
	)
	return &types.MsgCreateTokenResponse{}, nil
}

// ! Function will return error for there is no implementation at the moment
func (k msgServer) UpdateToken(goCtx context.Context, msg *types.MsgUpdateToken) (*types.MsgUpdateTokenResponse, error) {

	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "operation not available")

	// ctx := sdk.UnwrapSDKContext(goCtx)

	// pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	// if !pass {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unauthorized")
	// }

	// // Check if the value exists
	// foundToken, isFound := k.GetToken(
	// 	ctx,
	// 	msg.Name,
	// )
	// if !isFound {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	// }

	// foundToken.Name = msg.Name
	// foundToken.MaxSupply = msg.MaxSupply

	// k.SetToken(ctx, foundToken)

	// return &types.MsgUpdateTokenResponse{}, nil
}

func (k msgServer) DeleteToken(goCtx context.Context, msg *types.MsgDeleteToken) (*types.MsgDeleteTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value exists
	_, isFound := k.GetToken(
		ctx,
		msg.Name,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveToken(
		ctx,
		msg.Name,
	)

	return &types.MsgDeleteTokenResponse{}, nil
}
