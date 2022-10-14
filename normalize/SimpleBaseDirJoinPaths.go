package normalize

import (
	"gitlab.com/evatix-go/core/osconsts"
)

func SimpleBaseDirJoinPaths(baseDir string, locations ...string) string {
	if len(locations) == 0 {
		return baseDir
	}

	combinedPath := baseDir +
		osconsts.PathSeparator +
		SimpleJoinPaths(locations...)

	return combinedPath
}
