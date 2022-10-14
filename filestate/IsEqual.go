package filestate

import "gitlab.com/evatix-go/pathhelper/ispath"

func IsEqual(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	left, right *Info,
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

	if left.IsInvalid != right.IsInvalid {
		return false
	}

	if left.Size != right.Size {
		return false
	}

	if left.IsFile != right.IsFile {
		return false
	}

	if ispath.NotEqualString(
		true,
		left.FullPath,
		right.FullPath) {
		return false
	}

	isChecksumNotEqual := !IsChecksumEqual(
		isIgnoreCompareOnAnyEmpty,
		left.HexContentChecksum,
		right.HexContentChecksum)

	if isChecksumNotEqual {
		return false
	}

	if !isIgnoreChmodCompare && left.Chmod != right.Chmod {
		return false
	}

	if !isIgnoreChownCompare && !left.UserGroupId.IsEqual(right.UserGroupId) {
		return false
	}

	if !isIgnoreModifiedTimeCompare && !left.IsLastModifiedEqual(right) {
		return false
	}

	return true
}
