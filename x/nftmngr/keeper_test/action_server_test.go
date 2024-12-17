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
				t.Fatal()
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
			t.Fatal()
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

func TestCrossSchemaAction(t *testing.T){
	keeperTest, ctx := keepertest.NftmngrKeeper(t)

	_, schemaInputA := keepertest.InitSchema(t, "../../../resources/schemas/divineelite-nft-schema.json")	
	metadataA := keepertest.InitMetadata(t, "../../../resources/metadatas/divine_elite/nft-data_10_years.json")
	metadataA.TokenId = "1"
	err := keeperTest.CreateNftSchemaKeeper(ctx, schemaInputA.Owner, schemaInputA)
	require.NoError(t, err)

	schemaA, found := keeperTest.GetNFTSchema(ctx, schemaInputA.Code)
	require.True(t, found, "Schema A should exist")
	
	err = keeperTest.CreateNewMetadataKeeper(ctx, schemaA.Owner, schemaA.Code, metadataA.TokenId, metadataA)
	require.NoError(t, err)
	
	tokenDataA, found := keeperTest.GetNftData(ctx, schemaA.Code, metadataA.TokenId)
	require.True(t, found, "Metadata A should exist")
	

	mapExistingAttributesA := make(map[string]bool)
	for _, attribute := range tokenDataA.OnchainAttributes {
		mapExistingAttributesA[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schemaA.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributesA[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				t.Fatal("No default value")
			}
			// Add attribute to nftdata with default value
			tokenDataA.OnchainAttributes = append(tokenDataA.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	var map_converted_schema_attributesA []*types.NftAttributeValue
	global_attributesA := schemaA.OnchainData.NftAttributes
	attributeMapA := make(map[string]bool)

	for _, schema_attribute := range global_attributesA {
		// Check if the attribute has already been added
		if attributeMapA[schema_attribute.Name] {
			continue
		}

		nftAttributeValue, found := keeperTest.GetSchemaAttribute(ctx, schemaA.Code, schema_attribute.Name)

		if !found {
			t.Fatal("No default value")
		}

		// Add the attribute to the map
		attributeMapA[schema_attribute.Name] = true

		nftAttributeValue_ := keeper.ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue)
		map_converted_schema_attributesA = append(map_converted_schema_attributesA, nftAttributeValue_)
	}


	_, schemaInputB := keepertest.InitSchema(t, "../../../resources/schemas/membership-nft-schema.json")
	metadataB := keepertest.InitMetadata(t, "../../../resources/metadatas/membership/junior/nft-data_10_years.json")
	metadataB.TokenId = "1"
	err = keeperTest.CreateNftSchemaKeeper(ctx, schemaInputB.Owner, schemaInputB)
	require.NoError(t, err)

	schemaB, found := keeperTest.GetNFTSchema(ctx, schemaInputB.Code)
	require.True(t, found, "Schema B should exist")

	err = keeperTest.CreateNewMetadataKeeper(ctx, schemaB.Owner, schemaB.Code, metadataB.TokenId, metadataB)
	require.NoError(t, err)

	tokenDataB, found := keeperTest.GetNftData(ctx, schemaB.Code, metadataB.TokenId)
	require.True(t, found, "Metadata B should exist")


	mapExistingAttributesB := make(map[string]bool)
	for _, attribute := range tokenDataB.OnchainAttributes {
		mapExistingAttributesB[attribute.Name] = true
	}

	// Loop over schema.TokenAttributes to check if exists in nftdata
	for _, attribute := range schemaInputB.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributesB[attribute.Name]; !ok {
			if attribute.DefaultMintValue == nil {
				t.Fatal("No default value")
			}
			// Add attribute to nftdata with default value
			tokenDataB.OnchainAttributes = append(tokenDataB.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	var map_converted_schema_attributesB []*types.NftAttributeValue
	global_attributesB := schemaB.OnchainData.NftAttributes
	attributeMapB := make(map[string]bool)

	for _, schema_attribute := range global_attributesB {
		// Check if the attribute has already been added
		if attributeMapB[schema_attribute.Name] {
			continue
		}

		nftAttributeValue, found := keeperTest.GetSchemaAttribute(ctx, schemaB.Code, schema_attribute.Name)

		if !found {
			t.Fatal("No default value")
		}

		// Add the attribute to the map
		attributeMapB[schema_attribute.Name] = true

		nftAttributeValue_ := keeper.ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue)
		map_converted_schema_attributesB = append(map_converted_schema_attributesB, nftAttributeValue_)
	}
	
	schemaList := []*types.NFTSchema{&schemaA, &schemaB}
	tokenDataList := []*types.NftData{&tokenDataA, &tokenDataB}
	crossSchemaOveride := types.CrossSchemaAttributeOverriding{
		schemaA.Code: types.AttributeOverriding_CHAIN,
		schemaB.Code: types.AttributeOverriding_CHAIN,
	}

	schemaGlobalAttributes := types.CrossSchemaGlobalAttributes{
		schemaA.Code: map_converted_schema_attributesA,
		schemaB.Code: map_converted_schema_attributesB,
	}

	crossMetadata := types.NewCrossSchemaMetadata(schemaList, tokenDataList, crossSchemaOveride, schemaGlobalAttributes)
	attriNumber := crossMetadata.GetNumber(schemaA.Code, "service_3")
	require.Equal(t, int64(9999), attriNumber)

}