package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) ConvertToAtto(goCtx context.Context, msg *types.MsgConvertToAtto) (*types.MsgConvertToAttoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := msg.Amount.Denom
	convertAmount := sdk.NewCoins(msg.Amount)

	// Check is this creator is exist
	signer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// check that receiver is cosmos address or ethereum address
	var addr []byte
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		if err := ethermint.ValidateAddress(msg.Receiver); err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address is not cosmos or ethereum address")
		}

		if common.IsHexAddress(msg.Receiver){
			addr = common.HexToAddress(msg.Receiver).Bytes()
		}
		receiver = sdk.AccAddress(addr)
	}

	// Check is this token is exist in token list
	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	// accept only usix token to convert to atto or asix
	if token.Base != "usix" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token is not usix")
	}

	// Check is this amount is valid
	if msg.Amount.Amount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// Check token balance of creator
	if balance := k.bankKeeper.GetBalance(ctx, signer, denom); balance.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	// Check token supply
	supply := k.bankKeeper.GetSupply(ctx, denom)
	if supply.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	//send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, convertAmount); err != nil {
		return nil, sdkerrors.Wrap(types.ErrSendCoinsFromAccountToModule, "Amount of token is too high than current balance due"+err.Error())
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, convertAmount); err != nil {
		return nil, err
	}

	attoAmount := sdk.NewCoin("asix", msg.Amount.Amount.MulRaw(1000000000000))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(attoAmount)); err != nil {
		return nil, err
	}

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(attoAmount),
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrSendCoinsFromAccountToModule, "unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account:"+err.Error())
	}

	return &types.MsgConvertToAttoResponse{
		Amount: attoAmount,
	}, nil
}
