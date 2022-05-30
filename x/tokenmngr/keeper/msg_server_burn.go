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
		Token:   msg.Token,
		Amount:  msg.Amount,
	}

	token, foundToken := k.GetToken(ctx, msg.Token)
	_ = token
	burner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if msg.Amount == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	// if uint64(token.MaxSupply) < msg.Amount{
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than maximum supply")
	// }

	supply := k.bankKeeper.GetSupply(ctx, msg.Token)
	if supply.Amount.Uint64() < msg.Amount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current supply")
	}

	var burnAmount uint64 = msg.Amount

	tokens := sdk.Coin{
		Denom:  msg.Token,
		Amount: sdk.NewIntFromUint64(burnAmount),
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

	prev, found := k.GetTokenBurn(ctx, msg.Token)
	if !found {
		var burn = types.TokenBurn{
			Token: msg.Token,
			Amount: msg.Amount,
		}
		k.SetTokenBurn(ctx, burn)
	}else{
		var burn = types.TokenBurn{
			Token: msg.Token,
			Amount: msg.Amount + prev.Amount,
		}
		k.SetTokenBurn(ctx, burn)
	}

	id := k.UpdateBurn(ctx, burn)

	return &types.MsgBurnResponse{Id: id}, nil
}
