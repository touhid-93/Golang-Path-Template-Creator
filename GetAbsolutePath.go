package pathhelper

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/errcore"

	"gitlab.com/evatix-go/pathhelper/ispath"
)

func GetAbsolutePath(
	basePath,
	relativePath string,
	isLongPathFix, isNormalize bool,
) string {
	if ispath.Empty(basePath) || ispath.Empty(relativePath) {
		panic(errcore.InvalidEmptyPathType)
	}

	return GetCombinedPath(
		constants.PathSeparator,
		false,
		isLongPathFix,
		isNormalize,
		basePath,
		relativePath,
	)
}
