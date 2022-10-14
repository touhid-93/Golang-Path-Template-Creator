package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinIf isNormalizePlusLongPathFix if true then for windows add UNC Location fix
func JoinIf(isNormalizePlusLongPathFix bool, path1, path2 string) string {
	finalPath := path1 + osconsts.PathSeparator + path2

	if !isNormalizePlusLongPathFix {
		return finalPath
	}

	return normalize.Path(finalPath)
}
