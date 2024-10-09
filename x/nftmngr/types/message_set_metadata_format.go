package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMetadataFormat = "set_metadata_format"

var _ sdk.Msg = &MsgSetMetadataFormat{}

func NewMsgSetMetadataFormat(creator string, schemaCode string, newFormat string) *MsgSetMetadataFormat {
	return &MsgSetMetadataFormat{
		Creator:    creator,
		SchemaCode: schemaCode,
		NewFormat:  newFormat,
	}
}

func (msg *MsgSetMetadataFormat) Route() string {
	return RouterKey
}

func (msg *MsgSetMetadataFormat) Type() string {
	return TypeMsgSetMetadataFormat
}

func (msg *MsgSetMetadataFormat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMetadataFormat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMetadataFormat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
