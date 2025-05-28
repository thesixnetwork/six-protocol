package types

import (
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultPeriod time.Duration = time.Hour * 24 * 2 // 2 days
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(votingPeriod time.Duration) Params {
	return Params{
		VotingPeriod: &votingPeriod,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultPeriod)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
