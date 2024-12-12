package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDisableVirtualSchemaProposal = "disable_virtual_schema_proposal"

var _ sdk.Msg = &MsgDisableVirtualSchemaProposal{}

func NewMsgDisableVirtualSchemaProposal(creator string, virtualNftSchemaCode string) *MsgDisableVirtualSchemaProposal {
	return &MsgDisableVirtualSchemaProposal{
		Creator:              creator,
		VirtualNftSchemaCode: virtualNftSchemaCode,
	}
}

func (msg *MsgDisableVirtualSchemaProposal) Route() string {
	return RouterKey
}

func (msg *MsgDisableVirtualSchemaProposal) Type() string {
	return TypeMsgDisableVirtualSchemaProposal
}

func (msg *MsgDisableVirtualSchemaProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDisableVirtualSchemaProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableVirtualSchemaProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
