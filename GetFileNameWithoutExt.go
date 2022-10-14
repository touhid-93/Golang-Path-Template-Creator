package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

// GetFileNameWithoutExt invalid ext should return empty string.
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithoutExt(currentPath string) (filename string) {
	return splitinternal.GetFileNameWithoutExt(
		currentPath)
}
