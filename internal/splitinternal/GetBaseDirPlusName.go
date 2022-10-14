package splitinternal

import "gitlab.com/evatix-go/core/constants"

// GetBaseDirPlusName
//
// No slash at the end
// reference example : https://play.golang.org/p/BJRR0Wk7GhJ
func GetBaseDirPlusName(currentPath string) (baseDir, baseDirName string) {
	baseDirPath := GetBaseDir(currentPath)

	i := LastSlash(
		baseDirPath)

	if i > constants.InvalidNotFoundCase {
		return baseDirPath, baseDirPath[i+1:]
	}

	return baseDirPath, constants.EmptyString
}
