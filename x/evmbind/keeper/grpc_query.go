package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/evmbind/types"
)

var _ types.QueryServer = Keeper{}
