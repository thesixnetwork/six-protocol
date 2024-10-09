package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAction = "add_action"

var _ sdk.Msg = &MsgAddAction{}

func NewMsgAddAction(creator string, code string, base64NewAction string) *MsgAddAction {
	return &MsgAddAction{
		Creator:         creator,
		Code:            code,
		Base64NewAction: base64NewAction,
	}
}

func (msg *MsgAddAction) Route() string {
	return RouterKey
}

func (msg *MsgAddAction) Type() string {
	return TypeMsgAddAction
}

func (msg *MsgAddAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
