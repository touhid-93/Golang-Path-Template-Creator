package filestate

func IsNotEqual(
	isIgnoreModifiedTimeCompare,
	isIgnoreChmodCompare,
	isIgnoreChownCompare,
	isIgnoreCompareOnAnyEmpty bool,
	left, right *Info,
) bool {
	return !IsEqual(
		isIgnoreModifiedTimeCompare,
		isIgnoreChmodCompare,
		isIgnoreChownCompare,
		isIgnoreCompareOnAnyEmpty,
		left,
		right)
}
