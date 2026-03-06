package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgGrantPermission = "grant_permission"

var _ sdk.Msg = &MsgGrantPermission{}

func NewMsgGrantPermission(creator string, name string, grantee string) *MsgGrantPermission {
	return &MsgGrantPermission{
		Creator: creator,
		Name:    name,
		Grantee: grantee,
	}
}

func (msg *MsgGrantPermission) Route() string {
	return RouterKey
}

func (msg *MsgGrantPermission) Type() string {
	return TypeMsgGrantPermission
}

func (msg *MsgGrantPermission) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgGrantPermission) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgGrantPermission) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
