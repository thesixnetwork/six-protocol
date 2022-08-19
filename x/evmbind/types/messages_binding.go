package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBinding = "create_binding"
	TypeMsgUpdateBinding = "update_binding"
	TypeMsgDeleteBinding = "delete_binding"
)

var _ sdk.Msg = &MsgCreateBinding{}

func NewMsgCreateBinding(
	creator string,
	ethAddress string,
	ethSignature string,
	signMessage string,

) *MsgCreateBinding {
	return &MsgCreateBinding{
		Creator:      creator,
		EthAddress:   ethAddress,
		EthSignature: ethSignature,
		SignMessage:  signMessage,
	}
}

func (msg *MsgCreateBinding) Route() string {
	return RouterKey
}

func (msg *MsgCreateBinding) Type() string {
	return TypeMsgCreateBinding
}

func (msg *MsgCreateBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBinding{}

func NewMsgUpdateBinding(
	creator string,
	ethAddress string,
	ethSignature string,
	signMessage string,

) *MsgUpdateBinding {
	return &MsgUpdateBinding{
		Creator:      creator,
		EthAddress:   ethAddress,
		EthSignature: ethSignature,
		SignMessage:  signMessage,
	}
}

func (msg *MsgUpdateBinding) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBinding) Type() string {
	return TypeMsgUpdateBinding
}

func (msg *MsgUpdateBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBinding{}

func NewMsgDeleteBinding(
	creator string,
	ethAddress string,

) *MsgDeleteBinding {
	return &MsgDeleteBinding{
		Creator:    creator,
		EthAddress: ethAddress,
	}
}
func (msg *MsgDeleteBinding) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBinding) Type() string {
	return TypeMsgDeleteBinding
}

func (msg *MsgDeleteBinding) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBinding) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBinding) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
