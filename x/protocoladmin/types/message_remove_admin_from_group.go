package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveAdminFromGroup = "remove_admin_from_group"

var _ sdk.Msg = &MsgRemoveAdminFromGroup{}

func NewMsgRemoveAdminFromGroup(creator string, name string, address string) *MsgRemoveAdminFromGroup {
	return &MsgRemoveAdminFromGroup{
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgRemoveAdminFromGroup) Route() string {
	return RouterKey
}

func (msg *MsgRemoveAdminFromGroup) Type() string {
	return TypeMsgRemoveAdminFromGroup
}

func (msg *MsgRemoveAdminFromGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveAdminFromGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveAdminFromGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
