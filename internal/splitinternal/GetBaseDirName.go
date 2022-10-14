package splitinternal

import "gitlab.com/evatix-go/core/constants"

func GetBaseDirName(currentPath string) (baseDirName string) {
	baseDirPath := GetBaseDir(currentPath)

	i := LastSlash(
		baseDirPath)

	if i > constants.InvalidNotFoundCase {
		return baseDirPath[i+1:]
	}

	return currentPath
}
