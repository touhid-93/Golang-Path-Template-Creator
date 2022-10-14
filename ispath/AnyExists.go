package ispath

// AnyExists Returns true if any of the paths exists
func AnyExists(paths ...string) bool {
	for _, path := range paths {
		if Exists(path) {
			return true
		}
	}

	return false
}
