package normalize

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
)

func SimpleJoinPath3(path1, path2, path3 string) string {
	if path1 == constants.EmptyString &&
		path2 == constants.EmptyString &&
		path3 == constants.EmptyString {
		return constants.EmptyString
	}

	if path1 == constants.EmptyString {
		return SimpleJoinPath(path2, path3)
	}

	if path2 == constants.EmptyString {
		return SimpleJoinPath(path1, path3)
	}

	if path3 == constants.EmptyString {
		return SimpleJoinPath(path1, path2)
	}

	joined := path1 +
		osconsts.PathSeparator +
		path2 +
		osconsts.PathSeparator +
		path3

	return joined
}
