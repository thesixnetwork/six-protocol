package keeper_test

import (
	"fmt"
	"testing"

	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestAction(t *testing.T) {
	_, schemainput := keepertest.InitSchema(t, "../simulation/schema.json")
	metadata := keepertest.InitMetadata(t, "../simulation/meta.json")
	keeperTest, ctx := keepertest.NftmngrKeeper(t)

	err := keeperTest.CreateNftSchemaKeeper(ctx, schemainput.Owner, schemainput)
	require.NoError(t, err)

	schema, found := keeperTest.GetNFTSchema(ctx, schemainput.Code)
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
