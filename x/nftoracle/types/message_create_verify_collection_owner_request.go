package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateVerifyCollectionOwnerRequest = "create_verify_collection_owner_request"

var _ sdk.Msg = &MsgCreateVerifyCollectionOwnerRequest{}

func NewMsgCreateVerifyCollectionOwnerRequest(creator string, nftSchemaCode string, base64VerifyRequestorSignature string, requiredConfirm uint64) *MsgCreateVerifyCollectionOwnerRequest {
	return &MsgCreateVerifyCollectionOwnerRequest{
		Creator:                        creator,
		NftSchemaCode:                  nftSchemaCode,
		Base64VerifyRequestorSignature: base64VerifyRequestorSignature,
		RequiredConfirm:                requiredConfirm,
	}
}

func (msg *MsgCreateVerifyCollectionOwnerRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateVerifyCollectionOwnerRequest) Type() string {
	return TypeMsgCreateVerifyCollectionOwnerRequest
}

func (msg *MsgCreateVerifyCollectionOwnerRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVerifyCollectionOwnerRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVerifyCollectionOwnerRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
