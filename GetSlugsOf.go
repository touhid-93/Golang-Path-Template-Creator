package pathhelper

func GetSlugsOf(
	isDotValid bool,
	slugFixer,
	spaceSlugFixer rune,
	paths []string,
) []string {
	return GetAsyncProcessed(
		paths,
		func(
			index int,
			currentPath string,
		) (result string) {
			return GetSlug(
				isDotValid,
				slugFixer,
				spaceSlugFixer,
				currentPath,
			)
		})
}
