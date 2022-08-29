package keeper

import (
	"context"
	"fmt"
	"time"
	// "strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
	// tkmtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) UseNft(goCtx context.Context, msg *types.MsgUseNft) (*types.MsgUseNftResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, foundToken := k.tokenmngrKeeper.GetToken(ctx, msg.Token)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	last_spend, _ := k.GetNftUsed(ctx, msg.Token, msg.Creator)
	last_spend_time, _ := time.Parse(time.RFC3339, last_spend.UpdateAt) // to UTC
    expiredAt := last_spend_time.Add(5 * time.Minute)

	msg_timestamp_UTC, error := time.Parse(time.RFC3339, msg.Timestamp) // to UTC
	if error != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid timestamp")
	}

	if last_spend.Creator == msg.Creator && msg_timestamp_UTC.Before(expiredAt) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "you can only use this nft once every 5 minutes")
	}	

	// if msg_timestamp_UTC.Before(expiredAt) {
	// 	return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Token is used recently")
	// }

	var spend = types.UseNft{
		Creator:   msg.Creator,
		Token:     msg.Token,
		Timestamp: msg.Timestamp,
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
	prev, found := k.GetNftUsed(ctx, msg.Token , msg.Creator)
	if !found {
		var tokenBurn = types.NftUsed{
			Token:  msg.Token,
			Amount: 1,
			UpdateAt: msg.Timestamp,
			Creator: msg.Creator,
		}
		k.SetNftUsed(ctx, tokenBurn)
	} else {
		var tokenBurn = types.NftUsed{
			Token:  msg.Token,
			Amount: 1 + prev.Amount,
			UpdateAt: msg.Timestamp,
			Creator: msg.Creator,
		}
		k.SetNftUsed(ctx, tokenBurn)
	}

	// Update to history
	id := k.UpdateUseNft(ctx, spend)

	return &types.MsgUseNftResponse{Id: id}, nil
}
