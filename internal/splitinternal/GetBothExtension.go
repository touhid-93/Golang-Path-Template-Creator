package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func GetBothExtension(currentPath string) (dotExt, ext string) {
	i := LastSlash(currentPath)
	fileName := currentPath[i+1:]
	dotExt = ""
	ext = ""

	index := strings.LastIndex(
		fileName,
		constants.Dot)

	if index > constants.InvalidNotFoundCase {
		dotExt = fileName[index:]
		ext = fileName[index+1:]
	}

	return dotExt, ext
}
