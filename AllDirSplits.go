package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func AllDirSplits(currentPath string) (baseDirNames *[]string) {
	return splitinternal.GetAllSplits(currentPath)
}
