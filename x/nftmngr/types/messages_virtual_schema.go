package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVirtualSchemaProposal = "create_virtual_schema_proposal"
	TypeMsgDeleteVirtualSchemaProposal = "delete_virtual_schema_proposal"
	TypeMsgEnableVirtualSchemaProposal = "enable_virtual_schema_proposal"
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
	return TypeMsgDeleteVirtualSchemaProposal
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

var _ sdk.Msg = &MsgEnableVirtualSchemaProposal{}

func NewMsgEnableVirtualSchema(creator string, virtualNftSchemaCode string) *MsgEnableVirtualSchemaProposal {
	return &MsgEnableVirtualSchemaProposal{
		Creator:              creator,
		VirtualNftSchemaCode: virtualNftSchemaCode,
	}
}

func (msg *MsgEnableVirtualSchemaProposal) Route() string {
	return RouterKey
}

func (msg *MsgEnableVirtualSchemaProposal) Type() string {
	return TypeMsgDeleteVirtualSchemaProposal
}

func (msg *MsgEnableVirtualSchemaProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEnableVirtualSchemaProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableVirtualSchemaProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
