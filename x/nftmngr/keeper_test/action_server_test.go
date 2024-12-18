package keeper_test

import (
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

type actionTestCase struct {
	name           string
	schemaPath     string
	metadataPath   string
	tokenId        string
	selectAction   string
	expectedValues map[string]interface{}
}

func runActionTest(t *testing.T, k *keeper.Keeper, ctx sdk.Context, testCase actionTestCase) {
	schema, tokenData, convertedSchemaAttributes := setupSchemaAndMetadata(t,
		k,
		ctx,
		testCase.schemaPath,
		testCase.metadataPath,
		testCase.tokenId,
	)

	// Create metadata object with initial attributes
	nftdata := types.NewMetadata(schema, tokenData, types.AttributeOverriding_CHAIN, convertedSchemaAttributes)

	// Process the action on the metadata
	actionParams_ := []*types.ActionParameter{}
	for _, action := range schema.OnchainData.Actions {
		if action.Name == testCase.selectAction {
			keeper.ProcessAction(nftdata, action, actionParams_)
			break
		}
	}

	for key, expectedValue := range testCase.expectedValues {
		switch v := expectedValue.(type) {
		case float64:
			value, err := nftdata.MustGetFloat(key)
			require.NoError(t, err)
			require.Equal(t, v, value)
		case bool:
			value, err := nftdata.MustGetBool(key)
			require.NoError(t, err)
			require.Equal(t, v, value)
		case int64:
			value, err := nftdata.MustGetNumber(key)
			require.NoError(t, err)
			require.Equal(t, v, value)
		}
	}
}

func TestAction(t *testing.T) {
	keeperTest, ctx := keepertest.NftmngrKeeper(t)

	testCases := []actionTestCase{
		{
			name:         "Check in action",
			schemaPath:   "../simulation/schema.json",
			metadataPath: "../simulation/meta.json",
			tokenId:      "1",
			selectAction: "check_in",
			expectedValues: map[string]interface{}{
				"points":        float64(50),
				"is_checked_in": true,
				"all_check_in":  int64(1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runActionTest(t, keeperTest, ctx, tc)
		})
	}
}

type virtualActionTestCase struct {
	name           string
	action         types.Action
	actionParams   []*types.ActionParameter
	expectedValues map[string]struct {
		schemaCode string
		attribute  string
		value      interface{}
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
		switch v := expected.value.(type) {
		case int64:
			attriNumber := crossMetadata.GetNumber(expected.schemaCode, expected.attribute)
			require.Equal(t, v, attriNumber)
		case float64:
			attriFloat := crossMetadata.GetFloat(expected.schemaCode, expected.attribute)
			require.Equal(t, v, attriFloat)
		case bool:
			attriBool := crossMetadata.GetBoolean(expected.schemaCode, expected.attribute)
			require.Equal(t, v, attriBool)
		default:
			t.Fatalf("unsupported type: %T", v)
		}
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
				value      interface{}
			}{
				"service_1": {schemaCode: schemaB.Code, attribute: "service_1", value: int64(9999)},
				"service_3": {schemaCode: schemaA.Code, attribute: "service_3", value: int64(0)},
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
				value      interface{}
			}{
				"service_2": {schemaCode: schemaB.Code, attribute: "service_2", value: int64(20)},
				"service_4": {schemaCode: schemaA.Code, attribute: "service_4", value: int64(0)},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			runVirtualActionTest(t, keeperTest, ctx, crossMetadata, &virtualSchema, tc)
		})
	}
}

