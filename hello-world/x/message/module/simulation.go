package message

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"hello-world/testutil/sample"
	messagesimulation "hello-world/x/message/simulation"
	"hello-world/x/message/types"
)

// avoid unused import issue
var (
	_ = messagesimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateMessage = "op_weight_msg_create_message"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMessage int = 100

	opWeightMsgCreateSupplychain = "op_weight_msg_create_supplychain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSupplychain int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	messageGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&messageGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateMessage int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMessage, &weightMsgCreateMessage, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMessage = defaultWeightMsgCreateMessage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMessage,
		messagesimulation.SimulateMsgCreateMessage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSupplychain int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateSupplychain, &weightMsgCreateSupplychain, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSupplychain = defaultWeightMsgCreateSupplychain
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSupplychain,
		messagesimulation.SimulateMsgCreateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMessage,
			defaultWeightMsgCreateMessage,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				messagesimulation.SimulateMsgCreateMessage(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateSupplychain,
			defaultWeightMsgCreateSupplychain,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				messagesimulation.SimulateMsgCreateSupplychain(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
