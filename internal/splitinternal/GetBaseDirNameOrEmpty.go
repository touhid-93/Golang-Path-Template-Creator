package splitinternal

import "gitlab.com/evatix-go/core/constants"

// GetBaseDirNameOrEmpty reference example : https://play.golang.org/p/BJRR0Wk7GhJ
func GetBaseDirNameOrEmpty(currentPath string) (baseDirName string) {
	baseDirPath := GetBaseDir(currentPath)

	i := LastSlash(
		baseDirPath)

	if i > constants.InvalidNotFoundCase {
		return baseDirPath[i+1:]
	}

	return constants.EmptyString
}
