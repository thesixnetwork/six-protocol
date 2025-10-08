package types

import (
	"fmt"
	time "time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamspace default name for parameter store
	DefaultMintRequestActiveDuration      = 120 * time.Second
	DefaultActionRequestActiveDuration    = 120 * time.Second
	DefaultVerifyRequestActiveDuration    = 120 * time.Second
	DefaultActionSignerActiveDuration     = 30 * (24 * time.Hour) // 30 days
	DefaultSyncActionSignerActiveDuration = 300 * time.Second
)

var (
	KeyMintRequestActiveDuration      = []byte("MintRequestActiveDuration")
	KeyActionRequestActiveDuration    = []byte("ActionRequestActiveDuration")
	KeyVerifyRequestActiveDuration    = []byte("VerifyRequestActiveDuration")
	KeyActionSignerActiveDuration     = []byte("ActionSignerActiveDuration")
	KeySyncActionSignerActiveDuration = []byte("SyncActionSignerActiveDuration")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(mintRequestActiveDuration time.Duration, actionRequestActiveDuration time.Duration, verifyRequestActiveDuration time.Duration, actionSignerActiveDuration time.Duration, syncActionSignerActiveDuration time.Duration) Params {
	return Params{
		MintRequestActiveDuration:      mintRequestActiveDuration,
		ActionRequestActiveDuration:    actionRequestActiveDuration,
		VerifyRequestActiveDuration:    verifyRequestActiveDuration,
		ActionSignerActiveDuration:     actionSignerActiveDuration,
		SyncActionSignerActiveDuration: syncActionSignerActiveDuration,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(DefaultMintRequestActiveDuration, DefaultActionRequestActiveDuration, DefaultVerifyRequestActiveDuration, DefaultActionSignerActiveDuration, DefaultSyncActionSignerActiveDuration)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintRequestActiveDuration, &p.MintRequestActiveDuration, validateMintRequestActiveDuration),
		paramtypes.NewParamSetPair(KeyActionRequestActiveDuration, &p.ActionRequestActiveDuration, validateActionRequestActiveDuration),
		paramtypes.NewParamSetPair(KeyVerifyRequestActiveDuration, &p.VerifyRequestActiveDuration, validateRequestActiveDuration),
		paramtypes.NewParamSetPair(KeyActionSignerActiveDuration, &p.ActionSignerActiveDuration, validateRequestActiveDuration),
		paramtypes.NewParamSetPair(KeySyncActionSignerActiveDuration, &p.SyncActionSignerActiveDuration, validateRequestSyncActiveDuration),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMintRequestActiveDuration(p.MintRequestActiveDuration); err != nil {
		return err
	}
	if err := validateActionRequestActiveDuration(p.ActionRequestActiveDuration); err != nil {
		return err
	}
	if err := validateRequestActiveDuration(p.VerifyRequestActiveDuration); err != nil {
		return err
	}
	if err := validateRequestActiveDuration(p.ActionSignerActiveDuration); err != nil {
		return err
	}
	if err := validateRequestActiveDuration(p.SyncActionSignerActiveDuration); err != nil {
		return err
	}
	return nil
}

func validateMintRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateActionRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateRequestActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}

func validateRequestSyncActiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unbonding time must be positive: %d", v)
	}

	return nil
}
