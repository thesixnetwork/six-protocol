package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateGroup = "create_group"
	TypeMsgUpdateGroup = "update_group"
	TypeMsgDeleteGroup = "delete_group"
)

var _ sdk.Msg = &MsgCreateGroup{}

func NewMsgCreateGroup(
	creator string,
	name string,

) *MsgCreateGroup {
	return &MsgCreateGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgCreateGroup) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroup) Type() string {
	return TypeMsgCreateGroup
}

func (msg *MsgCreateGroup) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgCreateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateGroup{}

func NewMsgUpdateGroup(
	creator string,
	name string,

) *MsgUpdateGroup {
	return &MsgUpdateGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgUpdateGroup) Route() string {
	return RouterKey
}

func (msg *MsgUpdateGroup) Type() string {
	return TypeMsgUpdateGroup
}

func (msg *MsgUpdateGroup) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgUpdateGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteGroup{}

func NewMsgDeleteGroup(
	owner string,
	name string,

) *MsgDeleteGroup {
	return &MsgDeleteGroup{
		Creator: owner,
		Name:    name,
	}
}
func (msg *MsgDeleteGroup) Route() string {
	return RouterKey
}

func (msg *MsgDeleteGroup) Type() string {
	return TypeMsgDeleteGroup
}

func (msg *MsgDeleteGroup) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgDeleteGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
