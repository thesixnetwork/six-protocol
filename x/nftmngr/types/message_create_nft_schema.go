package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateNFTSchema = "create_nft_schema"

var _ sdk.Msg = &MsgCreateNFTSchema{}

func NewMsgCreateNFTSchema(creator string, nftSchemaBase64 string) *MsgCreateNFTSchema {
	return &MsgCreateNFTSchema{
		Creator:         creator,
		NftSchemaBase64: nftSchemaBase64,
	}
}

func (msg *MsgCreateNFTSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateNFTSchema) Type() string {
	return TypeMsgCreateNFTSchema
}

func (msg *MsgCreateNFTSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNFTSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNFTSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
