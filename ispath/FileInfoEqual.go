package ispath

import "os"

func FileInfoEqual(
	left, right os.FileInfo,
) bool {
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

	return true
}
