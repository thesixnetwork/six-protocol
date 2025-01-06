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

	opWeightMsgCreateVirtualAction = "op_weight_msg_virtual"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVirtualAction int = 100

	opWeightMsgUpdateVirtualAction = "op_weight_msg_virtual"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateVirtualAction int = 100

	opWeightMsgDeleteVirtualAction = "op_weight_msg_virtual"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVirtualAction int = 100

	opWeightMsgCreateVirtualSchema = "op_weight_msg_vir_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateVirtualSchema int = 100

	opWeightMsgDeleteVirtualSchema = "op_weight_msg_vir_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteVirtualSchema int = 100

	opWeightMsgVoteCreateVirtualSchema = "op_weight_msg_vote_create_virtual_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteCreateVirtualSchema int = 100

	opWeightMsgDisableVirtualSchemaProposal = "op_weight_msg_disable_virtual_schema_proposal"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDisableVirtualSchemaProposal int = 100

	opWeightMsgPerformVirtualAction = "op_weight_msg_perform_virtual_action"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPerformVirtualAction int = 100

	opWeightMsgEnableVirtualSchema = "op_weight_msg_enable_virtual_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEnableVirtualSchema int = 100

	opWeightMsgVoteDisableVirtualSchema = "op_weight_msg_vote_disable_virtual_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteDisableVirtualSchema int = 100

	opWeightMsgVoteEnableVirtualSchema = "op_weight_msg_vote_enable_virtual_schema"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteEnableVirtualSchema int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	nftmngrGenesis := types.GenesisState{
		Params:                  types.DefaultParams(),
		NFTSchemaList:           []types.NFTSchema{},
		NftDataList:             []types.NftData{},
		ActionByRefIdList:       []types.ActionByRefId{},
		OrganizationList:        []types.Organization{},
		NFTSchemaByContractList: []types.NFTSchemaByContract{},
		NftFeeConfig:            &types.NFTFeeConfig{},
		NFTFeeBalance:           &types.NFTFeeBalance{},
		MetadataCreatorList:     []types.MetadataCreator{},
		NftCollectionList:       []types.NftCollection{},
		ActionExecutorList:      []types.ActionExecutor{},
		SchemaAttributeList:     []types.SchemaAttribute{},
		ActionOfSchemaList:      []types.ActionOfSchema{},
		ExecutorOfSchemaList:    []types.ExecutorOfSchema{},
		VirtualActionList: []types.VirtualAction{
			{
				VirtualNftSchemaCode:   "0",
				Name:            "",
				Desc:            "",
				Disable:         false,
				When:            "",
				Then:            []string{},
				AllowedActioner: 0,
				Params:          []*types.ActionParams{},
			},
			{
				VirtualNftSchemaCode:   "1",
				Name:            "",
				Desc:            "",
				Disable:         false,
				When:            "",
				Then:            []string{},
				AllowedActioner: 0,
				Params:          []*types.ActionParams{},
			},
		},
		VirtualSchemaList: []types.VirtualSchema{
			{
				VirtualNftSchemaCode: "0",
			},
			{
				VirtualNftSchemaCode: "1",
			},
		},
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
	var weightMsgCreateVirtualAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVirtualAction, &weightMsgCreateVirtualAction, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVirtualAction = defaultWeightMsgCreateVirtualAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVirtualAction,
		nftmngrsimulation.SimulateMsgCreateVirtualAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateVirtualAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateVirtualAction, &weightMsgUpdateVirtualAction, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateVirtualAction = defaultWeightMsgUpdateVirtualAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateVirtualAction,
		nftmngrsimulation.SimulateMsgUpdateVirtualAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVirtualAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVirtualAction, &weightMsgDeleteVirtualAction, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVirtualAction = defaultWeightMsgDeleteVirtualAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVirtualAction,
		nftmngrsimulation.SimulateMsgDeleteVirtualAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateVirtualSchema, &weightMsgCreateVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgCreateVirtualSchema = defaultWeightMsgCreateVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateVirtualSchema,
		nftmngrsimulation.SimulateMsgCreateVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteVirtualSchema, &weightMsgDeleteVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteVirtualSchema = defaultWeightMsgDeleteVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteVirtualSchema,
		nftmngrsimulation.SimulateMsgDeleteVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteCreateVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteCreateVirtualSchema, &weightMsgVoteCreateVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgVoteCreateVirtualSchema = defaultWeightMsgVoteCreateVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteCreateVirtualSchema,
		nftmngrsimulation.SimulateMsgVoteCreateVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDisableVirtualSchemaProposal int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDisableVirtualSchemaProposal, &weightMsgDisableVirtualSchemaProposal, nil,
		func(_ *rand.Rand) {
			weightMsgDisableVirtualSchemaProposal = defaultWeightMsgDisableVirtualSchemaProposal
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDisableVirtualSchemaProposal,
		nftmngrsimulation.SimulateMsgDisableVirtualSchemaProposal(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPerformVirtualAction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgPerformVirtualAction, &weightMsgPerformVirtualAction, nil,
		func(_ *rand.Rand) {
			weightMsgPerformVirtualAction = defaultWeightMsgPerformVirtualAction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPerformVirtualAction,
		nftmngrsimulation.SimulateMsgPerformVirtualAction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEnableVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEnableVirtualSchema, &weightMsgEnableVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgEnableVirtualSchema = defaultWeightMsgEnableVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEnableVirtualSchema,
		nftmngrsimulation.SimulateMsgEnableVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteDisableVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteDisableVirtualSchema, &weightMsgVoteDisableVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgVoteDisableVirtualSchema = defaultWeightMsgVoteDisableVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteDisableVirtualSchema,
		nftmngrsimulation.SimulateMsgVoteDisableVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteEnableVirtualSchema int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteEnableVirtualSchema, &weightMsgVoteEnableVirtualSchema, nil,
		func(_ *rand.Rand) {
			weightMsgVoteEnableVirtualSchema = defaultWeightMsgVoteEnableVirtualSchema
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteEnableVirtualSchema,
		nftmngrsimulation.SimulateMsgVoteEnableVirtualSchema(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
