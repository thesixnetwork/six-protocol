package evmante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	ibcante "github.com/cosmos/ibc-go/v3/modules/core/ante"
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"

	evmtypes "github.com/thesixnetwork/six-protocol/x/evm/types"
)

// HandlerOptions extend the SDK's AnteHandler options by requiring the IBC
// channel keeper, EVM Keeper and Fee Market Keeper.
type HandlerOptions struct {
	AccountKeeper   evmtypes.AccountKeeper
	BankKeeper      evmtypes.BankKeeper
	IBCKeeper       *ibckeeper.Keeper
	FeeMarketKeeper evmtypes.FeeMarketKeeper
	EVMKeeper       EVMKeeper
	FeegrantKeeper  ante.FeegrantKeeper
	SignModeHandler authsigning.SignModeHandler
	SigGasConsumer  func(meter sdk.GasMeter, sig signing.SignatureV2, params authtypes.Params) error
	MaxTxGasWanted  uint64
}

func (options HandlerOptions) Validate() error {
	if options.AccountKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for AnteHandler")
	}
	if options.BankKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for AnteHandler")
	}
	if options.SignModeHandler == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	if options.FeeMarketKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "fee market keeper is required for AnteHandler")
	}
	if options.EVMKeeper == nil {
		return sdkerrors.Wrap(sdkerrors.ErrLogic, "evm keeper is required for AnteHandler")
	}
	return nil
}

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		NewEthSetUpContextDecorator(options.EVMKeeper),                         // outermost AnteDecorator. SetUpContext must be called first
		NewEthMempoolFeeDecorator(options.EVMKeeper),                           // Check eth effective gas price against minimal-gas-prices
		NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EVMKeeper), // Check eth effective gas price against the global MinGasPrice
		NewEthValidateBasicDecorator(options.EVMKeeper),
		NewEthSigVerificationDecorator(options.EVMKeeper),
		NewEthAccountVerificationDecorator(options.AccountKeeper, options.EVMKeeper),
		NewCanTransferDecorator(options.EVMKeeper),
		NewEthGasConsumeDecorator(options.EVMKeeper, options.MaxTxGasWanted),
		NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		NewGasWantedDecorator(options.EVMKeeper, options.FeeMarketKeeper),
		NewEthEmitEventDecorator(options.EVMKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		// NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EVMKeeper),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
		NewGasWantedDecorator(options.EVMKeeper, options.FeeMarketKeeper),
	)
}

func newCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		RejectMessagesDecorator{}, // reject MsgEthereumTxs
		ante.NewSetUpContextDecorator(),
		// NOTE: extensions option decorator removed
		// ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		NewMinGasPriceDecorator(options.FeeMarketKeeper, options.EVMKeeper),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		NewEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		ibcante.NewAnteDecorator(options.IBCKeeper),
		NewGasWantedDecorator(options.EVMKeeper, options.FeeMarketKeeper),
	)
}
