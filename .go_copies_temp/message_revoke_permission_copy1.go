package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokePermission = "revoke_permission"

var _ sdk.Msg = &MsgRevokePermission{}

func NewMsgRevokePermission(creator string, name string, revokee string) *MsgRevokePermission {
	return &MsgRevokePermission{
		Creator: creator,
		Name:    name,
		Revokee: revokee,
	}
}

func (msg *MsgRevokePermission) Route() string {
	return RouterKey
}

func (msg *MsgRevokePermission) Type() string {
	return TypeMsgRevokePermission
}

func (msg *MsgRevokePermission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokePermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokePermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
