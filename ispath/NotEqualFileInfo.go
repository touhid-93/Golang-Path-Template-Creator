package ispath

import "os"

func NotEqualFileInfo(
	left, right os.FileInfo,
) bool {
	return !FileInfoEqual(left, right)
}
