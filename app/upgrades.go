package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	storetypes "cosmossdk.io/store/types"
	circuittypes "cosmossdk.io/x/circuit/types"
	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensusparamkeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	ratelimittypes "github.com/cosmos/ibc-apps/modules/rate-limiting/v8/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	nftmngrtypes "github.com/thesixnetwork/six-protocol/v4/x/nftmngr/types"

	"github.com/creachadair/tomledit"

	srvmig "github.com/thesixnetwork/six-protocol/v4/server/config/migration"
)

const UpgradeName = "v4.0.0"
const UpgradeNameHotfix = "v4.0.0-hotfix-2" // HOTFIX ONLY ON FIVENET

func (app *App) RegisterUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeName, func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		// First run the standard module migrations
		newVM, err := app.ModuleManager.RunMigrations(ctx, app.configurator, vm)
		if err != nil {
			return newVM, err
		}

		// setup nftmngr params
		app.NftmngrKeeper.SetParams(ctx, nftmngrtypes.DefaultParams())

		// ONLY during upgrade execution: migrate app.toml configuration to v0.50 format
		// This ensures all nodes get the updated configuration automatically
		if err := app.migrateAppConfig(); err != nil {
			app.Logger().Error("Failed to migrate app.toml config", "error", err)
			// Log error but don't fail upgrade - operators can migrate manually if needed
		} else {
			app.Logger().Info("Successfully migrated app.toml to v0.50 format")
		}

		return newVM, nil
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}
	if upgradeInfo.Name == UpgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				crisistypes.StoreKey,
				consensusparamkeeper.StoreKey,
				circuittypes.StoreKey,
				icahosttypes.StoreKey,
				ibcfeetypes.StoreKey,
				icacontrollertypes.StoreKey,

				ratelimittypes.StoreKey,
			},
			Deleted: []string{},
		}
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	// hot fix for fivenet
	app.UpgradeKeeper.SetUpgradeHandler(UpgradeNameHotfix, func(ctx context.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		newVM, err := app.ModuleManager.RunMigrations(ctx, app.configurator, vm)
		if err != nil {
			return newVM, err
		}

		// setup nftmngr params
		app.NftmngrKeeper.SetParams(ctx, nftmngrtypes.DefaultParams())

		return newVM, nil
	})

	if upgradeInfo.Name == UpgradeNameHotfix && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added:   []string{},
			Deleted: []string{},
		}
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

// migrateAppConfig migrates the app.toml configuration file to v0.50 format
func (app *App) migrateAppConfig() error {
	// NOTE:: use DAEMONE_HOME because some node might has different path of home such as /dev/ssd2/chaindata/.six/data
	homeDirs := []string{
		os.Getenv("DAEMON_HOME"),

		// NOTE:: might be use on local development
		"/root/.six",
	}

	var configPath string

	// Find the config file in possible locations
	for _, dir := range homeDirs {
		if dir == "" {
			continue
		}
		testPath := filepath.Join(dir, "config", "app.toml")
		if _, err := os.Stat(testPath); err == nil {
			configPath = testPath
			break
		}
	}

	if configPath == "" {
		return fmt.Errorf("app.toml not found in any of the expected locations: %v", homeDirs)
	}

	backupPath := configPath + ".pre-v4.0.0"
	if _, err := os.Stat(backupPath); err == nil {
		app.Logger().Info("App config migration already completed, backup exists", "backup", backupPath)
		return nil
	}

	if err := copyFile(configPath, backupPath); err != nil {
		return fmt.Errorf("failed to create backup: %w", err)
	}

	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	doc, err := tomledit.Parse(strings.NewReader(string(configBytes)))
	if err != nil {
		return fmt.Errorf("failed to parse current config: %w", err)
	}

	// Apply migration using the existing migration logic
	plan := srvmig.PlanBuilder(doc, "")

	ctx := context.Background()
	for _, step := range plan {
		if err := step.T.Apply(ctx, doc); err != nil {
			return fmt.Errorf("failed to apply migration step '%s': %w", step.Desc, err)
		}
	}

	var buf bytes.Buffer
	if err := tomledit.Format(&buf, doc); err != nil {
		return fmt.Errorf("failed to format migrated config: %w", err)
	}

	// Write to file with proper permissions
	if err := os.WriteFile(configPath, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("failed to write migrated config: %w", err)
	}

	app.Logger().Info("App config migrated successfully", "path", configPath, "backup", backupPath)
	return nil
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Close()
}
