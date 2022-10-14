package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/internal/splitinternal"
)

func GetFilenamePlusBothExtensions(
	currentPath string,
) (
	fileName,
	dotExt,
	ext string,
) {
	return splitinternal.
		GetFilenamePlusBothExtensions(
			currentPath)
}
