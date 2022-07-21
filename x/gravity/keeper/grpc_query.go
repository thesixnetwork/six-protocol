package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/gravity/types"
)

var _ types.QueryServer = Keeper{}
