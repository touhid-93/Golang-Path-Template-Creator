package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func TrimReplacePrefixRootJoinSuffix(
	isNormalize,
	isLongPathForceFix bool,
	source,
	trimRootPrefix,
	rootReplacer,
	suffixJoin string,
) string {
	relativeWithoutRoot := normalize.TrimPrefixRoot(
		source,
		trimRootPrefix)

	joined := JoinSimpleThree(
		rootReplacer,
		relativeWithoutRoot,
		suffixJoin)

	if !isNormalize && !isLongPathForceFix {
		return joined
	}

	return normalize.PathUsingSeparatorIf(
		isLongPathForceFix,
		true,
		isNormalize,
		osconsts.PathSeparator,
		joined)
}
