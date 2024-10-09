package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAction = "update_action"

var _ sdk.Msg = &MsgUpdateAction{}

func NewMsgUpdateAction(creator string, nftSchemaCode string, base64UpdateAction string) *MsgUpdateAction {
	return &MsgUpdateAction{
		Creator:            creator,
		NftSchemaCode:      nftSchemaCode,
		Base64UpdateAction: base64UpdateAction,
	}
}

func (msg *MsgUpdateAction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAction) Type() string {
	return TypeMsgUpdateAction
}

func (msg *MsgUpdateAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
