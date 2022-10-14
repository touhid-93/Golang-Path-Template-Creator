package splitinternal

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// GetFileNameWithoutExt
//
// invalid ext should return empty string.
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetFileNameWithoutExt(currentPath string) (filename string) {
	i := LastSlash(
		currentPath)

	if i <= constants.Zero {
		indexOfDot := strings.Index(
			currentPath,
			constants.Dot)

		if indexOfDot > -1 {
			return currentPath[:indexOfDot]
		}

		return currentPath
	}

	filename = currentPath[i+1:]
	indexOfDot := strings.Index(
		filename,
		constants.Dot)

	if indexOfDot > constants.InvalidNotFoundCase {
		return filename[:indexOfDot]
	}

	return filename
}
