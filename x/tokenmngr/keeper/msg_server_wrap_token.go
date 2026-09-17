package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	evmostype "github.com/evmos/evmos/v20/types"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func (k msgServer) WrapToken(goCtx context.Context, msg *types.MsgWrapToken) (*types.MsgWrapTokenResponse, error) {
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
		if err := evmostype.ValidateAddress(msg.Receiver); err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "receiver address is not cosmos or ethereum address")
		}

		if common.IsHexAddress(msg.Receiver) {
			addr = common.HexToAddress(msg.Receiver).Bytes()
		}
		receiver = sdk.AccAddress(addr)
	}

	// Bind the wrapped denom to the real base coin. The peg mints asix against
	// the usix that gets locked here, so the incoming denom MUST be usix itself
	// — checking an arbitrary token record's Base is not enough, since a
	// TOKEN_ADMIN could create a decoy token named "foo" with Base=="usix" and
	// wrap worthless "foo" to mint real asix (and later drain the usix reserve).
	if denom != DefaultMicroDenom {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "only %s can be wrapped, got %s", DefaultMicroDenom, denom)
	}

	// Check is this token is exist in token list
	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	// accept only usix token to convert to atto or asix
	if token.Base != DefaultMicroDenom {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "token is not usix")
	}

	// Check is this amount is valid
	if msg.Amount.Amount.IsZero() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// Check token balance of creator
	if balance := k.bankKeeper.GetBalance(ctx, signer, denom); balance.Amount.LT(msg.Amount.Amount) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	// Check token supply
	supply := k.bankKeeper.GetSupply(ctx, denom)
	if supply.Amount.LT(msg.Amount.Amount) {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	// send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, signer, types.ModuleName, convertAmount); err != nil {
		return nil, errorsmod.Wrap(types.ErrSendCoinsFromAccountToModule, "Amount of token is too high than current balance due"+err.Error())
	}

	// if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, convertAmount); err != nil {
	// 	return nil, err
	// }

	attoAmount := sdk.NewCoin("asix", msg.Amount.Amount.MulRaw(int64(DefaultAttoToMicroDiff)))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(attoAmount)); err != nil {
		return nil, err
	}

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(attoAmount),
	); err != nil {
		return nil, errorsmod.Wrap(types.ErrSendCoinsFromAccountToModule, "unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account:"+err.Error())
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypesConvertCoinToWei),
			sdk.NewAttribute(types.AttributeKeyDestAddress, receiver.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
	})

	return &types.MsgWrapTokenResponse{
		Amount: attoAmount,
	}, nil
}
