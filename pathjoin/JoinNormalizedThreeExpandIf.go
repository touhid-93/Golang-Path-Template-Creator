package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalizedThreeExpandIf normalized and expand applied if condition meets
func JoinNormalizedThreeExpandIf(
	isExpandNormalize bool,
	first, second, third string,
) string {
	joined := normalize.SimpleJoinPath3(
		first,
		second,
		third)

	if !isExpandNormalize {
		return joined
	}

	expand := expandpath.ExpandVariablesIf(isExpandNormalize, joined)

	return normalize.Path(expand)
}
