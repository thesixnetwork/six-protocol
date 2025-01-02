package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMultiMetadata = "create_multi_metadata"

var _ sdk.Msg = &MsgCreateMultiMetadata{}

func NewMsgCreateMultiMetadata(creator string, nftSchemaCode string, tokenId []string, base64NFTData string) *MsgCreateMultiMetadata {
	return &MsgCreateMultiMetadata{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Base64NFTData: base64NFTData,
	}
}

func (msg *MsgCreateMultiMetadata) Route() string {
	return RouterKey
}

func (msg *MsgCreateMultiMetadata) Type() string {
	return TypeMsgCreateMultiMetadata
}

func (msg *MsgCreateMultiMetadata) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMultiMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMultiMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
