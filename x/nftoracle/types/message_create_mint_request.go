package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateMintRequest = "create_mint_request"

var _ sdk.Msg = &MsgCreateMintRequest{}

func NewMsgCreateMintRequest(creator string, nftSchemaCode string, tokenId string, requiredConfirm uint64) *MsgCreateMintRequest {
	return &MsgCreateMintRequest{
		Creator:         creator,
		NftSchemaCode:   nftSchemaCode,
		TokenId:         tokenId,
		RequiredConfirm: requiredConfirm,
	}
}

func (msg *MsgCreateMintRequest) Route() string {
	return RouterKey
}

func (msg *MsgCreateMintRequest) Type() string {
	return TypeMsgCreateMintRequest
}

func (msg *MsgCreateMintRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMintRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMintRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s) ,the address is %v", err, msg.Creator)
	}
	return nil
}
