package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetModsAvailable returns /etc/apache/mods-available as a string
func GetModsAvailable() string {
	if osconsts.IsWindows {
		panic("Location only available for Unix OS")
	}

	return knowndir.ModsAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
