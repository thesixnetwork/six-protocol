package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/consume/types"
)

var _ types.QueryServer = Keeper{}
