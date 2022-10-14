package pathjoin

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinFixWithoutLongPath
//
// Items are applied with
// normalize,
// skip on empty path given
func JoinFixWithoutLongPath(
	baseDir string,
	paths ...string,
) string {
	if len(paths) == 0 && baseDir == "" {
		return constants.EmptyString
	}

	joined := JoinBaseDirWithSep(
		true,
		false,
		false,
		osconsts.PathSeparator,
		baseDir,
		paths...)

	return normalize.PathFixWithoutLongPathIf(
		true,
		joined)
}
