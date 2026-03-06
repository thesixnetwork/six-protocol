package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateOptions{}

func NewMsgCreateOptions(creator string, defaultMintee string) *MsgCreateOptions {
	return &MsgCreateOptions{
		Creator:       creator,
		DefaultMintee: defaultMintee,
	}
}

func (msg *MsgCreateOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateOptions{}

func NewMsgUpdateOptions(creator string, defaultMintee string) *MsgUpdateOptions {
	return &MsgUpdateOptions{
		Creator:       creator,
		DefaultMintee: defaultMintee,
	}
}

func (msg *MsgUpdateOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteOptions{}

func NewMsgDeleteOptions(creator string) *MsgDeleteOptions {
	return &MsgDeleteOptions{
		Creator: creator,
	}
}

func (msg *MsgDeleteOptions) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
