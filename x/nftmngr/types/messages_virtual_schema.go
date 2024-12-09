package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVirtualSchema = "create_vir_schema"
	TypeMsgUpdateVirtualSchema = "update_vir_schema"
	TypeMsgDeleteVirtualSchema = "delete_vir_schema"
)

var _ sdk.Msg = &MsgCreateVirtualSchema{}

func NewMsgCreateVirtualSchema(
	creator string,
	index string,
	code string,

) *MsgCreateVirtualSchema {
	return &MsgCreateVirtualSchema{
		Creator: creator,
		Index:   index,
		Code:    code,
	}
}

func (msg *MsgCreateVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateVirtualSchema) Type() string {
	return TypeMsgCreateVirtualSchema
}

func (msg *MsgCreateVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVirtualSchema{}

func NewMsgUpdateVirtualSchema(
	creator string,
	index string,
	code string,

) *MsgUpdateVirtualSchema {
	return &MsgUpdateVirtualSchema{
		Creator: creator,
		Index:   index,
		Code:    code,
	}
}

func (msg *MsgUpdateVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVirtualSchema) Type() string {
	return TypeMsgUpdateVirtualSchema
}

func (msg *MsgUpdateVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVirtualSchema{}

func NewMsgDeleteVirtualSchema(
	creator string,
	index string,

) *MsgDeleteVirtualSchema {
	return &MsgDeleteVirtualSchema{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVirtualSchema) Type() string {
	return TypeMsgDeleteVirtualSchema
}

func (msg *MsgDeleteVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
