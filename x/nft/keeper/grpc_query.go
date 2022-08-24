package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nft/types"
)

var _ types.QueryServer = Keeper{}
