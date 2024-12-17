package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/testutil/nullify"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNFTSchema(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.NFTSchema {
	items := make([]types.NFTSchema, n)
	for i := range items {
		items[i].Code = strconv.Itoa(i)

		keeper.SetNFTSchema(ctx, items[i])
	}
	return items
}

func TestNFTSchemaGet(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNFTSchemaRemove(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNFTSchema(ctx,
			item.Code,
		)
		_, found := keeper.GetNFTSchema(ctx,
			item.Code,
		)
		require.False(t, found)
	}
}

func TestNFTSchemaGetAll(t *testing.T) {
	keeper, ctx := keepertest.NftmngrKeeper(t)
	items := createNNFTSchema(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNFTSchema(ctx)),
	)
}

func TestCreateSchema(t *testing.T) {
	_, schemaInput := keepertest.InitSchema(t, "../simulation/schema.json")

	keeperTest, ctx := keepertest.NftmngrKeeper(t)
	_, err := keeper.ValidateNFTSchema(&schemaInput)
	if err != nil {
		t.Fatal(err)
	}

	// if mint_authorization is empty then set system to default
	if len(schemaInput.MintAuthorization) == 0 || schemaInput.MintAuthorization != types.KeyMintPermissionOnlySystem && schemaInput.MintAuthorization != types.KeyMintPermissionAll {
		schemaInput.MintAuthorization = types.KeyMintPermissionOnlySystem
	}

	// Check if the schemaInput already exists
	_, found := keeperTest.GetNFTSchema(ctx, schemaInput.Code)
	if found {
		t.Fatal(err)
	}

	keeper.MergeAllAttributesAndAlterOrderIndex(schemaInput.OriginData.OriginAttributes, schemaInput.OnchainData.NftAttributes, schemaInput.OnchainData.TokenAttributes)

	// parse schemaInput to NFTSchema
	schema := types.NFTSchema{
		Code:        schemaInput.Code,
		Name:        schemaInput.Name,
		Owner:       schemaInput.Owner,
		Description: schemaInput.Description,
		OriginData:  schemaInput.OriginData,
		OnchainData: &types.OnChainData{
			TokenAttributes: schemaInput.OnchainData.TokenAttributes,
			NftAttributes:   schemaInput.OnchainData.NftAttributes,
			Actions:         schemaInput.OnchainData.Actions,
			Status:          schemaInput.OnchainData.Status,
		},
		IsVerified:        schemaInput.IsVerified,
		MintAuthorization: schemaInput.MintAuthorization,
	}

	for _, schemaDefaultMintAttribute := range schemaInput.OnchainData.NftAttributes {
		// parse DefaultMintValue to SchemaAttributeValue
		schmaAttributeValue, err := keeper.ConvertDefaultMintValueToSchemaAttributeValue(schemaDefaultMintAttribute.DefaultMintValue)
		if err != nil {
			t.Fatal(err)
		}

		keeperTest.SetSchemaAttribute(ctx, types.SchemaAttribute{
			NftSchemaCode: schemaInput.Code,
			Name:          schemaDefaultMintAttribute.Name,
			DataType:      schemaDefaultMintAttribute.DataType,
			CurrentValue:  schmaAttributeValue,
			Creator:       "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq",
		})
	}

	keeperTest.SetNFTSchema(ctx, schema)


	schemaFromKeeper, found := keeperTest.GetNFTSchema(ctx, schema.Code)
	if !found {
		t.Fatal("Schema not found")
	} else {
		require.Equal(t, schema, schemaFromKeeper)
	}
}
