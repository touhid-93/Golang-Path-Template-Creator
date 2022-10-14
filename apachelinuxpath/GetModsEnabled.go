package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetModsEnabled returns /etc/apache/mods-enabled as a string
func GetModsEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModsEnabled.CombineWith(knowndirget.ApacheLinuxPath())
}
