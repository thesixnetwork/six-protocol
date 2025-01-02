package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgResyncAttributes = "resync_attributes"

var _ sdk.Msg = &MsgResyncAttributes{}

func NewMsgResyncAttributes(creator string, nftSchemaCode string, tokenId string) *MsgResyncAttributes {
	return &MsgResyncAttributes{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
	}
}

func (msg *MsgResyncAttributes) Route() string {
	return RouterKey
}

func (msg *MsgResyncAttributes) Type() string {
	return TypeMsgResyncAttributes
}

func (msg *MsgResyncAttributes) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgResyncAttributes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResyncAttributes) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
