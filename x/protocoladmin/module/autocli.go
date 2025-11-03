package protocoladmin

import (
	modulev1 "github.com/thesixnetwork/six-protocol/api/sixprotocol/protocoladmin"

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
					RpcMethod: "AdminAll",
					Use:       "list-admin",
					Short:     "List all admin",
				},
				{
					RpcMethod:      "Admin",
					Use:            "show-admin [id]",
					Short:          "Shows a admin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "group"}, {ProtoField: "admin"}},
				},
				{
					RpcMethod: "GroupAll",
					Use:       "list-group",
					Short:     "List all group",
				},
				{
					RpcMethod:      "Group",
					Use:            "show-group [id]",
					Short:          "Shows a group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "ListAdminOfGroup",
					Use:            "list-admin-of-group [group]",
					Short:          "Query listAdminOfGroup",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "group"}},
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
					RpcMethod:      "CreateGroup",
					Use:            "create-group [name]",
					Short:          "Create a new group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "UpdateGroup",
					Use:            "update-group [name]",
					Short:          "Update group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "DeleteGroup",
					Use:            "delete-group [name]",
					Short:          "Delete group",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod:      "AddAdminToGroup",
					Use:            "add-admin-to-group [name] [address]",
					Short:          "Send a AddAdminToGroup tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "address"}},
				},
				{
					RpcMethod:      "RemoveAdminFromGroup",
					Use:            "remove-admin-from-group [name] [address]",
					Short:          "Send a RemoveAdminFromGroup tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "address"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
