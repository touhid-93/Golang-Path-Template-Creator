package normalize

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/core/osconsts"
)

// pathUsingSeparator
//
// Always returns path without the separator at the end.
// Separator must be one char.
// Long path fix will not be applied other than Windows operating system.
func pathUsingSeparator(
	isLongPathFix bool,
	isForceLongPath bool,
	pathSeparator,
	givenPath string,
) string {
	if givenPath == "" {
		return givenPath
	}

	var finalResult string
	if pathSeparator == osconsts.PathSeparator {
		givenPath = unixFix(givenPath)
		finalResult = filepath.Clean(givenPath)
	} else {
		finalResult = removeAndFixDoubleSeparatorToFinalSeparator(
			pathSeparator,
			givenPath)
	}

	if len(finalResult) == 0 {
		return finalResult
	}

	isApplyLongPathFix := isLongPathFix && pathSeparator == WindowsPathSeparator
	if isApplyLongPathFix {
		if finalResult[constants.Zero] == pathSeparator[constants.Zero] {
			finalResult = finalResult[constants.One:]
		}

		finalResult = getLongPathFixedUsingSeparator(
			isForceLongPath,
			pathSeparator,
			finalResult,
		)
	}

	lastIndex := len(finalResult) - 1

	if finalResult[lastIndex] == pathSeparator[coreindexes.First] {
		// removing last path separator
		return finalResult[:lastIndex]
	}

	return finalResult
}
