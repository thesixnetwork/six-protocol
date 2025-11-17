# SIX Protocol Go Workspace

This repository uses Go workspaces to manage multiple Go modules in a single repository. The workspace includes the main SIX Protocol blockchain implementation and the SixClient Go SDK.

## 📁 Repository Structure

```
six-protocol/
├── go.work                    # Go workspace configuration
├── go.mod                     # Main six-protocol module
├── Makefile                   # Main blockchain build commands
├── Makefile.workspace         # Workspace management commands
├── test_workspace.go          # Workspace validation script
├── WORKSPACE.md              # This documentation
│
├── app/                      # Blockchain application
├── x/                        # Custom blockchain modules
│   ├── nftmngr/             # NFT management module
│   ├── tokenmngr/           # Token management module
│   └── ...
│
└── sixclient/               # Go SDK for SIX Protocol
    ├── go.mod               # SixClient module dependencies
    ├── README.md            # SixClient documentation
    ├── client.go            # Main client implementation
    ├── bank.go              # Banking operations
    ├── nft.go               # NFT operations
    ├── token.go             # Token operations
    └── examples/            # Usage examples
        └── basic_usage.go
```

## 🚀 Quick Start

### 1. Initialize Workspace

```bash
# Initialize the Go workspace
make -f Makefile.workspace workspace-init

# Or manually
go work init . ./sixclient
```

### 2. Test Workspace Setup

```bash
# Test that everything is working
make -f Makefile.workspace test-workspace

# Or run directly
go run test_workspace.go
```

### 3. Build All Modules

```bash
# Build both main module and sixclient
make -f Makefile.workspace workspace-build
```

## 🔧 Development Commands

### Workspace Management

```bash
# Show help with all available commands
make -f Makefile.workspace help

# Initialize workspace
make -f Makefile.workspace workspace-init

# Sync all dependencies
make -f Makefile.workspace workspace-sync

# Test all modules
make -f Makefile.workspace workspace-test

# Build all modules
make -f Makefile.workspace workspace-build

# Clean all artifacts
make -f Makefile.workspace workspace-clean

# Show workspace status
make -f Makefile.workspace workspace-status
```

### SixClient SDK Commands

```bash
# Test only the sixclient module
make -f Makefile.workspace sixclient-test

# Build only the sixclient module
make -f Makefile.workspace sixclient-build

# Run sixclient examples
export MNEMONIC="your mnemonic here"
make -f Makefile.workspace sixclient-example
```

### Main Blockchain Commands

```bash
# Test main blockchain module
make -f Makefile.workspace main-test

# Build main blockchain module
make -f Makefile.workspace main-build

# Use the main Makefile for blockchain-specific commands
make build        # Build sixd binary
make install      # Install sixd
make test         # Run blockchain tests
```

## 📦 Module Dependencies

### Main Module (`github.com/thesixnetwork/six-protocol`)

- **Purpose**: SIX Protocol blockchain implementation
- **Dependencies**: Cosmos SDK, CometBFT, custom modules
- **Build Target**: `sixd` binary

### SixClient Module (`github.com/thesixnetwork/six-protocol/sixclient`)

- **Purpose**: Go SDK for SIX Protocol interaction
- **Dependencies**: Cosmos SDK types, SIX Protocol modules
- **Usage**: Import as Go library

## 🔄 Dependency Management

The workspace uses Go modules with local replace directives:

```go
// sixclient/go.mod
module github.com/thesixnetwork/six-protocol/sixclient

require (
    github.com/thesixnetwork/six-protocol v0.0.0
    // ... other deps
)

replace github.com/thesixnetwork/six-protocol => ../
```

### Updating Dependencies

```bash
# Update all dependencies
make -f Makefile.workspace deps-update

# Verify all dependencies
make -f Makefile.workspace deps-verify

# Sync workspace after changes
make -f Makefile.workspace workspace-sync
```

## 🧪 Testing

### Workspace Test Script

The `test_workspace.go` script validates:

- ✅ SixClient package imports correctly
- ✅ Client creation with valid/invalid mnemonics
- ✅ Network configurations work
- ✅ Utility functions work
- ✅ Error handling works properly

```bash
# Run workspace validation
go run test_workspace.go

# With your real mnemonic
export MNEMONIC="your twelve word mnemonic"
go run test_workspace.go
```

### Module-Specific Tests

```bash
# Test main blockchain module
make -f Makefile.workspace main-test

# Test sixclient module only
make -f Makefile.workspace sixclient-test

# Test all modules
make -f Makefile.workspace workspace-test
```

## 🛠 Development Setup

### First Time Setup

```bash
# Complete development setup
make -f Makefile.workspace dev-setup

# This runs:
# - workspace-init
# - workspace-sync
# - install-tools
```

### Environment Variables

```bash
# Required for examples and testing
export MNEMONIC="your twelve word mnemonic phrase here"

# Optional: specify network
export NETWORK="testnet"  # or "mainnet", "local"
```

### IDE Configuration

For VS Code, create `.vscode/settings.json`:

```json
{
  "go.useLanguageServer": true,
  "go.buildOnSave": "workspace",
  "go.lintOnSave": "workspace",
  "go.vetOnSave": "workspace",
  "gopls": {
    "experimentalWorkspaceModule": true
  }
}
```

For GoLand/IntelliJ:
1. Open the root directory
2. Go will automatically detect the workspace
3. Enable Go modules in Settings → Go → Go Modules

## 📝 Code Organization

### Import Paths

```go
// Importing the main blockchain types
import (
    "github.com/thesixnetwork/six-protocol/x/nftmngr/types"
    "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
    "github.com/thesixnetwork/six-protocol/app"
)

// Importing the sixclient SDK
import (
    "github.com/thesixnetwork/six-protocol/sixclient"
)
```

### Adding New Modules

1. Create module directory
2. Initialize with `go mod init`
3. Add to `go.work`:
   ```bash
   go work use ./new-module
   ```
4. Add replace directive if needed

## 🔍 Troubleshooting

### Common Issues

#### 1. Module Not Found

```
Error: module not found
```

**Solution:**
```bash
go work sync
go mod tidy
```

#### 2. Import Cycle

```
Error: import cycle not allowed
```

**Solution:** Check module dependencies and restructure imports.

#### 3. Version Conflicts

```
Error: version conflict
```

**Solution:**
```bash
# Clean and rebuild
make -f Makefile.workspace workspace-clean
go clean -modcache
make -f Makefile.workspace workspace-sync
```

### Debug Commands

```bash
# Show module information
go list -m all

# Show workspace modules
go work vendor

# Verify workspace integrity
make -f Makefile.workspace workspace-status
```

### Reset Workspace

```bash
# Complete reset
make -f Makefile.workspace workspace-clean
rm go.work go.work.sum
make -f Makefile.workspace workspace-init
make -f Makefile.workspace workspace-sync
```

## 🏗 Building and Deployment

### Development Build

```bash
# Build all modules for development
make -f Makefile.workspace workspace-build
```

### Release Build

```bash
# Create optimized release build
make -f Makefile.workspace release
```

### Docker Build

```bash
# Build using main Dockerfile
docker build -t six-protocol .

# Build development image with workspace
docker build -f Dockerfile.dev -t six-protocol:dev .
```

## 📚 Documentation

### Generated Documentation

```bash
# Generate API documentation
make -f Makefile.workspace docs
```

### Manual Documentation

- **Main Module**: See blockchain-specific documentation in `/docs`
- **SixClient SDK**: See `sixclient/README.md`
- **Examples**: See `sixclient/examples/`

## 🤝 Contributing

### Workflow

1. **Setup workspace**: `make -f Makefile.workspace dev-setup`
2. **Create feature branch**: `git checkout -b feature/your-feature`
3. **Make changes**: Edit code in appropriate module
4. **Test changes**: `make -f Makefile.workspace workspace-test`
5. **Commit changes**: Follow conventional commits
6. **Submit PR**: Create pull request

### Code Standards

```bash
# Format code
make -f Makefile.workspace format

# Run linters
make -f Makefile.workspace lint

# Run all checks
make -f Makefile.workspace workspace-test
```

### Adding Examples

1. Create example in `sixclient/examples/`
2. Update `examples/README.md`
3. Test example works: `make -f Makefile.workspace sixclient-example`

## 🔗 Related Resources

- **SIX Protocol Documentation**: [docs.sixprotocol.net](https://docs.sixprotocol.net)
- **Cosmos SDK Documentation**: [docs.cosmos.network](https://docs.cosmos.network)
- **Go Workspaces**: [go.dev/ref/mod#workspaces](https://go.dev/ref/mod#workspaces)
- **SixClient SDK**: [sixclient/README.md](sixclient/README.md)

## 📞 Support

- **Discord**: [SIX Protocol Community](https://discord.gg/sixprotocol)
- **Issues**: [GitHub Issues](https://github.com/thesixnetwork/six-protocol/issues)
- **Documentation**: [SIX Protocol Docs](https://docs.sixprotocol.net)

---

**Happy coding with SIX Protocol! 🚀**