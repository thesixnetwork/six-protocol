package keeper

import (
	"github.com/thesixnetwork/six-protocol/v4/x/protocoladmin/types"
)

var _ types.QueryServer = Keeper{}
