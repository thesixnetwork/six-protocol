package keeper

import (
	"fmt"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
	tkmtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) UseNft(goCtx context.Context, msg *types.MsgUseNft) (*types.MsgUseNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var spend = tkmtypes.Burn{
		Creator: msg.Creator,
		Token:   msg.Token,
		Amount:  1,
	}

	// Chect is this creator is exist
	spender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	supply := k.bankKeeper.GetSupply(ctx, msg.Token)
	if supply.Amount.Uint64() < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	var burnAmount uint64 = 1

	tokens := sdk.Coin{
		Denom:  msg.Token,
		Amount: sdk.NewIntFromUint64(burnAmount),
	}

	if balance := k.bankKeeper.GetBalance(ctx, spender, msg.Token); balance.Amount.Uint64() < 1 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	//send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, spender, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from account to module despite previously burning coins to module account: %v", err))
	}

	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return nil, err
	}

	// Get burning history
	prev, found := k.tokenmngrKeeper.GetTokenBurn(ctx, msg.Token)
	if !found {
		var tokenBurn = tkmtypes.TokenBurn{
			Token:  msg.Token,
			Amount: 1,
		}
		k.tokenmngrKeeper.SetTokenBurn(ctx, tokenBurn)
	} else {
		var tokenBurn = tkmtypes.TokenBurn{
			Token:  msg.Token,
			Amount: 1 + prev.Amount,
		}
		k.tokenmngrKeeper.SetTokenBurn(ctx, tokenBurn)
	}

	// Update to history
	id := k.tokenmngrKeeper.UpdateBurn(ctx, spend)

	return &types.MsgUseNftResponse{Id: id}, nil
}
