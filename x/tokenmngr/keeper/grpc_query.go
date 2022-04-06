package keeper

import (
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var _ types.QueryServer = Keeper{}
