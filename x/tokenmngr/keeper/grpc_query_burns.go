package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func (k Keeper) Burns(c context.Context, req *types.QueryBurnsRequest) (*types.QueryBurnsResponse, error) {
	// Throw an error if request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	// Define a variable that will store a list of burns
	var burns []*types.Burn
	// Get context with the information about the environment
	ctx := sdk.UnwrapSDKContext(c)
	// Get the key-value module store using the store key (in our case store key is "chain")
	store := ctx.KVStore(k.storeKey)
	// Get the part of the store that keeps burns (using burn key, which is "Burn-value-")
	burnStore := prefix.NewStore(store, []byte(types.BurnKey))
	// Paginate the burns store based on PageRequest
	pageRes, err := query.Paginate(burnStore, req.Pagination, func(key []byte, value []byte) error {
		var burn types.Burn
		if err := k.cdc.Unmarshal(value, &burn); err != nil {
			return err
		}
		burns = append(burns, &burn)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	// Return a struct containing a list of burns and pagination info
	return &types.QueryBurnsResponse{Burn: burns, Pagination: pageRes}, nil
}
