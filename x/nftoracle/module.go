package nftoracle

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/thesixnetwork/six-protocol/x/nftoracle/client/cli"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/keeper"
	"github.com/thesixnetwork/six-protocol/x/nftoracle/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the capability module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the capability module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the capability module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the capability module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterRESTRoutes registers the capability module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the capability module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the capability module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the capability module.
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// Name returns the capability module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the capability module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// QuerierRoute returns the capability module's query routing key.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns the capability module's Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the capability module's invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the capability module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the capability module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 2 }

// BeginBlock executes all ABCI BeginBlock logic respective to the capability module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock executes all ABCI EndBlock logic respective to the capability module. It
// returns no validator updates.
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	am.keeper.IterateActiveMintRequestsQueue(ctx, ctx.BlockHeader().Time, func(mintRequest types.MintRequest) (stop bool) {
		if mintRequest.Status == types.RequestStatus_PENDING {
			mintRequest.Status = types.RequestStatus_EXPIRED
			mintRequest.ExpiredHeight = ctx.BlockHeight()
			am.keeper.SetMintRequest(ctx, mintRequest)
			am.keeper.RemoveFromActiveMintRequestQueue(ctx, mintRequest.Id, ctx.BlockHeader().Time)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeMintRequest,
					sdk.NewAttribute(types.AttributeKeyMintRequestStatus, types.RequestStatus_EXPIRED.String()),
				),
			)
		}
		return false
	})
	am.keeper.IterateActiveActionRequestsQueue(ctx, ctx.BlockHeader().Time, func(actionRequest types.ActionOracleRequest) (stop bool) {
		if actionRequest.Status == types.RequestStatus_PENDING {
			actionRequest.Status = types.RequestStatus_EXPIRED
			actionRequest.ExpiredHeight = ctx.BlockHeight()
			am.keeper.SetActionRequest(ctx, actionRequest)
			am.keeper.RemoveFromActiveActionRequestQueue(ctx, actionRequest.Id, ctx.BlockHeader().Time)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeActionRequest,
					sdk.NewAttribute(types.AttributeKeyActionRequestStatus, types.RequestStatus_EXPIRED.String()),
				),
			)
		}
		return false
	})
	am.keeper.IterateActiveVerifyCollectionOwnersQueue(ctx, ctx.BlockHeader().Time, func(verifyCollectionOwnerRequest types.CollectionOwnerRequest) (stop bool) {
		if verifyCollectionOwnerRequest.Status == types.RequestStatus_PENDING {
			verifyCollectionOwnerRequest.Status = types.RequestStatus_EXPIRED
			verifyCollectionOwnerRequest.ExpiredHeight = ctx.BlockHeight()
			am.keeper.SetCollectionOwnerRequest(ctx, verifyCollectionOwnerRequest)
			am.keeper.RemoveFromActiveVerifyCollectionOwnerQueue(ctx, verifyCollectionOwnerRequest.Id, ctx.BlockHeader().Time)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeVerificationRequest,
					sdk.NewAttribute(types.AttributeKeyVerificationRequestStatus, types.RequestStatus_EXPIRED.String()),
				),
			)
		}
		return false
	})
	am.keeper.IterateActiveSyncActionSignerQueue(ctx, ctx.BlockHeader().Time, func(syncRequest types.SyncActionSigner) (stop bool) {
		if syncRequest.Status == types.RequestStatus_PENDING {
			syncRequest.Status = types.RequestStatus_EXPIRED
			syncRequest.ExpiredHeight = ctx.BlockHeight()
			am.keeper.SetSyncActionSigner(ctx, syncRequest)
			am.keeper.RemoveFromActiveSyncActionSignerQueue(ctx, syncRequest.Id, ctx.BlockHeader().Time)
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeSyncActionSigner,
					sdk.NewAttribute(types.EventTypeSyncActionSignerRequestStatus, types.RequestStatus_EXPIRED.String()),
				),
			)
		}
		return false
	})

	return []abci.ValidatorUpdate{}
}
