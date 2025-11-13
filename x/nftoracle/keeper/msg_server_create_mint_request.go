package keeper

import (
	"context"
	"strconv"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// create mint request
func (k msgServer) CreateMintRequest(goCtx context.Context, msg *types.MsgCreateMintRequest) (*types.MsgCreateMintRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if nft_schema_code exists
	_, found := k.nftmngrKeeper.GetNFTSchema(ctx, msg.NftSchemaCode)
	if !found {
		return nil, errormod.Wrap(types.ErrNFTSchemaNotFound, msg.NftSchemaCode)
	}
	// Check if the token is already minted
	_, found = k.nftmngrKeeper.GetNftData(ctx, msg.NftSchemaCode, msg.TokenId)
	if found {
		return nil, errormod.Wrap(types.ErrMetadataAlreadyExists, msg.NftSchemaCode)
	}

	createdAt := ctx.BlockTime()
	endTime := createdAt.Add(k.MintRequestActiveDuration(ctx))

	// Get Oracle Config
	oracleConfig, found := k.GetOracleConfig(ctx)
	if !found {
		return nil, errormod.Wrap(types.ErrOracleConfigNotFound, "no oracle config found")
	}

	// Verify msg.RequiredConfirmations is less than or equal to oracleConfig.MinimumConfirmation
	if int32(msg.RequiredConfirm) < oracleConfig.MinimumConfirmation {
		return nil, errormod.Wrap(types.ErrRequiredConfirmTooLess, strconv.Itoa(int(oracleConfig.MinimumConfirmation)))
	}

	id_ := k.Keeper.AppendMintRequest(ctx, types.MintRequest{
		NftSchemaCode:   msg.NftSchemaCode,
		TokenId:         msg.TokenId,
		RequiredConfirm: msg.RequiredConfirm,
		Status:          types.RequestStatus_PENDING,
		CurrentConfirm:  0,
		CreatedAt:       createdAt,
		ValidUntil:      endTime,
		Confirmers:      make([]string, 0),
		DataHashes:      make([]*types.DataHash, 0),
	})

	k.Keeper.InsertActiveMintRequestQueue(ctx, id_, endTime)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeMintRequestCreated,
			sdk.NewAttribute(types.AttributeKeyMintRequestID, strconv.FormatUint(id_, 10)),
			sdk.NewAttribute(types.AttributeKeyNftSchemaCode, msg.NftSchemaCode),
			sdk.NewAttribute(types.AttributeKeyTokenID, msg.TokenId),
			sdk.NewAttribute(types.AttributeKeyRequiredConfirm, strconv.FormatUint(msg.RequiredConfirm, 10)),
		),
	})

	return &types.MsgCreateMintRequestResponse{
		Id:            id_,
		NftSchemaCode: msg.NftSchemaCode,
		TokenId:       msg.TokenId,
	}, nil
}
