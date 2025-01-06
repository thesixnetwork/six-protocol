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

func (k Keeper) NFTSchemaAll(c context.Context, req *types.QueryAllNFTSchemaRequest) (*types.QueryAllNFTSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTSchemas []types.NFTSchemaQueryResult
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nFTSchemaStore := prefix.NewStore(store, types.KeyPrefix(types.NFTSchemaKeyPrefix))

	pageRes, err := query.Paginate(nFTSchemaStore, req.Pagination, func(key []byte, value []byte) error {
		var nFTSchema types.NFTSchema
		if err := k.cdc.Unmarshal(value, &nFTSchema); err != nil {
			return err
		}

		nFTSchemas = append(nFTSchemas, nFTSchema.ResultWithEmptyVirtualAction())
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaResponse{NFTSchema: nFTSchemas, Pagination: pageRes}, nil
}

func (k Keeper) NFTSchema(c context.Context, req *types.QueryGetNFTSchemaRequest) (*types.QueryGetNFTSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	schema, found := k.GetNFTSchema(ctx, req.Code)
	if !found {
		return nil, status.Error(codes.NotFound, "NFT schema not found")
	}

	virtualSchemas := k.getVirtualSchemasForNFTSchema(ctx, req.Code)
	virtualActions := k.getVirtualActionsForSchemas(ctx, virtualSchemas)

	result := schema.ResultWithEmptyVirtualAction()

	if len(virtualActions) > 0 {
		for _, action := range virtualActions {
			result.OnchainData.VirtualActions = append(result.OnchainData.VirtualActions, &action)
		}
	}

	return &types.QueryGetNFTSchemaResponse{NFTSchema: result}, nil
}

func (k Keeper) NFTSchemaByContractAll(c context.Context, req *types.QueryAllNFTSchemaByContractRequest) (*types.QueryAllNFTSchemaByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var nFTSchemaByContracts []types.NFTSchemaByContract
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	nFTSchemaByContractStore := prefix.NewStore(store, types.KeyPrefix(types.NFTSchemaByContractKeyPrefix))

	pageRes, err := query.Paginate(nFTSchemaByContractStore, req.Pagination, func(key []byte, value []byte) error {
		var nFTSchemaByContract types.NFTSchemaByContract
		if err := k.cdc.Unmarshal(value, &nFTSchemaByContract); err != nil {
			return err
		}

		nFTSchemaByContracts = append(nFTSchemaByContracts, nFTSchemaByContract)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllNFTSchemaByContractResponse{NFTSchemaByContract: nFTSchemaByContracts, Pagination: pageRes}, nil
}

func (k Keeper) NFTSchemaByContract(c context.Context, req *types.QueryGetNFTSchemaByContractRequest) (*types.QueryGetNFTSchemaByContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetNFTSchemaByContract(
		ctx,
		req.OriginContractAddress,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetNFTSchemaByContractResponse{NFTSchemaByContract: val}, nil
}

func (k Keeper) getVirtualSchemasForNFTSchema(ctx sdk.Context, nftSchemaCode string) []types.VirtualSchema {
	//pre-index all schemas by nftSchemaCode
	virtualSchemaIndex := make(map[string][]types.VirtualSchema)
	allVirtualSchemas := k.GetAllVirtualSchema(ctx)

	for _, schema := range allVirtualSchemas {
		for _, reg := range schema.Registry {
			virtualSchemaIndex[reg.NftSchemaCode] = append(virtualSchemaIndex[reg.NftSchemaCode], schema)
		}
	}

	return virtualSchemaIndex[nftSchemaCode]
}

func (k Keeper) getVirtualActionsForSchemas(ctx sdk.Context, virtualSchemas []types.VirtualSchema) []types.VirtualAction {
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
