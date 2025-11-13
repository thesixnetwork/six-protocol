package keeper

import (
	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"
)

var _ types.QueryServer = Keeper{}
