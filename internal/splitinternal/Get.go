package splitinternal

func Get(currentPath string) (baseDir, fileName string) {
	i := LastSlash(currentPath)

	return currentPath[:i+1], currentPath[i+1:]
}
