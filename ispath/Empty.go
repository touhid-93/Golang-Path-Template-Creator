package ispath

import "strings"

func Empty(path string) bool {
	return len(path) == 0
}

func EmptyPtr(path *string) bool {
	return path == nil || Empty(*path)
}

func Equal(
	isCaseSensitive bool,
	leftPath, rightPath string,
) bool {
	if isCaseSensitive {
		return leftPath == rightPath
	}

	return strings.EqualFold(leftPath, rightPath)
}
