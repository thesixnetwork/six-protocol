package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/nftadmin/types"
)

var _ types.QueryServer = Keeper{}
