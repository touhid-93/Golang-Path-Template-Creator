package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Pictures directory path as a string.
func UserPicturesPath() string {
	return knowndir.Pictures.CombineWith(UserPath())
}
