package keeper

import (
	"context"
	"fmt"
	"os"
	"testing"

	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	common "github.com/thesixnetwork/six-protocol/precompiles/common"
	tokenmngrkeeper "github.com/thesixnetwork/six-protocol/x/tokenmngr/keeper"
	tokenmngrtypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func TestMain(m *testing.M) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("6x", "6xpub")
	config.Seal()
	os.Exit(m.Run())
}

func StakingKeeper(t testing.TB) (*stakingkeeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	tkey := sdk.NewTransientStoreKey("transient_staking")
	accountKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	authTransientKey := sdk.NewTransientStoreKey("auth_transient")
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)

	db := tmdb.NewMemDB()
	stateStore, err := store.NewCommitMultiStore(db, log.NewNopLogger(), nil)
	require.NoError(t, err)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, nil)
	stateStore.MountStoreWithDB(accountKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(authTransientKey, storetypes.StoreTypeTransient, nil)
	stateStore.MountStoreWithDB(bankKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	authSubspace := typesparams.NewSubspace(
		cdc,
		authtypes.ModuleCdc.LegacyAmino,
		accountKey,
		authTransientKey,
		"AuthParams",
	)

	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		"mint":                         {authtypes.Minter},
	}

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		accountKey,
		authSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	blockedAddrs := map[string]bool{
		authtypes.FeeCollectorName:     true,
		stakingtypes.BondedPoolName:    true,
		stakingtypes.NotBondedPoolName: true,
	}

	paramsSubspace := typesparams.NewSubspace(
		cdc,
		stakingtypes.ModuleCdc.LegacyAmino,
		storeKey,
		tkey,
		"StakingParams",
	).WithKeyTable(stakingtypes.ParamKeyTable())

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		bankKey,
		accountKeeper,
		paramsSubspace,
		blockedAddrs,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	bondedPool := authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName, authtypes.Burner, authtypes.Staking)
	notBondedPool := authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName, authtypes.Burner, authtypes.Staking)

	accountKeeper.SetModuleAccount(ctx, bondedPool)
	accountKeeper.SetModuleAccount(ctx, notBondedPool)

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	k := stakingkeeper.NewKeeper(
		cdc,
		storeKey,
		accountKeeper,
		bankKeeper,
		paramsSubspace,
	)

	k.SetParams(ctx, stakingtypes.DefaultParams())

	return k, ctx
}

// NewAllKeepers initializes staking, bank, and tokenmngr dependencies.
func NewAllKeepers(t testing.TB) (sdk.Context, bankkeeper.BaseKeeper, *stakingkeeper.Keeper, *tokenmngrkeeper.Keeper, authkeeper.AccountKeeper) {
	// Keys
	storeKey := sdk.NewKVStoreKey(stakingtypes.StoreKey)
	tkey := sdk.NewTransientStoreKey("transient_staking")
	accountKey := sdk.NewKVStoreKey(authtypes.StoreKey)
	authTransientKey := sdk.NewTransientStoreKey("auth_transient")
	bankKey := sdk.NewKVStoreKey(banktypes.StoreKey)
	tokenmngrStoreKey := sdk.NewKVStoreKey(tokenmngrtypes.StoreKey)
	tokenmngrMemStoreKey := storetypes.NewMemoryStoreKey(tokenmngrtypes.MemStoreKey)
	upgradeStoreKey := sdk.NewKVStoreKey("upgrade")

	// DB
	db := tmdb.NewMemDB()
	stateStore, err := store.NewCommitMultiStore(db, log.NewNopLogger(), nil)
	require.NoError(t, err)

	// Mount all required stores
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkey, storetypes.StoreTypeTransient, nil)
	stateStore.MountStoreWithDB(accountKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(authTransientKey, storetypes.StoreTypeTransient, nil)
	stateStore.MountStoreWithDB(bankKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tokenmngrStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tokenmngrMemStoreKey, storetypes.StoreTypeMemory, nil)
	stateStore.MountStoreWithDB(upgradeStoreKey, storetypes.StoreTypeIAVL, db)

	require.NoError(t, stateStore.LoadLatestVersion())

	// Codec
	registry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry) // ✅ Required for staking validator pubkeys
	appCodec := codec.NewProtoCodec(registry)

	// Auth Subspace
	authSubspace := typesparams.NewSubspace(
		appCodec,
		authtypes.ModuleCdc.LegacyAmino,
		accountKey,
		authTransientKey,
		"AuthParams",
	)

	// Module account permissions
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:     nil,
		stakingtypes.ModuleName:        {authtypes.Burner, authtypes.Staking, authtypes.Minter},
		"mint":                         {authtypes.Minter},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		tokenmngrtypes.ModuleName:      {authtypes.Minter, authtypes.Burner}, // ✅ Required
	}

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	accountKeeper := authkeeper.NewAccountKeeper(
		appCodec,
		accountKey,
		authSubspace,
		authtypes.ProtoBaseAccount,
		maccPerms,
	)

	blockedAddrs := map[string]bool{
		authtypes.FeeCollectorName:     true,
		stakingtypes.BondedPoolName:    true,
		stakingtypes.NotBondedPoolName: true,
	}

	paramsSubspace := typesparams.NewSubspace(
		appCodec,
		stakingtypes.ModuleCdc.LegacyAmino,
		storeKey,
		tkey,
		"StakingParams",
	).WithKeyTable(stakingtypes.ParamKeyTable())

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	bankKeeper := bankkeeper.NewBaseKeeper(
		appCodec,
		bankKey,
		accountKeeper,
		paramsSubspace,
		blockedAddrs,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// 🔧 Set module accounts
	accountKeeper.SetModuleAccount(ctx, authtypes.NewEmptyModuleAccount(stakingtypes.BondedPoolName, authtypes.Burner, authtypes.Staking))
	accountKeeper.SetModuleAccount(ctx, authtypes.NewEmptyModuleAccount(stakingtypes.NotBondedPoolName, authtypes.Burner, authtypes.Staking))
	accountKeeper.SetModuleAccount(ctx, authtypes.NewEmptyModuleAccount("mint", authtypes.Minter))

	// 🔧 ✅ Add tokenmngr module account
	accountKeeper.SetModuleAccount(ctx, authtypes.NewEmptyModuleAccount(tokenmngrtypes.ModuleName, authtypes.Minter, authtypes.Burner))

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		storeKey,
		accountKeeper,
		bankKeeper,
		paramsSubspace,
	)
	stakingKeeper.SetParams(ctx, stakingtypes.DefaultParams())

	// Upgrade keeper
	skipUpgradeHeights := make(map[int64]bool)
	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	upgradeKeeper := upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		upgradeStoreKey,
		appCodec,
		"",
		nil,
	)

	// Tokenmngr setup
	capabilityKeeper := capabilitykeeper.NewKeeper(appCodec, tokenmngrStoreKey, tokenmngrMemStoreKey)

	tokenmngrParamsSubspace := typesparams.NewSubspace(
		appCodec,
		tokenmngrtypes.Amino,
		tokenmngrStoreKey,
		tokenmngrMemStoreKey,
		"TokenmngrParams",
	)

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	ibcKeeper := ibckeeper.NewKeeper(
		appCodec,
		tokenmngrStoreKey,
		tokenmngrParamsSubspace,
		stakingKeeper,
		upgradeKeeper,
		capabilityKeeper.ScopeToModule("TokenmngrIBCKeeper"),
	)

	// NOTE: This is an outdated constructor. It needs to be updated for Cosmos SDK v0.50+
	tokenmngrKeeper := tokenmngrkeeper.NewKeeper(
		appCodec,
		tokenmngrStoreKey,
		tokenmngrMemStoreKey,
		tokenmngrParamsSubspace,
		ibcKeeper.ChannelKeeper,
		&ibcKeeper.PortKeeper,
		capabilityKeeper.ScopeToModule("TokenmngrScopedKeeper"),
		bankKeeper,
		accountKeeper,
		nil,
		nil,
	)

	initialAsixBalance := sdk.NewCoins(sdk.NewCoin("asix", sdk.NewInt(1_000_000_000_000))) // 1M ASIX

	require.NoError(t, bankKeeper.MintCoins(ctx, tokenmngrtypes.ModuleName, initialAsixBalance))

	// Ensure it's spendable (safely move to spendable balance)
	require.NoError(t, bankKeeper.SendCoinsFromModuleToModule(ctx, tokenmngrtypes.ModuleName, tokenmngrtypes.ModuleName, initialAsixBalance))

	// Mint usix to burn
	usixAmount := sdk.NewCoins(sdk.NewCoin("usix", sdk.NewInt(1_000_000_000_000_000_000))) // 1 ASIX in atto
	require.NoError(t, bankKeeper.MintCoins(ctx, tokenmngrtypes.ModuleName, usixAmount))
	require.NoError(t, bankKeeper.SendCoinsFromModuleToModule(ctx, tokenmngrtypes.ModuleName, tokenmngrtypes.ModuleName, usixAmount))

	// Mint asix to send back
	asixAmount := sdk.NewCoins(sdk.NewCoin("asix", sdk.NewInt(1_000_000_000_000))) // matching amount in micro
	require.NoError(t, bankKeeper.MintCoins(ctx, tokenmngrtypes.ModuleName, asixAmount))

	return ctx, bankKeeper, stakingKeeper, tokenmngrKeeper, accountKeeper
}

// NewStakingQuerier returns a concrete StakingQuerier using the keeper
func NewStakingQuerier(k *stakingkeeper.Keeper) common.StakingQuerier {
	return &stakingQuerier{keeper: k}
}

// internal type that implements common.StakingQuerier
type stakingQuerier struct {
	keeper *stakingkeeper.Keeper
}

func (q *stakingQuerier) Validator(ctx sdk.Context, valAddr sdk.ValAddress) (stakingtypes.Validator, error) {
	validator, found := q.keeper.GetValidator(ctx, valAddr)
	if !found {
		return stakingtypes.Validator{}, fmt.Errorf("validator %s not found", valAddr.String())
	}
	return validator, nil
}

func (q *stakingQuerier) Delegation(c context.Context, req *stakingtypes.QueryDelegationRequest) (*stakingtypes.QueryDelegationResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	delAddr, err := sdk.AccAddressFromBech32(req.DelegatorAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid delegator address: %w", err)
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, fmt.Errorf("invalid validator address: %w", err)
	}

	delegation, found := q.keeper.GetDelegation(ctx, delAddr, valAddr)
	if !found {
		return nil, fmt.Errorf("delegation not found for delegator %s and validator %s", req.DelegatorAddr, req.ValidatorAddr)
	}

	// Type assert to stakingtypes.Delegation if needed, or construct DelegationResponse manually
	resp := stakingtypes.NewDelegationResp(
		delegation.GetDelegatorAddr(),
		delegation.GetValidatorAddr(),
		delegation.GetShares(),
		sdk.NewCoin("token", sdk.ZeroInt()), // placeholder for balance
	)

	return &stakingtypes.QueryDelegationResponse{
		DelegationResponse: &resp,
	}, nil
}
