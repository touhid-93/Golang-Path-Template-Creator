package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func GetFilenamePlusExt(currentPath string) (fileName string, ext string) {
	i := LastSlash(currentPath)
	fileName = currentPath[i+1:]
	ext = ""

	index := strings.LastIndex(
		fileName,
		constants.Dot)

	if index > constants.InvalidNotFoundCase {
		ext = fileName[index+1:]
	}

	return fileName, ext
}
