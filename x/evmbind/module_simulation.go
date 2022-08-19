package evmbind

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	evmbindsimulation "github.com/thesixnetwork/six-protocol/x/evmbind/simulation"
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = evmbindsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateBinding = "op_weight_msg_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateBinding int = 100

	opWeightMsgUpdateBinding = "op_weight_msg_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBinding int = 100

	opWeightMsgDeleteBinding = "op_weight_msg_binding"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteBinding int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	evmbindGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		BindingList: []types.Binding{
			{
				Creator:    sample.AccAddress(),
				EthAddress: "0",
			},
			{
				Creator:    sample.AccAddress(),
				EthAddress: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&evmbindGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateBinding, &weightMsgCreateBinding, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBinding = defaultWeightMsgCreateBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBinding,
		evmbindsimulation.SimulateMsgCreateBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBinding, &weightMsgUpdateBinding, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBinding = defaultWeightMsgUpdateBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBinding,
		evmbindsimulation.SimulateMsgUpdateBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteBinding int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteBinding, &weightMsgDeleteBinding, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBinding = defaultWeightMsgDeleteBinding
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBinding,
		evmbindsimulation.SimulateMsgDeleteBinding(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
