package normalize

func DbPath(path string) string {
	return PathFixWithoutLongPathIf(
		true,
		path)
}
