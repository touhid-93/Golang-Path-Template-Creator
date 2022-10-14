package pathhelper

func GetAbsolutePaths(
	isLongPathFix, isNormalize bool,
	basePath string,
	relativePaths []string,
) []string {
	return GetAsyncProcessed(
		relativePaths,
		func(
			index int,
			relativePath string,
		) (result string) {
			return GetAbsolutePath(
				basePath,
				relativePath,
				isLongPathFix,
				isNormalize)
		})
}
