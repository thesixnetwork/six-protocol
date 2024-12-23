package keeper

import (
	"os"
	"testing"

	"github.com/gogo/protobuf/jsonpb"

	sim "github.com/thesixnetwork/six-protocol/x/nftmngr/simulation"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("6x", "6xpub")
	config.Seal()
}

func NftmngrKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"NftmngrParams",
	)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		nil,
		nil,
		nil,
		nil,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}

func InitSchema(t *testing.T, schemaJSONFilePath string) (types.NFTSchema, types.NFTSchemaINPUT) {
	// init schema
	schemaJSON, err := os.ReadFile(schemaJSONFilePath)
	if err != nil {
		panic(err)
	}

	schemaInput := types.NFTSchemaINPUT{}
	err = jsonpb.UnmarshalString(string(schemaJSON), &schemaInput)
	if err != nil {
		panic(err)
	}

	schema := sim.GenNFTSchemaFromInput(schemaInput)

	keeper, ctx := NftmngrKeeper(t)
	err = keeper.CreateNftSchemaKeeper(ctx, "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq", schemaInput)
	if err != nil {
		t.Fatal(err)
	}

	return schema, schemaInput
}

func InitMetadata(t *testing.T, metadataJSONFilePath string) types.NftData {
	// init metadata
	metaJSON, err := os.ReadFile(metadataJSONFilePath)
	if err != nil {
		panic(err)
	}

	metaInput := types.NftData{}
	err = jsonpb.UnmarshalString(string(metaJSON), &metaInput)
	if err != nil {
		panic(err)
	}

	keeper, ctx := NftmngrKeeper(t)
	keeper.SetNftData(ctx, metaInput)

	return metaInput
}
