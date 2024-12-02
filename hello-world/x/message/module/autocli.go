package message

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "hello-world/api/helloworld/message"
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
					RpcMethod:      "CreateMessage",
					Use:            "create-message [text]",
					Short:          "Send a create-message tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "text"}},
				},
				{
					RpcMethod:      "CreateSupplychain",
					Use:            "create-supplychain [item-code] [quantity] [item-desc]",
					Short:          "Send a create-supplychain tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "itemCode"}, {ProtoField: "quantity"}, {ProtoField: "itemDesc"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
