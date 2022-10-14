package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns downloads directory path as a string.
func UserDownloadsPath() string {
	return knowndir.Downloads.CombineWith(UserPath())
}
