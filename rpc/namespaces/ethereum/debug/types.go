package debug

import (
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/thesixnetwork/six-protocol/rpc/types"
)

type TraceCallConfig struct {
	evmtypes.TraceConfig
	StateOverrides *types.StateOverride
	BlockOverrides *types.BlockOverrides
}
