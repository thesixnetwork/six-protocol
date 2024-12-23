package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"
)

func (k msgServer) SubmitVerifyCollectionOwner(goCtx context.Context, msg *types.MsgSubmitVerifyCollectionOwner) (*types.MsgSubmitVerifyCollectionOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// only oracle can submit verify collection owner request
	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// check if creator or oracle has permission
	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracle, oracle)
	if !granted {
		return nil, sdkerrors.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	// get request for verification
	verifyRequest, found := k.GetCollectionOwnerRequest(ctx, msg.VerifyRequestID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrVerifyRequestNotFound, strconv.FormatUint(msg.VerifyRequestID, 10))
	}

	// check if request is still pending
	if verifyRequest.Status != types.RequestStatus_PENDING {
		return nil, sdkerrors.Wrap(types.ErrVerifyRequestNotPending, strconv.FormatUint(msg.VerifyRequestID, 10))
	}

	// check if current confirmation count is less than required confirmation count
	if verifyRequest.CurrentConfirm >= verifyRequest.RequiredConfirm {
		return nil, sdkerrors.Wrap(types.ErrVerifyRequestConfirmedAlreadyComplete, strconv.FormatUint(msg.VerifyRequestID, 10))
	}

	// Convert msg.base64OriginContractInfo to bytes
	_msgBase64OriginContractInfo, err := base64.StdEncoding.DecodeString(msg.Base64OriginContractInfo)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	contractOrigingParam := types.OriginContractParam{}
	err = k.cdc.(*codec.ProtoCodec).UnmarshalJSON(_msgBase64OriginContractInfo, &contractOrigingParam)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParsingBase64, err.Error())
	}

	if verifyRequest.CurrentConfirm == 0 {
		dataHash := sha256.Sum256(_msgBase64OriginContractInfo)
		// hash schema
		verifyRequest.ContractInfo = append(verifyRequest.ContractInfo, &types.OriginContractInfo{
			ContractOriginDataInfo: &contractOrigingParam,
			Hash:                   dataHash[:],
			Confirmers:             []string{msg.Creator},
		})
	} else {
		// // check if creator has alraedy confirmed
		// if _, ok := verifyRequest.Confirmers[msg.Creator]; ok {
		// 	return nil, sdkerrors.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.VerifyRequestID, 10)+", "+msg.Creator)
		// }
		for _, confirmer := range verifyRequest.Confirmers {
			if confirmer == msg.Creator {
				return nil, sdkerrors.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.VerifyRequestID, 10))
			}
		}

		// Compare data hash with previous data
		dataHash := sha256.Sum256(_msgBase64OriginContractInfo)
		dataHashMatch := false
		for _, origin_tx := range verifyRequest.ContractInfo {
			if res := bytes.Compare(dataHash[:], origin_tx.Hash); res == 0 {
				dataHashMatch = true
				origin_tx.Confirmers = append(origin_tx.Confirmers, msg.Creator)
				break
			}
		}
		if !dataHashMatch {
			verifyRequest.ContractInfo = append(verifyRequest.ContractInfo, &types.OriginContractInfo{
				ContractOriginDataInfo: &contractOrigingParam,
				Hash:                   dataHash[:],
				Confirmers:             []string{msg.Creator},
			})
		}
	}

	if verifyRequest.Confirmers == nil {
		verifyRequest.Confirmers = make([]string, 0)
	}

	// mark craetor as confirmed
	// verifyRequest.Confirmers[msg.Creator] = true
	verifyRequest.Confirmers = append(verifyRequest.Confirmers, msg.Creator)

	// increase actionRequest.CurrentConfirm
	verifyRequest.CurrentConfirm++

	if verifyRequest.CurrentConfirm == verifyRequest.RequiredConfirm {

		// // chaek if signer and deployer address is same
		// if verifyRequest.Signer != contractOrigingParam.ContractOwner {
		// 	verifyRequest.Status = types.RequestStatus_FAILED_REJECT_BY_CONSENSUS
		// } else {
		// 	verifyRequest.Status = types.RequestStatus_SUCCESS_WITH_CONSENSUS
		// }

		// Verify collection owner
		// Check if there is only one data hash
		if len(verifyRequest.ContractInfo) > 1 {
			// Update verifyRequest.Status to be FAILED
			verifyRequest.Status = types.RequestStatus_FAILED_WITHOUT_CONSENSUS
		} else if verifyRequest.Signer != contractOrigingParam.ContractOwner {
			// Update verfiyRequest.Status to be Rejected because information is not match
			verifyRequest.Status = types.RequestStatus_FAILED_REJECT_BY_CONSENSUS
		} else {
			// Update verifyRequest.Status to be SUCCESS
			verifyRequest.Status = types.RequestStatus_SUCCESS_WITH_CONSENSUS
		}

		// Emit event a consensus result
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeMintRequestConfirmed,
				sdk.NewAttribute(types.AttributeKeyVerifyRequestID, strconv.FormatUint(verifyRequest.Id, 10)),
				sdk.NewAttribute(types.AttributeKeyVerificationRequestStatus, verifyRequest.Status.String()),
			),
		)

		if verifyRequest.Status == types.RequestStatus_SUCCESS_WITH_CONSENSUS {
			err = nil
			// Udpate NFT Data
			schema, found := k.nftmngrKeeper.GetNFTSchema(ctx, verifyRequest.NftSchemaCode)
			if !found {
				return nil, sdkerrors.Wrap(types.ErrMetaDataNotFound, verifyRequest.NftSchemaCode)
			}

			// Set Collection Owner is verified
			schema.IsVerified = true
			k.nftmngrKeeper.SetNFTSchema(ctx, schema)
		}

	}

	// Update Mint Request Back to storage
	k.SetCollectionOwnerRequest(ctx, verifyRequest)

	return &types.MsgSubmitVerifyCollectionOwnerResponse{}, nil
}
