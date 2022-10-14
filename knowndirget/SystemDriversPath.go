package knowndirget

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns path to System drivers directory on different platforms.
func GetSystemDriversPath() string {
	if osconsts.IsWindows {
		return knowndir.Drivers.CombineWith(System32())
	}

	return knowndir.DriversUnix.Value()
}
