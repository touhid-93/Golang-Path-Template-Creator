package normalizeinternal

func PathsCombineDirect(
	root string,
	combinedPaths ...string,
) (first string, allCombinedPaths []string) {
	return PathsCombine(root, combinedPaths)
}
