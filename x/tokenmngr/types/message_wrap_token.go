package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWrapToken = "wrap_token"

var _ sdk.Msg = &MsgWrapToken{}

func NewMsgWrapToken(creator string, amount sdk.Coin, receiver string) *MsgWrapToken {
	return &MsgWrapToken{
		Creator:  creator,
		Amount:   amount,
		Receiver: receiver,
	}
}

func (msg *MsgWrapToken) Route() string {
	return RouterKey
}

func (msg *MsgWrapToken) Type() string {
	return TypeMsgWrapToken
}

func (msg *MsgWrapToken) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWrapToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWrapToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
