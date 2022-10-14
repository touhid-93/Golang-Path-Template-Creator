package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetModulesEnabled
//
//  returns /etc/nginx/modules-enabled as a string
func GetModulesEnabled() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.ModulesEnabled.CombineWith(knowndirget.NginxLinuxPath())
}
