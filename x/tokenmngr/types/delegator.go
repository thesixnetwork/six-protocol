package types

import (
	"cosmossdk.io/math"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// DelegatorStakingInfo contains comprehensive staking information for a delegator
type DelegatorStakingInfo struct {
	DelegatorAddress     string                             `json:"delegator_address"`
	Delegations          []stakingtypes.Delegation          `json:"delegations"`
	UnbondingDelegations []stakingtypes.UnbondingDelegation `json:"unbonding_delegations"`
	Redelegations        []stakingtypes.Redelegation        `json:"redelegations"`
	TotalBonded          math.Int                           `json:"total_bonded"`
	TotalUnbonding       math.Int                           `json:"total_unbonding"`
}
