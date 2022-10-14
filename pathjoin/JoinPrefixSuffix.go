package pathjoin

// JoinPrefixSuffix
//
// only adds prefix if not exist in main
// only adds suffix if not exist in main
func JoinPrefixSuffix(
	prefix, main, suffix string,
) string {
	return JoinPrefixSuffixIf(
		true,
		false,
		prefix,
		main,
		suffix)
}
