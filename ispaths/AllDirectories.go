package ispaths

func AllDirectories(fullPaths ...string) bool {
	return AllDirectoriesPtr(
		&fullPaths)
}
