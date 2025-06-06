package keeper

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/thesixnetwork/six-protocol/x/nftmngr/types"

	"cosmossdk.io/store/prefix"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (k Keeper) NFTSchema(ctx context.Context, req *types.QueryGetNFTSchemaRequest) (*types.QueryGetNFTSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetNftschema(
		ctx,
		req.Code,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	virtualSchemas := k.getVirtualSchemasForNFTSchema(ctx, req.Code)
	virtualActions := k.getVirtualActionsForSchemas(ctx, virtualSchemas)

	result := val.ResultWithEmptyVirtualAction()

	if len(virtualActions) > 0 {
		for _, action := range virtualActions {
			result.OnchainData.VirtualActions = append(result.OnchainData.VirtualActions, &action)
		}
	}

	return &types.QueryGetNFTSchemaResponse{NFTSchema: result}, nil
}

func (k Keeper) NFTSchemaAll(ctx context.Context, req *types.QueryAllNFTSchemaRequest) (*types.QueryAllNFTSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nftschemas []types.NFTSchemaQueryResult

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	nftschemaStore := prefix.NewStore(store, types.KeyPrefix(types.NftschemaKeyPrefix))

	pageRes, err := query.Paginate(nftschemaStore, req.Pagination, func(key []byte, value []byte) error {
		var nftschema types.NFTSchema
		if err := k.cdc.Unmarshal(value, &nftschema); err != nil {
			return err
		}

		nftschemas = append(nftschemas, nftschema.ResultWithEmptyVirtualAction())
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaResponse{NFTSchema: nftschemas, Pagination: pageRes}, nil
}

func (k Keeper) getVirtualSchemasForNFTSchema(ctx context.Context, nftSchemaCode string) []types.VirtualSchema {
	// pre-index all schemas by nftSchemaCode
	virtualSchemaIndex := make(map[string][]types.VirtualSchema)
	allVirtualSchemas := k.GetAllVirtualSchema(ctx)

	for _, schema := range allVirtualSchemas {
		for _, reg := range schema.Registry {
			virtualSchemaIndex[reg.NftSchemaCode] = append(virtualSchemaIndex[reg.NftSchemaCode], schema)
		}
	}

	return virtualSchemaIndex[nftSchemaCode]
}

func (k Keeper) getVirtualActionsForSchemas(ctx context.Context, virtualSchemas []types.VirtualSchema) []types.VirtualAction {
	var virtualActions []types.VirtualAction
	allVirtualActions := k.GetAllVirtualAction(ctx)

	actionMap := make(map[string][]types.VirtualAction)
	for _, action := range allVirtualActions {
		actionMap[action.VirtualNftSchemaCode] = append(actionMap[action.VirtualNftSchemaCode], action)
	}

	for _, schema := range virtualSchemas {
		if actions, found := actionMap[schema.VirtualNftSchemaCode]; found {
			virtualActions = append(virtualActions, actions...)
		}
	}

	return virtualActions
}
