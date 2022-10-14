package ispath

import (
	"os"
	"path/filepath"
)

// FileInfoDetailedEqual
//
// Conditions :
//  - isQuickVerifyOnPathEqual : true represents quick return true if path is equal (in terms of string)
//  - isPathMustMatchIfDir : true represents must match path if it is a dir.
//  - isVerifyContent : true represents contents bytes will be verified.
func FileInfoDetailedEqual(
	isQuickVerifyOnPathEqual,
	isPathMustMatchIfDir,
	isVerifyContent bool,
	leftFullPath, rightFullPath string,
	left, right os.FileInfo,
) bool {
	if isQuickVerifyOnPathEqual && leftFullPath == rightFullPath {
		// exact same path should be same
		return true
	}

	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left == right {
		return true
	}

	if left.Name() != right.Name() {
		return false
	}

	if left.IsDir() != right.IsDir() {
		return false
	}

	if left.Size() != right.Size() {
		return false
	}

	if left.Mode() != right.Mode() {
		return false
	}

	if !left.ModTime().Equal(right.ModTime()) {
		return false
	}

	isDir := left.IsDir()
	leftCleanedPath := filepath.Clean(leftFullPath)
	rightCleanedPath := filepath.Clean(rightFullPath)

	if isPathMustMatchIfDir && isDir {
		return leftCleanedPath == rightCleanedPath
	} else if isDir {
		return true
	}

	// if no content verification then both are same
	if !isVerifyContent || leftCleanedPath == rightCleanedPath {
		// if same path then no need to verify contents of it
		return true
	}

	return FileContentEqualWithoutCheckingPathEqual(leftFullPath, rightFullPath)
}
