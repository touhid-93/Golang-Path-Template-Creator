package splitinternal

// LastSlash
//
// is strings.LastIndex(s, "/" or "\\")
func LastSlash(s string) int {
	i := len(s) - 1

	for i >= 0 && !(s[i] == '/' || s[i] == '\\') {
		i--
	}

	return i
}
