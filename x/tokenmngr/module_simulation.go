package tokenmngr

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
	tokenmngrsimulation "github.com/thesixnetwork/six-protocol/x/tokenmngr/simulation"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = tokenmngrsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateToken = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateToken int = 100

	opWeightMsgUpdateToken = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateToken int = 100

	opWeightMsgDeleteToken = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteToken int = 100

	opWeightMsgCreateMintperm = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMintperm int = 100

	opWeightMsgUpdateMintperm = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMintperm int = 100

	opWeightMsgDeleteMintperm = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMintperm int = 100

	opWeightMsgMint = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMint int = 100

	opWeightMsgCreateOptions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOptions int = 100

	opWeightMsgUpdateOptions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOptions int = 100

	opWeightMsgDeleteOptions = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOptions int = 100

	opWeightMsgBurn = "op_weight_msg_burn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurn int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenmngrGenesis := types.GenesisState{
		TokenList: []types.Token{
			{
				Creator: sample.AccAddress(),
				Name:    "0",
			},
			{
				Creator: sample.AccAddress(),
				Name:    "1",
			},
		},
		MintpermList: []types.Mintperm{
			{
				Creator: sample.AccAddress(),
				Token:   "0",
				Address: "0",
			},
			{
				Creator: sample.AccAddress(),
				Token:   "1",
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenmngrGenesis)
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

	var weightMsgCreateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateToken, &weightMsgCreateToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateToken = defaultWeightMsgCreateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateToken,
		tokenmngrsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateToken, &weightMsgUpdateToken, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateToken = defaultWeightMsgUpdateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateToken,
		tokenmngrsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteToken, &weightMsgDeleteToken, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteToken = defaultWeightMsgDeleteToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteToken,
		tokenmngrsimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMintperm int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMintperm, &weightMsgCreateMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMintperm = defaultWeightMsgCreateMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMintperm,
		tokenmngrsimulation.SimulateMsgCreateMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMintperm int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMintperm, &weightMsgUpdateMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMintperm = defaultWeightMsgUpdateMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMintperm,
		tokenmngrsimulation.SimulateMsgUpdateMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMintperm int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMintperm, &weightMsgDeleteMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMintperm = defaultWeightMsgDeleteMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMintperm,
		tokenmngrsimulation.SimulateMsgDeleteMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = defaultWeightMsgMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMint,
		tokenmngrsimulation.SimulateMsgMint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOptions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateOptions, &weightMsgCreateOptions, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOptions = defaultWeightMsgCreateOptions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOptions,
		tokenmngrsimulation.SimulateMsgCreateOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOptions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOptions, &weightMsgUpdateOptions, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOptions = defaultWeightMsgUpdateOptions
		},
	)
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgUpdateOptions,
	// 	tokenmngrsimulation.SimulateMsgUpdateOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	var weightMsgDeleteOptions int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteOptions, &weightMsgDeleteOptions, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOptions = defaultWeightMsgDeleteOptions
		},
	)
	// operations = append(operations, simulation.NewWeightedOperation(
	// 	weightMsgDeleteOptions,
	// 	tokenmngrsimulation.SimulateMsgDeleteOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	// ))

	var weightMsgBurn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurn, &weightMsgBurn, nil,
		func(_ *rand.Rand) {
			weightMsgBurn = defaultWeightMsgBurn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurn,
		tokenmngrsimulation.SimulateMsgBurn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
