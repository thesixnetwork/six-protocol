package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitVerifyCollectionOwner = "submit_verify_collection_owner"

var _ sdk.Msg = &MsgSubmitVerifyCollectionOwner{}

func NewMsgSubmitVerifyCollectionOwner(creator string, verifyRequestID uint64, schemaCode string, base64OriginContractInfo string) *MsgSubmitVerifyCollectionOwner {
	return &MsgSubmitVerifyCollectionOwner{
		Creator:                  creator,
		VerifyRequestID:          verifyRequestID,
		NftSchemaCode:            schemaCode,
		Base64OriginContractInfo: base64OriginContractInfo,
	}
}

func (msg *MsgSubmitVerifyCollectionOwner) Route() string {
	return RouterKey
}

func (msg *MsgSubmitVerifyCollectionOwner) Type() string {
	return TypeMsgSubmitVerifyCollectionOwner
}

func (msg *MsgSubmitVerifyCollectionOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitVerifyCollectionOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitVerifyCollectionOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
