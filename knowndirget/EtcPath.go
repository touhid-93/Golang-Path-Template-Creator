package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// EtcPath
//
// Returns path to etc directory on different platforms.
func EtcPath() string {
	if osconsts.IsWindows {
		return knowndir.Etc.CombineWith(GetSystemDriversPath())
	}

	return knowndir.Etc.CombineWith(UnixRoot())
}
