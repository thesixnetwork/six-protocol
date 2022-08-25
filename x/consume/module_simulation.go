package consume

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/thesixnetwork/six-protocol/testutil/sample"
	consumesimulation "github.com/thesixnetwork/six-protocol/x/consume/simulation"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = consumesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgUseNft = "op_weight_msg_use_nft"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUseNft int = 100

	opWeightMsgUseNftByEVM = "op_weight_msg_use_nft_by_evm"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUseNftByEVM int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	consumeGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&consumeGenesis)
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

	var weightMsgUseNft int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUseNft, &weightMsgUseNft, nil,
		func(_ *rand.Rand) {
			weightMsgUseNft = defaultWeightMsgUseNft
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUseNft,
		consumesimulation.SimulateMsgUseNft(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUseNftByEVM int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUseNftByEVM, &weightMsgUseNftByEVM, nil,
		func(_ *rand.Rand) {
			weightMsgUseNftByEVM = defaultWeightMsgUseNftByEVM
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUseNftByEVM,
		consumesimulation.SimulateMsgUseNftByEVM(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
