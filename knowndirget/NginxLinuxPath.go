package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// "/etc/nginx/"
func NginxLinuxPath() string {
	if osconsts.IsWindows {
		return ""
	}

	return knowndir.NginxLinuxPath.Value()
}
