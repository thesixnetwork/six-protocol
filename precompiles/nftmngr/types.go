package nftmngr

// is TransactionMetadata
const (
	AddAction            = "addAction"
	AddAttribute         = "addAttribute"
	ChangeOrgOwner       = "changeOrgOwner"
	ChangeSchemaOwner    = "changeSchemaOwner"
	CreateMetadata       = "createMetadata"
	CreateSchema         = "createSchema"
	ActionByAdmin        = "actionByAdmin"
	ResyncAttribute      = "resyncAttribute"
	UpdateAttribute      = "updateSchemaAttribute"
	AttributeOveride     = "attributeOveride"
	SetBaseURI           = "setBaseURI"
	SetMetadataFormat    = "setMetadataFormat"
	SetMintAuth          = "setMintAuth"
	SetOriginChain       = "setOriginChain"
	SetOriginContract    = "setOriginContract"
	SetUriRetreival      = "setUriRetreival"
	ShowAttribute        = "showAttribute"
	ToggleAction         = "toggleAction"
	UpdateAction         = "updateAction"
	AddActionExecutor    = "addActionExecutor"
	RemoveActionExecutor = "removeActionExecutor"

	// virtual schema and virtual action
	PerformVirtualAction  = "virtualAction"
	VoteVirtualSchema     = "voteVirtualSchema"
	VirtualSchemaProposal = "virtualSchemaProposal"
)

// readonly
const (
	IsActionExecutor  = "isActionExecutor"
	IsSchemaOwner     = "isSchemaOwner"
	GetAttributeValue = "getAttributeValue"
)

type TransactionMetadata struct {
	// RequiresAuth  bool
	// ModifiesState bool
	// MinGas        uint64
	Description string
}

var transactionMethods map[string]TransactionMetadata

func init() {
	transactionMethods = map[string]TransactionMetadata{
		ActionByAdmin: {
			Description: "Perform action",
		},
		AddAction: {
			Description: "Add new action to schema",
		},
		AddAttribute: {
			Description: "Add new attribute to schema",
		},
		ChangeOrgOwner: {
			Description: "Change organization owner",
		},
		ChangeSchemaOwner: {
			Description: "Change schema owner",
		},
		CreateMetadata: {
			Description: "Create new metadata",
		},
		CreateSchema: {
			Description: "Create new NFT schema",
		},
		ResyncAttribute: {
			Description: "Resynchronize attribute",
		},
		UpdateAttribute: {
			Description: "Update existing attribute",
		},
		AttributeOveride: {
			Description: "Override attribute properties",
		},
		SetBaseURI: {
			Description: "Set base URI for NFTs",
		},
		SetMetadataFormat: {
			Description: "Set metadata format",
		},
		SetMintAuth: {
			Description: "Set minting authorization",
		},
		SetOriginChain: {
			Description: "Set origin chain",
		},
		SetOriginContract: {
			Description: "Set origin contract",
		},
		SetUriRetreival: {
			Description: "Set URI retrieval method",
		},
		ShowAttribute: {
			Description: "Show attribute details",
		},
		ToggleAction: {
			Description: "Toggle action state",
		},
		UpdateAction: {
			Description: "Update existing action",
		},
		AddActionExecutor: {
			Description: "Add new action executor",
		},
		RemoveActionExecutor: {
			Description: "Remove action executor",
		},
		VirtualSchemaProposal: {
			Description: "Proposal to create or edit virtual schema",
		},
		VoteVirtualSchema: {
			Description: "Vote to Accept or reject virtual schema proposal",
		},
		PerformVirtualAction: {
			Description: "Perform action that allow to make cross some value of schemas",
		},
	}
}
