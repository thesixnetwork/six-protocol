package keeper_test

import (
	//"encoding/json"
	"fmt"
	//"math/rand"
	"os"
	//"strconv"
	"testing"

	keeperTestify "github.com/thesixnetwork/six-protocol/testutil/keeper"
	// utils "github.com/thesixnetwork/six-protocol/testutil/utils"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	sim "github.com/thesixnetwork/six-protocol/x/nftmngr/simulation"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"github.com/gogo/protobuf/jsonpb"
	//"github.com/stretchr/testify/assert"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("6x", "6xpub")
	config.Seal()
}

func initSchema(t *testing.T, schemaJSONFilePath string) (types.NFTSchema, types.NFTSchemaINPUT) {
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

	keeper, ctx := keeperTestify.NftmngrKeeper(t)
	err = keeper.CreateNftSchemaKeeper(ctx, "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq", schemaInput)
	if err != nil {
		t.Fatal(err)
	}

	return schema, schemaInput
}

func initMetadata(t *testing.T, metadataJSONFilePath string) types.NftData {
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

	keeper, ctx := keeperTestify.NftmngrKeeper(t)
	keeper.SetNftData(ctx, metaInput)

	return metaInput
}

func TestCreateMetadata(t *testing.T) {
	_, _ = initSchema(t, "../../../resources/nft-schema.json")
	meta := initMetadata(t, "../../../resources/nft-data.json")

	keeper, ctx := keeperTestify.NftmngrKeeper(t)
	keeper.SetNftData(ctx, meta)

	_, found := keeper.GetNftData(ctx, meta.NftSchemaCode, meta.TokenId)
	if !found {
		fmt.Println("Metadata not found")
	} else {
		require.True(t, found)
	}
}

func TestCreateSchema(t *testing.T) {
	_, schemaInput := initSchema(t, "../simulation/schema.json")

	keeperTest, ctx := keeperTestify.NftmngrKeeper(t)
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
}

func TestAction(t *testing.T) {
	_, schemainput := initSchema(t, "../simulation/schema.json")
	metadata := initMetadata(t, "../simulation/meta.json")
	keeperTest, ctx := keeperTestify.NftmngrKeeper(t)

  err := keeperTest.CreateNftSchemaKeeper(ctx, schemainput.Owner, schemainput)
  require.NoError(t, err)

	schema, found := keeperTest.GetNFTSchema(ctx, schemainput.Code); 
  if !found {
		fmt.Println("Schema not found")
	} else {
		require.True(t, found)
	}

	err = keeperTest.CreateNewMetadataKeeper(ctx, schema.Owner, schema.Code, metadata.TokenId, metadata)
  require.NoError(t, err)

	tokenData, found := keeperTest.GetNftData(ctx, schema.Code, metadata.TokenId)
	if !found {
		fmt.Println("Metadata not found")
	} else {
		require.True(t, found)
	}

	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range tokenData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
        t.Fatal("No default value")
			}
			// Add attribute to nftdata with default value
			tokenData.OnchainAttributes = append(tokenData.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	var map_converted_schema_attributes []*types.NftAttributeValue
	global_attributes := schema.OnchainData.NftAttributes
	attributeMap := make(map[string]bool)

	for _, schema_attribute := range global_attributes {
		// Check if the attribute has already been added
		if attributeMap[schema_attribute.Name] {
			continue
		}

		nftAttributeValue, found := keeperTest.GetSchemaAttribute(ctx, schema.Code, schema_attribute.Name)

		if !found {
      t.Fatal("No default value")
		}

		// Add the attribute to the map
		attributeMap[schema_attribute.Name] = true

		nftAttributeValue_ := keeper.ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue)
		map_converted_schema_attributes = append(map_converted_schema_attributes, nftAttributeValue_)
	}

	// Create metadata object with initial attributes
	nftdata := types.NewMetadata(&schema, &tokenData, types.AttributeOverriding_CHAIN, map_converted_schema_attributes)

	selectAction := "check_in"
	// Process the action on the metadata
	actionParams_ := []*types.ActionParameter{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == selectAction {
			keeper.ProcessAction(nftdata, action, actionParams_)
			break
		}
	}

	// Test 1: Check that the points attribute was updated correctly
	newPoints, err := nftdata.MustGetFloat("points")
	require.NoError(t, err)
	require.Equal(t, float64(50), newPoints)

	// Test 2: Check that the is_checked_in attribute was updated correctly
	isCheckIn, err := nftdata.MustGetBool("is_checked_in")
	require.NoError(t, err)
	require.True(t, isCheckIn)

	// Test 3: Check that the all_check_in attribute was updated correctly
	allCheckInNumber, err := nftdata.MustGetNumber("all_check_in")
	require.NoError(t, err)
	require.Equal(t, int64(1), allCheckInNumber)
}
