package keeper

import (
	"github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"
)

var _ types.QueryServer = Keeper{}
