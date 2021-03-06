package upgrade

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"
)

const (
	RouterKey  = upgrade.RouterKey
	ModuleName = upgrade.ModuleName
	StoreKey   = upgrade.StoreKey
)

var (
	// function aliases
	NewAppModule                      = upgrade.NewAppModule
	NewSoftwareUpgradeProposalHandler = upgrade.NewSoftwareUpgradeProposalHandler
)

type (
	AppModuleBasic = upgrade.AppModuleBasic
)
