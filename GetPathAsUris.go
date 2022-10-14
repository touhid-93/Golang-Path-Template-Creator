package pathhelper

func GetPathAsUris(
	isNormalizePath bool,
	paths []string,
) []string {
	return GetAsyncProcessed(
		paths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return GetPathAsUri(
				currentPath,
				isNormalizePath)
		})
}
