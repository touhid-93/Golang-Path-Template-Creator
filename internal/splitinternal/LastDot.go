package splitinternal

func LastDot(s string) int {
	i := len(s) - 1

	for i >= 0 && s[i] != '.' {
		i--
	}

	return i
}
