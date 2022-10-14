package pathhelper

import "gitlab.com/evatix-go/pathhelper/internal/splitinternal"

func GetBaseDirOrEmpty(currentPath string) (baseDir string) {
	return splitinternal.GetBaseDirNameOrEmpty(
		currentPath)
}
