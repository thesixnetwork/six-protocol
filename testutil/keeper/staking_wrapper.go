package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type StakingKeeperWrapper struct {
	MsgServer stakingtypes.MsgServer
	Keeper    *stakingkeeper.Keeper
}

func (w StakingKeeperWrapper) Delegate(goCtx context.Context, msg *stakingtypes.MsgDelegate) (*stakingtypes.MsgDelegateResponse, error) {
	return w.MsgServer.Delegate(goCtx, msg)
}

func (w StakingKeeperWrapper) BeginRedelegate(goCtx context.Context, msg *stakingtypes.MsgBeginRedelegate) (*stakingtypes.MsgBeginRedelegateResponse, error) {
	return w.MsgServer.BeginRedelegate(goCtx, msg)
}

func (w StakingKeeperWrapper) Undelegate(goCtx context.Context, msg *stakingtypes.MsgUndelegate) (*stakingtypes.MsgUndelegateResponse, error) {
	return w.MsgServer.Undelegate(goCtx, msg)
}

func (w StakingKeeperWrapper) GetRedelegation(ctx sdk.Context, delAddr sdk.AccAddress, srcValAddr, dstValAddr sdk.ValAddress) (stakingtypes.Redelegation, bool) {
    return w.Keeper.GetRedelegation(ctx, delAddr, srcValAddr, dstValAddr)
}


func (w StakingKeeperWrapper) MaxEntries(ctx sdk.Context) uint32 {
	return w.Keeper.MaxEntries(ctx)
}
