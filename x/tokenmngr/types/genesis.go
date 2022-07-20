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
		PortId:        PortID,
		TokenList:     []Token{},
		MintpermList:  []Mintperm{},
		Options:       nil,
		TokenBurnList: []TokenBurn{},
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
	// Check for duplicated index in token
	tokenIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenList {
		index := string(TokenKey(elem.Name))
		if _, ok := tokenIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for token")
		}
		tokenIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in mintperm
	mintpermIndexMap := make(map[string]struct{})

	for _, elem := range gs.MintpermList {
		index := string(MintpermKey(elem.Token, elem.Address))
		if _, ok := mintpermIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for mintperm")
		}
		mintpermIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in tokenBurn
	tokenBurnIndexMap := make(map[string]struct{})

	for _, elem := range gs.TokenBurnList {
		index := string(TokenBurnKey(elem.Token))
		if _, ok := tokenBurnIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenBurn")
		}
		tokenBurnIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
