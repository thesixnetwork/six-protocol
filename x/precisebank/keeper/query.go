package keeper

import (
	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

var _ types.QueryServer = Keeper{}
