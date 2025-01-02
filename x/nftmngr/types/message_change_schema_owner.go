package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

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
