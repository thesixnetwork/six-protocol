package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NFTSchemaList:     []NFTSchema{},
		NftDataList:       []NftData{},
		ActionByRefIdList: []ActionByRefId{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in nftschema
	nftschemaIndexMap := make(map[string]struct{})

	for _, elem := range gs.NFTSchemaList {
		index := string(NFTSchemaKey(elem.Code))
		if _, ok := nftschemaIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for nftschema")
		}
		nftschemaIndexMap[index] = struct{}{}
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
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
