package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, hasPerm := k.GetMintperm(ctx, msg.Token, msg.Creator)
	if !hasPerm {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "mint require mint permission")
	}

	token, foundToken := k.GetToken(ctx, msg.Token)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	supply := k.bankKeeper.GetSupply(ctx, msg.Token)

	if supply.Amount.Uint64() >= uint64(token.MaxSupply) && token.MaxSupply != 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token reach max supply")
	}

	var minAmount uint64
	newTotalSupply := supply.Amount.Uint64() + msg.Amount
	if newTotalSupply > uint64(token.MaxSupply) && token.MaxSupply != 0 {
		minAmount = uint64(token.MaxSupply) - supply.Amount.Uint64()
	} else {
		minAmount = msg.Amount
	}

	tokens := sdk.Coin{
		Denom:  msg.Token,
		Amount: sdk.NewIntFromUint64(minAmount),
	}

	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return nil, err
	}

	// k.Logger(ctx).Info(fmt.Sprintf("receiver: %d", sdk.AccAddress()))

	var minteeAddress string
	if token.Mintee == "" {
		options, found := k.GetOptions(ctx)
		if found {
			minteeAddress = options.DefaultMintee
		}
	} else {
		minteeAddress = token.Mintee
	}

	mintee, err := sdk.AccAddressFromBech32(minteeAddress)
	if err != nil {
		return nil, err
	}

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, mintee, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account despite previously minting coins to module account: %v", err))
	}

	return &types.MsgMintResponse{}, nil
}
