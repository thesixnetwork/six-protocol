package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateActionSigner = "create_action_signer"
	TypeMsgUpdateActionSigner = "update_action_signer"
	TypeMsgDeleteActionSigner = "delete_action_signer"
)

var _ sdk.Msg = &MsgCreateActionSigner{}

func NewMsgCreateActionSigner(
	creator string,
	bse64EncodedSetSignerAction string,
) *MsgCreateActionSigner {
	return &MsgCreateActionSigner{
		Creator:                      creator,
		Base64EncodedSetSignerAction: bse64EncodedSetSignerAction,
	}
}

func (msg *MsgCreateActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgCreateActionSigner) Type() string {
	return TypeMsgCreateActionSigner
}

func (msg *MsgCreateActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateActionSigner{}

func NewMsgUpdateActionSigner(
	creator string,
	base64EncodedSetSignerAction string,

) *MsgUpdateActionSigner {
	return &MsgUpdateActionSigner{
		Creator:                      creator,
		Base64EncodedSetSignerAction: base64EncodedSetSignerAction,
	}
}

func (msg *MsgUpdateActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgUpdateActionSigner) Type() string {
	return TypeMsgUpdateActionSigner
}

func (msg *MsgUpdateActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActionSigner{}

func NewMsgDeleteActionSigner(
	creator string,
	base64EncodedSetSignerAction string,

) *MsgDeleteActionSigner {
	return &MsgDeleteActionSigner{
		Creator:                      creator,
		Base64EncodedSetSignerAction: base64EncodedSetSignerAction,
	}
}
func (msg *MsgDeleteActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgDeleteActionSigner) Type() string {
	return TypeMsgDeleteActionSigner
}

func (msg *MsgDeleteActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
