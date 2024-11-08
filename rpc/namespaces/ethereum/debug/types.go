package debug

import (
	"github.com/thesixnetwork/six-protocol/rpc/types"
	evmtypes "github.com/thesixnetwork/six-protocol/x/evm/types"
)

type TraceCallConfig struct {
	evmtypes.TraceConfig
	StateOverrides *types.StateOverride
	BlockOverrides *types.BlockOverrides
}
