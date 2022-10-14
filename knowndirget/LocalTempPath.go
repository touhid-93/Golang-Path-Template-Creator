package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to local temp directory. If directory doesn't exist it still returns the path as a string.
func LocalTempPath() string {
	if osconsts.IsWindows {
		return knowndir.LocalTempWin.CombineWith(AppDataPath())
	}

	return knowndir.LocalTempUnix.CombineWith(UnixRoot())
}
