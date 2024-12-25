package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
