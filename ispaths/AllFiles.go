package ispaths

func AllFiles(fullPaths ...string) bool {
	if fullPaths == nil {
		return false
	}

	return AllFilesPtr(&fullPaths)
}
