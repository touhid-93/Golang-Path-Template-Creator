package ispathinternal

import "path/filepath"

// FileContentEqual
//
// Returns false if contents are not equal at any point
//
// Reference : https://stackoverflow.com/a/30038571
func FileContentEqual(leftFullPath string, rightFullPath string) bool {
	leftCleanedPath := filepath.Clean(leftFullPath)
	rightCleanedPath := filepath.Clean(rightFullPath)

	if leftCleanedPath == rightCleanedPath {
		// if same path then no need to verify contents of it
		return true
	}

	return FileContentEqualWithoutCheckingPathEqual(
		leftFullPath,
		rightFullPath)
}
