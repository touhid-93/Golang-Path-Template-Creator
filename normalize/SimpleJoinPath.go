package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func SimpleJoinPath(path1, path2 string) string {
	if path1 == constants.EmptyString && path2 == constants.EmptyString {
		return constants.EmptyString
	}

	if path1 == constants.EmptyString {
		return path2
	}

	if path2 == constants.EmptyString {
		return path1
	}

	joined := path1 +
		osconsts.PathSeparator +
		path2

	return joined
}
