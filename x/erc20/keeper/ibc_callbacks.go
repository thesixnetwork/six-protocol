package keeper

import (
	"strings"

	errorsmod "cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/hashicorp/go-metrics"

	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"

	"github.com/evmos/evmos/v20/ibc"

	"github.com/thesixnetwork/six-protocol/v4/x/erc20/types"
)

// OnRecvPacket performs the ICS20 middleware receive callback for automatically
// converting an IBC Coin to their ERC20 representation.
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	// If ERC20 module is disabled no-op
	if !k.IsERC20Enabled(ctx) {
		return ack
	}

	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		err = errorsmod.Wrapf(errortypes.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data")
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// use a zero gas config to avoid extra costs for the relayers
	ctx = ctx.
		WithKVGasConfig(storetypes.GasConfig{}).
		WithTransientKVGasConfig(storetypes.GasConfig{})

	sender, recipient, _, _, err := ibc.GetTransferSenderRecipient(data)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	evmParams := k.evmKeeper.GetParams(ctx)

	// If we received an IBC message from a non-EVM channel, the sender and receiver
	// accounts should be different.
	if sender.Equals(recipient) && !evmParams.IsEVMChannel(packet.DestinationChannel) {
		return channeltypes.NewErrorAcknowledgement(types.ErrInvalidIBC)
	}

	receiverAcc := k.accountKeeper.GetAccount(ctx, recipient)

	// return acknowledgement without conversion if receiver is a module account
	if types.IsModuleAccount(receiverAcc) {
		return ack
	}

	// parse the transferred denom
	coin := ibc.GetReceivedCoin(
		packet.SourcePort, packet.SourceChannel,
		packet.DestinationPort, packet.DestinationChannel,
		data.Denom, data.Amount,
	)

	// If the coin denom starts with `factory/` then it is a token factory coin, skip it
	if strings.HasPrefix(data.Denom, "factory/") {
		return ack
	}

	// check if the coin is a native staking token
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}
	if coin.Denom == bondDenom {
		// no-op, received coin is the staking denomination
		return ack
	}

	pairID := k.GetTokenPairID(ctx, coin.Denom)
	pair, found := k.GetTokenPair(ctx, pairID)
	switch {
	// Case 1. token pair is not registered and is a single hop IBC Coin
	case !found && strings.HasPrefix(coin.Denom, "ibc/") && ibc.IsBaseDenomFromSourceChain(data.Denom):
		tokenPair, err := k.RegisterERC20Extension(ctx, coin.Denom)
		if err != nil {
			return channeltypes.NewErrorAcknowledgement(err)
		}

		ctx.EventManager().EmitEvents(
			sdk.Events{
				sdk.NewEvent(
					types.EventTypeRegisterERC20Extension,
					sdk.NewAttribute(types.AttributeCoinSourceChannel, packet.SourceChannel),
					sdk.NewAttribute(types.AttributeKeyERC20Token, tokenPair.Erc20Address),
					sdk.NewAttribute(types.AttributeKeyCosmosCoin, tokenPair.Denom),
				),
			},
		)
		return ack

	// Case 2. native ERC20 token
	case found && pair.IsNativeERC20():
		// Token pair is disabled -> return
		if !pair.Enabled {
			return ack
		}

		balance := k.bankKeeper.GetBalance(ctx, recipient, coin.Denom)
		if err := k.ConvertCoinNativeERC20(ctx, pair, balance.Amount, common.BytesToAddress(recipient.Bytes()), recipient); err != nil {
			return channeltypes.NewErrorAcknowledgement(err)
		}

		// Telemetry for successful conversion.
		telemetry.IncrCounterWithLabels(
			[]string{types.ModuleName, "ibc", "on_recv", "total"},
			1,
			[]metrics.Label{
				telemetry.NewLabel("denom", coin.Denom),
				telemetry.NewLabel("source_channel", packet.SourceChannel),
				telemetry.NewLabel("source_port", packet.SourcePort),
			},
		)
	}

	return ack
}

// OnAcknowledgementPacket responds to the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementPacket(
	ctx sdk.Context, _ channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	ack channeltypes.Acknowledgement,
) error {
	switch ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:
		return k.ConvertCoinToERC20FromPacket(ctx, data)
	default:
		return nil
	}
}

// OnTimeoutPacket converts the IBC coin to ERC20 after refunding the sender
// since the original packet sent was never received and has been timed out.
func (k Keeper) OnTimeoutPacket(ctx sdk.Context, _ channeltypes.Packet, data transfertypes.FungibleTokenPacketData) error {
	return k.ConvertCoinToERC20FromPacket(ctx, data)
}

// ConvertCoinToERC20FromPacket converts the IBC coin to ERC20 after refunding the sender.
// This function is only executed when IBC timeout or an Error ACK happens.
func (k Keeper) ConvertCoinToERC20FromPacket(ctx sdk.Context, data transfertypes.FungibleTokenPacketData) error {
	sender, err := sdk.AccAddressFromBech32(data.Sender)
	if err != nil {
		return err
	}

	pairID := k.GetTokenPairID(ctx, data.Denom)
	pair, found := k.GetTokenPair(ctx, pairID)
	if !found {
		// no-op, token pair is not registered
		return nil
	}

	coin := ibc.GetSentCoin(data.Denom, data.Amount)

	switch {

	// Case 1. if pair is native coin -> no-op
	case pair.IsNativeCoin():
		return nil

	// Case 2. if pair is native ERC20 -> unescrow
	case pair.IsNativeERC20():
		ctx = ctx.
			WithKVGasConfig(storetypes.GasConfig{}).
			WithTransientKVGasConfig(storetypes.GasConfig{})

		params := k.GetParams(ctx)
		if !params.EnableErc20 || !k.IsDenomRegistered(ctx, coin.Denom) {
			return nil
		}

		// assume that all module accounts need to have their tokens in the
		// IBC representation as opposed to ERC20
		senderAcc := k.accountKeeper.GetAccount(ctx, sender)
		if types.IsModuleAccount(senderAcc) {
			return nil
		}

		// Convert from Coin to ERC20
		if err := k.ConvertCoinNativeERC20(ctx, pair, coin.Amount, common.BytesToAddress(sender), sender); err != nil {
			defer func() {
				telemetry.IncrCounter(1, types.ModuleName, "ibc", "error", "total")
			}()
			return err
		}
	}

	return nil
}
