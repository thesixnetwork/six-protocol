package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgWrapToken{}

func NewMsgWrapToken(creator string, amount sdk.Coin, receiver string) *MsgWrapToken {
	return &MsgWrapToken{
		Creator:  creator,
		Amount:   amount,
		Receiver: receiver,
	}
}

func (msg *MsgWrapToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
