package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetConfEnabled returns /etc/apache/conf-enabled as a string
func GetConfEnabled() string {
	if osconsts.IsWindows {
		panic("Location only available for Unix OS")
	}

	return knowndir.ConfEnabled.CombineWith(knowndirget.ApacheLinuxPath())
}
