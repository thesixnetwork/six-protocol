package keeper

import (
	"context"
	"fmt"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	
)

func (k msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, hasPerm := k.GetMintperm(ctx, msg.Amount.Denom, msg.Creator)
	if !hasPerm {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "mint require mint permission")
	}

	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	total_supply := k.bankKeeper.GetSupply(ctx, msg.Amount.Denom)

	if total_supply.Amount.GTE(token.MaxSupply.Amount) && !token.MaxSupply.Amount.IsZero() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token reach max supply")
	}

	var mintAmount sdkmath.Int

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

	return &types.MsgMintResponse{}, nil
}
