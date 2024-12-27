package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeOrgOwner = "chage_org_owner"

var _ sdk.Msg = &MsgChangeOrgOwner{}

func NewMsgChangeOrgOwner(creator string, orgName string, toNewOwner string) *MsgChangeOrgOwner {
	return &MsgChangeOrgOwner{
		Creator:    creator,
		OrgName:    orgName,
		ToNewOwner: toNewOwner,
	}
}

func (msg *MsgChangeOrgOwner) Route() string {
	return RouterKey
}

func (msg *MsgChangeOrgOwner) Type() string {
	return TypeMsgChangeOrgOwner
}

func (msg *MsgChangeOrgOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeOrgOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeOrgOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgChangeSchemaOwner = "change_schema_owner"

var _ sdk.Msg = &MsgChangeSchemaOwner{}

func NewMsgChangeSchemaOwner(creator string, nftSchemaCode string, newOwner string) *MsgChangeSchemaOwner {
	return &MsgChangeSchemaOwner{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		NewOwner:      newOwner,
	}
}

func (msg *MsgChangeSchemaOwner) Route() string {
	return RouterKey
}

func (msg *MsgChangeSchemaOwner) Type() string {
	return TypeMsgChangeSchemaOwner
}

func (msg *MsgChangeSchemaOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeSchemaOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeSchemaOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
