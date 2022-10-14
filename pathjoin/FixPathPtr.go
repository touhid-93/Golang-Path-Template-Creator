package pathjoin

// FixPathPtr normalized or expand or both apply based on condition
func FixPathPtr(
	isNormalizePlusLogPathFix,
	isExpand bool,
	location string,
) *string {
	fixedPath := FixPath(
		isNormalizePlusLogPathFix,
		isExpand,
		location)

	return &fixedPath
}
