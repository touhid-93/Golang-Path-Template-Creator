package pathjoin

import "gitlab.com/evatix-go/core/osconsts"

// JoinNormalizedExpandIf normalized and expand applied if condition meets
func JoinNormalizedExpandIf(
	isExpandNormalize bool, path1, path2 string,
) string {
	if !isExpandNormalize {
		finalPath := path1 +
			osconsts.PathSeparator +
			path2

		return finalPath
	}

	return JoinNormalizedExpand(path1, path2)
}
