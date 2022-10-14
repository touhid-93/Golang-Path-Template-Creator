package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetSitesEnabled
//
//  returns /etc/nginx/sites-enabled as a string
func GetSitesEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.SitesEnabled.CombineWith(knowndirget.NginxLinuxPath())
}
