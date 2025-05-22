package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

var _ types.QueryServer = Keeper{}
