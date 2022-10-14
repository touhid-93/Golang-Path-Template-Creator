package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to local directory in windows. Otherwise returns Users directory.
func LocalPath() string {
	if osconsts.IsWindows {
		return knowndir.Local.CombineWith(AppDataPath())
	}

	return UserPath()
}
