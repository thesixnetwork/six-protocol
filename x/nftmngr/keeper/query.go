package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

var _ types.QueryServer = Keeper{}
