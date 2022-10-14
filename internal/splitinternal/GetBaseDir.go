package splitinternal

import (
	"gitlab.com/evatix-go/core/constants"
)

// GetBaseDir
//
// No slash at the end
// reference example : https://play.golang.org/p/BJRR0Wk7GhJ
func GetBaseDir(currentPath string) (baseDir string) {
	i := LastSlash(
		currentPath)

	if i > constants.InvalidNotFoundCase {
		return currentPath[:i]
	}

	return currentPath
}
