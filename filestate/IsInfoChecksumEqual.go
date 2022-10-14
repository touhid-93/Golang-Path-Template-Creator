package filestate

func IsInfoChecksumEqual(
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

	return IsChecksumEqual(
		isIgnoreCompareOnAnyEmpty,
		left.HexContentChecksum,
		right.HexContentChecksum)
}
