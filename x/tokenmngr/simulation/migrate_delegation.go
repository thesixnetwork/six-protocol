package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/keeper"
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func SimulateMsgMigrateDelegation(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMigrateDelegation{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MigrateDelegation simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "MigrateDelegation simulation not implemented"), nil, nil
	}
}
