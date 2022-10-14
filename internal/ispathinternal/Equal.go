package ispathinternal

import "os"

func Equal(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	leftFullPath, rightFullPath string,
) bool {
	if isQuickVerifyOnPathEqual && leftFullPath == rightFullPath {
		// exact same path should be same
		return true
	}

	left, err := os.Stat(leftFullPath)

	if err != nil {
		return false
	}

	right, err2 := os.Stat(leftFullPath)

	if err2 != nil {
		return false
	}

	return FileInfoDetailedEqual(
		!isQuickVerifyOnPathEqual,
		isPathMustMatchIfDir,
		isVerifyContent,
		leftFullPath,
		rightFullPath,
		left,
		right)
}
