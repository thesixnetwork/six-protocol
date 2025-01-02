package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SchemaAttributeAll(c context.Context, req *types.QueryAllSchemaAttributeRequest) (*types.QueryAllSchemaAttributeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var schemaAttributes []types.SchemaAttribute
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	schemaAttributeStore := prefix.NewStore(store, types.KeyPrefix(types.SchemaAttributeKeyPrefix))

	pageRes, err := query.Paginate(schemaAttributeStore, req.Pagination, func(key []byte, value []byte) error {
		var schemaAttribute types.SchemaAttribute
		if err := k.cdc.Unmarshal(value, &schemaAttribute); err != nil {
			return err
		}

		schemaAttributes = append(schemaAttributes, schemaAttribute)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSchemaAttributeResponse{SchemaAttribute: schemaAttributes, Pagination: pageRes}, nil
}

func (k Keeper) SchemaAttribute(c context.Context, req *types.QueryGetSchemaAttributeRequest) (*types.QueryGetSchemaAttributeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetSchemaAttribute(
		ctx,
		req.NftSchemaCode,
		req.Name,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSchemaAttributeResponse{SchemaAttribute: val}, nil
}
