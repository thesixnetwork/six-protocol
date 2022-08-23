package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

)

func (k msgServer) EthSend(goCtx context.Context, msg *types.MsgEthSend) (*types.MsgEthSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	eth_reciever, found := k.GetBinding(
		ctx,
		msg.ToEth,
	)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "eth_reciever address not found")
	}

	eth_sender, found := k.GetBinding(
		ctx,
		msg.FromEth,
	)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "eth_sender address not found")
	}

	// Convert amount strings to sdk.Coins
    price, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid amount")
	}

	// Map eth address with 6x address
	six_source := eth_sender.Creator // creator is the owner of mapping address ex. { "6xblahblahblah" : "ethblahblahblah" }
	six_destination := eth_reciever.Creator

	// Account from string to sdk.AccAddress
	sender, _ := sdk.AccAddressFromBech32(six_source)
	reciever, _ := sdk.AccAddressFromBech32(six_destination)

	k.bankKeeper.SendCoins(ctx, sender, reciever, price)
	return &types.MsgEthSendResponse{}, nil
}
