package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateOptions = "create_options"
	TypeMsgUpdateOptions = "update_options"
	TypeMsgDeleteOptions = "delete_options"
)

var _ sdk.Msg = &MsgCreateOptions{}

func NewMsgCreateOptions(creator string, defaultMintee string) *MsgCreateOptions {
	return &MsgCreateOptions{
		Creator:       creator,
		DefaultMintee: defaultMintee,
	}
}

func (msg *MsgCreateOptions) Route() string {
	return RouterKey
}

func (msg *MsgCreateOptions) Type() string {
	return TypeMsgCreateOptions
}

func (msg *MsgCreateOptions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateOptions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOptions{}

func NewMsgUpdateOptions(creator string, defaultMintee string) *MsgUpdateOptions {
	return &MsgUpdateOptions{
		Creator:       creator,
		DefaultMintee: defaultMintee,
	}
}

func (msg *MsgUpdateOptions) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOptions) Type() string {
	return TypeMsgUpdateOptions
}

func (msg *MsgUpdateOptions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOptions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOptions{}

func NewMsgDeleteOptions(creator string) *MsgDeleteOptions {
	return &MsgDeleteOptions{
		Creator: creator,
	}
}
func (msg *MsgDeleteOptions) Route() string {
	return RouterKey
}

func (msg *MsgDeleteOptions) Type() string {
	return TypeMsgDeleteOptions
}

func (msg *MsgDeleteOptions) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteOptions) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
