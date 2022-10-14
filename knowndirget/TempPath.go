package knowndirget

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/internal/ispathinternal"
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns temp directory. After checking in on all possible locations, if directory doesn't exist, then doesnt create the directory
// and returns the path as a string.
func TempPath() string {
	if osconsts.IsWindows {
		return os.Getenv(knowndir.Temp.Value())
	}

	desiredTempPathUnix := knowndir.TempDir.CombineWith(UserPath())

	// Checking if temp directory is available
	tempUnix := os.Getenv(knowndir.TempDir.Value())

	if ispathinternal.AllNotExists(desiredTempPathUnix, tempUnix) {
		return desiredTempPathUnix
	}

	return tempUnix
}
