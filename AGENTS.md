# AGENTS.md — Working with the SIX Protocol Chain

This repository is the **SIX Protocol blockchain node** (`sixd`): a Cosmos SDK 0.50 chain with full EVM compatibility (via Evmos v20). It runs dual Cosmos/EVM consensus, hosts custom on-chain NFT and token-management modules, and exposes Ethereum-compatible JSON-RPC endpoints. Most agent work will touch Go chain code, custom `x/*` modules, EVM precompiles, or protobuf definitions.

> **Only document what is observed.** If something here is not visible in the code/config, it is omitted intentionally.

---

## Table of Contents

1. [Quick Start](#quick-start)
2. [Project Type & Stack](#project-type--stack)
3. [Essential Commands](#essential-commands)
4. [Repository Layout](#repository-layout)
5. [Application Architecture](#application-architecture)
6. [Custom Modules](#custom-modules)
7. [EVM Precompiles](#evm-precompiles)
8. [Protobuf & Code Generation](#protobuf--code-generation)
9. [Testing Patterns](#testing-patterns)
10. [Local Development & Docker](#local-development--docker)
11. [Conventions & Gotchas](#conventions--gotchas)
12. [Commit Message Convention](#commit-message-convention)
13. [CI/CD & Release](#cicd--release)

---

## Quick Start

```bash
# Build the sixd binary
make build

# Install into $GOPATH/bin
make install

# Run unit tests
make test

# Spin up a local single-node testnet
scripts/test_node.sh
```

Default binary name: `sixd`  
Default node home: `~/.six`

---

## Project Type & Stack

- **Chain framework:** Cosmos SDK v0.50 with CometBFT consensus (`github.com/cometbft/cometbft v0.38.12`).
- **EVM compatibility:** Evmos v20 (`github.com/evmos/evmos/v20`). Provides `x/evm`, `x/feemarket`, Ethereum JSON-RPC, and the EVM state machine.
- **Go version:** 1.25.7 (see `go.mod` and `Dockerfile`).
- **Module path:** `github.com/thesixnetwork/six-protocol/v4`.
- **Protobuf/codegen:** `buf` + `gocosmos` / `grpc-gateway` / `pulsar` plugins; historically scaffolded with Ignite, still referenced in `config.yml` and `proto/` files.
- **Smart-contract toolkit:** Foundry (`contracts/`, root `foundry.toml`).
- **Front-end clients:** Vue, TypeScript, React hooks generated under `vue/` and (by config) `ts-client/`, `react/`, etc.

---

## Essential Commands

| Command | What it does |
|---------|--------------|
| `make build` | Build `build/sixd` with `netgo` + `ledger` tags (Windows build is explicitly blocked). |
| `make install` | `go install ./cmd/sixd` with the same build tags/ldflags. |
| `make test` / `make test-unit` | Run all Go unit tests with `ledger test_ledger_mock` tags. |
| `make test-race` | Race-enabled unit tests. |
| `make test-cover` | Long-running race + coverage run (30 m timeout). |
| `make benchmark` | Run Go benchmarks. |
| `make test-sim-import-export` | App import/export simulation via `runsim`. |
| `make lint` | Run `golangci-lint` + `gofumpt -d -s`. Requires tools installed. |
| `make format` | Run `gofumpt`, `misspell`, and `goimports -local github.com/thesixnetwork/six-protocol` on non-generated Go files. |
| `make proto-all` | Dockerized proto format, lint, generation, then `make format`. |
| `make proto-gen` | Dockerized gogo/pulsar proto generation via `scripts/protocgen.sh`. |
| `make proto-lint` | Dockerized `buf lint`. |
| `make proto-check-breaking` | Dockerized `buf breaking` against `main` branch. |
| `make proto-go` | Run `ignite g proto-go -y`. |
| `make clean` | Remove `build/` and `snapcraft-local.yaml`. |

Build tags used by default: `netgo` and, on non-OpenBSD Unix with `gcc` present, `ledger`. `LEDGER_ENABLED=false` disables ledger support. `WITH_CLEVELDB=yes` adds `gcc` tag and links CLevelDB.

> The `make lint` target passes `--tests=false` to `golangci-lint`; tests are not linted by that rule. `make format` deliberately skips `*.pb.go`, `*.pb.gw.go`, and `*.pulsar.go`.

---

## Repository Layout

```text
.
├── app/                  # Cosmos app wiring, module keepers, ante handlers, genesis, upgrades
├── cmd/sixd/             # sixd CLI entry point
├── x/                    # Custom Cosmos SDK modules
│   ├── nftadmin/         # On-chain authorization / permission registry
│   ├── nftmngr/          # NFT schema, metadata, actions, rule engine
│   ├── nftoracle/        # Oracle proposal/voting for cross-chain NFT state
│   ├── protocoladmin/    # Protocol-level admin permissions
│   └── tokenmngr/        # Token wrapping (usix/asix), conversion, tokenfactory
├── precompiles/          # EVM precompile implementations exposed to Solidity
├── server/               # Custom `start` command, JSON-RPC, indexer wiring
├── client/               # CLI helpers, docs, keys
├── proto/                # Protobuf definitions and buf templates
├── contracts/            # Foundry Solidity contracts/scripts/tests
├── broadcast/            # Deployment/broadcast scripts (Solidity)
├── scripts/              # Chain init, test node, proto generation, test helpers
├── docker/               # Docker helper scripts
├── resources/            # JSON fixtures for schemas, actions, NFT metadata, gov props
├── testutil/             # Keeper test harnesses and sample data
├── docs/                 # Markdown docs including zero-gas oracle concept
└── vue/                  # Vue front-end app
```

### Generated code locations

- Go protobuf outputs are emitted into the same package directories under `x/*/types/` (`*.pb.go`, `*.pb.gw.go`, query/tx/genesis pb files).
- Pulsar outputs land in `api/sixprotocol/...`.
- Swagger/OpenAPI is generated to `docs/static/openapi.yml`.

**Do not hand-edit generated `*.pb.go`, `*.pb.gw.go`, or `*.pulsar.go` files.** Regenerate via `make proto-gen` or `make proto-go`.

---

## Application Architecture

`app/app.go` wires a standard Cosmos SDK 0.50 app with Evmos EVM support.

### Key app constants (`app/app.go:178-184`)

```go
AccountAddressPrefix = "6x"
Name                 = "six"
Bech32Prefix         = "6x"
BaseDenomUnit        = 6
BaseDenom            = "usix"
```

- Cosmos bech32 prefix: `6x`
- Native staking token: `usix` (6 decimals)
- EVM token: `asix` (18 decimals, 1 asix = 10^12 usix)

### Module ordering matters

`app.go` declares module `moduleManager` and `BasicModuleManager` lists. Keeper initialization order and `SetOrderInitGenesis` / `SetOrderExportGenesis` are fixed. When adding a new module you must:

1. Add it to the `App` struct.
2. Initialize it in `New` after its dependencies.
3. Register it in `ModuleManager` and `BasicModuleManager`.
4. Add it to genesis/export order slices.

Several keepers are initialized conditionally for EVM/IBC (e.g. `EVMKeeper`, `IBCKeeper`, `RatelimitKeeper`).

### Dual ante handler

`app/ante/ante.go` routes transactions by extension option:

- `/ethermint.evm.v1.ExtensionOptionsEthereumTx` → `newMonoEVMAnteHandler`
- `/ethermint.types.v1.ExtensionOptionDynamicFeeTx` → `newCosmosAnteHandler`
- Otherwise → normal Cosmos SDK ante handler.

So the chain accepts both Ethereum-format and Cosmos-format transactions; fee logic differs by path.

### Denom / account model

- Cosmos accounts use the `6x` bech32 prefix.
- EVM accounts are derived from the same key material via Evmos `ethsecp256k1`.
- `usix` is the Cosmos fee/staking denom; `asix` is the EVM denom. `x/tokenmngr` provides wrap/unwrap conversion between them.

---

## Custom Modules

All custom modules follow the standard Cosmos SDK `x/<module>/{keeper,types,module,client,simulation}` layout.

### `x/nftmngr` — NFT schema & rule engine

This is the largest custom module. It stores:

- **Schemas** (`NFTSchema`) defining NFT collections, attributes, and actions.
- **Metadata** (`NftData`) per token ID.
- **Actions** (`Action`) with GRL (Grule) rule logic executed on-chain.
- **Executors**, **organizations**, **virtual schemas**, and **proposals**.

Key design points:

- Actions are encoded as base64 JSON inside transaction messages (`Base64NewAction`, `Base64UpdateAction`) and then parsed into `types.Action` in the message server.
- Rule execution uses `github.com/hyperjumptech/grule-rule-engine`. See `x/nftmngr/keeper/engine.go`.
- A **process-local rule library cache** (`ruleLibrary`) avoids re-parsing identical GRL rules. It is guarded by `ruleLibraryMu`. The cache key is derived from the action name + SHA-256 of the rule text. This is safe for consensus because building is deterministic and each execution clones the knowledge base (`NewKnowledgeBaseInstance`).
- `engineMaxCycle` caps rule evaluation at 100 cycles to prevent runaway rules.
- Schema codes and token IDs are used as store prefixes; see `x/nftmngr/types/keys.go` for the byte layout.

### `x/nftadmin`

Authorization registry. Other modules check `NftadminKeeper.HasPermission(ctx, name, addr)` before privileged operations. Admin list is module state, not hard-coded.

### `x/protocoladmin`

Protocol-level admin permissions, similar role but scoped to protocol/token operations.

### `x/tokenmngr`

Wraps/unswaps `usix` ↔ `asix` and supports tokenfactory-style denom creation. Heavy interop with `x/evm` and `x/bank`.

### `x/nftoracle`

Oracle voting/proposals for cross-chain NFT state verification. Used by the broader "zero-gas oracle" concept documented in `docs/zero-gas-oracle/`.

### Module params pattern

Custom modules still use the legacy `x/params` subspace pattern (`ParamKeyTable`, `ParamSetPairs`). The Cosmos SDK deprecates this in favor of `MsgUpdateParams` governance messages, but the codebase has not migrated. New params should follow the existing pattern unless a deliberate migration is undertaken.

---

## EVM Precompiles

The chain exposes several Cosmos modules as Ethereum precompiled contracts. `precompiles/setup.go` initializes them once and mutates global `vm.PrecompiledContracts*` maps.

| Precompile | Address | Solidity interface | Wraps |
|------------|---------|-------------------|-------|
| Bank | `0x000...00001001` | `Bank.sol` | Bank send/balance queries |
| Staking | `0x000...00001005` | `Staking.sol` | Delegate, redelegate, undelegate |
| Distribution | `0x000...00001007` | `Distribution.sol` | Rewards/withdraw address |
| NFT Manager | `0x000...00001055` | `INFTManager.sol` | NFT schema/action/attribute ops |
| Token Factory | `0x000...00001069` | `TokenFactory.sol` | Wrap/unwrap, tokenfactory denom |

### Precompile addresses and static precompiles

For local testnets, `scripts/test_node.sh` populates `app_state.evm.params.active_static_precompiles` with the canonical Evmos addresses plus these custom addresses. On mainnet these must be enabled by governance param update or chain upgrade.

### Adding or changing a precompile

1. Implement the contract in `precompiles/<name>/` following the existing pattern: `NewPrecompile`, `GetABI`, `Address`, `GetName`.
2. Add expected-keeper interfaces in `precompiles/common/expected_keepers.go`.
3. Register it in `precompiles/setup.go:InitializePrecompiles`.
4. Provide/refresh the Solidity interface (`*.sol`) and ABI JSON (`abi.json`).
5. If you change ABI JSON, regenerate any TypeScript bindings.
6. **Critical:** `addPrecompileToVM` mutates global EVM precompile maps. It uses `appendAddressUnique` because duplicate addresses in the global slices cause EVM validation/sorting failures. Be careful in tests that create multiple app instances.

---

## Protobuf & Code Generation

### Tooling

- `buf` (Docker image `bufbuild/buf@sha256:...` pinned in `Makefile`).
- Generation script: `scripts/protocgen.sh`.
- Custom buf templates: `proto/buf.gen.gogo.yaml`, `buf.gen.pulsar.yaml`, etc.
- Ignite still owns `proto/buf.gen.gogo.yaml` (header comment says "auto-generated from Ignite").

### Generating code

```bash
# Regenerate all Go/Pulsar proto code, format proto, lint, then Go format
make proto-all

# Just generate Go/Pulsar
make proto-gen

# Generate using ignite CLI
make proto-go
```

`scripts/protocgen.sh` does the following:

1. Iterates over `proto/` directories.
2. For files with `option go_package` **not** matching `github.com/thesixnetwork/six-protocol/api`, runs `buf generate --template buf.gen.gogo-nonignite.yaml <file>`.
3. Runs `buf generate --template buf.gen.pulsar-nonignite.yaml` for pulsar outputs.
4. Copies generated `github.com/thesixnetwork/six-protocol/*` into the repo root and removes the `github.com/` directory.
5. Runs `go mod tidy`.

### Proto package layout

Proto packages are under `proto/sixprotocol/<module>/`:

- `genesis.proto`
- `params.proto`
- `query.proto`
- `tx.proto`
- Domain-specific files (`action.proto`, `nft_schema.proto`, etc.)

When adding a message/query/tx to a module, update the corresponding `*.proto` files and regenerate. Do not add messages by only editing `.pb.go` files.

---

## Testing Patterns

### Unit tests

- Standard `go test` with `-mod=readonly` and tags `ledger test_ledger_mock`.
- Keepers are tested with in-memory `dbm.NewMemDB()` stores via `testutil/keeper/<module>.go` helpers.
- Example: `testutil/keeper/nftmngr.go` constructs a minimal `keeper.Keeper` with `nil` external keepers for unit-level tests.

### Integration / network tests

- `testutil/network/` exists for Cosmos SDK in-process network tests.
- `evm-migration-tests/` holds Go-based EVM integration tests with its own `Makefile`. Run via `make` in that directory or `run_evm_tests.sh`.

### Simulations

- Each custom module has a `simulation/` package.
- App-level simulations are in `app/sim_test.go` and `app/sim_bench_test.go`.
- `make test-sim-import-export` and `make test-sim-multi-seed-short` require the `runsim` tool (installed via `contrib/devtools/Makefile`).

### EVM tests

- Foundry tests live in `contracts/test/`.
- Run with `forge test` from the repo root (uses root `foundry.toml`) or inside `contracts/`.

---

## Local Development & Docker

### One-liner local testnet

```bash
CHAIN_ID="localchain_9000-1" HOME_DIR="~/.sixprotocol" BLOCK_TIME="5s" CLEAN=true sh scripts/test_node.sh
```

The script:

- Installs `sixd` (`make install`).
- Wipes the node home (if `CLEAN=true`).
- Adds two test keys from fixed mnemonics (`acc0`, `acc1`).
- Customizes `genesis.json` for EVM precompiles, fee market, staking denom, etc.
- Starts the node.

### Docker

```bash
docker build -t sixnode .
```

- `Dockerfile` builds a static-ish `sixd` with `LEDGER_ENABLED=false BUILD_TAGS=muslc` on `golang:1.25.7-bookworm`.
- Runtime image is `cgr.dev/chainguard/wolfi-base` with `tzdata` set to `Asia/Bangkok`.
- Exposes ports: `1317` (REST), `26656` (P2P), `26657` (RPC). Other EVM/JSON-RPC/gRPC ports can be published as needed.
- `docker-compose.yml` and `docker-compose.db.yml` orchestrate a multi-node or DB-backed setup.
- `docker/*.sh` scripts handle reset, init, cosmovisor setup, and startup.

### Useful node ports

- Tendermint RPC: `26657`
- Cosmos REST: `1317`
- gRPC: `9090`
- gRPC-Web: `9091`
- Ethereum JSON-RPC: `8545` (default Evmos config)

---

## Conventions & Gotchas

### Go code style

- `gofumpt` + `goimports` with `-local github.com/thesixnetwork/six-protocol` grouping.
- Standard groups: stdlib → third-party → `github.com/thesixnetwork/six-protocol/...`.
- `misspell` is run by `make format`.
- Build with `-mod=readonly` everywhere.

### Keeper interface pattern

Each module's `types/expected_keepers.go` defines narrow interfaces for dependencies. This is how `x/nftmngr` depends on `x/nftadmin` without importing the concrete keeper. When you add a cross-module dependency, define the interface there and pass the real keeper at app-wiring time.

### Module authority

Custom module keepers receive an `authority string` (usually the governance module account). Many state-changing messages require `creator` signer == authority or a permissioned admin from `nftadmin`/`protocoladmin`.

### Bech32 and EVM address conversion

- Cosmos address: `6x1...`
- Evmos-derived Ethereum address: `0x...`
- Use `sixutils` helpers or Evmos `evmostypes` address utilities; do not hand-roll conversion logic.

### EVM denom note

The EVM side uses `asix` as `evm_denom`. The fee market `base_fee` is in wei-scale (`100000000000` = 100 Gwei in `asix` units). Check `config.yml` and `scripts/test_node.sh` for local defaults.

### Vote extensions

Local genesis enables vote extensions at height 1 (`consensus.params.abci.vote_extensions_enable_height="1"`).

### `attosix` / `usix` / `asix` naming

- `usix` = micro SIX, Cosmos base unit (10^0 for Cosmos, 10^-6 display).
- `asix` = atto SIX, EVM base unit (10^0 for EVM, 10^-18 display).
- `asix` is also the on-chain denom name; 1 `asix` = 10^12 `usix`.

When converting, `x/tokenmngr` has `AttoCoinConverter`. Be careful which unit an amount is in before sending to bank/evm.

### Solidity / Foundry

- `foundry.toml` uses solc `0.8.20`, optimizer on, 200 runs.
- `contracts/lib/` for dependencies; `contracts/src/` for source; `contracts/test/` for tests.
- `remappings.txt` and `contracts/remappings.txt` should stay in sync.
- The root-level `broadcast/` directory contains deployment scripts (`*.s.sol`) that may reference precompile addresses.

### Zero-gas oracle concept

There is extensive documentation under `docs/zero-gas-oracle/` and a sample script under `scripts/zero-gas-oracle-voting/`. This is an off-chain/on-chain hybrid: relayers submit oracle votes that are validated without normal gas fees for authorized operations.

---

## Commit Message Convention

The repository follows the convention in `git-commit-convention.md`.

Format:

```text
<type>(<scope>): <short summary>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `docs-api`, `style`, `refactor`, `perf`, `test`, `chore`, `ci`, `proto`, `revert`, `rename`.

Subject rules:

- Imperative, present tense.
- Lowercase first word.
- No trailing period.
- Aim for < 50 characters.

Examples:

```text
feat(nftmngr): add virtual schema proposal voting

fix(precompiles): guard against duplicate precompile addresses
test(tokenmngr): add wrap/unwrap edge cases
proto(nftmngr): regenerate pb files for new attribute fields
```

---

## CI/CD & Release

### `.github/workflows/build.yml`

Triggered on PRs, merge groups, and pushes to `main`/`development`.

1. Builds `sixd` for `linux/amd64` (`GOOS=linux GOARCH=amd64 make build`).
2. Verifies `./build/sixd version`.
3. Builds a Docker image and runs Trivy vulnerability scans (Critical/High/Medium).
4. Parses scan results and sends a summary to Google Chat.

### `.github/workflows/release.yml`

Triggered on GitHub releases with tags `v*`.

1. Builds a statically linked `linux/amd64` binary in `golang:1.25-alpine` with musl/static ldflags.
2. Uploads the binary as a GitHub artifact and to GCS (`secrets.SIXNET_BINARIES_BUCKET`).

### Branching

Primary branches: `main` and `development`. Feature work targets `development`; releases cut from `main`.

---

## Working Checklist for Agents

Before submitting changes:

- [ ] `make build` succeeds.
- [ ] `make test` passes (or the targeted subset of tests).
- [ ] If you touched `.proto` files, ran `make proto-gen` and committed generated files.
- [ ] If you touched Go source, ran `make format`.
- [ ] If you added a precompile, registered it in `precompiles/setup.go` and updated Solidity interfaces.
- [ ] If you changed module state, added/updated keeper tests via `testutil/keeper/<module>.go`.
- [ ] If you changed EVM behavior, ran the `evm-migration-tests` or Foundry contract tests.
- [ ] Commit messages follow `git-commit-convention.md`.
