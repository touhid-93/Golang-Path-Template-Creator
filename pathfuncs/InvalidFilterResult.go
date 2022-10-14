package pathfuncs

func InvalidFilterResult(
	fullPath string,
) *FilterResult {
	return &FilterResult{
		FullPath: fullPath,
	}
}
