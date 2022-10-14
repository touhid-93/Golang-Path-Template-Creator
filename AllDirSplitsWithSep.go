package pathhelper

import "strings"

func AllDirSplitsWithSep(currentPath, separator string) (baseDirNames *[]string) {
	splits := strings.Split(currentPath, separator)

	return &splits
}
