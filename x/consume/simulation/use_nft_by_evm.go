package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/thesixnetwork/six-protocol/x/consume/keeper"
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

func SimulateMsgUseNftByEVM(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgUseNftByEVM{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the UseNftByEVM simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "UseNftByEVM simulation not implemented"), nil, nil
	}
}
