package keeper

import (
	"context"
	"fmt"

	// "encoding/binary"
	// "github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	var burn = types.Burn{
		Creator: msg.Creator,
		Amount:  msg.Amount,
	}
	// Check is token exist
	// _, foundToken := k.GetToken(ctx, msg.Token)
	// if !foundToken {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	// }

	// Chect is this creator is exist
	burner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if msg.Amount.Amount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}
	//TODO:: Make sure MaxSupply and totalSupply is Dupplicate or not
	// if uint64(token.MaxSupply) < msg.Amount{
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than maximum supply")
	// }

	supply := k.bankKeeper.GetSupply(ctx, msg.Amount.Denom)
	if supply.Amount.LT(msg.Amount.Amount){
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	var burnAmount uint64 = msg.Amount.Amount.Uint64()

	tokens := sdk.Coin{
		Denom:  msg.Amount.Denom,
		Amount: sdk.NewIntFromUint64(burnAmount),
	}

	if balance := k.bankKeeper.GetBalance(ctx, burner, msg.Amount.Denom); balance.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	//send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, burner, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from account to module despite previously burning coins to module account: %v", err))
	}

	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return nil, err
	}

	// Get burning history
	prev, found := k.GetTokenBurn(ctx, msg.Amount.Denom)
	if !found {
		var tokenBurn = types.TokenBurn{
			Amount: msg.Amount,
		}
		k.SetTokenBurn(ctx, tokenBurn)
	} else {
		new_burn_amount := prev.Amount.Amount.Add(msg.Amount.Amount)
		new_burn_coin := sdk.Coin{
			Denom:  msg.Amount.Denom,
			Amount: new_burn_amount,
		}
		var tokenBurn = types.TokenBurn{
			Amount: new_burn_coin,
		}
		k.SetTokenBurn(ctx, tokenBurn)
	}

	// Update to history
	id := k.UpdateBurn(ctx, burn)

	return &types.MsgBurnResponse{Id: id}, nil
}
