package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMintperm{}

func NewMsgCreateMintperm(
	creator string,
	token string,
	address string,
) *MsgCreateMintperm {
	return &MsgCreateMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}

func (msg *MsgCreateMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMintperm{}

func NewMsgUpdateMintperm(
	creator string,
	token string,
	address string,
) *MsgUpdateMintperm {
	return &MsgUpdateMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}

func (msg *MsgUpdateMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMintperm{}

func NewMsgDeleteMintperm(
	creator string,
	token string,
	address string,
) *MsgDeleteMintperm {
	return &MsgDeleteMintperm{
		Creator: creator,
		Token:   token,
		Address: address,
	}
}

func (msg *MsgDeleteMintperm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
