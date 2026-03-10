package tokenmngr

import (
	"math/rand"

	"github.com/thesixnetwork/six-protocol/v4/testutil/sample"
	tokenmngrsimulation "github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/simulation"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = tokenmngrsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateToken int = 100

	opWeightMsgUpdateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateToken int = 100

	opWeightMsgDeleteToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteToken int = 100

	opWeightMsgCreateMintperm = "op_weight_msg_mintperm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMintperm int = 100

	opWeightMsgUpdateMintperm = "op_weight_msg_mintperm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMintperm int = 100

	opWeightMsgDeleteMintperm = "op_weight_msg_mintperm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMintperm int = 100

	opWeightMsgCreateOptions = "op_weight_msg_options"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateOptions int = 100

	opWeightMsgUpdateOptions = "op_weight_msg_options"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOptions int = 100

	opWeightMsgDeleteOptions = "op_weight_msg_options"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteOptions int = 100

	opWeightMsgBurn = "op_weight_msg_burn"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurn int = 100

	opWeightMsgWrapToken = "op_weight_msg_wrap_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWrapToken int = 100

	opWeightMsgUnwrapToken = "op_weight_msg_unwrap_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnwrapToken int = 100

	opWeightMsgSendWrapToken = "op_weight_msg_send_wrap_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSendWrapToken int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenmngrGenesis := types.GenesisState{
		Params: types.DefaultParams(),
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

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateToken int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateToken, &weightMsgCreateToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateToken = defaultWeightMsgCreateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateToken,
		tokenmngrsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateToken int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateToken, &weightMsgUpdateToken, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateToken = defaultWeightMsgUpdateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateToken,
		tokenmngrsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteToken int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteToken, &weightMsgDeleteToken, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteToken = defaultWeightMsgDeleteToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteToken,
		tokenmngrsimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateMintperm int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMintperm, &weightMsgCreateMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMintperm = defaultWeightMsgCreateMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMintperm,
		tokenmngrsimulation.SimulateMsgCreateMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMintperm int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMintperm, &weightMsgUpdateMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMintperm = defaultWeightMsgUpdateMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMintperm,
		tokenmngrsimulation.SimulateMsgUpdateMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMintperm int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMintperm, &weightMsgDeleteMintperm, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMintperm = defaultWeightMsgDeleteMintperm
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMintperm,
		tokenmngrsimulation.SimulateMsgDeleteMintperm(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateOptions int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateOptions, &weightMsgCreateOptions, nil,
		func(_ *rand.Rand) {
			weightMsgCreateOptions = defaultWeightMsgCreateOptions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateOptions,
		tokenmngrsimulation.SimulateMsgCreateOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateOptions int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateOptions, &weightMsgUpdateOptions, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOptions = defaultWeightMsgUpdateOptions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOptions,
		tokenmngrsimulation.SimulateMsgUpdateOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteOptions int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteOptions, &weightMsgDeleteOptions, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteOptions = defaultWeightMsgDeleteOptions
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteOptions,
		tokenmngrsimulation.SimulateMsgDeleteOptions(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurn int
	simState.AppParams.GetOrGenerate(opWeightMsgBurn, &weightMsgBurn, nil,
		func(_ *rand.Rand) {
			weightMsgBurn = defaultWeightMsgBurn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurn,
		tokenmngrsimulation.SimulateMsgBurn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWrapToken int
	simState.AppParams.GetOrGenerate(opWeightMsgWrapToken, &weightMsgWrapToken, nil,
		func(_ *rand.Rand) {
			weightMsgWrapToken = defaultWeightMsgWrapToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWrapToken,
		tokenmngrsimulation.SimulateMsgWrapToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnwrapToken int
	simState.AppParams.GetOrGenerate(opWeightMsgUnwrapToken, &weightMsgUnwrapToken, nil,
		func(_ *rand.Rand) {
			weightMsgUnwrapToken = defaultWeightMsgUnwrapToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnwrapToken,
		tokenmngrsimulation.SimulateMsgUnwrapToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSendWrapToken int
	simState.AppParams.GetOrGenerate(opWeightMsgSendWrapToken, &weightMsgSendWrapToken, nil,
		func(_ *rand.Rand) {
			weightMsgSendWrapToken = defaultWeightMsgSendWrapToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSendWrapToken,
		tokenmngrsimulation.SimulateMsgSendWrapToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateToken,
			defaultWeightMsgCreateToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateToken,
			defaultWeightMsgUpdateToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteToken,
			defaultWeightMsgDeleteToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgDeleteToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateMintperm,
			defaultWeightMsgCreateMintperm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgCreateMintperm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateMintperm,
			defaultWeightMsgUpdateMintperm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgUpdateMintperm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteMintperm,
			defaultWeightMsgDeleteMintperm,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgDeleteMintperm(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateOptions,
			defaultWeightMsgCreateOptions,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgCreateOptions(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateOptions,
			defaultWeightMsgUpdateOptions,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgUpdateOptions(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteOptions,
			defaultWeightMsgDeleteOptions,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgDeleteOptions(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBurn,
			defaultWeightMsgBurn,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgBurn(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWrapToken,
			defaultWeightMsgWrapToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgWrapToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUnwrapToken,
			defaultWeightMsgUnwrapToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgUnwrapToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSendWrapToken,
			defaultWeightMsgSendWrapToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokenmngrsimulation.SimulateMsgSendWrapToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
