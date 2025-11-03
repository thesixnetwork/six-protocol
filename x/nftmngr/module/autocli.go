package nftmngr

import (
	modulev1 "github.com/thesixnetwork/six-protocol/api/sixprotocol/nftmngr"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
// func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
func (am AppModule) _() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "NFTSchemaAll",
					Use:       "list-nft-schema",
					Short:     "List all NFTSchema",
				},
				{
					RpcMethod:      "NFTSchema",
					Use:            "show-nft-schema [id]",
					Short:          "Shows a NFTSchema",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "code"}},
				},
				{
					RpcMethod: "NftDataAll",
					Use:       "list-nft-data",
					Short:     "List all NftData",
				},
				{
					RpcMethod:      "NftData",
					Use:            "show-nft-data [id]",
					Short:          "Shows a NftData",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftSchemaCode"}, {ProtoField: "tokenId"}},
				},
				{
					RpcMethod: "ActionByRefIdAll",
					Use:       "list-action-by-ref-id",
					Short:     "List all ActionByRefId",
				},
				{
					RpcMethod:      "ActionByRefId",
					Use:            "show-action-by-ref-id [id]",
					Short:          "Shows a ActionByRefId",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "refId"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateNFTSchema",
					Use:            "create-nft-schema [nft-schema-base-64]",
					Short:          "Send a createNFTSchema tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "nftSchemaBase64"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
