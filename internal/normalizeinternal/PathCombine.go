package normalizeinternal

import (
	"gitlab.com/evatix-go/core/constants"
)

func PathsCombine(
	root string, combinedPaths []string,
) (first string, allCombinedPaths []string) {
	length := len(combinedPaths)

	results := make(
		[]string,
		len(combinedPaths))

	if length == 0 {
		return constants.EmptyString, results
	}

	for i, currentPath := range combinedPaths {
		combinedPath := JoinFixIf(
			true,
			root,
			currentPath)

		results[i] = combinedPath
	}

	return results[constants.Zero], results
}
