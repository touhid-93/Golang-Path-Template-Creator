package pathhelper

func GetAbsoluteFromExecutableDirectoryPath(relativePath string, isLongPathFix, isNormalize bool) string {
	basePath := GetExecutablePath()

	return GetAbsolutePath(
		basePath,
		relativePath,
		isLongPathFix,
		isNormalize)
}
