package pathhelper

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/pathjoin"
)

func PathCombineWithBase(
	isNormalize,
	isExpandEnv bool,
	baseDir,
	relativePath string,
) string {
	combinedPath := pathjoin.JoinNormalizedIf(
		isNormalize,
		baseDir,
		relativePath)

	if isExpandEnv {
		combinedPath = expandpath.ExpandVariables(combinedPath)
	}

	return combinedPath
}
