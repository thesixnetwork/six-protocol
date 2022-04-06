package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMintperm = "create_mintperm"
	TypeMsgUpdateMintperm = "update_mintperm"
	TypeMsgDeleteMintperm = "delete_mintperm"
)

var _ sdk.Msg = &MsgCreateMintperm{}

func NewMsgCreateMintperm(
	creator string,
	token string,
	address string,

) *MsgCreateMintperm {
	return &MsgCreateMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}

func (msg *MsgCreateMintperm) Route() string {
	return RouterKey
}

func (msg *MsgCreateMintperm) Type() string {
	return TypeMsgCreateMintperm
}

func (msg *MsgCreateMintperm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMintperm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMintperm{}

func NewMsgUpdateMintperm(
	creator string,
	token string,
	address string,

) *MsgUpdateMintperm {
	return &MsgUpdateMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}

func (msg *MsgUpdateMintperm) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMintperm) Type() string {
	return TypeMsgUpdateMintperm
}

func (msg *MsgUpdateMintperm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMintperm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMintperm{}

func NewMsgDeleteMintperm(
	creator string,
	token string,
	address string,

) *MsgDeleteMintperm {
	return &MsgDeleteMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}
func (msg *MsgDeleteMintperm) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMintperm) Type() string {
	return TypeMsgDeleteMintperm
}

func (msg *MsgDeleteMintperm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMintperm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
