package splitinternal

import "strings"

func GetAllSplitsWithSep(
	currentPath, separator string,
) (baseDirNames *[]string) {
	splits := strings.Split(
		currentPath,
		separator)

	return &splits
}
