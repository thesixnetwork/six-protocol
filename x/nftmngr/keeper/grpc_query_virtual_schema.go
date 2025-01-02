package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
)

func (k Keeper) VirtualSchemaProposalAll(c context.Context, req *types.QueryAllVirtualSchemaProposalRequest) (*types.QueryAllVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virtualSchemaProposals []types.VirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(virtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var virtualSchemaProposal types.VirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &virtualSchemaProposal); err != nil {
			return err
		}

		virtualSchemaProposals = append(virtualSchemaProposals, virtualSchemaProposal)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualSchemaProposalResponse{VirtualSchemaProposal: virtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) VirtualSchemaProposal(c context.Context, req *types.QueryGetVirtualSchemaProposalRequest) (*types.QueryGetVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualSchemaProposalResponse{VirtualSchemaProposal: val}, nil
}

func (k Keeper) VirtualSchemaAll(c context.Context, req *types.QueryAllVirtualSchemaRequest) (*types.QueryAllVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var virSchemas []types.VirtualSchema
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	virSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.VirtualSchemaKeyPrefix))

	pageRes, err := query.Paginate(virSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var virSchema types.VirtualSchema
		if err := k.cdc.Unmarshal(value, &virSchema); err != nil {
			return err
		}

		virSchemas = append(virSchemas, virSchema)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVirtualSchemaResponse{VirtualSchema: virSchemas, Pagination: pageRes}, nil
}

func (k Keeper) VirtualSchema(c context.Context, req *types.QueryGetVirtualSchemaRequest) (*types.QueryGetVirtualSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetVirtualSchema(
		ctx,
		req.NftSchemaCode,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetVirtualSchemaResponse{VirtualSchema: val}, nil
}

func (k Keeper) DisableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllDisableVirtualSchemaProposalRequest) (*types.QueryAllDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var disableVirtualSchemas []types.DisableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	disableVirtualSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.DisableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(disableVirtualSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var disableVirtualSchema types.DisableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &disableVirtualSchema); err != nil {
			return err
		}

		disableVirtualSchemas = append(disableVirtualSchemas, disableVirtualSchema)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDisableVirtualSchemaProposalResponse{DisableVirtualSchemaProposal: disableVirtualSchemas, Pagination: pageRes}, nil
}

func (k Keeper) DisableVirtualSchemaProposal(c context.Context, req *types.QueryGetDisableVirtualSchemaProposalRequest) (*types.QueryGetDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetDisableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetDisableVirtualSchemaProposalResponse{DisableVirtualSchemaProposal: val}, nil
}

func (k Keeper) ActiveDisableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllActiveDisableVirtualSchemaProposalRequest) (*types.QueryAllActiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeDisableVirtualSchemaProposals []types.ActiveDisableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeDisableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveDisableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(activeDisableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var activeDisableVirtualSchemaProposal types.ActiveDisableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &activeDisableVirtualSchemaProposal); err != nil {
			return err
		}

		activeDisableVirtualSchemaProposals = append(activeDisableVirtualSchemaProposals, activeDisableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveDisableVirtualSchemaProposalResponse{ActiveDisableVirtualSchemaProposal: activeDisableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) ActiveDisableVirtualSchemaProposal(c context.Context, req *types.QueryGetActiveDisableVirtualSchemaProposalRequest) (*types.QueryGetActiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveDisableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActiveDisableVirtualSchemaProposalResponse{ActiveDisableVirtualSchemaProposal: val}, nil
}

func (k Keeper) InactiveEnableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllInactiveEnableVirtualSchemaProposalRequest) (*types.QueryAllInactiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inactiveEnableVirtualSchemaProposals []types.InactiveEnableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inactiveEnableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.InactiveEnableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(inactiveEnableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var inactiveEnableVirtualSchemaProposal types.InactiveEnableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &inactiveEnableVirtualSchemaProposal); err != nil {
			return err
		}

		inactiveEnableVirtualSchemaProposals = append(inactiveEnableVirtualSchemaProposals, inactiveEnableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInactiveEnableVirtualSchemaProposalResponse{InactiveEnableVirtualSchemaProposal: inactiveEnableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) InactiveEnableVirtualSchemaProposal(c context.Context, req *types.QueryGetInactiveEnableVirtualSchemaProposalRequest) (*types.QueryGetInactiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInactiveEnableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInactiveEnableVirtualSchemaProposalResponse{InactiveEnableVirtualSchemaProposal: val}, nil
}

func (k Keeper) InactiveDisableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllInactiveDisableVirtualSchemaProposalRequest) (*types.QueryAllInactiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var inactiveDisableVirtualSchemaProposals []types.InactiveDisableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	inactiveDisableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.InactiveDisableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(inactiveDisableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var inactiveDisableVirtualSchemaProposal types.InactiveDisableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &inactiveDisableVirtualSchemaProposal); err != nil {
			return err
		}

		inactiveDisableVirtualSchemaProposals = append(inactiveDisableVirtualSchemaProposals, inactiveDisableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: inactiveDisableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) InactiveDisableVirtualSchemaProposal(c context.Context, req *types.QueryGetInactiveDisableVirtualSchemaProposalRequest) (*types.QueryGetInactiveDisableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetInactiveDisableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetInactiveDisableVirtualSchemaProposalResponse{InactiveDisableVirtualSchemaProposal: val}, nil
}

func (k Keeper) EnableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllEnableVirtualSchemaProposalRequest) (*types.QueryAllEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var enableVirtualSchemaProposals []types.EnableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	enableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.EnableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(enableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var enableVirtualSchemaProposal types.EnableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &enableVirtualSchemaProposal); err != nil {
			return err
		}

		enableVirtualSchemaProposals = append(enableVirtualSchemaProposals, enableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllEnableVirtualSchemaProposalResponse{EnableVirtualSchemaProposal: enableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) EnableVirtualSchemaProposal(c context.Context, req *types.QueryGetEnableVirtualSchemaProposalRequest) (*types.QueryGetEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetEnableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetEnableVirtualSchemaProposalResponse{EnableVirtualSchemaProposal: val}, nil
}


func (k Keeper) ActiveEnableVirtualSchemaProposalAll(c context.Context, req *types.QueryAllActiveEnableVirtualSchemaProposalRequest) (*types.QueryAllActiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var activeEnableVirtualSchemaProposals []types.ActiveEnableVirtualSchemaProposal
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	activeEnableVirtualSchemaProposalStore := prefix.NewStore(store, types.KeyPrefix(types.ActiveEnableVirtualSchemaProposalKeyPrefix))

	pageRes, err := query.Paginate(activeEnableVirtualSchemaProposalStore, req.Pagination, func(key []byte, value []byte) error {
		var activeEnableVirtualSchemaProposal types.ActiveEnableVirtualSchemaProposal
		if err := k.cdc.Unmarshal(value, &activeEnableVirtualSchemaProposal); err != nil {
			return err
		}

		activeEnableVirtualSchemaProposals = append(activeEnableVirtualSchemaProposals, activeEnableVirtualSchemaProposal)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllActiveEnableVirtualSchemaProposalResponse{ActiveEnableVirtualSchemaProposal: activeEnableVirtualSchemaProposals, Pagination: pageRes}, nil
}

func (k Keeper) ActiveEnableVirtualSchemaProposal(c context.Context, req *types.QueryGetActiveEnableVirtualSchemaProposalRequest) (*types.QueryGetActiveEnableVirtualSchemaProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetActiveEnableVirtualSchemaProposal(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetActiveEnableVirtualSchemaProposalResponse{ActiveEnableVirtualSchemaProposal: val}, nil
}
