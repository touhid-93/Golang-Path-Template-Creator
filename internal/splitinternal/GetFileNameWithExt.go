package splitinternal

import "gitlab.com/evatix-go/core/constants"

// GetName
//
// Can be dir name, file name with ext
//
// reference example : https://play.golang.org/p/oT6eWNZAeEi
func GetName(currentPath string) (fileName string) {
	i := LastSlash(
		currentPath)

	if i <= constants.Zero {
		return currentPath
	}

	return currentPath[i+1:]
}
