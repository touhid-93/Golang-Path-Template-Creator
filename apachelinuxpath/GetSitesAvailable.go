package apachelinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetSitesAvailable returns /etc/apache/sites-available as a string
func GetSitesAvailable() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.SitesAvailable.CombineWith(knowndirget.ApacheLinuxPath())
}
