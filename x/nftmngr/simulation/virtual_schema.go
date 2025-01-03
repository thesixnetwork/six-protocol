package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func SimulateMsgCreateVirtualSchema(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateVirtualSchemaProposal{
			Creator:              simAccount.Address.String(),
			VirtualNftSchemaCode: strconv.Itoa(i),
		}

		_, found := k.GetVirtualSchema(ctx, msg.VirtualNftSchemaCode)
		if found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "VirtualSchema already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteVirtualSchema(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount       = simtypes.Account{}
			virSchema        = types.VirtualSchema{}
			msg              = &types.MsgDeleteVirtualSchema{}
			allVirtualSchema = k.GetAllVirtualSchema(ctx)
			found            = false
		)
		for _, obj := range allVirtualSchema {
			simAccount, found = FindAccount(accs, obj.VirtualNftSchemaCode)
			if found {
				virSchema = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "virSchema creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()

		msg.VirtualNftSchemaCode = virSchema.VirtualNftSchemaCode

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}


func SimulateMsgDisableVirtualSchemaProposal(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgDisableVirtualSchemaProposal{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the DisableVirtualSchemaProposal simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "DisableVirtualSchemaProposal simulation not implemented"), nil, nil
	}
}

func SimulateMsgEnableVirtualSchema(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEnableVirtualSchemaProposal{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the EnableVirtualSchema simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "EnableVirtualSchema simulation not implemented"), nil, nil
	}
}
