package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSyncActionSigner = "create_sync_action_signer"

var _ sdk.Msg = &MsgCreateSyncActionSigner{}

func NewMsgCreateSyncActionSigner(creator string, chain string, actorAddress string, ownerAddress string, requiredConfirm uint64) *MsgCreateSyncActionSigner {
	return &MsgCreateSyncActionSigner{
		Creator:         creator,
		Chain:           chain,
		ActorAddress:    actorAddress,
		OwnerAddress:    ownerAddress,
		RequiredConfirm: requiredConfirm,
	}
}

func (msg *MsgCreateSyncActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgCreateSyncActionSigner) Type() string {
	return TypeMsgCreateSyncActionSigner
}

func (msg *MsgCreateSyncActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSyncActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSyncActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
