package pathjoin

import (
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinConditionalNormalizedThreeExpandIf normalized or expand or both apply based on condition
func JoinConditionalNormalizedThreeExpandIf(
	isNormalizePlusLongPathFix,
	isExpand bool,
	first, second, third string,
) string {
	combined := normalize.SimpleJoinPath3(
		first,
		second,
		third)

	expand := expandpath.ExpandVariablesIf(
		isExpand,
		combined)

	return normalize.PathUsingSeparatorUsingSingleIf(
		isNormalizePlusLongPathFix,
		osconsts.PathSeparator,
		expand)
}
