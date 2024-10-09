package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAttribute = "add_attribute"

var _ sdk.Msg = &MsgAddAttribute{}

func NewMsgAddAttribute(creator string, code string, location AttributeLocation, newAttibute string) *MsgAddAttribute {
	return &MsgAddAttribute{
		Creator:                     creator,
		Code:                        code,
		Location:                    location,
		Base64NewAttriuteDefenition: newAttibute,
	}
}

func (msg *MsgAddAttribute) Route() string {
	return RouterKey
}

func (msg *MsgAddAttribute) Type() string {
	return TypeMsgAddAttribute
}

func (msg *MsgAddAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
