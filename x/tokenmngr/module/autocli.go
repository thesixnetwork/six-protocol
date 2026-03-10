package tokenmngr

import (
	modulev1 "github.com/thesixnetwork/six-protocol/v4/api/sixprotocol/tokenmngr"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
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
					RpcMethod: "TokenAll",
					Use:       "list-token",
					Short:     "List all token",
				},
				{
					RpcMethod:      "Token",
					Use:            "show-token [id]",
					Short:          "Shows a token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod: "MintpermAll",
					Use:       "list-mintperm",
					Short:     "List all mintperm",
				},
				{
					RpcMethod:      "Mintperm",
					Use:            "show-mintperm [id]",
					Short:          "Shows a mintperm",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token"}, {ProtoField: "address"}},
				},
				{
					RpcMethod: "TokenBurnAll",
					Use:       "list-token-burn",
					Short:     "List all tokenBurn",
				},
				{
					RpcMethod:      "TokenBurn",
					Use:            "show-token-burn [id]",
					Short:          "Shows a tokenBurn",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "Options",
					Use:       "show-options",
					Short:     "show options",
				},
				{
					RpcMethod:      "ListPrecompile",
					Use:            "precompiles",
					Short:          "Query Precompile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
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
					RpcMethod:      "CreateToken",
					Use:            "create-token [name] [max-supply] [mintee] [denom-metadata]",
					Short:          "Create a new token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "maxSupply"}, {ProtoField: "mintee"}, {ProtoField: "denomMetaData"}},
				},
				{
					RpcMethod:      "UpdateToken",
					Use:            "update-token [name] [max-supply]",
					Short:          "Update token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "maxSupply"}},
				},
				{
					RpcMethod:      "DeleteToken",
					Use:            "delete-token [name]",
					Short:          "Delete token",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "CreateMintperm",
					Use:            "create-mintperm [token] [address]",
					Short:          "Create a new mintperm",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "UpdateMintperm",
					Use:            "update-mintperm [token] [address]",
					Short:          "Update mintperm",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "DeleteMintperm",
					Use:            "delete-mintperm [token] [address]",
					Short:          "Delete mintperm",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "token"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "CreateOptions",
					Use:            "create-options [defaultMintee]",
					Short:          "Create options",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "defaultMintee"}},
				},
				{
					RpcMethod:      "UpdateOptions",
					Use:            "update-options [defaultMintee]",
					Short:          "Update options",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "defaultMintee"}},
				},
				{
					RpcMethod: "DeleteOptions",
					Use:       "delete-options",
					Short:     "Delete options",
				},
				{
					RpcMethod:      "Burn",
					Use:            "burn [amount]",
					Short:          "Send a burn tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod:      "Mint",
					Use:            "mint [amount]",
					Short:          "Send a mint tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod:      "WrapToken",
					Use:            "wrap-token [amount] [receiver]",
					Short:          "Send a wrapToken tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "UnwrapToken",
					Use:            "unwrap-token [amount] [receiver]",
					Short:          "Send a unwrapToken tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "receiver"}},
				},
				{
					RpcMethod:      "SendWrapToken",
					Use:            "send-wrap-token [eth-address] [amount]",
					Short:          "Send a sendWrapToken tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethAddress"}, {ProtoField: "amount"}},
				},
				{
					RpcMethod:      "MigrateDelegation",
					Use:            "migrate-delegation [eth-address]",
					Short:          "Migrate Delegation to evm address tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "ethAddress"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
