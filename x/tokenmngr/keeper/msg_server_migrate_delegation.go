package keeper

import (
	"context"

	// errormod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"

	// "github.com/ethereum/go-ethereum/common"
	// evmostypes "github.com/evmos/evmos/v20/types"
)

func (k msgServer) MigrateDelegation(goCtx context.Context, msg *types.MsgMigrateDelegation) (*types.MsgMigrateDelegationResponse, error) {
	_ = sdk.UnwrapSDKContext(goCtx)

	// // check that receiver is cosmos address or ethereum address
	// var addr []byte
	// var receiver sdk.AccAddress

	// if err := evmostypes.ValidateAddress(msg.EthAddress); err != nil {
	// 	return nil, errormod.Wrap(sdkerrors.ErrInvalidAddress, "receiver address must be ethereum address")
	// }

	// if common.IsHexAddress(msg.EthAddress) {
	// 	addr = common.HexToAddress(msg.EthAddress).Bytes()
	// }

	// receiver = sdk.AccAddress(addr)

	// // Check is this creator is exist
	// sender, err := sdk.AccAddressFromBech32(msg.Creator)
	// if err != nil {
	// 	return nil, err
	// }

	// if err = k.ChangeDelegatorAddress(ctx, sender, receiver); err != nil {
	// 	return nil, err
	// }

	return &types.MsgMigrateDelegationResponse{}, nil
}
