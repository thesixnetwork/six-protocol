package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateVirtual = "create_virtual_action"
	TypeMsgUpdateVirtual = "update_virtual_action"
	TypeMsgDeleteVirtual = "delete_virtual_action"
)

var _ sdk.Msg = &MsgCreateVirtualAction{}

func NewMsgCreateVirtualAction(
	creator string,
	code string,
	newActions []*Action,
) *MsgCreateVirtualAction {
	return &MsgCreateVirtualAction{
		Creator:       creator,
		NftSchemaCode: code,
		NewActions:    newActions,
	}
}

func (msg *MsgCreateVirtualAction) Route() string {
	return RouterKey
}

func (msg *MsgCreateVirtualAction) Type() string {
	return TypeMsgCreateVirtual
}

func (msg *MsgCreateVirtualAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVirtualAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVirtualAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgCreateVirtualAction) CheckDuplicateAction() error {
	actionNames := make(map[string]bool)
	for _, action := range msg.NewActions {
		if _, existed := actionNames[action.Name]; existed {
			return sdkerrors.Wrapf(ErrDuplicateActionName, "duplicate action: %s", action)
		}

		paramName := make(map[string]bool)
		for _, param := range action.Params {
			if _, existed := paramName[param.Name]; existed {
				return sdkerrors.Wrapf(ErrInvalidParameter, "duplicate action name: %s", action)
			}
		}
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVirtualAction{}

func NewMsgUpdateVirtual(
	creator string,
	code string,
	newActions []*Action,
) *MsgUpdateVirtualAction {
	return &MsgUpdateVirtualAction{
		Creator:       creator,
		NftSchemaCode: code,
		NewActions:    newActions,
	}
}

func (msg *MsgUpdateVirtualAction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVirtualAction) Type() string {
	return TypeMsgUpdateVirtual
}

func (msg *MsgUpdateVirtualAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVirtualAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVirtualAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

func (msg *MsgUpdateVirtualAction) CheckDuplicateAction() error {
	actionNames := make(map[string]bool)
	for _, action := range msg.NewActions {
		if _, existed := actionNames[action.Name]; existed {
			return sdkerrors.Wrapf(ErrDuplicateActionName, "duplicate action: %s", action)
		}

		paramName := make(map[string]bool)
		for _, param := range action.Params {
			if _, existed := paramName[param.Name]; existed {
				return sdkerrors.Wrapf(ErrInvalidParameter, "duplicate action name: %s", action)
			}
		}
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteVirtualAction{}

func NewMsgDeleteVirtual(
	creator string,
	code string,
	actionName string,
) *MsgDeleteVirtualAction {
	return &MsgDeleteVirtualAction{
		Creator:       creator,
		NftSchemaCode: code,
		Name:          actionName,
	}
}

func (msg *MsgDeleteVirtualAction) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVirtualAction) Type() string {
	return TypeMsgDeleteVirtual
}

func (msg *MsgDeleteVirtualAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVirtualAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVirtualAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
