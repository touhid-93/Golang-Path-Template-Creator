package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// FontsPath
//
// Returns path to Fonts directory on different platforms.
func FontsPath() string {
	if osconsts.IsWindows {
		return knowndir.Fonts.CombineWith(WidowsDirectory())
	}

	return knowndir.FontsUnix.CombineWith(UnixRoot())
}
