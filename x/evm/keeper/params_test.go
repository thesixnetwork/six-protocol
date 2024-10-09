// Copyright (C) 2024 SIX Network
// This file is part of the modified EVM module from Ethermint, 
// and is licensed under the terms of the GNU Lesser General Public License v3
package keeper_test

import (
	"github.com/thesixnetwork/six-protocol/x/evm/types"
)

func (suite *KeeperTestSuite) TestParams() {
	params := suite.app.EVMKeeper.GetParams(suite.ctx)
	suite.Require().Equal(types.DefaultParams(), params)
	params.EvmDenom = "inj"
	suite.app.EVMKeeper.SetParams(suite.ctx, params)
	newParams := suite.app.EVMKeeper.GetParams(suite.ctx)
	suite.Require().Equal(newParams, params)
}
