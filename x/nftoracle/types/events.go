package types

const (
	EventTypeMintRequestCreated = "mint_request_created"
	AttributeKeyMintRequestID   = "mint_request_id"
	AttributeKeyNftSchemaCode   = "nft_schema_code"
	AttributeKeyTokenID         = "token_id"
	AttributeKeyRequiredConfirm = "required_confirm"

	EventTypeMintRequest          = "mint_request"
	AttributeKeyMintRequestStatus = "status"

	EventTypeMintRequestConfirmed = "mint_request_confirmed"

	EventTypeActionRequestCreated = "action_request_created"
	AttributeKeyActionRequestID   = "action_request_id"

	EventTypeActionRequest          = "action_request"
	AttributeKeyActionRequestStatus = "status"

	EventMessage                    = "action"
	EventTypeRunAction              = "run_action"
	AttributeKeyRunActionResult     = "run_action_result"
	AttributeKeyRunActionChangeList = "run_action_changelist"

	EventTypeVerificationRequestCreated   = "verify_request_created"
	AttributeKeyVerifyRequestID           = "verify_request_id"
	AttributeKeyVerificationRequestStatus = "status"

	EventTypeVerificationRequest       = "verify_request"
	EventTypeVerificationRequestStatus = "status"

	EventTypeSyncActionSigner              = "sync_signer_request"
	EventTypeSyncActionSignerRequestStatus = "status"

	AttributeRequestorAddress = "requestor_address"

	EventTypeSetMinimumConfirmation = "set_minimum_confirmation"
	AttributeKeyMinimumConfirmation = "minimum_confirmation"

	EventTypeActionSigner      = "set_action_signer"
	AttributeKeySignerOwner    = "action_signer_owner"
	AttributeKeySignerActor    = "action_signer_actor"
	AttributeKeySginerExpireAt = "action_signer_expire_at"

	EventTypeCreateActionSignerConfig = "create_action_signer_config"
	EventTypeSetActionSignerConfig    = "set_action_signer_config"
	EventTypeDeleteActionSignerConfig = "delete_action_signer_config"
	AttributeKeyChain                 = "chain_name"
	AttributeKeyNewContract           = "new_config_contract"
	AttributeKeyOldContract           = "old_config_contract"

	EventTypeSyncActionReqeustCreated         = "sync_action_request_created"
	AttributeKeySyncActionRequestID           = "sync_action_request_id"
	AttributeKeySyncActionRequestChain        = "sync_action_request_chain"
	AttributeKeySyncActionRequestOwnerAddress = "sync_action_request_owner_address"
	AttributeKeySyncActionRequestActorAddress = "sync_action_request_actor_address"

	EventTypeSyncRequestConfirmed = "sync_request_confirmed"
	AttributeKeySyncRequestID     = "sync_request_id"
	AttributeKeySyncRequestStatus = "sync_request_status"
)
