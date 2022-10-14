package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetBaseDirNames(currentPath string) (baseDirNames *[]string) {
	return splitinternal.GetBaseDirNames(currentPath)
}
