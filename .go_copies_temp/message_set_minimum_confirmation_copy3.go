package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMinimumConfirmation = "set_minimum_confirmation"

var _ sdk.Msg = &MsgSetMinimumConfirmation{}

func NewMsgSetMinimumConfirmation(creator string, newConfirmation string) *MsgSetMinimumConfirmation {
	return &MsgSetMinimumConfirmation{
		Creator:         creator,
		NewConfirmation: newConfirmation,
	}
}

func (msg *MsgSetMinimumConfirmation) Route() string {
	return RouterKey
}

func (msg *MsgSetMinimumConfirmation) Type() string {
	return TypeMsgSetMinimumConfirmation
}

func (msg *MsgSetMinimumConfirmation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMinimumConfirmation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMinimumConfirmation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
