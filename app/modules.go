package app

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	protocoladminmoduletypes "github.com/thesixnetwork/six-protocol/x/protocoladmin/types"
	tokenmngrmoduletypes "github.com/thesixnetwork/six-protocol/x/tokenmngr/types"
)

func orderBeginBlockers() []string {
	return []string{
		// Upgrades should be run _very_ first
		upgradetypes.ModuleName,
		// Note: epochs' begin should be "real" start of epochs, we keep epochs beginblock at the beginning
		capabilitytypes.ModuleName,
		minttypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		ibchost.ModuleName,
		ibctransfertypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName, vestingtypes.ModuleName,
		feegrantmodule.ModuleName,
		wasm.ModuleName,
		protocoladminmoduletypes.ModuleName,
		tokenmngrmoduletypes.ModuleName,
		// superfluid must come after distribution and epochs
	}
}

var orderEndBlockers = []string{
	crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
	capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName, distrtypes.ModuleName,
	slashingtypes.ModuleName, minttypes.ModuleName,
	genutiltypes.ModuleName, evidencetypes.ModuleName, authz.ModuleName,
	paramstypes.ModuleName, upgradetypes.ModuleName, vestingtypes.ModuleName,
	ibchost.ModuleName, ibctransfertypes.ModuleName,
	feegrantmodule.ModuleName,
	wasm.ModuleName,
	protocoladminmoduletypes.ModuleName,
	tokenmngrmoduletypes.ModuleName,
	// Note: epochs' endblock should be "real" end of epochs, we keep epochs endblock at the end
}

var orderInitGenesis = []string{
	capabilitytypes.ModuleName,
	authtypes.ModuleName,
	banktypes.ModuleName,
	distrtypes.ModuleName,
	stakingtypes.ModuleName,
	slashingtypes.ModuleName,
	govtypes.ModuleName,
	minttypes.ModuleName,
	crisistypes.ModuleName,
	ibchost.ModuleName,
	genutiltypes.ModuleName,
	evidencetypes.ModuleName,
	ibctransfertypes.ModuleName,
	paramstypes.ModuleName, upgradetypes.ModuleName, vestingtypes.ModuleName,
	feegrantmodule.ModuleName,
	wasm.ModuleName,
	protocoladminmoduletypes.ModuleName,
	tokenmngrmoduletypes.ModuleName,
}
