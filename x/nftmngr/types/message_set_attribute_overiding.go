package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetAttributeOveriding = "set_attribute_overiding"

var _ sdk.Msg = &MsgSetAttributeOveriding{}

func NewMsgSetAttributeOveriding(creator string, schemaCode string, newOveridingType int32) *MsgSetAttributeOveriding {
	return &MsgSetAttributeOveriding{
		Creator:          creator,
		SchemaCode:       schemaCode,
		NewOveridingType: newOveridingType,
	}
}

func (msg *MsgSetAttributeOveriding) Route() string {
	return RouterKey
}

func (msg *MsgSetAttributeOveriding) Type() string {
	return TypeMsgSetAttributeOveriding
}

func (msg *MsgSetAttributeOveriding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetAttributeOveriding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetAttributeOveriding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
