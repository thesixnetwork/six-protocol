package keeper

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	sixerc20keeper "github.com/thesixnetwork/six-protocol/v4/x/erc20/keeper"
	sixerc20types "github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// Erc20Keeper creates a minimal erc20 Keeper backed by an in-memory store.
// Dependency keepers (AccountKeeper, BankKeeper, EVMKeeper, StakingKeeper) are
// nil/zero-value — sufficient for unit tests covering params and token-pair CRUD,
// which do not invoke those dependencies.
func Erc20Keeper(t testing.TB) (sixerc20keeper.Keeper, sdk.Context) {
	t.Helper()

	storeKey := storetypes.NewKVStoreKey(sixerc20types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	k := sixerc20keeper.NewKeeper(
		storeKey,
		cdc,
		authority,
		nil,                  // accountKeeper — not needed for unit tests
		nil,                  // bankKeeper    — not needed for unit tests
		nil,                  // evmKeeper     — not needed for unit tests
		nil,                  // stakingKeeper — not needed for unit tests
		authzkeeper.Keeper{}, // authzKeeper   — zero-value struct
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	return k, ctx
}
