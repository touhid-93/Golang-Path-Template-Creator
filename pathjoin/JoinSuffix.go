package pathjoin

// JoinSuffix
//
//  only adds suffix if not exist in main ending
func JoinSuffix(
	main, suffix string,
) string {
	return JoinPrefixSuffixIf(
		true,
		false,
		"",
		main,
		suffix)
}
