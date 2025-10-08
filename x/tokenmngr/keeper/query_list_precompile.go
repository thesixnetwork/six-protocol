package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPrecompile(goCtx context.Context, req *types.QueryListPrecompileRequest) (*types.QueryListPrecompileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// params := k.GetParams(ctx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryListPrecompileResponse{}, nil
}
