package pathhelper

import "unicode"

func isLetter(path string) bool {
	for _, r := range path {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}
