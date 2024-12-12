package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NFTSchemaList:             []NFTSchema{},
		NftDataList:               []NftData{},
		ActionByRefIdList:         []ActionByRefId{},
		OrganizationList:          []Organization{},
		NFTSchemaByContractList:   []NFTSchemaByContract{},
		NftFeeConfig:              nil,
		NFTFeeBalance:             nil,
		MetadataCreatorList:       []MetadataCreator{},
		NftCollectionList:         []NftCollection{},
		ActionExecutorList:        []ActionExecutor{},
		SchemaAttributeList:       []SchemaAttribute{},
		ActionOfSchemaList:        []ActionOfSchema{},
		ExecutorOfSchemaList:      []ExecutorOfSchema{},
		VirtualActionList:         []VirtualAction{},
		VirtualSchemaList:         []VirtualSchema{},
		DisableVirtualSchemaList:  []DisableVirtualSchema{},
		VirtualSchemaProposalList: []VirtualSchemaProposal{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in nFTSchema
	nFTSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.NFTSchemaList {
		index := string(NFTSchemaKey(elem.Code))
		if _, ok := nFTSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nFTSchema")
		}
		nFTSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nftData
	nftDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.NftDataList {
		index := string(NftDataKey(elem.NftSchemaCode, elem.TokenId))
		if _, ok := nftDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nftData")
		}
		nftDataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionByRefId
	actionByRefIdIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionByRefIdList {
		index := string(ActionByRefIdKey(elem.RefId))
		if _, ok := actionByRefIdIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionByRefId")
		}
		actionByRefIdIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in organization
	organizationIndexMap := make(map[string]struct{})

	for _, elem := range gs.OrganizationList {
		index := string(OrganizationKey(elem.Name))
		if _, ok := organizationIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for organization")
		}
		organizationIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nFTSchemaByContract
	nFTSchemaByContractIndexMap := make(map[string]struct{})

	for _, elem := range gs.NFTSchemaByContractList {
		index := string(NFTSchemaByContractKey(elem.OriginContractAddress))
		if _, ok := nFTSchemaByContractIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nFTSchemaByContract")
		}
		nFTSchemaByContractIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in metadataCreator
	metadataCreatorIndexMap := make(map[string]struct{})

	for _, elem := range gs.MetadataCreatorList {
		index := string(MetadataCreatorKey(elem.NftSchemaCode))
		if _, ok := metadataCreatorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for metadataCreator")
		}
		metadataCreatorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in nftCollection
	nftCollectionIndexMap := make(map[string]struct{})

	for _, elem := range gs.NftCollectionList {
		index := string(NftCollectionKey(elem.NftSchemaCode))
		if _, ok := nftCollectionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nftCollection")
		}
		nftCollectionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionExecutor
	actionExecutorIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionExecutorList {
		index := string(ActionExecutorKey(elem.NftSchemaCode, elem.ExecutorAddress))
		if _, ok := actionExecutorIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionExecutor")
		}
		actionExecutorIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in schemaAttribute
	schemaAttributeIndexMap := make(map[string]struct{})

	for _, elem := range gs.SchemaAttributeList {
		index := string(SchemaAttributeKey(elem.NftSchemaCode, elem.Name))
		if _, ok := schemaAttributeIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for schemaAttribute")
		}
		schemaAttributeIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in actionOfSchema
	actionOfSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.ActionOfSchemaList {
		index := string(ActionOfSchemaKey(elem.NftSchemaCode, elem.Name))
		if _, ok := actionOfSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for actionOfSchema")
		}
		actionOfSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in executorOfSchema
	executorOfSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.ExecutorOfSchemaList {
		index := string(ExecutorOfSchemaKey(elem.NftSchemaCode))
		if _, ok := executorOfSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for executorOfSchema")
		}
		executorOfSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in VirtualAction
	VirtualActionIndexMap := make(map[string]struct{})

	for _, elem := range gs.VirtualActionList {
		index := string(VirtualActionKey(elem.NftSchemaCode, elem.Name))
		if _, ok := VirtualActionIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for VirtualAction")
		}
		VirtualActionIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in virSchema
	virSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.VirtualSchemaList {
		index := string(VirtualSchemaKey(elem.VirtualNftSchemaCode))
		if _, ok := virSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for virSchema")
		}
		virSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in disableVirtualSchema
	disableVirtualSchemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.DisableVirtualSchemaList {
		index := string(DisableVirtualSchemaKey(elem.Id))
		if _, ok := disableVirtualSchemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for disableVirtualSchema")
		}
		disableVirtualSchemaIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in virtualSchemaProposal
	virtualSchemaProposalIndexMap := make(map[string]struct{})

	for _, elem := range gs.VirtualSchemaProposalList {
		index := string(VirtualSchemaProposalKey(elem.Id))
		if _, ok := virtualSchemaProposalIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for virtualSchemaProposal")
		}
		virtualSchemaProposalIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
