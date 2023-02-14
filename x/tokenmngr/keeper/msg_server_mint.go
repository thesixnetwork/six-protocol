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

	_, hasPerm := k.GetMintperm(ctx, msg.Amount.Denom, msg.Creator)
	if !hasPerm {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "mint require mint permission")
	}

	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	total_supply := k.bankKeeper.GetSupply(ctx, msg.Amount.Denom)

	if total_supply.Amount.GTE(token.MaxSupply.Amount) && !token.MaxSupply.Amount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token reach max supply")
	}

	var mintAmount sdk.Int
	newTotalSupply := total_supply.Amount.Add(msg.Amount.Amount)
	if newTotalSupply.GT(token.MaxSupply.Amount) && !token.MaxSupply.Amount.IsZero() {
		mintAmount = token.MaxSupply.Amount.Sub(total_supply.Amount)
	} else {
		mintAmount = msg.Amount.Amount
	}

	tokens := sdk.Coin{
		Denom:  msg.Amount.Denom,
		Amount: mintAmount,
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
		panic(fmt.Sprintf("unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account: %v", err))
	}
	// for _, msg.Amount := range msg.Amount {
	// 	// balance := k.bankKeeper.GetBalance(ctx, sdk.AccAddress(msg.Creator), msg.Amount.Denom)

	// }

	return &types.MsgMintResponse{}, nil
}
