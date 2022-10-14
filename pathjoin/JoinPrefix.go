package pathjoin

// JoinPrefix
//
// only adds prefix if not exist in main
func JoinPrefix(
	prefix, main string,
) string {
	return JoinPrefixSuffixIf(
		true,
		false,
		prefix,
		main,
		"")
}
