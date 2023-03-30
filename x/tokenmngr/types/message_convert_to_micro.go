package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConvertToMicro = "convert_to_micro"

var _ sdk.Msg = &MsgConvertToMicro{}

func NewMsgConvertToMicro(creator string, amount sdk.Coin, receiver string) *MsgConvertToMicro {
	return &MsgConvertToMicro{
		Creator: creator,
		Amount:  amount,
		Receiver: receiver,
	}
}

func (msg *MsgConvertToMicro) Route() string {
	return RouterKey
}

func (msg *MsgConvertToMicro) Type() string {
	return TypeMsgConvertToMicro
}

func (msg *MsgConvertToMicro) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConvertToMicro) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConvertToMicro) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
