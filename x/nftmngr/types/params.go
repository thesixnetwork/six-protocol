package types

import (
	fmt "fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultPeriod time.Duration = time.Hour * 24 * 2 // 2 days
)

var ParamsStoreKeyVotingPeriod = []byte("VotingPeriod")

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
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamsStoreKeyVotingPeriod, &p.VotingPeriod, validateDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if p.VotingPeriod == nil {
		return fmt.Errorf("voting period must not be nil: %d", p.VotingPeriod)
	}
	if p.VotingPeriod.Seconds() <= 0 {
		return fmt.Errorf("voting period must be positive: %s", p.VotingPeriod)
	}
	return nil
}

func validateDuration(i interface{}) error {
	_, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
