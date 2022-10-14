package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func LinuxTouchLocationFile(parentPath, fileName string) *errorwrapper.Wrapper {
	joinedPath := pathjoin.JoinNormalized(
		parentPath,
		fileName)

	return LinuxTouchFile(joinedPath)
}
