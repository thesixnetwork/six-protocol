package simulation

import (
	"math/rand"

	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddAdminToGroup(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddAdminToGroup{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddAdminToGroup simulation
		_, foundGroup := k.GetGroup(ctx, msg.Creator)
		if !foundGroup {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "unable to find group "), nil, nil
		}

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AddAdminToGroup simulation not implemented"), nil, nil
	}
}
