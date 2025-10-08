package precompiles

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ecommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/thesixnetwork/six-protocol/precompiles/bank"
	"github.com/thesixnetwork/six-protocol/precompiles/common"
	"github.com/thesixnetwork/six-protocol/precompiles/distribution"
	"github.com/thesixnetwork/six-protocol/precompiles/nftmngr"
	"github.com/thesixnetwork/six-protocol/precompiles/staking"
	"github.com/thesixnetwork/six-protocol/precompiles/tokenfactory"
)

var (
	SetupMtx    = &sync.Mutex{}
	Initialized = false
)

type PrecompileInfo struct {
	ABI     abi.ABI
	Address ecommon.Address
}

// PrecompileNamesToInfo is Populated by InitializePrecompiles
var PrecompileNamesToInfo = map[string]PrecompileInfo{}

type IPrecompile interface {
	vm.PrecompiledContract
	GetABI() abi.ABI
	GetName() string
	Address() ecommon.Address
}

func InitializePrecompiles(
	dryRun bool,
	cdc codec.BinaryCodec,
	bankKeeper common.BankKeeper,
	accountKeeper common.AccountKeeper,
	tokenmngrKeeper common.TokenmngrKeeper,
	tokenmngrMsgServer common.TokenmngrMsgServer,
	nftmngrKeeper common.NftmngrKeeper,
	stakingKeeper common.StakingMsgServer,
	stakingQuerier common.StakingQuerier,
	distrKeeper common.DistributionKeeper,
	distrQuerier common.DistributionQuerier,
) error {
	SetupMtx.Lock()
	defer SetupMtx.Unlock()
	// if Initialized {
	// 	panic("precompiles already initialized")
	// }
	bankp, err := bank.NewPrecompile(bankKeeper)
	if err != nil {
		return err
	}
	tokenfactoryp, err := tokenfactory.NewPrecompile(bankKeeper, accountKeeper, tokenmngrKeeper, tokenmngrMsgServer)
	if err != nil {
		return err
	}

	nftmngrp, err := nftmngr.NewPrecompile(nftmngrKeeper, accountKeeper, bankKeeper)
	if err != nil {
		return err
	}

	stakingp, err := staking.NewPrecompile(stakingKeeper, stakingQuerier, bankKeeper, tokenmngrKeeper)
	if err != nil {
		return err
	}

	distributionp, err := distribution.NewPrecompile(distrKeeper, tokenmngrKeeper)
	if err != nil {
		return err
	}

	PrecompileNamesToInfo[bankp.GetName()] = PrecompileInfo{ABI: bankp.GetABI(), Address: bankp.Address()}
	PrecompileNamesToInfo[tokenfactoryp.GetName()] = PrecompileInfo{ABI: tokenfactoryp.GetABI(), Address: tokenfactoryp.Address()}
	PrecompileNamesToInfo[nftmngrp.GetName()] = PrecompileInfo{ABI: nftmngrp.GetABI(), Address: nftmngrp.Address()}
	PrecompileNamesToInfo[stakingp.GetName()] = PrecompileInfo{ABI: stakingp.GetABI(), Address: stakingp.Address()}
	PrecompileNamesToInfo[distributionp.GetName()] = PrecompileInfo{ABI: distributionp.GetABI(), Address: distributionp.Address()}

	if !dryRun {
		addPrecompileToVM(bankp)
		addPrecompileToVM(tokenfactoryp)
		addPrecompileToVM(nftmngrp)
		addPrecompileToVM(stakingp)
		addPrecompileToVM(distributionp)
		Initialized = true
	}

	return nil
}

func GetPrecompileInfo(name string) PrecompileInfo {
	if !Initialized {
		// Precompile Info does not require any keeper state
		_ = InitializePrecompiles(true, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	}
	i, ok := PrecompileNamesToInfo[name]
	if !ok {
		panic(name + "doesn't exist as a precompile")
	}
	return i
}

// This function modifies global variable in `vm` module. It should only be called once
// per precompile during initialization
func addPrecompileToVM(p IPrecompile) {
	vm.PrecompiledContractsHomestead[p.Address()] = p
	vm.PrecompiledContractsByzantium[p.Address()] = p
	vm.PrecompiledContractsIstanbul[p.Address()] = p
	vm.PrecompiledContractsBerlin[p.Address()] = p
	vm.PrecompiledContractsCancun[p.Address()] = p
	vm.PrecompiledContractsBLS[p.Address()] = p
	vm.PrecompiledAddressesHomestead = append(vm.PrecompiledAddressesHomestead, p.Address())
	vm.PrecompiledAddressesByzantium = append(vm.PrecompiledAddressesByzantium, p.Address())
	vm.PrecompiledAddressesIstanbul = append(vm.PrecompiledAddressesIstanbul, p.Address())
	vm.PrecompiledAddressesBerlin = append(vm.PrecompiledAddressesBerlin, p.Address())
	vm.PrecompiledAddressesCancun = append(vm.PrecompiledAddressesCancun, p.Address())
}
