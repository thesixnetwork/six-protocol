package keeper

import (
	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

var _ types.QueryServer = Keeper{}
