package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:    PortID,
		GroupList: []Group{},
		AdminList: []Admin{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
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
	// Check for duplicated index in admin
	adminIndexMap := make(map[string]struct{})

	for _, elem := range gs.AdminList {
		index := string(AdminKey(elem.Group, elem.Admin))
		if _, ok := adminIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for admin")
		}
		adminIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
