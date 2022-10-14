package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Videos directory path as a string.
func UserVideosPath() string {
	return knowndir.Videos.CombineWith(UserPath())
}
