package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/nftoracle module sentinel errors
var (
	ErrInvalidSigner = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample        = sdkerrors.Register(ModuleName, 1101, "sample error")

	// 1 - 99
	ErrNFTSchemaNotFound                   = sdkerrors.Register(ModuleName, 1, "NFTSchema not found")
	ErrMetadataAlreadyExists               = sdkerrors.Register(ModuleName, 2, "Metadata already exists")
	ErrNoOraclePermission                  = sdkerrors.Register(ModuleName, 3, "No oracle permission")
	ErrMintRequestNotFound                 = sdkerrors.Register(ModuleName, 4, "Mint request not found")
	ErrMintRequestNotPending               = sdkerrors.Register(ModuleName, 5, "Mint request not pending")
	ErrMintRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 6, "Mint request confirmed already complete")
	ErrParsingBase64                       = sdkerrors.Register(ModuleName, 7, "Error parsing base64")
	ErrParsingOriginData                   = sdkerrors.Register(ModuleName, 8, "Error parsing origin data")
	ErrDataHashMismatch                    = sdkerrors.Register(ModuleName, 9, "Data hash mismatch")
	ErrOracleConfirmedAlready              = sdkerrors.Register(ModuleName, 10, "Oracle confirmed already")
	ErrUnsupportedMetadataFormat           = sdkerrors.Register(ModuleName, 11, "Unsupported metadata format")
	ErrInvalidBase64                       = sdkerrors.Register(ModuleName, 12, "Invalid base64")
	ErrParsingActionSignature              = sdkerrors.Register(ModuleName, 13, "Error parsing action signature")
	ErrParsingActionParam                  = sdkerrors.Register(ModuleName, 14, "Error parsing action params")
	ErrVerifyingSignature                  = sdkerrors.Register(ModuleName, 15, "Error verify signature")
	ErrMetadataNotExists                   = sdkerrors.Register(ModuleName, 16, "Metadata not exists")
	ErrInvalidSigningOnBehalfOf            = sdkerrors.Register(ModuleName, 17, "Invalid signing on behalf of")
	ErrHexDecode                           = sdkerrors.Register(ModuleName, 18, "Error hex decode")
	ErrEcrecover                           = sdkerrors.Register(ModuleName, 19, "Cannot ecrecover pubkey")
	ErrUnmarshalPubkey                     = sdkerrors.Register(ModuleName, 20, "Error Get address from pubkey")
	ErrValidateSignature                   = sdkerrors.Register(ModuleName, 21, "Error To Validate Signature")
	ErrInvalidExpireEpoch                  = sdkerrors.Register(ModuleName, 22, "Invalid Expire Epoch Input")

	// 100 - 199
	ErrActionRequestNotFound                 = sdkerrors.Register(ModuleName, 100, "Action request not found")
	ErrActionRequestNotPending               = sdkerrors.Register(ModuleName, 101, "Action request not pending")
	ErrActionRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 102, "Action request confirmed already complete")
	ErrMetaDataNotFound                      = sdkerrors.Register(ModuleName, 103, "Metadata not found")
	ErrUnauthorizedCaller                    = sdkerrors.Register(ModuleName, 104, "Unauthorized caller")
	ErrEmptyChangeList                       = sdkerrors.Register(ModuleName, 105, "There's no changes from action")
	ErrRefIdAlreadyExists                    = sdkerrors.Register(ModuleName, 106, "RefId already exists")
	ErrNFTSchemaAttributeNotFound            = sdkerrors.Register(ModuleName, 107, "NFTSchema attribute not found")

	// 200 - 299
	ErrParsingCollectionOwnerSignature = sdkerrors.Register(ModuleName, 200, "Error parsing collection owner signature")
	ErrNFTSchemaAlreadyVerified        = sdkerrors.Register(ModuleName, 201, "NFTSchema already verified")
	ErrNotCollectionOwner              = sdkerrors.Register(ModuleName, 202, "Not collection owner")
	ErrNotContractOwner                = sdkerrors.Register(ModuleName, 203, "Not Contract owner")

	// 300 - 399
	ErrVerifyRequestNotFound                 = sdkerrors.Register(ModuleName, 300, "Verify request not found")
	ErrVerifyRequestNotPending               = sdkerrors.Register(ModuleName, 301, "Verify request not pending")
	ErrVerifyRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 302, "Verify request confirmed already complete")
	ErrOracleRejectVerifyRequest             = sdkerrors.Register(ModuleName, 303, "Oracle reject verify request")
	ErrNoOracleAdminPermission               = sdkerrors.Register(ModuleName, 304, "No oracle admin permission")
	ErrOracleConfigNotFound                  = sdkerrors.Register(ModuleName, 305, "Oracle config not found")
	ErrRequiredConfirmTooLess                = sdkerrors.Register(ModuleName, 306, "Required confirm too less")

	// 400 - 499
	ErrParsingSetSignerSignature     = sdkerrors.Register(ModuleName, 400, "Error parsing collection owner signature")
	ErrNoPermissionAdminSignerConfig = sdkerrors.Register(ModuleName, 401, "No permission admin signer config")
	ErrActionSignerNotFound          = sdkerrors.Register(ModuleName, 402, "Action signer not found")
	ErrDeltedActionSigner            = sdkerrors.Register(ModuleName, 403, "Failed to delete action signer")

	// 500 - 599
	ErrSyncActionSignerRequestNotFound                 = sdkerrors.Register(ModuleName, 500, "Sync action signer request not found")
	ErrSyncActionSignerRequestNotPending               = sdkerrors.Register(ModuleName, 501, "Sync action signer request not pending")
	ErrSyncActionSignerRequestConfirmedAlreadyComplete = sdkerrors.Register(ModuleName, 502, "Sync action signer request confirmed already completed")
	ErrGetQueryActionSigner                            = sdkerrors.Register(ModuleName, 503, "Cannot query action signer")
	ErrActionSignerConfigNotFound                      = sdkerrors.Register(ModuleName, 504, "Action signer config not found")

	// 600 - 699 Gasless Oracle Voting errors
	ErrOracleAlreadyVoted               = sdkerrors.Register(ModuleName, 600, "Oracle already voted")
	ErrOracleSpamPrevention             = sdkerrors.Register(ModuleName, 601, "Oracle spam prevention triggered")
	ErrCollectionOwnerRequestNotFound   = sdkerrors.Register(ModuleName, 602, "Collection owner request not found")
	ErrCollectionOwnerRequestNotPending = sdkerrors.Register(ModuleName, 603, "Collection owner request not pending")
	ErrInvalidGaslessTransaction        = sdkerrors.Register(ModuleName, 604, "Invalid gasless transaction")

	// 1000 - 1099 The error code is reserved for SIXLINK
	ErrDialToEndpoint    = sdkerrors.Register(ModuleName, 1000, "Dial with endpoint")
	ErrParseAbi          = sdkerrors.Register(ModuleName, 1001, "Parse abi")
	ErrConvertingTokenId = sdkerrors.Register(ModuleName, 1002, "Converting token id")
	ErrCallContract      = sdkerrors.Register(ModuleName, 1003, "Call contract")
)
