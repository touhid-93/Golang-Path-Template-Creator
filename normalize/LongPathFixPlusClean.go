package normalize

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func LongPathFixPlusClean(
	isForceLongPathFix bool,
	givenPath string,
) string {
	if givenPath == "" {
		return givenPath
	}

	givenPath = TrimPrefixUncPath(
		givenPath)
	givenPath = unixFix(givenPath)
	finalResult := filepath.Clean(givenPath)
	if len(finalResult) == 0 {
		return finalResult
	}

	if osconsts.IsWindows {
		if finalResult[constants.Zero] == PathSeparatorChar {
			finalResult = finalResult[constants.One:]
		}

		finalResult = getLongPathFixedUsingSeparator(
			isForceLongPathFix,
			osconsts.PathSeparator,
			finalResult,
		)
	}

	lastIndex := len(finalResult) - 1

	if finalResult[lastIndex] == PathSeparatorChar {
		// removing last path separator
		return finalResult[:lastIndex]
	}

	return finalResult
}
