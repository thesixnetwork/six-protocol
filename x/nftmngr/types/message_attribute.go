package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAddAttribute   = "add_attribute"
	TypeMsgShowAttributes = "show_attributes"
)

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
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgShowAttributes{}

func NewMsgShowAttributes(creator string, nftSchemaCode string, show bool, attributeNames []string) *MsgShowAttributes {
	return &MsgShowAttributes{
		Creator:        creator,
		NftSchemaCode:  nftSchemaCode,
		Show:           show,
		AttributeNames: attributeNames,
	}
}

func (msg *MsgShowAttributes) Route() string {
	return RouterKey
}

func (msg *MsgShowAttributes) Type() string {
	return TypeMsgShowAttributes
}

func (msg *MsgShowAttributes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgShowAttributes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgShowAttributes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgResyncAttributes = "resync_attributes"

var _ sdk.Msg = &MsgResyncAttributes{}

func NewMsgResyncAttributes(creator string, nftSchemaCode string, tokenId string) *MsgResyncAttributes {
	return &MsgResyncAttributes{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
	}
}

func (msg *MsgResyncAttributes) Route() string {
	return RouterKey
}

func (msg *MsgResyncAttributes) Type() string {
	return TypeMsgResyncAttributes
}

func (msg *MsgResyncAttributes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResyncAttributes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResyncAttributes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
