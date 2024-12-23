package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVirtualSchemaProposal = "create_virtual_schema"
	TypeMsgDeleteVirtualSchema         = "delete_virtual_schema"
)

var _ sdk.Msg = &MsgCreateVirtualSchemaProposal{}

func NewMsgCreateVirtualSchemaProposal(
	creator string,
	code string,
	request []VirtualSchemaRegistryRequest,
) *MsgCreateVirtualSchemaProposal {
	return &MsgCreateVirtualSchemaProposal{
		Creator:              creator,
		VirtualNftSchemaCode: code,
		Registry:             request,
	}
}

func (msg *MsgCreateVirtualSchemaProposal) Route() string {
	return RouterKey
}

func (msg *MsgCreateVirtualSchemaProposal) Type() string {
	return TypeMsgCreateVirtualSchemaProposal
}

func (msg *MsgCreateVirtualSchemaProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVirtualSchemaProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVirtualSchemaProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVirtualSchema{}

func NewMsgDeleteVirtualSchema(
	creator string,
	code string,
) *MsgDeleteVirtualSchema {
	return &MsgDeleteVirtualSchema{
		Creator:              creator,
		VirtualNftSchemaCode: code,
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
