package nftmngr_test

import (
	// "encoding/json"

	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/suite"

	prenftmgr "github.com/thesixnetwork/six-protocol/precompiles/nftmngr"
	testkeeper "github.com/thesixnetwork/six-protocol/testutil/keeper"
	nftmngrkeeper "github.com/thesixnetwork/six-protocol/x/nftmngr/keeper"
	nftmngrtype "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

type NftmngrPrecompileTestSuite struct {
	suite.Suite
	ctx           sdk.Context
	nftmngrKeeper *nftmngrkeeper.Keeper
	bankKeeper    bankkeeper.Keeper
	cdc           codec.Codec
	nftprecompile *prenftmgr.PrecompileExecutor
}

func TestNftmngrPrecompile(t *testing.T) {
	suite.Run(t, new(NftmngrPrecompileTestSuite))
}

func (suite *NftmngrPrecompileTestSuite) SetupTest() {
	// Initialize codec
	registry := codectypes.NewInterfaceRegistry()
	suite.cdc = codec.NewProtoCodec(registry)

	// Setup keeper and context
	keeper, ctx := testkeeper.NftmngrKeeper(suite.T())
	suite.nftmngrKeeper = keeper
	suite.ctx = ctx

	// Create precompile
	// precompile, err := prenftmgr.NewPrecompile(suite.nftmngrKeeper, suite.bankKeeper)
	precompile, err := prenftmgr.NewExecutor(suite.nftmngrKeeper, suite.bankKeeper)
	suite.Require().NoError(err)
	suite.nftprecompile = precompile
}

// Add utility method tests
func (suite *NftmngrPrecompileTestSuite) TestActionParameterParsing() {
	testParams := []*nftmngrtype.ActionParameter{
		{
			Name:  "service_name",
			Value: "binchotan",
		},
		{
			Name:  "amount",
			Value: "1",
		},
	}

	result, err := suite.nftprecompile.ParametersFromJSONString(`[{"name":"service_name","value":"binchotan"},{"name":"amount","value":"1"}]`)
	suite.Require().NoError(err)
	suite.Require().Len(result, 2)

	// Basic value checks
	suite.Require().Equal(testParams[0].Name, result[0].Name)
	suite.Require().Equal(testParams[0].Value, result[0].Value)
	suite.Require().Equal(testParams[1].Name, result[1].Name)
	suite.Require().Equal(testParams[1].Value, result[1].Value)
}
