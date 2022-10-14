package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func GetExtensions(
	currentPath string,
) (
	dotExt, ext string,
) {
	return splitinternal.GetBothExtension(
		currentPath)
}
