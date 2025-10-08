package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteVirtualSchemaProposal = "vote_virtual_schema_proposal"

var _ sdk.Msg = &MsgVoteVirtualSchemaProposal{}

func NewMsgVoteVirtualSchemaProposal(creator, id, nftSchemaCode string, option RegistryStatus) *MsgVoteVirtualSchemaProposal {
	return &MsgVoteVirtualSchemaProposal{
		Creator:       creator,
		Id:            id,
		NftSchemaCode: nftSchemaCode,
		Option:        option,
	}
}

func (msg *MsgVoteVirtualSchemaProposal) Route() string {
	return RouterKey
}

func (msg *MsgVoteVirtualSchemaProposal) Type() string {
	return TypeMsgVoteVirtualSchemaProposal
}

func (msg *MsgVoteVirtualSchemaProposal) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteVirtualSchemaProposal) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteVirtualSchemaProposal) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
