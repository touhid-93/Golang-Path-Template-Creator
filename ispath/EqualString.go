package ispath

import (
	"strings"

	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/pathhelper/normalize"
)

func EqualString(
	isApplyNormalize bool,
	left, right string,
) bool {
	if left == right {
		return true
	}

	if !isApplyNormalize && osconsts.IsWindows {
		return strings.EqualFold(left, right)
	}

	leftClean := normalize.Path(left)
	rightClean := normalize.Path(right)

	if leftClean == rightClean {
		return true
	}

	if osconsts.IsWindows {
		// can be case : insensitive
		return strings.EqualFold(left, right)
	}

	return false
}
