package ispath

// AllNotExist returns false if any of the paths exists
func AllNotExist(paths ...string) bool {
	for _, path := range paths {
		if Exists(path) {
			return false
		}
	}

	return true
}
