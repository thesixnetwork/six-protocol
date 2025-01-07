package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// general
	cdc.RegisterConcrete(&MsgCreateNFTSchema{}, "nftmngr/CreateNFTSchema", nil)
	cdc.RegisterConcrete(&MsgCreateMetadata{}, "nftmngr/CreateMetadata", nil)
	cdc.RegisterConcrete(&MsgSetBaseUri{}, "nftmngr/SetBaseUri", nil)
	cdc.RegisterConcrete(&MsgChangeSchemaOwner{}, "nftmngr/ChangeSchemaOwner", nil)
	cdc.RegisterConcrete(&MsgSetFeeConfig{}, "nftmngr/SetFeeConfig", nil)
	cdc.RegisterConcrete(&MsgSetMintauth{}, "nftmngr/SetMintauth", nil)
	cdc.RegisterConcrete(&MsgChangeOrgOwner{}, "nftmngr/ChageOrgOwner", nil)
	cdc.RegisterConcrete(&MsgSetUriRetrievalMethod{}, "nftmngr/SetUriRetrievalMethod", nil)
	cdc.RegisterConcrete(&MsgSetOriginChain{}, "nftmngr/SetOriginChain", nil)
	cdc.RegisterConcrete(&MsgSetOriginContract{}, "nftmngr/SetOriginContract", nil)
	cdc.RegisterConcrete(&MsgSetAttributeOveriding{}, "nftmngr/SetAttributeOveriding", nil)
	cdc.RegisterConcrete(&MsgSetMetadataFormat{}, "nftmngr/SetMetadataFormat", nil)

	// attributes
	cdc.RegisterConcrete(&MsgAddAttribute{}, "nftmngr/AddAttribute", nil)
	cdc.RegisterConcrete(&MsgResyncAttributes{}, "nftmngr/ResyncAttributes", nil)
	cdc.RegisterConcrete(&MsgShowAttributes{}, "nftmngr/ShowAttributes", nil)
	cdc.RegisterConcrete(&MsgUpdateSchemaAttribute{}, "nftmngr/UpdateSchemaAttribute", nil)
	// action
	cdc.RegisterConcrete(&MsgPerformActionByAdmin{}, "nftmngr/PerformActionByAdmin", nil)
	cdc.RegisterConcrete(&MsgAddAction{}, "nftmngr/AddAction", nil)
	cdc.RegisterConcrete(&MsgToggleAction{}, "nftmngr/ToggleAction", nil)
	cdc.RegisterConcrete(&MsgUpdateAction{}, "nftmngr/UpdateAction", nil)
	// executor
	cdc.RegisterConcrete(&MsgCreateActionExecutor{}, "nftmngr/CreateActionExecutor", nil)
	cdc.RegisterConcrete(&MsgUpdateActionExecutor{}, "nftmngr/UpdateActionExecutor", nil)
	cdc.RegisterConcrete(&MsgDeleteActionExecutor{}, "nftmngr/DeleteActionExecutor", nil)
	// virtual schema
	// virtual action
	cdc.RegisterConcrete(&MsgPerformVirtualAction{}, "nftmngr/PerformVirtualAction", nil)
	cdc.RegisterConcrete(&MsgCreateVirtualAction{}, "nftmngr/CreateVirtualAction", nil)
	cdc.RegisterConcrete(&MsgUpdateVirtualAction{}, "nftmngr/UpdateVirtualAction", nil)
	cdc.RegisterConcrete(&MsgDeleteVirtualAction{}, "nftmngr/DeleteVirtualAction", nil)
	// proposal
	cdc.RegisterConcrete(&MsgProposalVirtualSchema{}, "nftmngr/ProposalVirtualSchema", nil)
	cdc.RegisterConcrete(&MsgVoteVirtualSchemaProposal{}, "nftmngr/VoteVirtualSchemaProposal", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// general
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetFeeConfig{},
		&MsgCreateMetadata{},
		&MsgSetBaseUri{},
		&MsgChangeSchemaOwner{},
		&MsgSetMintauth{},
		&MsgChangeOrgOwner{},
		&MsgSetUriRetrievalMethod{},
		&MsgSetOriginChain{},
		&MsgSetOriginContract{},
		&MsgSetAttributeOveriding{},
		&MsgSetMetadataFormat{},
	)
	// attributes
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddAttribute{},
		&MsgResyncAttributes{},
		&MsgShowAttributes{},
		&MsgUpdateSchemaAttribute{},
	)
	// action
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPerformActionByAdmin{},
		&MsgAddAction{},
		&MsgToggleAction{},
		&MsgUpdateAction{},
	)
	// executor
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateActionExecutor{},
		&MsgUpdateActionExecutor{},
		&MsgDeleteActionExecutor{},
	)
	// virtual schema
	// virtual action
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVirtualAction{},
		&MsgUpdateVirtualAction{},
		&MsgDeleteVirtualAction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPerformVirtualAction{},
	)
	// proposal
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposalVirtualSchema{},
		&MsgVoteVirtualSchemaProposal{},
	)
	// this line is used by starport scaffolding # 3
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
