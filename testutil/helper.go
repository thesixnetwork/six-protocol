package testutil

import (


	sdk "github.com/cosmos/cosmos-sdk/types"
	
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

)

// TestAccount holds both SDK and EVM addresses for testing
type TestAccount struct {
	Priv      *secp256k1.PrivKey
	Address   sdk.AccAddress
	EvmAddr   [20]byte
	PubKey    cryptotypes.PubKey
}

// TestValidator holds the validator operator address for testing
type TestValidator struct {
	Address sdk.ValAddress
}

// CreateRandomAccount generates a random account with a private key and address
func CreateRandomAccount() TestAccount {
	priv := secp256k1.GenPrivKey()
	addr := sdk.AccAddress(priv.PubKey().Address())
	evmAddr := [20]byte{}
	copy(evmAddr[:], addr.Bytes()[:20])

	return TestAccount{
		Priv:    priv,
		Address: addr,
		EvmAddr: evmAddr,
		PubKey:  priv.PubKey(),
	}
}

// FundAccount sends tokens to the given address in the context
func FundAccount(bankKeeper bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, "mint", coins); err != nil {
		return err
	}
	return bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", addr, coins)
}

// CreateValidator registers a validator with self-delegation
func CreateValidator(
	ctx sdk.Context,
	stakingKeeper *stakingkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	moniker string,
	stakeAmount sdk.Int,
) TestValidator {
	priv := secp256k1.GenPrivKey()
	valAddr := sdk.ValAddress(priv.PubKey().Address())
	accAddr := sdk.AccAddress(priv.PubKey().Address())

	coins := sdk.NewCoins(sdk.NewCoin("usix", stakeAmount))
	_ = bankKeeper.MintCoins(ctx, "mint", coins)
	_ = bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", accAddr, coins)

	description := stakingtypes.NewDescription(moniker, "", "", "", "")
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	msg, err := stakingtypes.NewMsgCreateValidator(
		valAddr,
		accAddr.String(), // ✅ Added this line
		priv.PubKey(),
		sdk.NewCoin("usix", stakeAmount),
		description,
		commission,
		sdk.OneInt(),
	)
	if err != nil {
		panic(err)
	}

	msgServer := stakingkeeper.NewMsgServerImpl(*stakingKeeper)
	_, err = msgServer.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	if err != nil {
		panic(err)
	}

	return TestValidator{Address: valAddr}
}



// GeneratePrivKeyAndValAddress creates a new secp256k1 private key and returns the key and validator address
func GeneratePrivKeyAndValAddress(moniker string) (*secp256k1.PrivKey, sdk.ValAddress) {
	priv := secp256k1.GenPrivKey()
	valAddr := sdk.ValAddress(priv.PubKey().Address())
	return priv, valAddr
}

// CreateValidatorWithPrivKey registers a validator with a given private key (useful for pre-approving)
func CreateValidatorWithPrivKey(
	ctx sdk.Context,
	stakingKeeper *stakingkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
	moniker string,
	stakeAmount sdk.Int,
	priv *secp256k1.PrivKey,
) TestValidator {
	valAddr := sdk.ValAddress(priv.PubKey().Address())
	accAddr := sdk.AccAddress(priv.PubKey().Address())

	coins := sdk.NewCoins(sdk.NewCoin("usix", stakeAmount))
	_ = bankKeeper.MintCoins(ctx, "mint", coins)
	_ = bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", accAddr, coins)

	description := stakingtypes.NewDescription(moniker, "", "", "", "")
	commission := stakingtypes.NewCommissionRates(sdk.ZeroDec(), sdk.ZeroDec(), sdk.ZeroDec())

	msg, err := stakingtypes.NewMsgCreateValidator(
		valAddr,
		accAddr.String(),
		priv.PubKey(),
		sdk.NewCoin("usix", stakeAmount),
		description,
		commission,
		sdk.OneInt(),
	)
	if err != nil {
		panic(err)
	}

	msgServer := stakingkeeper.NewMsgServerImpl(*stakingKeeper)
	_, err = msgServer.CreateValidator(sdk.WrapSDKContext(ctx), msg)
	if err != nil {
		panic(err)
	}

	return TestValidator{Address: valAddr}
}



