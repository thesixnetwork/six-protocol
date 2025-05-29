package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	errormod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	"github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

var (
	DefaultAttoDenom       = "asix"
	DefaultMicroDenom      = "usix"
	DefaultAttoToMicroDiff = 1_000_000_000_000
)

func (k Keeper) AttoCoinConverter(ctx sdk.Context, sender sdk.AccAddress, receiver sdk.AccAddress, amount sdkmath.Int) error {
	attoCoin := sdk.Coin{
		Denom:  DefaultAttoDenom,
		Amount: amount,
	}

	convertAmount := sdk.NewCoins(attoCoin)

	token, foundToken := k.GetToken(ctx, DefaultAttoDenom)
	if !foundToken {
		return errormod.Wrap(sdkerrors.ErrKeyNotFound, "token does not exist")
	}

	// unnecessary
	if token.Base != DefaultAttoDenom {
		return errormod.Wrap(sdkerrors.ErrInvalidRequest, "token is not asix")
	}

	if amount.IsZero() {
		return errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount must be more than zero")
	}

	if balance := k.bankKeeper.GetBalance(ctx, sender, DefaultAttoDenom); balance.Amount.LT(amount) {
		return errormod.Wrap(sdkerrors.ErrInvalidRequest, "Amount of token is too high than current balance")
	}

	supply := k.bankKeeper.GetSupply(ctx, DefaultAttoDenom)
	if supply.Amount.LT(amount) {
		return errormod.Wrap(sdkerrors.ErrInvalidRequest, "amount of token is higher than current total supply")
	}

	// send to module
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, convertAmount); err != nil {
		return errormod.Wrap(types.ErrSendCoinsFromAccountToModule, "Amount of token is too high than current balance due"+err.Error())
	}

	// module burn atto coin
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, convertAmount); err != nil {
		return err
	}

	microSix := sdk.NewCoin(DefaultMicroDenom, amount.QuoRaw(int64(DefaultAttoToMicroDiff)))

	// get the module account balance
	tokenmngrModuleAccount := k.accountKeeper.GetModuleAddress(types.ModuleName)
	moduleBalance := k.bankKeeper.GetBalance(ctx, tokenmngrModuleAccount, DefaultMicroDenom)

	// check if module account balance is enough to send
	if moduleBalance.Amount.LT(microSix.Amount) {
		return errormod.Wrap(sdkerrors.ErrInsufficientFunds, "module account balance is not enough to send")
	}

	// send back locked micro coin
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(microSix),
	); err != nil {
		return errormod.Wrap(types.ErrSendCoinsFromAccountToModule, "unable to send msg.Amounts from module to account despite previously minting msg.Amounts to module account: "+err.Error())
	}

	return nil
}