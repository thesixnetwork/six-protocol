package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
)

var _ types.QueryServer = Keeper{}
