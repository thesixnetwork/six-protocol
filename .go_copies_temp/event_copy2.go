package types

const (
	EventTypeMint          = "mint_token"
	AttributeKeyMintStatus = "status"

	EventTypeBurn          = "burn_token"
	AttributeKeyBurnStatus = "status"

	EventTypeGrantPermission          = "grant_permission"
	AttributeKeyGrantPermissionStatus = "status"
	AttributeKeyPermissionType        = "permission_type"
	AttributeKeyPermissionAddress     = "permission_address"

	EventTypeRevokePermission           = "revoke_permission"
	AttributeKeyRevokePermissionStatus  = "status"
	AttributeKeyRevokePermissionType    = "revoke_type"
	AttributeKeyRevokePermissionAddress = "revokee_address"

	AttributeKeyToken  = "token"
	AttributeKeyAmount = "amount"
)
