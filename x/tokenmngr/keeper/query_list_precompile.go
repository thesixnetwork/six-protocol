package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/tokenmngr/types"
)

func (k Keeper) ListPrecompile(goCtx context.Context, req *types.QueryListPrecompileRequest) (*types.QueryListPrecompileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	map_precompiles := k.evmKeeper.GetPrecompiles()

	var precompiles []string

	for _, v := range map_precompiles {
		precompiles = append(precompiles, v.Address().Hex())
	}

	return &types.QueryListPrecompileResponse{Precompiles: precompiles}, nil
}
