package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateActionExecutor = "create_action_executor"
	TypeMsgUpdateActionExecutor = "update_action_executor"
	TypeMsgDeleteActionExecutor = "delete_action_executor"
)

var _ sdk.Msg = &MsgCreateActionExecutor{}

func NewMsgCreateActionExecutor(
	creator string,
	nftSchemaCode string,
	executorAddress string,

) *MsgCreateActionExecutor {
	return &MsgCreateActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaCode,
		ExecutorAddress: executorAddress,
	}
}

func (msg *MsgCreateActionExecutor) Route() string {
	return RouterKey
}

func (msg *MsgCreateActionExecutor) Type() string {
	return TypeMsgCreateActionExecutor
}

func (msg *MsgCreateActionExecutor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateActionExecutor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateActionExecutor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateActionExecutor{}

func NewMsgUpdateActionExecutor(
	creator string,
	nftSchemaCode string,
	executorAddress string,

) *MsgUpdateActionExecutor {
	return &MsgUpdateActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaCode,
		ExecutorAddress: executorAddress,
	}
}

func (msg *MsgUpdateActionExecutor) Route() string {
	return RouterKey
}

func (msg *MsgUpdateActionExecutor) Type() string {
	return TypeMsgUpdateActionExecutor
}

func (msg *MsgUpdateActionExecutor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateActionExecutor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateActionExecutor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteActionExecutor{}

func NewMsgDeleteActionExecutor(
	creator string,
	nftSchemaCode string,
	executorAddress string,

) *MsgDeleteActionExecutor {
	return &MsgDeleteActionExecutor{
		Creator:         creator,
		NftSchemaCode:   nftSchemaCode,
		ExecutorAddress: executorAddress,
	}
}
func (msg *MsgDeleteActionExecutor) Route() string {
	return RouterKey
}

func (msg *MsgDeleteActionExecutor) Type() string {
	return TypeMsgDeleteActionExecutor
}

func (msg *MsgDeleteActionExecutor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteActionExecutor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteActionExecutor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
