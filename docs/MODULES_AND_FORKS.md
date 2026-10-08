# SIX Protocol Chain — Module & Fork Customization Summary

> Hand-off reference for the SIX Protocol blockchain node (`sixd`).
> Last updated: 2026-10-08.

The `sixd` node is a **Cosmos SDK v0.50 chain with full EVM** (Evmos v20 + go-ethereum core),
running dual Cosmos/EVM consensus. Three upstream dependencies are **hard-forked** by SIX
(consumed via `replace` in `go.mod`), and five custom `x/` modules plus a custom precompile
layer sit on top.

```
go.mod replaces (what's forked):
  cosmos-sdk      → thesixnetwork/cosmos-sdk  v0.50.10-sixpatch-3
  evmos/v20       → thesixnetwork/evmos/v20   v20.0.0-sixpatch-2
  go-ethereum     → thesixnetwork/go-ethereum v1.13.6-...ba2bad3ed2da
  cosmossdk.io/api, cosmossdk.io/store → thesixnetwork forks
```

---

## Part 1 — Custom `x/` modules

| Module | Purpose | State-changing Msgs (highlights) |
|---|---|---|
| **nftmngr** | Largest module. NFT schema + metadata + on-chain **rule/action engine** (Grule GRL). Schemas, attributes, actions, executors, orgs, virtual-schema governance. | `CreateNFTSchema`, `CreateMetadata`, `AddAttribute`, `AddAction`/`UpdateAction`/`ToggleAction`, `PerformActionByAdmin`, `SetFeeConfig`, `SetMintauth`, `Create/DeleteActionExecutor`, `ProposalVirtualSchema`/`VoteVirtualSchemaProposal`/`PerformVirtualAction`, origin-chain setters |
| **tokenmngr** | Token defs + mint permissions; **wrap/unwrap `usix` ↔ wrapped/EVM token**; delegation migration. Heavy Bank + EVM + Staking/Distribution interop. | `Create/Update/DeleteToken`, `Create/Update/DeleteMintperm`, `Mint`/`Burn`, `WrapToken`/`UnwrapToken`/`SendWrapToken`, `MigrateDelegation`, Options msgs |
| **nftoracle** | Multi-sig **oracle consensus** for cross-chain NFT verification (mint/action/collection-owner) + action-signer registry with expiry & cross-chain sync. | `CreateMintRequest`/`SubmitMintResponse`, `CreateActionRequest`/`SubmitActionResponse`, `VerifyCollectionOwner` msgs, `SetMinimumConfirmation`, `Create/Update/Delete ActionSigner` (+Config, +Sync) |
| **nftadmin** | Root-admin + named-permission ACL. Permission gate used by nftmngr/nftoracle. | `GrantPermission`, `RevokePermission` |
| **protocoladmin** | Named admin groups for protocol-level authority. Used by tokenmngr for param authority. | `Create/Update/DeleteGroup`, `Add/RemoveAdminFromGroup` |

**Dependency flow:** `protocoladmin`/`nftadmin` (authority) → `tokenmngr` + `nftmngr`/`nftoracle`
→ SDK (Bank/Account/Staking/Distribution) + Evmos EVM keeper. `nftadmin` is the permission check
everyone calls; `nftoracle` feeds verified data back into `nftmngr`.

**Notes:**
- All five modules still use the **legacy `x/params` subspace** pattern (not migrated to `MsgUpdateParams`).
- `nftmngr`'s rule engine uses a process-local GRL cache (SHA-256 keyed, 100-cycle cap) that is
  consensus-safe because builds are deterministic and each execution clones the knowledge base.

### Precompiles (`precompiles/`) — Cosmos modules exposed to Solidity

| Precompile | Address | Bridges to |
|---|---|---|
| **bank** | `0x…1001` | SDK x/bank — balances, metadata, send |
| **staking** | `0x…1005` | SDK x/staking (+ tokenmngr denom conversion) — delegate/redelegate/undelegate |
| **distribution** | `0x…1007` | SDK x/distribution (+ tokenmngr) — rewards, withdraw address |
| **nftmngr** | `0x…1055` | x/nftmngr + x/nftadmin — full schema/metadata/action lifecycle |
| **tokenfactory** | `0x…1069` | x/tokenmngr — wrap/unwrap, cross-chain transfer, unwrap-stake |
| **common** | — | shared base: MultiStore snapshot/revert, balance tracking, gas, mutex |

---

## Part 2 — cosmos-sdk fork (base v0.50.10 → `v0.50.10-sixpatch-3`)

**Headline: nearly all SIX logic lives in `x/staking`.** No changes to x/gov, x/bank, x/auth,
x/tx, codec/unknownproto, or the ante handler — those concerns live in the app repo, not the SDK fork.

1. **Validator "modes"** — new `ValidatorMode` enum + fields (`min_delegation`, `delegation_increment`,
   `max_license`, `license_count`, `mode`, `enable_redelegation`) on `Validator`:
   - `NORMAL` — vanilla.
   - `LICENSE` — permissioned; delegations must be whole multiples of a "license unit", respect min,
     and total licenses ≤ `max_license`. New `LicenseCountInvariant`.
   - `FAST` — only whitelisted delegators; they get **instant unbonding** (self-bond still waits the
     full period so it stays slashable). Cannot convert an existing validator into fast mode.
2. **Validator approval** — new `MsgSetValidatorApproval`; `CreateValidator` requires a pre-granted
   operator address (consumed on use), bootstrapped from gov authority.
3. **Redelegation opt-in** — off by default; tri-state `RedelegationUpdate` on `MsgEditValidator`
   (UNSPECIFIED = leave unchanged). Fast mode can never redelegate.
4. **Whitelist delegators** — `MsgCreate/DeleteWhitelistDelegator` for fast mode; includes a
   **store-key re-keying migration** (`migrations/sixchain`) that must run on upgrade or whitelisted
   delegators silently lose status.
5. **Infra hooks (non-staking):** `store` gains `Copy()` deep-clone on CacheKV/CacheMulti stores
   (plumbing the EVM StateDB snapshot/revert depends on); `baseapp` exposes `StreamingManager()`.

**Key files:** `x/staking/keeper/{msg_server,delegation,whitelist,validator_approval,invariants,genesis}.go`,
`x/staking/migrations/sixchain/store.go`, `store/{cachekv,cachemulti}/store.go`.

> ⚠️ The local `../cosmos-sdk` checkout is on an in-progress `six/v0.53.6` rebase with uncommitted
> staking edits. The above describes shipped `v0.50.10-sixpatch-3`.

---

## Part 3 — evmos v20 fork (base v20.0.0 → `v20.0.0-sixpatch-2`, 80 commits)

Biggest theme: re-implements `x/evm/statedb` + precompile framework + EVM core to
**go-ethereum 1.13.x semantics** (Shanghai/Cancun, EIP-1153/6780, time-based signer), wires SIX
precompiles + native `usix`, and lands the audit fixes.

1. **statedb fixes** — all correspond to prior audit findings:
   - **logIndex bug**: logs now `map[txHash][]*Log` (multi-event txs no longer share an index).
   - **pendingStorage vs transientStorage** separated (was corrupting TSTORE/TLOAD); plus
     **pendingStorage revert-gap** journal entry (the v4.0.4 HIGH blocker).
   - **SubBalance underflow guard** (cosmos/evm #1176) — fail-closed to prevent unbacked `asix`/`usix`
     mint; the 18-decimal restriction deliberately omitted (6-decimal chain).
2. **Precompile framework rewrite** → `Executor` interface; latest `b767c121` adds **journaled
   cosmos-side state for precompiles** (snapshot MultiStore + events, `AddPrecompileFn`) — the fix
   for the ASA-2026-002-class missing-snapshot/revert bug. Plus method-filter / read-only
   write-protection and caller-is-precompile checks.
3. **SIX static precompile registration** — SIX address set (`…1001/1005/1007/1055/1069`) as new
   default; **address-sort fix** (the active list must be sorted); hardened `GetPrecompileInstance`.
4. **EVM params/chain config** — `ExtraEIPs` `[]string`→`[]int32` (geth-numeric); hardforks moved
   from **block numbers → timestamps** (Shanghai/Cancun/Prague/Verkle); signers use block time.
   Migrations v2→v7 registered.
5. **Ante / unordered-tx** — `IncrementNonce` can skip the nonce-equality check for EVM txs;
   **`DefaultEVMUnsafeOrderedTx = true`** ships on by default ⚠️ (the EVM-replay-risk item from the
   2026-08 security review — still as shipped).
6. **state_transition gas refund** reworked so leftover-gas refund can't over-refund / mint new EVM
   coin (same unbacked-mint class as the SubBalance guard).
7. **Server/CLI** — Rosetta removed, tendermint→cometbft naming, native `usix` gas/fee flags,
   `eth_call`/`estimateGas` state-overrides + `debug_traceBlock`.
8. **EIP-712/Web3Tx & 6-decimal**: no logic change here — those live in six-protocol (ante rejects
   Web3Tx; integrations must use `SIGN_MODE_DIRECT`; `usix`/`asix` mapping driven by app config).

---

## Part 4 — go-ethereum fork

> ⚠️ **Version drift:** `go.mod` pins **v1.13.6**, but the local `../go-ethereum` checkout is on
> **v1.17.2** (newer, in-progress). The notes below describe that local v1.17.2 branch (single SIX
> commit `fb40f1f25` "support custom precompile", +874/−146 over 20 files). The actively-built
> v1.13.6 version is slightly different but shares the same intent.

A single coherent feature: turn vanilla geth into an embeddable EVM with **stateful Cosmos
precompiles** (mirrors the cosmos/evm pattern):

1. **`PrecompiledContract` interface redesigned** — `Run(evm *EVM, contract *Contract, readonly bool)`
   + `Address()` + `Name()`. Precompiles now receive the live `*EVM` (→ Cosmos state access),
   caller/value, and the static-call flag. All ~21 builtins mechanically updated.
2. **Mutable per-EVM precompile registry** — `SetPrecompiles`/`WithPrecompiles`/`Precompile(addr)`/
   `ActivePrecompiles()`; all call/create paths use it, so the host app registers Cosmos precompiles
   at runtime. `ValidatePrecompiles` rejects dup/nil/zero-address **and address-vs-`Address()`
   mismatch** (the precompile-address-mismatch concern).
3. **OpCode hooks** — `OpCodeHooks{CallHook, CreateHook}` invoked before every Call/Create path (for
   Cosmos-side access-control / blocklist); `NewEVMWithHooks`.
4. Supporting glue: `Contract.isPrecompile` safety flag (no jumpdest/code on synthetic precompiles),
   `IsStorageEmpty` on StateDB, interpreter accessors, a **custom-EIP/opcode extension API**
   (`ExtendActivators`/`ExtendOperations`), and an EIP-712 zero-field type-encoding fix.
5. No changes to chain `params`, `core/types` tx/signing, gas mechanics, or precompile
   **address ranges** — addresses are entirely deferred to the host app.

> Caveat: this geth layer only snapshots **EVM** state. Cosmos KV writes made by a custom precompile
> are outside geth's snapshot — correct revert depends on the evmos-layer precompile journaling
> (Part 3 #2).

---

## Cross-cutting themes

- **Unbacked-mint safety** — SubBalance guard + gas-refund rework + wrap/unwrap base-denom lock. The
  chain goes to lengths so EVM↔Cosmos balance reconciliation can't mint `usix`/`asix` out of thin air.
- **Precompile snapshot/revert correctness** — the hardest-won fixes (pendingStorage journal,
  cosmos-side precompile journaling) span all three forks; a revert must unwind EVM *and* Cosmos
  state together.
- **Permissioned staking** — modes/license/approval/whitelist/redelegation; the entire cosmos-sdk
  fork exists for this.

## Known items shipped as-is

- `DefaultEVMUnsafeOrderedTx = true` (EVM replay risk flagged in the 2026-08 security review).
- go-ethereum version drift (v1.13.6 pinned in `go.mod`, v1.17.2 checked out locally).
