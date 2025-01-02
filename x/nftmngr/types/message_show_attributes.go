package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgShowAttributes = "show_attributes"

var _ sdk.Msg = &MsgShowAttributes{}

func NewMsgShowAttributes(creator string, nftSchemaCode string, show bool, attributeNames []string) *MsgShowAttributes {
	return &MsgShowAttributes{
		Creator:        creator,
		NftSchemaCode:  nftSchemaCode,
		Show:           show,
		AttributeNames: attributeNames,
	}
}

func (msg *MsgShowAttributes) Route() string {
	return RouterKey
}

func (msg *MsgShowAttributes) Type() string {
	return TypeMsgShowAttributes
}

func (msg *MsgShowAttributes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgShowAttributes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgShowAttributes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
