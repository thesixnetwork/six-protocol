package keeper_test

import (
	fmt "fmt"
	"os"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/stretchr/testify/require"
	sim "github.com/thesixnetwork/six-protocol/x/nftmngr/simulation"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func TestNewMetadata(t *testing.T) {
	schemaJSON, err := os.ReadFile("../simulation/schema.json")
	if err != nil {
		fmt.Println(err)
	}

	schemaInput := types.NFTSchemaINPUT{}
	err = jsonpb.UnmarshalString(string(schemaJSON), &schemaInput)
	if err != nil {
		fmt.Println(err)
	}

	metaJSON, err := os.ReadFile("../simulation/meta.json")
	if err != nil {
		fmt.Println(err)
	}

	metaInput := types.NftData{}
	err = jsonpb.UnmarshalString(string(metaJSON), &metaInput)
	if err != nil {
		fmt.Println(err)
	}


	schema_, data_, globalAttributeSchema_ := sim.SimulateCreateMetadata(schemaInput, metaInput)
	meta := types.NewMetadata(&schema_, &data_, types.AttributeOverriding_CHAIN, globalAttributeSchema_)

	require.Greater(t, len(meta.MapAllKey), len(metaInput.OnchainAttributes))
}
