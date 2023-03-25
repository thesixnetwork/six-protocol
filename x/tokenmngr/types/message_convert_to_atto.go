package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConvertToAtto = "convert_to_atto"

var _ sdk.Msg = &MsgConvertToAtto{}

func NewMsgConvertToAtto(creator string, amount sdk.Coin) *MsgConvertToAtto {
	return &MsgConvertToAtto{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgConvertToAtto) Route() string {
	return RouterKey
}

func (msg *MsgConvertToAtto) Type() string {
	return TypeMsgConvertToAtto
}

func (msg *MsgConvertToAtto) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConvertToAtto) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConvertToAtto) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
