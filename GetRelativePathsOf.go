package pathhelper

import (
	"gitlab.com/evatix-go/core"
)

func GetRelativePaths(
	rootPath string,
	fullPaths ...string,
) *[]string {
	if fullPaths == nil {
		return core.EmptyStringsPtr()
	}

	return GetRelativePathsOfPtr(
		rootPath,
		&fullPaths)
}
