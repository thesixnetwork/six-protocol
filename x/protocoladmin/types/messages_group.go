package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func (msg *MsgCreateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
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

func (msg *MsgUpdateGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteGroup{}

func NewMsgDeleteGroup(
	creator string,
	name string,
) *MsgDeleteGroup {
	return &MsgDeleteGroup{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDeleteGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
