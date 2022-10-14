package pathjoin

func TrimReplacePrefixRootNormalized(
	source,
	trimRootPrefix,
	rootReplacer string,
) string {
	return TrimReplacePrefixRoot(
		true,
		false,
		source,
		trimRootPrefix,
		rootReplacer)
}
