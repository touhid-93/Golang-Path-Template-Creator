package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinConditionalNormalizedExpandIf normalized or expand or both apply based on condition
func JoinConditionalNormalizedExpandIf(
	isNormalizePlusLongPathFix,
	isExpand bool,
	path1, path2 string,
) string {
	combined := JoinSimple(
		path1,
		path2,
	)

	expand := expandpath.ExpandVariablesIf(
		isExpand,
		combined)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
