package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetModulesAvailable
//
//  returns /etc/nginx/modules-available as a string
func GetModulesAvailable() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModulesAvailable.CombineWith(
		knowndirget.NginxLinuxPath())
}
