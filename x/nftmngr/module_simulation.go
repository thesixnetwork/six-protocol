package nftmngr

import (
	"math/rand"

	nftmngrsimulation "github.com/thesixnetwork/six-protocol/x/nftmngr/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nftmngrsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAddAttribute = "op_weight_msg_add_attribute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAttribute int = 100

	opWeightMsgAddTokenAttribute = "op_weight_msg_add_token_attribute"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddTokenAttribute int = 100

	opWeightMsgAddAction = "op_weight_msg_add_action"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAction int = 100

	opWeightMsgSetBaseUri = "op_weight_msg_set_base_uri"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetBaseUri int = 100
	opWeightMsgToggleAction        = "op_weight_msg_toggle_action"
	// TODO: Determine the simulation weight value
	defaultWeightMsgToggleAction int = 100

	_ = "op_weight_msg_set_schema_owner"
	// TODO: Determine the simulation weight value
	_ int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nftmngrGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nftmngrGenesis)
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

	var weightMsgAddAttribute int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddAttribute, &weightMsgAddAttribute, nil,
		func(_ *rand.Rand) {
			weightMsgAddAttribute = defaultWeightMsgAddAttribute
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAttribute,
		nftmngrsimulation.SimulateMsgAddAttribute(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddTokenAttribute int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddTokenAttribute, &weightMsgAddTokenAttribute, nil,
		func(_ *rand.Rand) {
			weightMsgAddTokenAttribute = defaultWeightMsgAddTokenAttribute
		},
	)

	var weightMsgAddAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddAction, &weightMsgAddAction, nil,
		func(_ *rand.Rand) {
			weightMsgAddAction = defaultWeightMsgAddAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAction,
		nftmngrsimulation.SimulateMsgAddAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetBaseUri int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSetBaseUri, &weightMsgSetBaseUri, nil,
		func(_ *rand.Rand) {
			weightMsgSetBaseUri = defaultWeightMsgSetBaseUri
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetBaseUri,
		nftmngrsimulation.SimulateMsgSetBaseUri(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	var weightMsgToggleAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgToggleAction, &weightMsgToggleAction, nil,
		func(_ *rand.Rand) {
			weightMsgToggleAction = defaultWeightMsgToggleAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgToggleAction,
		nftmngrsimulation.SimulateMsgToggleAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))
	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
