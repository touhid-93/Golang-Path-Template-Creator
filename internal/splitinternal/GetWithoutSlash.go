package splitinternal

func GetWithoutSlash(currentPath string) (baseDir, fileName string) {
	i := LastSlash(currentPath)

	return currentPath[:i], currentPath[i+1:]
}
