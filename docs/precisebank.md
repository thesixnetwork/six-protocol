# x/precisebank Module

## Overview

The `x/precisebank` module enables sub-integer precision for token balances in the SIX Protocol chain. It bridges the precision gap between:

- **`usix`** — the native Cosmos integer denomination (6 decimals relative to `six`)
- **`asix`** — the EVM-compatible extended denomination (18 decimals relative to `six`)

The conversion factor between these denominations is **10^12** (1 usix = 1,000,000,000,000 asix).

This module is adapted from [Kava's x/precisebank](https://github.com/Kava-Labs/kava/tree/master/x/precisebank) to fit the SIX Protocol chain architecture and denomination scheme.

## Problem Statement

Cosmos SDK's `x/bank` module stores coin balances as integers. SIX Protocol uses `usix` (micro‑SIX) as its integer coin, which has 6 decimal places relative to the display `six` denomination.

EVM operations (e.g., Solidity transfers) work with 18 decimal places (`asix`, atto‑SIX). This means 12 additional decimal places of precision are needed that the bank module cannot natively represent.

Without `x/precisebank`, any fractional `asix` amount smaller than 1 `usix` would be truncated or lost during EVM ↔ Cosmos transfers.

## Architecture

### Data Model

The module stores two types of data in its KV store:

1. **Fractional Balances** — Per-account fractional amounts (0 < amount < 10^12) stored under a prefix key `0x01`.
2. **Remainder** — A global remainder value stored under key `0x02` that accounts for rounding dust produced by mint/burn operations.

### Reserve Account

The module maintains a **reserve module account** (`precisebank`) with `Minter` and `Burner` permissions. This reserve holds the integer `usix` coin that backs all fractional balances.

**Invariant**: `sum(all fractional balances) + remainder ≡ 0 (mod 10^12)` and the reserve's `usix` balance = `(sum + remainder) / 10^12`.

### Extended Balance

An account's full extended `asix` balance is:

```
extended_balance = (bank_balance_usix * 10^12) + fractional_balance
```

### Coin Flow

#### Send (Account → Account)

When sending `asix` between accounts:

1. Split the amount into integer part (`usix`) and fractional part.
2. Subtract the fractional amount from the sender's fractional balance.
   - If the sender's fractional balance is insufficient, **borrow** 1 `usix` from the sender via `x/bank` and convert to fractional.
3. Add the fractional amount to the recipient's fractional balance.
   - If the recipient's fractional balance overflows (≥ 10^12), **carry** 1 `usix` to the recipient via `x/bank`.
4. Transfer the integer `usix` portion via `x/bank`.

#### Mint (Module → Account)

1. Split the mint amount into integer + fractional.
2. Add fractional to the target account's balance.
   - Handle carry to `usix` if fractional overflows.
3. Manage the reserve: mint integer `usix` to the `precisebank` module, then transfer to target.
4. Update the remainder for any rounding adjustments.

#### Burn (Account → Void)

1. Split the burn amount into integer + fractional.
2. Subtract fractional from the target account's balance.
   - Handle borrow from `usix` if fractional is insufficient.
3. Manage the reserve: transfer `usix` from `precisebank` to the burning module, then burn.
4. Update the remainder for any rounding adjustments.

## Module Integration in SIX Protocol

### Store Key

The module registers its own KV store with key `"precisebank"`.

### Module Account Permissions

The `precisebank` module account has `Minter` and `Burner` permissions, declared in `app_config.go`:

```go
precisebanktypes.ModuleName: {authtypes.Minter, authtypes.Burner},
```

### Keeper Dependencies

| Dependency | Interface | Purpose |
|---|---|---|
| `x/auth` | `AccountKeeper` | Retrieve module accounts and account sequences |
| `x/bank` | `BankKeeper` | Transfer, mint, burn `usix` coins; query balances |

### Keeper Initialization (app.go)

```go
app.PreciseBankKeeper = precisebankkeeper.NewKeeper(
    appCodec,
    runtime.NewKVStoreService(keys[precisebanktypes.StoreKey]),
    logger,
    app.BankKeeper,
    app.AccountKeeper,
)
```

### Module Registration

The module is registered in the `ModuleManager` and participates in:

- **Genesis ordering** — included in `genesisModuleOrder`
- **Begin/End block** — included in `beginBlockers` and `endBlockers` (currently no-ops)
- **Invariants** — 5 invariants registered to validate state consistency

## Invariants

| Invariant | Description |
|---|---|
| `ReserveBacksFractions` | Module's `usix` balance ≥ sum of fractional balances + remainder |
| `BalancedFractionalTotal` | (sum of fractional balances + remainder) mod 10^12 == 0 |
| `ValidFractionalAmounts` | All fractional balances are in range (0, 10^12) |
| `ValidRemainderAmount` | Remainder is in valid range [0, 10^12) |
| `FractionalDenomNotInBank` | No `asix` balances exist in `x/bank` (all `asix` tracked by this module) |

## gRPC Queries

The module exposes three query RPC methods defined in `proto/sixprotocol/precisebank/query.proto` and implemented in `keeper/grpc_query.go`.

| RPC | REST path | Description |
|---|---|---|
| `Remainder` | `GET /thesixnetwork/six-protocol/precisebank/remainder` | Returns the global remainder amount |
| `FractionalBalance` | `GET /thesixnetwork/six-protocol/precisebank/fractional_balance/{address}` | Returns fractional `asix` balance for one account |
| `FractionalBalances` | `GET /thesixnetwork/six-protocol/precisebank/fractional_balances` | Paginated list of all accounts with non-zero fractional balances |

### Remainder

```go
res, err := queryClient.Remainder(ctx, &types.QueryRemainderRequest{})
// res.Remainder — decimal string, e.g. "500000000000"
```

### FractionalBalance

```go
res, err := queryClient.FractionalBalance(ctx, &types.QueryFractionalBalanceRequest{
    Address: "6x1...",
})
// res.FractionalBalance — decimal string of the sub-integer asix amount
```

### FractionalBalances

```go
res, err := queryClient.FractionalBalances(ctx, &types.QueryFractionalBalancesRequest{
    Pagination: &query.PageRequest{Limit: 100},
})
// res.FractionalBalances — []FractionalBalanceEntry{Address, FractionalBalance}
// res.Pagination        — standard cosmos page response
```

## Genesis

### State

```json
{
  "balances": [
    { "address": "6x1...", "amount": "500000000000" }
  ],
  "remainder": "0"
}
```

### Validation

During `InitGenesis`:
1. All fractional balances must be valid (positive, < 10^12).
2. No duplicate addresses.
3. `sum(fractional balances) + remainder ≡ 0 (mod 10^12)`.
4. Module account must exist.
5. Module account's `usix` balance must equal `(sum + remainder) / 10^12`.

## File Structure

```
x/precisebank/
├── genesis.go                     # InitGenesis, ExportGenesis
├── module.go                      # AppModule, AppModuleBasic
├── keeper/
│   ├── keeper.go                  # Keeper struct, NewKeeper
│   ├── fractional_balance.go      # Get/Set/Delete fractional balances, remainder
│   ├── view.go                    # GetBalance, SpendableCoin
│   ├── send.go                    # SendCoins, SendCoinsFromAccountToModule, etc.
│   ├── mint.go                    # MintCoins
│   ├── burn.go                    # BurnCoins
│   ├── invariants.go              # RegisterInvariants, AllInvariants
│   ├── query.go                   # var _ types.QueryServer = Keeper{}
│   └── grpc_query.go              # Remainder, FractionalBalance, FractionalBalances handlers
├── types/
│   ├── keys.go                    # ModuleName, StoreKey, key prefixes
│   ├── constants.go               # IntegerCoinDenom, ExtendedCoinDenom
│   ├── fractional_balance.go      # FractionalBalance helpers, ConversionFactor
│   ├── fractional_balances.go     # FractionalBalances slice type
│   ├── extended_balance.go        # SumExtendedCoin helper
│   ├── expected_keepers.go        # AccountKeeper, BankKeeper interfaces
│   ├── genesis.go                 # GenesisState helpers (NewGenesisState, Validate, etc.)
│   ├── genesis.pb.go              # Generated: GenesisState, FractionalBalance protobuf types
│   ├── codec.go                   # Codec registration
│   ├── errors.go                  # Sentinel errors
│   ├── query.pb.go                # Generated: QueryServer interface, request/response types
│   └── query.pb.gw.go             # Generated: REST gateway handlers
proto/sixprotocol/precisebank/
├── genesis.proto                  # GenesisState and FractionalBalance message definitions
└── query.proto                    # gRPC Query service definition
```

## Differences from Kava's Implementation

| Aspect | Kava | SIX Protocol |
|---|---|---|
| Integer denomination | `ukava` | `usix` |
| Extended denomination | `akava` | `asix` |
| Conversion factor | 10^12 | 10^12 (same) |
| Store access | `storetypes.StoreKey` | `store.KVStoreService` (SDK v0.50) |
| Module pattern | Legacy `AppModule` | SDK v0.50 `appmodule` interfaces |
| Genesis codec | Protobuf | Protobuf (`genesis.proto` with `gogoproto.customtype` for `cosmossdk.io/math.Int`) |
| gRPC queries | Yes | Implemented (Remainder, FractionalBalance, FractionalBalances) |

## Usage

### Querying Extended Balance

The keeper provides `GetBalance(ctx, addr)` which returns the full `asix` balance:

```go
balance := app.PreciseBankKeeper.GetBalance(ctx, addr)
// balance.Denom == "asix"
// balance.Amount includes both usix * 10^12 + fractional
```

### Sending Extended Coins

```go
err := app.PreciseBankKeeper.SendCoins(ctx, fromAddr, toAddr, sdk.NewCoins(
    sdk.NewCoin("asix", sdkmath.NewInt(1_500_000_000_000)), // 1.5 usix
))
```

### Minting / Burning

```go
err := app.PreciseBankKeeper.MintCoins(ctx, "evm", sdk.NewCoins(
    sdk.NewCoin("asix", sdkmath.NewInt(500_000_000_000)), // 0.5 usix
))
```

## Module Workflow

This section describes how the precisebank module integrates into the Cosmos SDK chain lifecycle.

### 1. Chain Startup (InitGenesis)

```
app.go creates Keeper
    │
    ▼
InitGenesis(ctx, keeper, ak, bk, genesisState)
    │
    ├── Validate genesis state
    ├── Verify module account exists
    ├── Verify module usix balance matches sum of fractional balances + remainder
    ├── Store each FractionalBalance in KV store
    └── Store remainder in KV store
```

### 2. Runtime — EVM Transfer Flow

When an EVM transfer of `asix` occurs (e.g., a Solidity `transfer()`):

```
EVM calls precisebank.SendCoins(from, to, asix_amount)
    │
    ├── Split amount into integer_part (usix) + fractional_part
    │
    ├── Update sender's fractional balance
    │   └── If insufficient → borrow 1 usix from x/bank, convert to 10^12 fractional
    │
    ├── Update recipient's fractional balance
    │   └── If overflow ≥ 10^12 → carry 1 usix to x/bank
    │
    └── Transfer integer usix via x/bank.SendCoins()
```

### 3. Runtime — Mint Flow

```
Module calls precisebank.MintCoins(moduleName, asix_amount)
    │
    ├── Split into integer + fractional
    ├── Add fractional to target module's balance (handle carry)
    ├── Mint integer usix to precisebank reserve
    ├── Transfer from reserve to target module
    └── Update remainder for rounding
```

### 4. Runtime — Burn Flow

```
Module calls precisebank.BurnCoins(moduleName, asix_amount)
    │
    ├── Split into integer + fractional
    ├── Subtract fractional from target module's balance (handle borrow)
    ├── Transfer integer usix from target module to reserve
    ├── Burn from reserve
    └── Update remainder for rounding
```

### 5. Query Flow

```
Client → gRPC/REST request
    │
    ▼
module.RegisterServices() → types.RegisterQueryServer()
    │
    ▼
keeper.Remainder()            → reads remainder from KV store
keeper.FractionalBalance()    → reads one account's fractional balance
keeper.FractionalBalances()   → iterates all fractional balances (paginated)
```

### 6. Chain Export (ExportGenesis)

```
ExportGenesis(ctx, keeper)
    │
    ├── Iterate all fractional balances from KV store
    ├── Read remainder from KV store
    └── Return GenesisState{Balances, Remainder}
```

### 7. Invariant Checks

Invariants run periodically (or on demand) to verify:

```
RegisterInvariants()
    │
    ├── ReserveBacksFractions:     reserve usix == (sum + remainder) / 10^12
    ├── BalancedFractionalTotal:   (sum + remainder) % 10^12 == 0
    ├── ValidFractionalAmounts:    all balances in (0, 10^12)
    ├── ValidRemainderAmount:      remainder in [0, 10^12)
    └── FractionalDenomNotInBank:  x/bank has zero asix supply
```

## Proto Generation

The protobuf types are generated using the standard project pipeline:

```bash
# From the project root:
cd proto
buf generate --template buf.gen.gogo-nonignite.yaml sixprotocol/precisebank/genesis.proto
buf generate --template buf.gen.gogo-nonignite.yaml sixprotocol/precisebank/query.proto
cd ..
cp -r github.com/thesixnetwork/six-protocol/v4/* ./
rm -rf github.com
```

Or use the project's full generation script:

```bash
scripts/protocgen.sh
```

Generated files:
- `x/precisebank/types/genesis.pb.go` — from `genesis.proto`
- `x/precisebank/types/query.pb.go` — from `query.proto`
- `x/precisebank/types/query.pb.gw.go` — REST gateway from `query.proto`
