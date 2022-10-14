package pathhelper

func isAlphabetChar(eachChar rune) bool {
	return isLowerCase(eachChar) || isUpperCase(eachChar)
}
