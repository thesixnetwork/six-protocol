package types

import (
	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateToken{}

func NewMsgCreateToken(
	creator string,
	name string,
	maxSupply sdk.Coin,
	mintee string,
	denomMetaData string,
) *MsgCreateToken {
	return &MsgCreateToken{
		Creator:       creator,
		Name:          name,
		MaxSupply:     maxSupply,
		Mintee:        mintee,
		DenomMetaData: denomMetaData,
	}
}

func (msg *MsgCreateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateToken{}

func NewMsgUpdateToken(
	creator string,
	name string,
	maxSupply sdk.Coin,
) *MsgUpdateToken {
	return &MsgUpdateToken{
		Creator:   creator,
		Name:      name,
		MaxSupply: maxSupply,
	}
}

func (msg *MsgUpdateToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteToken{}

func NewMsgDeleteToken(
	creator string,
	name string,
) *MsgDeleteToken {
	return &MsgDeleteToken{
		Creator: creator,
		Name:    name,
	}
}

func (msg *MsgDeleteToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
