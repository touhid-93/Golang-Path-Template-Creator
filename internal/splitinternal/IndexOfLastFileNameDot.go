package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// returns -1 if last file name doesn't have any extension
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func IndexOfDotFromLastFileName(currentPath string) (fileName string, index int) {
	i := LastSlash(currentPath)
	fileName = currentPath[i+1:]

	return fileName, strings.LastIndex(
		fileName,
		constants.Dot)
}
