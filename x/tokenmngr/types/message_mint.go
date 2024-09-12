package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	v1types "github.com/thesixnetwork/six-protocol/x/tokenmngr/types/v1"
)

const TypeMsgMint = "mint"

var _ sdk.Msg = &MsgMint{}

func NewMsgMint(creator string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMint) Route() string {
	return RouterKey
}

func (msg *MsgMint) Type() string {
	return TypeMsgMint
}

func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// ConvertToLegacyMsgMint converts the new MsgMint to the v1.MsgMint format
func (msg *MsgMint) ConvertToLegacyMsgMint() *v1types.MsgMint {
	amount := msg.Amount.Amount.Uint64()
	return &v1types.MsgMint{
		Creator: msg.Creator,
		Amount:  amount,
		Token:   msg.Amount.Denom,
	}
}

func ConvertFromLegacyMsgMint(legacyMsg v1types.MsgMint) *MsgMint {
	return &MsgMint{
		Creator: legacyMsg.Creator,
		Amount:  sdk.NewCoin(legacyMsg.Token, sdk.NewIntFromUint64(legacyMsg.Amount)),
	}
}