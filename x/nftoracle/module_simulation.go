package nftoracle

import (
	"math/rand"

	nftoraclesimulation "github.com/thesixnetwork/six-protocol/x/nftoracle/simulation"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = nftoraclesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateActionSigner = "op_weight_msg_action_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateActionSigner int = 100

	opWeightMsgUpdateActionSigner = "op_weight_msg_action_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateActionSigner int = 100

	opWeightMsgDeleteActionSigner = "op_weight_msg_action_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteActionSigner int = 100

	opWeightMsgCreateSyncActionSigner = "op_weight_msg_create_sync_action_signer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSyncActionSigner int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nftoracleGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ActionSignerList: []types.ActionSigner{
			{
				OwnerAddress: sample.AccAddress(),
				ActorAddress: "0",
			},
			{
				OwnerAddress: sample.AccAddress(),
				ActorAddress: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&nftoracleGenesis)
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

	var weightMsgCreateActionSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateActionSigner, &weightMsgCreateActionSigner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateActionSigner = defaultWeightMsgCreateActionSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateActionSigner,
		nftoraclesimulation.SimulateMsgCreateActionSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateActionSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateActionSigner, &weightMsgUpdateActionSigner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateActionSigner = defaultWeightMsgUpdateActionSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateActionSigner,
		nftoraclesimulation.SimulateMsgUpdateActionSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteActionSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteActionSigner, &weightMsgDeleteActionSigner, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteActionSigner = defaultWeightMsgDeleteActionSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteActionSigner,
		nftoraclesimulation.SimulateMsgDeleteActionSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateSyncActionSigner int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateSyncActionSigner, &weightMsgCreateSyncActionSigner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSyncActionSigner = defaultWeightMsgCreateSyncActionSigner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateSyncActionSigner,
		nftoraclesimulation.SimulateMsgCreateSyncActionSigner(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
