package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func TrimReplacePrefixRoot(
	isNormalize,
	isLongPathForceFix bool,
	source,
	trimRootPrefix,
	rootReplacer string,
) string {
	relativeWithoutRoot := normalize.TrimPrefixRoot(
		source,
		trimRootPrefix)

	joined := JoinSimple(
		rootReplacer,
		relativeWithoutRoot)

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
