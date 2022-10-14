package normalize

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func unixFix(
	givenPath string,
) string {
	if osconsts.IsWindows {
		return givenPath
	}

	return strings.ReplaceAll(
		givenPath,
		constants.BackSlash,
		constants.ForwardSlash)
}
