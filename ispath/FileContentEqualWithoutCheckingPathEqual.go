package ispath

import "gitlab.com/evatix-go/pathhelper/internal/ispathinternal"

func FileContentEqualWithoutCheckingPathEqual(
	leftFullPath string,
	rightFullPath string,
) bool {
	return ispathinternal.FileContentEqualWithoutCheckingPathEqual(
		leftFullPath,
		rightFullPath)
}
