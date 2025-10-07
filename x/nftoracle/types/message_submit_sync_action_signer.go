package types

import (
	errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitSyncActionSigner = "submit_sync_action_signer"

var _ sdk.Msg = &MsgSubmitSyncActionSigner{}

func NewMsgSubmitSyncActionSigner(creator string, chain string, syncId uint64, actorAddress string, ownerAddress string, expireEpoch string) *MsgSubmitSyncActionSigner {
	return &MsgSubmitSyncActionSigner{
		Creator:      creator,
		SyncId:       syncId,
		Chain:        chain,
		ActorAddress: actorAddress,
		OwnerAddress: ownerAddress,
		ExpireEpoch:  expireEpoch,
	}
}

func (msg *MsgSubmitSyncActionSigner) Route() string {
	return RouterKey
}

func (msg *MsgSubmitSyncActionSigner) Type() string {
	return TypeMsgSubmitSyncActionSigner
}

func (msg *MsgSubmitSyncActionSigner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitSyncActionSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitSyncActionSigner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
