package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteCreateVirtualSchema = "vote_create_virtual_schema"

var _ sdk.Msg = &MsgVoteCreateVirtualSchema{}

func NewMsgVoteCreateVirtualSchema(creator string, id string, option RegistryStatus) *MsgVoteCreateVirtualSchema {
	return &MsgVoteCreateVirtualSchema{
		Creator: creator,
		Id:      id,
		Option:  option,
	}
}

func (msg *MsgVoteCreateVirtualSchema) Route() string {
	return RouterKey
}

func (msg *MsgVoteCreateVirtualSchema) Type() string {
	return TypeMsgVoteCreateVirtualSchema
}

func (msg *MsgVoteCreateVirtualSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteCreateVirtualSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteCreateVirtualSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
