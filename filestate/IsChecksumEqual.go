package filestate

func IsChecksumEqual(
	isIgnoreCompareOnAnyEmpty bool,
	left, right string,
) bool {
	if isIgnoreCompareOnAnyEmpty && left == "" || right == "" {
		return true
	}

	return left == right
}
