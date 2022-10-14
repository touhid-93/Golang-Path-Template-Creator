package ispath

// AllFiles returns false if is not file
func AllFiles(paths ...string) bool {
	for _, path := range paths {
		if !File(path) {
			return false
		}
	}

	return true
}
