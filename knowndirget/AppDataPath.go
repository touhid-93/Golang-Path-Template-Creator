package knowndirget

import (
	"os"

	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// AppDataPath Returns path to %AppData% in windows. If directory doesn't exist it still returns the path as a string.
// Requires further investigation for linux platform: https://stackoverflow.com/questions/1510104/where-to-store-application-data-non-user-specific-on-linux,
// https://stackoverflow.com/questions/17517131/appdata-in-non-windows
func AppDataPath() string {
	if osconsts.IsWindows {
		return os.Getenv(knowndir.AppData.Value())
	}

	return knowndir.AppDataUnix.Value()
}
