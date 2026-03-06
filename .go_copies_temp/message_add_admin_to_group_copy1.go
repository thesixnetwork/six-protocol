package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAdminToGroup{}

func NewMsgAddAdminToGroup(creator string, name string, address string) *MsgAddAdminToGroup {
	return &MsgAddAdminToGroup{
		Creator: creator,
		Name:    name,
		Address: address,
	}
}

func (msg *MsgAddAdminToGroup) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
