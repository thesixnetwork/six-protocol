package protocoladmin

import (
	"math/rand"

	"github.com/thesixnetwork/six-protocol/testutil/sample"
	protocoladminsimulation "github.com/thesixnetwork/six-protocol/x/protocoladmin/simulation"
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = protocoladminsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateGroup = "op_weight_msg_group"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGroup int = 100

	opWeightMsgUpdateGroup = "op_weight_msg_group"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateGroup int = 100

	opWeightMsgDeleteGroup = "op_weight_msg_group"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteGroup int = 100

	opWeightMsgAddAdminToGroup = "op_weight_msg_add_admin_to_group"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddAdminToGroup int = 100

	opWeightMsgRemoveAdminFromGroup = "op_weight_msg_remove_admin_from_group"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemoveAdminFromGroup int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	protocoladminGenesis := types.GenesisState{
		Params: types.DefaultParams(),
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

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateGroup, &weightMsgCreateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGroup = defaultWeightMsgCreateGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGroup,
		protocoladminsimulation.SimulateMsgCreateGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateGroup, &weightMsgUpdateGroup, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateGroup = defaultWeightMsgUpdateGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateGroup,
		protocoladminsimulation.SimulateMsgUpdateGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteGroup, &weightMsgDeleteGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteGroup = defaultWeightMsgDeleteGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteGroup,
		protocoladminsimulation.SimulateMsgDeleteGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddAdminToGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgAddAdminToGroup, &weightMsgAddAdminToGroup, nil,
		func(_ *rand.Rand) {
			weightMsgAddAdminToGroup = defaultWeightMsgAddAdminToGroup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddAdminToGroup,
		protocoladminsimulation.SimulateMsgAddAdminToGroup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemoveAdminFromGroup int
	simState.AppParams.GetOrGenerate(opWeightMsgRemoveAdminFromGroup, &weightMsgRemoveAdminFromGroup, nil,
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

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateGroup,
			defaultWeightMsgCreateGroup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				protocoladminsimulation.SimulateMsgCreateGroup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateGroup,
			defaultWeightMsgUpdateGroup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				protocoladminsimulation.SimulateMsgUpdateGroup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteGroup,
			defaultWeightMsgDeleteGroup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				protocoladminsimulation.SimulateMsgDeleteGroup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddAdminToGroup,
			defaultWeightMsgAddAdminToGroup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				protocoladminsimulation.SimulateMsgAddAdminToGroup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRemoveAdminFromGroup,
			defaultWeightMsgRemoveAdminFromGroup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				protocoladminsimulation.SimulateMsgRemoveAdminFromGroup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
