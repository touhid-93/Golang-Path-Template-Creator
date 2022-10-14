package ispath

import "os"

func NotEqualFileInfoDetailed(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	leftFullPath, rightFullPath string,
	left, right os.FileInfo,
) bool {
	return !FileInfoDetailedEqual(
		isQuickVerifyOnPathEqual,
		isPathMustMatchIfDir,
		isVerifyContent,
		leftFullPath, rightFullPath,
		left, right)
}
