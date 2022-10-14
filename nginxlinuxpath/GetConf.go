package nginxlinuxpath

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/knowndir"
	"gitlab.com/evatix-go/pathhelper/knowndirget"
)

// GetConf returns /etc/nginx/conf.d as a string
func GetConf() string {
	if osconsts.IsWindows {
		return constants.EmptyString
	}

	return knowndir.Conf.CombineWith(knowndirget.NginxLinuxPath())
}
