package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRemoveAdminFromGroup{}

func NewMsgRemoveAdminFromGroup(creator string, name string, address string) *MsgRemoveAdminFromGroup {
	return &MsgRemoveAdminFromGroup{
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgRemoveAdminFromGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
