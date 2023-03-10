package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) CreateMintperm(goCtx context.Context, msg *types.MsgCreateMintperm) (*types.MsgCreateMintpermResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	_, foundToken := k.GetToken(ctx, msg.Token)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not existed")
	}

	// ** Prepare this code for future but it might be bypass by bank module **
	// ** Need to test and verify **
	// ** But without this it is still working on testnet and mainnet **
	// ** So it is not a priority **
	// tokenmngr_token, foundToken := k.GetToken(ctx, msg.Token)
	// if !foundToken {
	// 	// find from bank module
	// 	token, found := k.bankKeeper.GetDenomMetaData(ctx, msg.Token)
	// 	if !found {
	// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token not found")
	// 	}

	// 	// get token admin
	// 	token_admin, found := k.protocoladminKeeper.GetGroup(ctx, TOKEN_ADMIN)
	// 	if !found {
	// 		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token admin not found")
	// 	}

	// 	// create token
	// 	maxsupply := sdk.NewIntFromUint64(0)
	// 	new_coin := sdk.NewCoin(token.Base, maxsupply)
	// 	tokenmngr_token = types.Token{
	// 		Name:      token.Display,
	// 		Base:      token.Base,
	// 		Mintee:    msg.Address,
	// 		Creator:   token_admin.Owner,
	// 		MaxSupply: new_coin,
	// 	}
	// 	k.SetToken(ctx, tokenmngr_token)
	// }

	// Check if the value already exists
	_, isFound := k.GetMintperm(
		ctx,
		msg.Token,
		msg.Address,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var mintperm = types.Mintperm{
		Creator: msg.Creator,
		Token:   msg.Token,
		Address: msg.Address,
	}

	k.SetMintperm(
		ctx,
		mintperm,
	)
	return &types.MsgCreateMintpermResponse{}, nil
}

func (k msgServer) UpdateMintperm(goCtx context.Context, msg *types.MsgUpdateMintperm) (*types.MsgUpdateMintpermResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value exists
	_, isFound := k.GetMintperm(
		ctx,
		msg.Token,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var mintperm = types.Mintperm{
		Creator: msg.Creator,
		Token:   msg.Token,
		Address: msg.Address,
	}

	k.SetMintperm(ctx, mintperm)

	return &types.MsgUpdateMintpermResponse{}, nil
}

func (k msgServer) DeleteMintperm(goCtx context.Context, msg *types.MsgDeleteMintperm) (*types.MsgDeleteMintpermResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	pass := k.protocoladminKeeper.Authenticate(ctx, TOKEN_ADMIN, msg.Creator)
	if !pass {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "message creator is not token admin or super admin")
	}

	// Check if the value exists
	_, isFound := k.GetMintperm(
		ctx,
		msg.Token,
		msg.Address,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveMintperm(
		ctx,
		msg.Token,
		msg.Address,
	)

	return &types.MsgDeleteMintpermResponse{}, nil
}
