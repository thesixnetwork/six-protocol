package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/v4/x/precisebank/types"
)

// Remainder implements types.QueryServer.
func (k Keeper) Remainder(goCtx context.Context, req *types.QueryRemainderRequest) (*types.QueryRemainderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	remainder := k.GetRemainderAmount(goCtx)

	return &types.QueryRemainderResponse{
		Remainder: remainder.String(),
	}, nil
}

// FractionalBalance implements types.QueryServer.
func (k Keeper) FractionalBalance(goCtx context.Context, req *types.QueryFractionalBalanceRequest) (*types.QueryFractionalBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address: "+err.Error())
	}

	fractionalBalance := k.GetFractionalBalance(goCtx, addr)

	return &types.QueryFractionalBalanceResponse{
		FractionalBalance: fractionalBalance.String(),
	}, nil
}

// FractionalBalances implements types.QueryServer.
func (k Keeper) FractionalBalances(goCtx context.Context, req *types.QueryFractionalBalancesRequest) (*types.QueryFractionalBalancesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	balanceStore := prefix.NewStore(store, types.FractionalBalancePrefix)

	var entries []types.FractionalBalanceEntry

	pageRes, err := query.Paginate(balanceStore, req.Pagination, func(key []byte, value []byte) error {
		addr := sdk.AccAddress(key)

		var amount sdkmath.Int
		if err := amount.Unmarshal(value); err != nil {
			return err
		}

		entries = append(entries, types.FractionalBalanceEntry{
			Address:           addr.String(),
			FractionalBalance: amount.String(),
		})
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFractionalBalancesResponse{
		FractionalBalances: entries,
		Pagination:         pageRes,
	}, nil
}
