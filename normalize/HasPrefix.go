package normalize

import "strings"

// HasPrefix returns true if contains prefix
func HasPrefix(cleanedPrefix, cleanAbsolutePath string) bool {
	prefixFix := TrimPrefixUncPath(cleanedPrefix)
	givenPath := TrimPrefixUncPath(cleanAbsolutePath)

	return strings.HasPrefix(givenPath, prefixFix)
}
