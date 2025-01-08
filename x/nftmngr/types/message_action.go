package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddAction = "add_action"

var _ sdk.Msg = &MsgAddAction{}

func NewMsgAddAction(creator string, code string, base64NewAction string) *MsgAddAction {
	return &MsgAddAction{
		Creator:         creator,
		Code:            code,
		Base64NewAction: base64NewAction,
	}
}

func (msg *MsgAddAction) Route() string {
	return RouterKey
}

func (msg *MsgAddAction) Type() string {
	return TypeMsgAddAction
}

func (msg *MsgAddAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgUpdateAction = "update_action"

var _ sdk.Msg = &MsgUpdateAction{}

func NewMsgUpdateAction(creator string, nftSchemaCode string, base64UpdateAction string) *MsgUpdateAction {
	return &MsgUpdateAction{
		Creator:            creator,
		NftSchemaCode:      nftSchemaCode,
		Base64UpdateAction: base64UpdateAction,
	}
}

func (msg *MsgUpdateAction) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAction) Type() string {
	return TypeMsgUpdateAction
}

func (msg *MsgUpdateAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgToggleAction = "toggle_action"

var _ sdk.Msg = &MsgToggleAction{}

func NewMsgToggleAction(creator string, code string, action string, status bool) *MsgToggleAction {
	return &MsgToggleAction{
		Creator: creator,
		Code:    code,
		Action:  action,
		Status:  status,
	}
}

func (msg *MsgToggleAction) Route() string {
	return RouterKey
}

func (msg *MsgToggleAction) Type() string {
	return TypeMsgToggleAction
}

func (msg *MsgToggleAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgToggleAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgToggleAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgPerformActionByAdmin = "perform_action_by_admin"

var _ sdk.Msg = &MsgPerformActionByAdmin{}

func NewMsgPerformActionByAdmin(creator, nftSchemaCode, tokenId, action, actionPrams, refId string) *MsgPerformActionByAdmin {
	// string to json object of  []*ActionParameter
	var actionPrams_ []*ActionParameter
	err := json.Unmarshal([]byte(actionPrams), &actionPrams_)
	if err != nil {
		panic(err)
	}

	return &MsgPerformActionByAdmin{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Action:        action,
		Parameters:    actionPrams_,
		RefId:         refId,
	}
}

func (msg *MsgPerformActionByAdmin) Route() string {
	return RouterKey
}

func (msg *MsgPerformActionByAdmin) Type() string {
	return TypeMsgPerformActionByAdmin
}

func (msg *MsgPerformActionByAdmin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformActionByAdmin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerformActionByAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgPerformVirtualAction = "perform_virtual_action"

var _ sdk.Msg = &MsgPerformVirtualAction{}

func NewMsgPerformVirtualAction(creator, nftSchemaName string, tokenIdsInput []*TokenIdMap, actionName string, paramInput []*ActionParameter, refId string) *MsgPerformVirtualAction {
	return &MsgPerformVirtualAction{
		Creator:       creator,
		NftSchemaName: nftSchemaName,
		TokenIdMap:    tokenIdsInput,
		Action:        actionName,
		Parameters:    paramInput,
		RefId:         refId,
	}
}

func (msg *MsgPerformVirtualAction) Route() string {
	return RouterKey
}

func (msg *MsgPerformVirtualAction) Type() string {
	return TypeMsgPerformVirtualAction
}

func (msg *MsgPerformVirtualAction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPerformVirtualAction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPerformVirtualAction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
