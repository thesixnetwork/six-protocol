// Copyright (C) 2024 SIX Network
// This file is part of the modified FeeMarket module from Ethermint (https://github.com/evmos/ethermint),
// and is licensed under the terms of the GNU Lesser General Public License v3
package keeper_test

import (
	"github.com/thesixnetwork/six-protocol/x/feemarket/types"
)

func (suite *KeeperTestSuite) TestSetGetParams() {
	params := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
	suite.Require().Equal(types.DefaultParams(), params)
	params.ElasticityMultiplier = 3
	suite.app.FeeMarketKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.FeeMarketKeeper.GetParams(suite.ctx)
	suite.Require().Equal(newParams, params)
}
