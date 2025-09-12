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
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// StakingKeeper creates a testing keeper for the x/staking module.
func StakingKeeper(t testing.TB) (*stakingkeeper.Keeper, sdk.Context) {
	// Define store keys
	storeKey := storetypes.NewKVStoreKey(stakingtypes.StoreKey)
	authStoreKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	bankStoreKey := storetypes.NewKVStoreKey(banktypes.StoreKey)

	// Create a simple in-memory database
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())

	// Mount the necessary stores
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(authStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(bankStoreKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	// Create codec and interface registry
	registry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	stakingtypes.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	// Authority address
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Module account permissions
	maccPerms := map[string][]string{
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	}

	// Create AccountKeeper
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(authStoreKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr),
		sdk.Bech32PrefixAccAddr,
		authority.String(),
	)

	// Create BankKeeper
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		runtime.NewKVStoreService(bankStoreKey),
		accountKeeper,
		map[string]bool{}, // blocked addresses (empty for tests)
		authority.String(),
		log.NewNopLogger(),
	)

	// Create StakingKeeper
	stakingKeeper := stakingkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		accountKeeper,
		bankKeeper,
		authority.String(),
		addresscodec.NewBech32Codec(sdk.Bech32PrefixValAddr),  // validator address codec
		addresscodec.NewBech32Codec(sdk.Bech32PrefixConsAddr), // consensus address codec
	)

	// Create context
	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Set default staking params
	if err := stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams()); err != nil {
		panic(err)
	}

	// Set module accounts
	accountKeeper.GetModuleAccount(ctx, stakingtypes.BondedPoolName)
	accountKeeper.GetModuleAccount(ctx, stakingtypes.NotBondedPoolName)

	return stakingKeeper, ctx
}

// StakingKeeperWithDeps creates a testing keeper for the x/staking module and returns all dependencies.
func StakingKeeperWithDeps(t testing.TB) (*stakingkeeper.Keeper, *authkeeper.AccountKeeper, sdk.Context) {
	// Define store keys
	storeKey := storetypes.NewKVStoreKey(stakingtypes.StoreKey)
	authStoreKey := storetypes.NewKVStoreKey(authtypes.StoreKey)
	bankStoreKey := storetypes.NewKVStoreKey(banktypes.StoreKey)

	// Create a simple in-memory database
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())

	// Mount the necessary stores
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(authStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(bankStoreKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	// Create codec and interface registry
	registry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	banktypes.RegisterInterfaces(registry)
	stakingtypes.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)

	// Authority address
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	// Module account permissions
	maccPerms := map[string][]string{
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	}

	// Create AccountKeeper
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(authStoreKey),
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.Bech32PrefixAccAddr),
		sdk.Bech32PrefixAccAddr,
		authority.String(),
	)

	// Create BankKeeper
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		runtime.NewKVStoreService(bankStoreKey),
		accountKeeper,
		map[string]bool{}, // blocked addresses (empty for tests)
		authority.String(),
		log.NewNopLogger(),
	)

	// Create StakingKeeper
	stakingKeeper := stakingkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		accountKeeper,
		bankKeeper,
		authority.String(),
		addresscodec.NewBech32Codec(sdk.Bech32PrefixValAddr),  // validator address codec
		addresscodec.NewBech32Codec(sdk.Bech32PrefixConsAddr), // consensus address codec
	)

	// Create context
	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Set default staking params
	if err := stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams()); err != nil {
		panic(err)
	}

	// Set module accounts
	accountKeeper.GetModuleAccount(ctx, stakingtypes.BondedPoolName)
	accountKeeper.GetModuleAccount(ctx, stakingtypes.NotBondedPoolName)

	return stakingKeeper, &accountKeeper, ctx
}
