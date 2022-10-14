package knowndirget

import (
	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// Returns Music directory path as a string.
func UserMusicPath() string {
	return knowndir.Music.CombineWith(UserPath())
}
