package keeper_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	keeperTestify "github.com/thesixnetwork/six-protocol/testutil/keeper"
	utils "github.com/thesixnetwork/six-protocol/testutil/utils"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	sim "github.com/thesixnetwork/six-protocol/x/nftmngr/simulation"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
)

func TestAction(t *testing.T) {
	// Read in schema and metadata from JSON files
	schemaJSON, err := os.ReadFile("../simulation/schema.json")
	require.NoError(t, err)
	schemaInput := types.NFTSchemaINPUT{}
	err = jsonpb.UnmarshalString(string(schemaJSON), &schemaInput)
	require.NoError(t, err)

	metaJSON, err := os.ReadFile("../simulation/meta.json")
	require.NoError(t, err)

	metaInput := types.NftData{}
	err = jsonpb.UnmarshalString(string(metaJSON), &metaInput)
	require.NoError(t, err)

	// Simulate create metadata
	schema_, data_, globalAttributeSchema_ := sim.SimulateCreateMetadata(schemaInput, metaInput)

	// ** START: NOTHING TO CARE ABOUT THIS PART **

	keeper_, ctx := keeperTestify.NftmngrKeeper(t)
	keeper_.SetNFTSchema(ctx, schema_)

	if _, found := keeper_.GetNFTSchema(ctx, schema_.Code); !found {
		fmt.Println("Schema not found")
	} else {
		require.True(t, found)
		require.NoError(t, err)
		// require.Equal(t, schema_, val)
	}

	// Create metadata object with initial attributes
	keeper_.SetNftData(ctx, data_)

	if val, found := keeper_.GetNftData(ctx, schema_.Code, "1"); !found {
		fmt.Println("Metadata not found")
	} else {
		require.True(t, found)
		require.NoError(t, err)
		require.Equal(t, data_, val)
	}
	// ** END: NOTHING TO CARE ABOUT THIS PART **

	// ** START: ACTUAL TEST **
	// Define the action to be processed
	selectAction := "check_in"

	// Create metadata object with initial attributes
	meta := types.NewMetadata(&schema_, &data_, types.AttributeOverriding_CHAIN, globalAttributeSchema_)

	// Process the action on the metadata
	actionParams_ := []*types.ActionParameter{}
	for _, action := range schema_.OnchainData.Actions {
		if action.Name == selectAction {
			keeper.ProcessAction(meta, action, actionParams_)
			break
		}
	}

	// Test 1: Check that the points attribute was updated correctly
	newPoints, err := meta.MustGetFloat("points")
	require.NoError(t, err)
	require.Equal(t, float64(50), newPoints)

	// Test 2: Check that the is_checked_in attribute was updated correctly
	isCheckIn, err := meta.MustGetBool("is_checked_in")
	require.NoError(t, err)
	require.True(t, isCheckIn)

	// Test 3: Check that the all_check_in attribute was updated correctly
	allCheckInNumber, err := meta.MustGetNumber("all_check_in")
	require.NoError(t, err)
	require.Equal(t, int64(1), allCheckInNumber)

	// fmt.Printf("CASE 1 SUCCESS WITH UPDATE VALUE \n newPoints: %v\n isCheckIn :%v\n allCheckInNumber: %v\n", newPoints, isCheckIn, allCheckInNumber)
	// // ** END: ACTUAL TEST **
}

func TestActionKeeper(t *testing.T) {
	var tokenId = "1"

	// Read in schema and metadata from JSON files
	schemaJSON, err := os.ReadFile("../simulation/nft_schema_test.json")
	require.NoError(t, err)
	schemaInput := types.NFTSchemaINPUT{}
	err = jsonpb.UnmarshalString(string(schemaJSON), &schemaInput)
	require.NoError(t, err)

	metaJSON, err := os.ReadFile("../simulation/nft_metadata_test.json")
	require.NoError(t, err)

	metaInput := types.NftData{}

	metaInput.TokenId = tokenId

	err = jsonpb.UnmarshalString(string(metaJSON), &metaInput)
	require.NoError(t, err)

	// ** START: NOTHING TO CARE ABOUT THIS PART **
	_, _, creator := utils.KeyTestPubAddr()

	testKeeper, ctx := keeperTestify.NftmngrKeeper(t)
	err = testKeeper.CreateNftSchemaKeeper(ctx, creator.String(), schemaInput)
	require.NoError(t, err)
	// testKeeper.SetNFTSchema(ctx, schema_)

	r := rand.New(rand.NewSource(99))

	_, found := testKeeper.GetNFTSchema(ctx, schemaInput.Code)
	if !found {
		fmt.Println("Schema not found")
	} else {
		require.True(t, found)
		require.NoError(t, err)
		// require.Equal(t, schema_, val)
	}

	err = testKeeper.CreateNewMetadataKeeper(ctx, creator.String(), schemaInput.Code, metaInput.TokenId, metaInput)
	require.NoError(t, err)

	_, found = testKeeper.GetNftData(ctx, schemaInput.Code, metaInput.TokenId)
	if !found {
		fmt.Println("Metadata not found")
	} else {
		require.True(t, found)
		require.NoError(t, err)
		// require.Equal(t, metaInput, data_)
	}

	// ** END: NOTHING TO CARE ABOUT THIS PART **

	// ** START: ACTUAL TEST **
	// Define the action to be processed
	selectAction := "use_service"

	// Process the action on the metadata
	actionParams_ := []*types.ActionParameter{}
	actionParams_ = append(actionParams_, &types.ActionParameter{
		Name:  "service_name",
		Value: "service_4",
	})

	intRand := r.Int()
	stringFromIntRand := strconv.FormatInt(int64(intRand), 10)

	changList, err := testKeeper.ActionByAdmin(ctx, creator.String(), schemaInput.Code, metaInput.TokenId, selectAction, stringFromIntRand, actionParams_)
	require.NoError(t, err)

	metaChangelist := []types.MetadataChange{}
	err = json.Unmarshal(changList, &metaChangelist)
	if err != nil {
		fmt.Printf("err: %v \n", err)
	}
	
	require.Equal(t, actionParams_[0].Value, metaChangelist[0].Key)
	require.Equal(t, "10", metaChangelist[0].PreviousValue)
	require.Equal(t, "9", metaChangelist[0].NewValue)
	// fmt.Printf("ChangeList: %+v \n", metaChangelist)
}
