package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

// GetBaseDir No slash at the end
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetBaseDir(currentPath string) (baseDir string) {
	return splitinternal.GetBaseDir(
		currentPath)
}
