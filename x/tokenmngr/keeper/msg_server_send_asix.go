package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	ethermint "github.com/evmos/ethermint/types"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k msgServer) SendWrapToken(goCtx context.Context, msg *types.MsgSendWrapToken) (*types.MsgSendWrapTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	denom := msg.Amount.Denom

	// check that receiver is cosmos address or ethereum address
	var addr []byte
	var receiver sdk.AccAddress
	if err := ethermint.ValidateAddress(msg.EthAddress); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "receiver address is not ethereum address")
	}
	if common.IsHexAddress(msg.EthAddress) {
		addr = common.HexToAddress(msg.EthAddress).Bytes()
	}
	receiver = sdk.AccAddress(addr)

	token, foundToken := k.GetToken(ctx, msg.Amount.Denom)
	if !foundToken {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	if token.Base != "asix" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "token is not asix")
	}

	if msg.Amount.Amount.IsZero() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is prohibit from module")
	}

	// Check is this creator is exist
	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if balance := k.bankKeeper.GetBalance(ctx, sender, denom); balance.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	supply := k.bankKeeper.GetSupply(ctx, denom)
	if supply.Amount.LT(msg.Amount.Amount) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	if !msg.Amount.IsValid() {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	// send amount to receiver
	if err = k.bankKeeper.SendCoins(ctx, sender, receiver, sdk.NewCoins(msg.Amount)); err != nil {
		return nil, err
	}

	return &types.MsgSendWrapTokenResponse{}, nil
}
