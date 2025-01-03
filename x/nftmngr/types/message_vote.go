package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgVoteCreateVirtualSchema = "vote_create_virtual_schema"
	TypeMsgVoteEnableVirtualSchema = "vote_enable_virtual_schema"
	TypeMsgVoteDisableVirtualSchema = "vote_disable_virtual_schema"
)

var _ sdk.Msg = &MsgVoteCreateVirtualSchema{}

func NewMsgVoteCreateVirtualSchema(creator, id, nftSchemaCode string, option RegistryStatus) *MsgVoteCreateVirtualSchema {
	return &MsgVoteCreateVirtualSchema{
		Creator:       creator,
		Id:            id,
		NftSchemaCode: nftSchemaCode,
		Option:        option,
	}
}

func (msg *MsgVoteCreateVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgVoteCreateVirtualSchema) Type() string {
	return TypeMsgVoteCreateVirtualSchema
}

func (msg *MsgVoteCreateVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteCreateVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteCreateVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func NewMsgVoteEnableVirtualSchema(creator, id, nftSchemaCode string, option RegistryStatus) *MsgVoteEnableVirtualSchema {
	return &MsgVoteEnableVirtualSchema{
		Creator:       creator,
		Id:            id,
		NftSchemaCode: nftSchemaCode,
		Option:        option,
	}
}

func (msg *MsgVoteEnableVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgVoteEnableVirtualSchema) Type() string {
	return TypeMsgVoteEnableVirtualSchema
}

func (msg *MsgVoteEnableVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteEnableVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteEnableVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgVoteDisableVirtualSchema{}

func NewMsgVoteDisableVirtualSchema(creator, id, nftSchemaCode string, option RegistryStatus) *MsgVoteDisableVirtualSchema {
	return &MsgVoteDisableVirtualSchema{
		Creator:       creator,
		Id:            id,
		NftSchemaCode: nftSchemaCode,
		Option:        option,
	}
}

func (msg *MsgVoteDisableVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgVoteDisableVirtualSchema) Type() string {
	return TypeMsgVoteDisableVirtualSchema
}

func (msg *MsgVoteDisableVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteDisableVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteDisableVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}