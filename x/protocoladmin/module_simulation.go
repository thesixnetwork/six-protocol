package protocoladmin

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
	protocoladminsimulation "github.com/thesixnetwork/six-protocol/x/protocoladmin/simulation"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = protocoladminsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGroup = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGroup int = 100

	opWeightMsgUpdateGroup = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGroup int = 100

	opWeightMsgDeleteGroup = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGroup int = 100

	opWeightMsgAddAdminToGroup = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAdminToGroup int = 100

	opWeightMsgRemoveAdminFromGroup = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveAdminFromGroup int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	protocoladminGenesis := types.GenesisState{
		GroupList: []types.Group{
			{
				Owner: sample.AccAddress(),
				Name:  "0",
			},
			{
				Owner: sample.AccAddress(),
				Name:  "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&protocoladminGenesis)
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

	var weightMsgCreateGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGroup, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = defaultWeightMsgCreateGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGroup,
		protocoladminsimulation.SimulateMsgCreateGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateGroup, &weightMsgUpdateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroup = defaultWeightMsgUpdateGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGroup,
		protocoladminsimulation.SimulateMsgUpdateGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteGroup, &weightMsgDeleteGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGroup = defaultWeightMsgDeleteGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGroup,
		protocoladminsimulation.SimulateMsgDeleteGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddAdminToGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddAdminToGroup, &weightMsgAddAdminToGroup, nil,
		func(_ *rand.Rand) {
			weightMsgAddAdminToGroup = defaultWeightMsgAddAdminToGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAdminToGroup,
		protocoladminsimulation.SimulateMsgAddAdminToGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveAdminFromGroup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemoveAdminFromGroup, &weightMsgRemoveAdminFromGroup, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveAdminFromGroup = defaultWeightMsgRemoveAdminFromGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemoveAdminFromGroup,
		protocoladminsimulation.SimulateMsgRemoveAdminFromGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
