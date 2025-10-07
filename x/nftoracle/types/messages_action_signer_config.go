package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateActionSignerConfig = "create_action_signer_config"
	TypeMsgUpdateActionSignerConfig = "update_action_signer_config"
	TypeMsgDeleteActionSignerConfig = "delete_action_signer_config"
)

var _ sdk.Msg = &MsgCreateActionSignerConfig{}

func NewMsgCreateActionSignerConfig(
	creator string,
	chain string,
	contractAddress string,
) *MsgCreateActionSignerConfig {
	return &MsgCreateActionSignerConfig{
		Creator:         creator,
		Chain:           chain,
		ContractAddress: contractAddress,
	}
}

func (msg *MsgCreateActionSignerConfig) Route() string {
	return RouterKey
}

func (msg *MsgCreateActionSignerConfig) Type() string {
	return TypeMsgCreateActionSignerConfig
}

func (msg *MsgCreateActionSignerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActionSignerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActionSignerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateActionSignerConfig{}

func NewMsgUpdateActionSignerConfig(
	creator string,
	chain string,
	contractAddress string,
) *MsgUpdateActionSignerConfig {
	return &MsgUpdateActionSignerConfig{
		Creator:         creator,
		Chain:           chain,
		ContractAddress: contractAddress,
	}
}

func (msg *MsgUpdateActionSignerConfig) Route() string {
	return RouterKey
}

func (msg *MsgUpdateActionSignerConfig) Type() string {
	return TypeMsgUpdateActionSignerConfig
}

func (msg *MsgUpdateActionSignerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateActionSignerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateActionSignerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActionSignerConfig{}

func NewMsgDeleteActionSignerConfig(
	creator string,
	chain string,
) *MsgDeleteActionSignerConfig {
	return &MsgDeleteActionSignerConfig{
		Creator: creator,
		Chain:   chain,
	}
}

func (msg *MsgDeleteActionSignerConfig) Route() string {
	return RouterKey
}

func (msg *MsgDeleteActionSignerConfig) Type() string {
	return TypeMsgDeleteActionSignerConfig
}

func (msg *MsgDeleteActionSignerConfig) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActionSignerConfig) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActionSignerConfig) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
