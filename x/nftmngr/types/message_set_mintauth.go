package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMintauth = "set_mintauth"

var _ sdk.Msg = &MsgSetMintauth{}

func NewMsgSetMintauth(creator string, nftSchemaCode string, authorizeTo AuthorizeTo) *MsgSetMintauth {
	return &MsgSetMintauth{
		Creator:       creator,
		NftSchemaCode: nftSchemaCode,
		AuthorizeTo:   authorizeTo,
	}
}

func (msg *MsgSetMintauth) Route() string {
	return RouterKey
}

func (msg *MsgSetMintauth) Type() string {
	return TypeMsgSetMintauth
}

func (msg *MsgSetMintauth) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMintauth) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMintauth) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
