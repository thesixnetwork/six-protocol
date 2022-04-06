package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAdminToGroup = "add_admin_to_group"

var _ sdk.Msg = &MsgAddAdminToGroup{}

func NewMsgAddAdminToGroup(creator string, name string, address string) *MsgAddAdminToGroup {
	return &MsgAddAdminToGroup{
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgAddAdminToGroup) Route() string {
	return RouterKey
}

func (msg *MsgAddAdminToGroup) Type() string {
	return TypeMsgAddAdminToGroup
}

func (msg *MsgAddAdminToGroup) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAdminToGroup) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAdminToGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
