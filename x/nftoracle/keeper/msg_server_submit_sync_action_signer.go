package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"strconv"
	"time"

	"github.com/thesixnetwork/six-protocol/v4/x/nftoracle/types"

	errormod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SubmitSyncActionSigner(goCtx context.Context, msg *types.MsgSubmitSyncActionSigner) (*types.MsgSubmitSyncActionSignerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	oracle, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	granted := k.nftadminKeeper.HasPermission(ctx, types.KeyPermissionOracle, oracle)
	if !granted {
		return nil, errormod.Wrap(types.ErrNoOraclePermission, msg.Creator)
	}

	SyncRequest, found := k.GetSyncActionSigner(ctx, msg.SyncId)
	if !found {
		return nil, errormod.Wrap(types.ErrSyncActionSignerRequestNotFound, strconv.FormatUint(msg.SyncId, 10))
	}

	// Check if requestId is still pending
	if SyncRequest.Status != types.RequestStatus_PENDING {
		return nil, errormod.Wrap(types.ErrSyncActionSignerRequestNotPending, strconv.FormatUint(msg.SyncId, 10))
	}

	// Check if current confirmation count is less than required confirmation count
	if SyncRequest.CurrentConfirm >= SyncRequest.RequiredConfirm {
		return nil, errormod.Wrap(types.ErrSyncActionSignerRequestConfirmedAlreadyComplete, strconv.FormatUint(msg.SyncId, 10))
	}
	// paramExpire, err := time.Parse(time.RFC3339, msg.ExpireAt)
	expireEpoch, err := strconv.ParseInt(msg.ExpireEpoch, 10, 64)
	if err != nil {
		return nil, errormod.Wrap(types.ErrInvalidExpireEpoch, strconv.FormatUint(msg.SyncId, 10))
	}
	paramExpire := time.Unix(expireEpoch, 0)

	param_info := types.ParameterSyncSignerByOracle{}
	// set param_info
	param_info.Chain = msg.Chain
	param_info.ActorAddress = msg.ActorAddress
	param_info.OwnerAddress = msg.OwnerAddress
	param_info.ExpireEpoch = msg.ExpireEpoch

	// byte of param_info
	paramDataBytes, err := k.cdc.Marshal(&param_info)
	if err != nil {
		return nil, err
	}

	// ! :: Check Deterministic Hash from concurrence response
	if SyncRequest.CurrentConfirm == 0 {
		contractParamHash := sha256.Sum256(paramDataBytes)
		SyncRequest.DataHashes = append(SyncRequest.DataHashes, &types.ContractInfoHash{
			ContractParam: &param_info,
			Hash:          contractParamHash[:],
			Confirmers:    []string{msg.Creator},
		})
	} else {
		// Check if creator has already confirmed this mint request
		// fetch confirmed data
		for _, confirmer := range SyncRequest.Confirmers {
			if confirmer == msg.Creator {
				return nil, errormod.Wrap(types.ErrOracleConfirmedAlready, strconv.FormatUint(msg.SyncId, 10))
			}
		}
		// Compare data hash with previous data hash
		contractParamHash := sha256.Sum256(paramDataBytes)
		dataHashMatch := false
		for _, hash := range SyncRequest.DataHashes {
			if res := bytes.Compare(contractParamHash[:], hash.Hash); res == 0 {
				dataHashMatch = true
				hash.Confirmers = append(hash.Confirmers, msg.Creator)
				break
			}
		}
		if !dataHashMatch {
			SyncRequest.DataHashes = append(SyncRequest.DataHashes, &types.ContractInfoHash{
				ContractParam: &param_info,
				Hash:          contractParamHash[:],
				Confirmers:    []string{msg.Creator},
			})
		}
	}

	if SyncRequest.Confirmers == nil {
		SyncRequest.Confirmers = make([]string, 0)
	}
	// Mark creator as confirmed
	// SyncRequest.Confirmers[msg.Creator] = true
	SyncRequest.Confirmers = append(SyncRequest.Confirmers, msg.Creator)

	// increase SyncRequest.CurrentConfirm
	SyncRequest.CurrentConfirm++

	if SyncRequest.CurrentConfirm == SyncRequest.RequiredConfirm {

		// Check if there is only one data hash
		if len(SyncRequest.DataHashes) > 1 {
			// Update SyncRequest.Status to be FAILED
			SyncRequest.Status = types.RequestStatus_FAILED_WITHOUT_CONSENSUS
		} else {
			// Update SyncRequest.Status to be SUCCESS
			SyncRequest.Status = types.RequestStatus_SUCCESS_WITH_CONSENSUS
		}

		// Emit event a consensus result
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeSyncRequestConfirmed,
				sdk.NewAttribute(types.AttributeKeySyncRequestID, strconv.FormatUint(SyncRequest.Id, 10)),
				sdk.NewAttribute(types.AttributeKeySyncRequestStatus, SyncRequest.Status.String()),
			),
		)

		if SyncRequest.Status == types.RequestStatus_SUCCESS_WITH_CONSENSUS {
			_, err := k.CreateSyncActionSignerByOracle(ctx, msg)
			if err != nil {
				SyncRequest.Status = types.RequestStatus_FAILED_ON_EXECUTION
				SyncRequest.ExecutionErrorMessage = err.Error()
			}

			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeActionSigner,
					sdk.NewAttribute(types.AttributeKeySignerOwner, msg.OwnerAddress),
					sdk.NewAttribute(types.AttributeKeySignerActor, msg.ActorAddress),
					sdk.NewAttribute(types.AttributeKeySginerExpireAt, paramExpire.UTC().Format(time.RFC3339)),
				),
			)
		}
	}

	k.SetSyncActionSigner(ctx, SyncRequest)

	return &types.MsgSubmitSyncActionSignerResponse{
		VerifyRequestID: SyncRequest.Id,
		ExpireAt:        SyncRequest.ValidUntil.Format(time.RFC3339),
	}, nil
}

// CreateSyncActionSignerByOracle
func (k msgServer) CreateSyncActionSignerByOracle(ctx sdk.Context, msg *types.MsgSubmitSyncActionSigner) (*types.MsgSubmitSyncActionSignerResponse, error) {
	SyncRequest, found := k.GetSyncActionSigner(ctx, msg.SyncId)
	if !found {
		return nil, errormod.Wrap(types.ErrSyncActionSignerRequestNotFound, strconv.FormatUint(msg.SyncId, 10))
	}

	_ExpireEpoch, _ := strconv.ParseInt(msg.ExpireEpoch, 10, 64)
	paramExpire := time.Unix(_ExpireEpoch, 0)

	if msg.ExpireEpoch == "" || msg.ExpireEpoch == "0" || msg.ExpireEpoch == " " {
		_, found := k.GetActionSigner(ctx, msg.ActorAddress, msg.OwnerAddress)
		if !found {
			return nil, errormod.Wrap(types.ErrActionSignerNotFound, msg.ActorAddress+","+msg.OwnerAddress)
		}

		k.RemoveActionSigner(ctx, msg.ActorAddress, msg.OwnerAddress)

		bindedList, found := k.GetBindedSigner(ctx, msg.OwnerAddress)
		if found {
			for _, bindedIndex := range bindedList.Signers {
				if bindedIndex.ActorAddress == msg.ActorAddress {
					k.RemoveSignerFromBindedSignerList(ctx, msg.OwnerAddress, bindedIndex.ActorAddress)
				}
			}
		}
	} else {
		_, found := k.GetActionSigner(ctx, msg.ActorAddress, msg.OwnerAddress)
		if !found {
			// create new action signer
			actionSigner := types.ActionSigner{
				ActorAddress: msg.ActorAddress,
				OwnerAddress: msg.OwnerAddress,
				ExpiredAt:    paramExpire,
				CreatedAt:    ctx.BlockTime(),
				Creator:      msg.Creator,
				CreationFlow: types.CreationFlow_ORACLE,
			}
			k.SetActionSigner(ctx, actionSigner)
			// add to binded signer list
			bindedList, found := k.GetBindedSigner(ctx, msg.OwnerAddress)
			if !found {
				bindedList = types.BindedSigner{
					OwnerAddress: msg.OwnerAddress,
					Signers:      make([]*types.BindedSignerParams, 0),
				}
			}

			// add the binded signer to the list
			bindedList.Signers = append(bindedList.Signers, &types.BindedSignerParams{
				ActorAddress: msg.ActorAddress,
				ExpiredAt:    paramExpire,
			})

			// set the binded signer
			bindedList.ActorCount = uint64(len(bindedList.Signers))
			k.SetBindedSigner(ctx, bindedList)
		} else {
			// update action signer
			actionSigner := types.ActionSigner{
				ActorAddress: msg.ActorAddress,
				OwnerAddress: msg.OwnerAddress,
				ExpiredAt:    paramExpire,
				CreatedAt:    ctx.BlockTime(),
				Creator:      msg.Creator,
				CreationFlow: types.CreationFlow_ORACLE,
			}
			k.SetActionSigner(ctx, actionSigner)

			bindedList, _ := k.GetBindedSigner(ctx, msg.OwnerAddress)
			for _, bindedIndex := range bindedList.Signers {
				if bindedIndex.ActorAddress == msg.ActorAddress {
					bindedIndex.ExpiredAt = paramExpire
				}
			}
			bindedList.ActorCount = uint64(len(bindedList.Signers))
			k.SetBindedSigner(ctx, bindedList)
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeActionSigner,
			sdk.NewAttribute(types.AttributeKeySignerOwner, msg.OwnerAddress),
			sdk.NewAttribute(types.AttributeKeySignerActor, msg.ActorAddress),
			sdk.NewAttribute(types.AttributeKeySginerExpireAt, paramExpire.UTC().Format(time.RFC3339)),
		),
	)
	return &types.MsgSubmitSyncActionSignerResponse{
		VerifyRequestID: SyncRequest.Id,
		ExpireAt:        SyncRequest.ValidUntil.Format(time.RFC3339),
	}, nil
}
