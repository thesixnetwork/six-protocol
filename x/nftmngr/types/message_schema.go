package types

import (
	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetUriRetrievalMethod = "set_uri_retrieval_method"

var _ sdk.Msg = &MsgSetUriRetrievalMethod{}

func NewMsgSetUriRetrievalMethod(creator string, schemaCode string, newMethod int32) *MsgSetUriRetrievalMethod {
	return &MsgSetUriRetrievalMethod{
		Creator:    creator,
		SchemaCode: schemaCode,
		NewMethod:  newMethod,
	}
}

func (msg *MsgSetUriRetrievalMethod) Route() string {
	return RouterKey
}

func (msg *MsgSetUriRetrievalMethod) Type() string {
	return TypeMsgSetUriRetrievalMethod
}

func (msg *MsgSetUriRetrievalMethod) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetUriRetrievalMethod) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetUriRetrievalMethod) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgCreateNFTSchema = "create_nft_schema"

var _ sdk.Msg = &MsgCreateNFTSchema{}

func NewMsgCreateNFTSchema(creator string, nftSchemaBase64 string) *MsgCreateNFTSchema {
	return &MsgCreateNFTSchema{
		Creator:         creator,
		NftSchemaBase64: nftSchemaBase64,
	}
}

func (msg *MsgCreateNFTSchema) Route() string {
	return RouterKey
}

func (msg *MsgCreateNFTSchema) Type() string {
	return TypeMsgCreateNFTSchema
}

func (msg *MsgCreateNFTSchema) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNFTSchema) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNFTSchema) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSetBaseUri = "set_base_uri"

var _ sdk.Msg = &MsgSetBaseUri{}

func NewMsgSetBaseUri(creator string, code string, newBaseUri string) *MsgSetBaseUri {
	return &MsgSetBaseUri{
		Creator:    creator,
		Code:       code,
		NewBaseUri: newBaseUri,
	}
}

func (msg *MsgSetBaseUri) Route() string {
	return RouterKey
}

func (msg *MsgSetBaseUri) Type() string {
	return TypeMsgSetBaseUri
}

func (msg *MsgSetBaseUri) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetBaseUri) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetBaseUri) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSetMetadataFormat = "set_metadata_format"

var _ sdk.Msg = &MsgSetMetadataFormat{}

func NewMsgSetMetadataFormat(creator string, schemaCode string, newFormat string) *MsgSetMetadataFormat {
	return &MsgSetMetadataFormat{
		Creator:    creator,
		SchemaCode: schemaCode,
		NewFormat:  newFormat,
	}
}

func (msg *MsgSetMetadataFormat) Route() string {
	return RouterKey
}

func (msg *MsgSetMetadataFormat) Type() string {
	return TypeMsgSetMetadataFormat
}

func (msg *MsgSetMetadataFormat) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetMetadataFormat) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMetadataFormat) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

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
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSetOriginChain = "set_origin_chain"

var _ sdk.Msg = &MsgSetOriginChain{}

func NewMsgSetOriginChain(creator string, schemaCode string, newOriginChain string) *MsgSetOriginChain {
	return &MsgSetOriginChain{
		Creator:        creator,
		SchemaCode:     schemaCode,
		NewOriginChain: newOriginChain,
	}
}

func (msg *MsgSetOriginChain) Route() string {
	return RouterKey
}

func (msg *MsgSetOriginChain) Type() string {
	return TypeMsgSetOriginChain
}

func (msg *MsgSetOriginChain) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetOriginChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetOriginChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

const TypeMsgSetOriginContract = "set_origin_contract"

var _ sdk.Msg = &MsgSetOriginContract{}

func NewMsgSetOriginContract(creator string, schemaCode string, newContractAddress string) *MsgSetOriginContract {
	return &MsgSetOriginContract{
		Creator:            creator,
		SchemaCode:         schemaCode,
		NewContractAddress: newContractAddress,
	}
}

func (msg *MsgSetOriginContract) Route() string {
	return RouterKey
}

func (msg *MsgSetOriginContract) Type() string {
	return TypeMsgSetOriginContract
}

func (msg *MsgSetOriginContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetOriginContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetOriginContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errormod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
