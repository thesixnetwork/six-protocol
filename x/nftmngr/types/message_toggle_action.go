package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgToggleAction = "toggle_action"

var _ sdk.Msg = &MsgToggleAction{}

func NewMsgToggleAction(creator string, code string, action string, status bool) *MsgToggleAction {
	return &MsgToggleAction{
		Creator: creator,
		Code:    code,
		Action:  action,
		Status:  status,
	}
}

func (msg *MsgToggleAction) Route() string {
	return RouterKey
}

func (msg *MsgToggleAction) Type() string {
	return TypeMsgToggleAction
}

func (msg *MsgToggleAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
