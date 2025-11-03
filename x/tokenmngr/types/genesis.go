package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TokenList:     []Token{},
		MintpermList:  []Mintperm{},
		TokenBurnList: []TokenBurn{},
		Options:       nil,
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
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
		index := string(TokenBurnKey(elem.Amount.Denom))
		if _, ok := tokenBurnIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tokenBurn")
		}
		tokenBurnIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
