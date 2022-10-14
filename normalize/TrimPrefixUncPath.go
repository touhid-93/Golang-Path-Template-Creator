package normalize

func TrimPrefixUncPath(
	givenPath string,
) string {
	return TrimPrefixes(
		givenPath,
		uncPrefixes...)
}
