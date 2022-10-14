package pathjoin

import (
	"path/filepath"

	"gitlab.com/evatix-go/core/osconsts"
)

// JoinSimpleIf doesn't apply normalize
func JoinSimpleIf(isUseLibraryFunc bool, path1, path2 string) string {
	if isUseLibraryFunc {
		return filepath.Join(path1, path2)
	}

	return path1 + osconsts.PathSeparator + path2
}
