package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/thesixnetwork/six-protocol/testutil/keeper"

	"github.com/stretchr/testify/require"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func setupSchemaAndMetadata(t *testing.T, k *keeper.Keeper, ctx sdk.Context, schemaPath string, metadataPath string, tokenId string) (*types.NFTSchema, *types.NftData, []*types.NftAttributeValue) {
	_, schemaInput := keepertest.InitSchema(t, schemaPath)
	metadata := keepertest.InitMetadata(t, metadataPath)
	metadata.TokenId = tokenId

	err := k.CreateNftSchemaKeeper(ctx, schemaInput.Owner, schemaInput)
	require.NoError(t, err)

	schema, found := k.GetNFTSchema(ctx, schemaInput.Code)
	require.True(t, found, "Schema should exist")

	err = k.CreateNewMetadataKeeper(ctx, schema.Owner, schema.Code, metadata.TokenId, metadata)
	require.NoError(t, err)

	tokenData, found := k.GetNftData(ctx, schema.Code, metadata.TokenId)
	require.True(t, found, "Metadata should exist")

	// Add missing attributes with default values
	mapExistingAttributes := make(map[string]bool)
	for _, attribute := range tokenData.OnchainAttributes {
		mapExistingAttributes[attribute.Name] = true
	}

	for _, attribute := range schema.OnchainData.TokenAttributes {
		if _, ok := mapExistingAttributes[attribute.Name]; !ok {
			require.NotNil(t, attribute.DefaultMintValue, "No default value")
			tokenData.OnchainAttributes = append(tokenData.OnchainAttributes, keeper.NewNFTAttributeValueFromDefaultValue(attribute.Name, attribute.DefaultMintValue))
		}
	}

	// Convert schema attributes
	var convertedSchemaAttributes []*types.NftAttributeValue
	attributeMap := make(map[string]bool)

	for _, schemaAttribute := range schema.OnchainData.NftAttributes {
		if attributeMap[schemaAttribute.Name] {
			continue
		}

		nftAttributeValue, found := k.GetSchemaAttribute(ctx, schema.Code, schemaAttribute.Name)
		require.True(t, found, "No default value")

		attributeMap[schemaAttribute.Name] = true
		convertedSchemaAttributes = append(convertedSchemaAttributes, keeper.ConverSchemaAttributeToNFTAttributeValue(&nftAttributeValue))
	}

	return &schema, &tokenData, convertedSchemaAttributes
}

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

type virtualActionTestCase struct {
    name           string
    action         types.Action
    actionParams   []*types.ActionParameter
    expectedValues map[string]struct {
        schemaCode string
        attribute  string
        value     int64
    }
}


func setupVirtualAction(t *testing.T, keeper *keeper.Keeper, ctx sdk.Context, virtualSchema *types.VirtualSchema, action types.Action) *types.VirtualAction {
    keeper.AddActionKeeper(ctx, "6x1myrlxmmasv6yq4axrxmdswj9kv5gc0ppx95rmq", virtualSchema.VirtualNftSchemaCode, action)
    
    virtualAction := types.VirtualAction{
        NftSchemaCode:   virtualSchema.VirtualNftSchemaCode,
        Name:            action.Name,
        Desc:            action.Desc,
        Disable:         action.Disable,
        When:            action.When,
        Then:            action.Then,
        AllowedActioner: action.AllowedActioner,
        Params:          action.Params,
    }
    keeper.SetVirtualAction(ctx, virtualAction)

    storedAction, found := keeper.GetVirtualAction(ctx, virtualSchema.VirtualNftSchemaCode, action.Name)
    require.True(t, found)
    require.Equal(t, action, *storedAction.ToAction())

    return &storedAction
}

func runVirtualActionTest(t *testing.T, k *keeper.Keeper, ctx sdk.Context, crossMetadata *types.CrossSchemaMetadata, virtualSchema *types.VirtualSchema, testCase virtualActionTestCase) {
    virtualAction := setupVirtualAction(t, k, ctx, virtualSchema, testCase.action)
    keeper.ProcessCrossSchemaAction(crossMetadata, virtualAction.ToAction(), testCase.actionParams)

    for _, expected := range testCase.expectedValues {
        attriNumber := crossMetadata.GetNumber(expected.schemaCode, expected.attribute)
        require.Equal(t, expected.value, attriNumber)
    }
}


func TestCrossSchemaAction(t *testing.T) {
	keeperTest, ctx := keepertest.NftmngrKeeper(t)

    // Setup Schema A
    schemaA, tokenDataA, convertedSchemaAttributesA := setupSchemaAndMetadata(t, 
        keeperTest, 
        ctx, 
        "../../../resources/schemas/divineelite-nft-schema.json",
        "../../../resources/metadatas/divine_elite/nft-data_10_years.json",
        "1",
	)

	// Setup Schema B
	schemaB, tokenDataB, convertedSchemaAttributesB := setupSchemaAndMetadata(t,
		keeperTest,
		ctx,
		"../../../resources/schemas/membership-nft-schema.json",
		"../../../resources/metadatas/membership/junior/nft-data_10_years.json",
		"1",
	)

	registrySchemaA := types.VirtualSchemaRegistry{
		NftSchemaCode:    schemaA.Code,
		SharedAttributes: []string{"service_3", "service_4"},
		Status:           types.RegistryStatus_ACCEPT,
	}

	registrySchemaB := types.VirtualSchemaRegistry{
		NftSchemaCode:    schemaB.Code,
		SharedAttributes: []string{"service_1", "service_2"},
		Status:           types.RegistryStatus_ACCEPT,
	}

	virtualSchema := types.VirtualSchema{
		VirtualNftSchemaCode: "divineXmembership",
		Registry: []*types.VirtualSchemaRegistry{
			&registrySchemaA, &registrySchemaB,
		},
		Enable:         false,
		ExpiredAtBlock: "0",
	}

	keeperTest.SetVirtualSchema(ctx, virtualSchema)

	schemaList := []*types.NFTSchema{schemaA, schemaB}
	tokenDataList := []*types.NftData{tokenDataA, tokenDataB}
	crossSchemaOveride := types.CrossSchemaAttributeOverriding{
		schemaA.Code: types.AttributeOverriding_CHAIN,
		schemaB.Code: types.AttributeOverriding_CHAIN,
	}

	schemaGlobalAttributes := types.CrossSchemaGlobalAttributes{
		schemaA.Code: convertedSchemaAttributesA,
		schemaB.Code: convertedSchemaAttributesB,
	}

	crossMetadata := types.NewCrossSchemaMetadata(schemaList, tokenDataList, crossSchemaOveride, schemaGlobalAttributes)
	testCases := []virtualActionTestCase{
        {
            name: "Bridge service 3 to 1",
            action: types.Action{
                Name:    "bridge_3_to_1",
                Desc:    "Bridge service 1 to service 4",
                Disable: false,
                When:    "true",
                Then: []string{
                    "ser3value = meta.GetNumber('sixprotocol.divine_elite','service_3')",
                    "ser1Value = meta.GetNumber('sixprotocol.membership','service_1')",
                    "toSetValue = ser3value + ser1Value",
                    "meta.SetNumber('sixprotocol.membership','service_1', toSetValue)",
                    "meta.SetNumber('sixprotocol.divine_elite','service_3', 0)",
                },
                AllowedActioner: 0,
                Params:          []*types.ActionParams{{}},
            },
            actionParams: []*types.ActionParameter{},
            expectedValues: map[string]struct {
                schemaCode string
                attribute  string
                value     int64
            }{
                "service_1": {schemaCode: schemaB.Code, attribute: "service_1", value: 9999},
                "service_3": {schemaCode: schemaA.Code, attribute: "service_3", value: 0},
            },
        },
        {
            name: "Bridge service 4 to 2",
            action: types.Action{
                Name:    "bridge_4_to_2",
                Desc:    "Bridge service 4 to service 2",
                Disable: false,
                When:    "meta.GetNumber('sixprotocol.divine_elite','service_4') >= params['amount'].GetNumber()",
                Then: []string{
                    "ser4value = meta.GetNumber('sixprotocol.divine_elite','service_4')",
                    "ser2Value = meta.GetNumber('sixprotocol.membership','service_2')",
                    "toSetValue = ser2Value + params['amount'].GetNumber()",
                    "meta.SetNumber('sixprotocol.membership','service_2', toSetValue)",
                    "meta.SetNumber('sixprotocol.divine_elite','service_4', ser4value - params['amount'].GetNumber())",
                },
                AllowedActioner: 0,
                Params: []*types.ActionParams{{
                    Name:         "amount",
                    DataType:     "number",
                    Desc:         "Service 4 Amount",
                    Required:     true,
                    DefaultValue: "0",
                }},
            },
            actionParams: []*types.ActionParameter{{
                Name:  "amount",
                Value: "10",
            }},
            expectedValues: map[string]struct {
                schemaCode string
                attribute  string
                value     int64
            }{
                "service_2": {schemaCode: schemaB.Code, attribute: "service_2", value: 20},
                "service_4": {schemaCode: schemaA.Code, attribute: "service_4", value: 0},
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            runVirtualActionTest(t, keeperTest, ctx, crossMetadata, &virtualSchema, tc)
        })
    }
}