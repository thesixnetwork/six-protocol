package keeper

import (
	"context"

	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func (k msgServer) UnwrapToken(goCtx context.Context, msg *types.MsgUnwrapToken) (*types.MsgUnwrapTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := msg.Amount.Denom
	convertAmount := sdk.NewCoins(msg.Amount) // asix

	// reject if denom is zero
	if msg.Amount.Amount.IsZero() {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// accept only zero after decimal point (atto)
	if !msg.Amount.Amount.ModRaw(int64(DefaultAttoToMicroDiff)).IsZero() {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// Check is this creator is exist
	signer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// check that receiver is cosmos address or ethereum address
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, "receiver address is not cosmos address")
	}

	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, errormod.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	if token.Base != DefaultAttoDenom {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "token is not asix")
	}

	if msg.Amount.Amount.IsZero() {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	if balance := k.bankKeeper.GetBalance(ctx, signer, denom); balance.Amount.LT(msg.Amount.Amount) {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	supply := k.bankKeeper.GetSupply(ctx, denom)
	if supply.Amount.LT(msg.Amount.Amount) {
		return nil, errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	// send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, convertAmount); err != nil {
		return nil, errormod.Wrap(types.ErrSendCoinsFromAccountToModule, "Amount of token is too high than current balance due"+err.Error())
	}

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, convertAmount); err != nil {
		return nil, err
	}

	microSix := sdk.NewCoin(DefaultMicroDenom, msg.Amount.Amount.QuoRaw(int64(DefaultAttoToMicroDiff)))

	// get the module account balance
	tokenmngrModuleAccount := k.accountKeeper.GetModuleAddress(types.ModuleName)
	moduleBalance := k.bankKeeper.GetBalance(ctx, tokenmngrModuleAccount, DefaultMicroDenom)

	// check if module account balance is enough to send
	if moduleBalance.Amount.LT(microSix.Amount) {
		return nil, errormod.Wrap(sdkerrors.ErrInsufficientFunds, "module account balance is not enough to send")
	}

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(microSix),
	); err != nil {
		return nil, errormod.Wrap(types.ErrSendCoinsFromAccountToModule, "unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account:"+err.Error())
	}

	return &types.MsgUnwrapTokenResponse{
		Amount: microSix,
	}, nil
}
