package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/sixgate/types"
)

var _ types.QueryServer = Keeper{}
