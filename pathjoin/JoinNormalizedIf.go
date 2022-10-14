package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
)

// JoinNormalizedIf normalized applied auto
func JoinNormalizedIf(isNormalize bool, path1, path2 string) string {
	if isNormalize {
		return JoinNormalized(path1, path2)
	}

	finalPath := path1 +
		osconsts.PathSeparator +
		path2

	return finalPath
}
