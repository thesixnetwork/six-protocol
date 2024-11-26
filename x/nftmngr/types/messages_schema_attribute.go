package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgUpdateSchemaAttribute = "update_schema_attribute"
	TypeMsgDeleteSchemaAttribute = "delete_schema_attribute"
)

var _ sdk.Msg = &MsgUpdateSchemaAttribute{}

func NewMsgUpdateSchemaAttribute(
	creator string,
	nftSchemaCode string,
	base64UpdateAttriuteDefenition string,

) *MsgUpdateSchemaAttribute {
	return &MsgUpdateSchemaAttribute{
		Creator:                        creator,
		NftSchemaCode:                  nftSchemaCode,
		Base64UpdateAttriuteDefenition: base64UpdateAttriuteDefenition,
	}
}

func (msg *MsgUpdateSchemaAttribute) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSchemaAttribute) Type() string {
	return TypeMsgUpdateSchemaAttribute
}

func (msg *MsgUpdateSchemaAttribute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateSchemaAttribute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSchemaAttribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
