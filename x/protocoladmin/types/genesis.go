package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AdminList: []Admin{},
		GroupList: []Group{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in admin
	adminIndexMap := make(map[string]struct{})

	for _, elem := range gs.AdminList {
		index := string(AdminKey(elem.Group, elem.Admin))
		if _, ok := adminIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for admin")
		}
		adminIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in group
	groupIndexMap := make(map[string]struct{})

	for _, elem := range gs.GroupList {
		index := string(GroupKey(elem.Name))
		if _, ok := groupIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for group")
		}
		groupIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
