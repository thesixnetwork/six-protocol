package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitMintResponse = "submit_mint_response"

var _ sdk.Msg = &MsgSubmitMintResponse{}

func NewMsgSubmitMintResponse(creator string, mintRequestID uint64, base64NftData string) *MsgSubmitMintResponse {
	return &MsgSubmitMintResponse{
		Creator:       creator,
		MintRequestID: mintRequestID,
		Base64NftData: base64NftData,
	}
}

func (msg *MsgSubmitMintResponse) Route() string {
	return RouterKey
}

func (msg *MsgSubmitMintResponse) Type() string {
	return TypeMsgSubmitMintResponse
}

func (msg *MsgSubmitMintResponse) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitMintResponse) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitMintResponse) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s) ,the address was %v", err, msg.Creator)
	}
	return nil
}
