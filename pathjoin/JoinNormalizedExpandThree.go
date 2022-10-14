package pathjoin

import (
	"gitlab.com/evatix-go/pathhelper/expandpath"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

// JoinNormalizedExpandThree normalized and expand applied auto
func JoinNormalizedExpandThree(
	first, second, third string,
) string {
	joined := JoinSimpleThree(first, second, third)
	expand := expandpath.ExpandVariables(joined)

	return normalize.Path(expand)
}
