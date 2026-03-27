package keeper

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"

	nftadminkeeper "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/keeper"
	nftadmintypes "github.com/thesixnetwork/six-protocol/v4/x/nftadmin/types"
	nftoraclekeeper "github.com/thesixnetwork/six-protocol/v4/x/nftoracle/keeper"
	nftoracletypes "github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// CombinedKeepers holds both oracle and admin keepers sharing the same context/store.
// Use this in tests that require both keepers to interact (e.g. gasless ante handler tests).
type CombinedKeepers struct {
	OracleKeeper nftoraclekeeper.Keeper
	AdminKeeper  nftadminkeeper.Keeper
	Ctx          sdk.Context
}

// NewCombinedKeepers creates nftoracle and nftadmin keepers that share a single CommitMultiStore
// so that both keepers operate on the same underlying context.
func NewCombinedKeepers(t testing.TB) CombinedKeepers {
	t.Helper()

	oracleStoreKey := storetypes.NewKVStoreKey(nftoracletypes.StoreKey)
	adminStoreKey := storetypes.NewKVStoreKey(nftadmintypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(oracleStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(adminStoreKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName).String()

	adminKeeper := nftadminkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(adminStoreKey),
		log.NewNopLogger(),
		authority,
		nil,
	)

	oracleKeeper := nftoraclekeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(oracleStoreKey),
		log.NewNopLogger(),
		authority,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	require.NoError(t, adminKeeper.SetParams(ctx, nftadmintypes.DefaultParams()))
	require.NoError(t, oracleKeeper.SetParams(ctx, nftoracletypes.DefaultParams()))

	return CombinedKeepers{
		OracleKeeper: oracleKeeper,
		AdminKeeper:  adminKeeper,
		Ctx:          ctx,
	}
}
