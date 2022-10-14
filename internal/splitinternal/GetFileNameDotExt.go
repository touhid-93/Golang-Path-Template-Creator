package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func GetFileNameDotExt(currentPath string) (fileName string, dotExt string) {
	i := LastSlash(currentPath)
	fileName = currentPath[i+1:]
	dotExt = ""

	index := strings.LastIndex(
		fileName,
		constants.Dot)

	if index > constants.InvalidNotFoundCase {
		dotExt = fileName[index:]
	}

	return fileName, dotExt
}
