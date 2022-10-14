package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

func GetPathAsUri(path string, isNormalizePath bool) string {
	if isNormalizePath {
		path = normalize.PathUsingSeparatorIf(
			false,
			false,
			true,
			osconsts.PathSeparator,
			path)
	}

	return constants.UriSchemePrefixStandard + strings.ReplaceAll(
		path,
		constants.BackSlash,
		constants.ForwardSlash)
}
