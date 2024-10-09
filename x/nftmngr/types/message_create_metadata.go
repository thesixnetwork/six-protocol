package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMetadata = "create_metadata"

var _ sdk.Msg = &MsgCreateMetadata{}

func NewMsgCreateMetadata(creator string, nftSchemaCode string, tokenId string, base64NFTData string) *MsgCreateMetadata {
	return &MsgCreateMetadata{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		TokenId:       tokenId,
		Base64NFTData: base64NFTData,
	}
}

func (msg *MsgCreateMetadata) Route() string {
	return RouterKey
}

func (msg *MsgCreateMetadata) Type() string {
	return TypeMsgCreateMetadata
}

func (msg *MsgCreateMetadata) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMetadata) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
