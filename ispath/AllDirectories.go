package ispath

// AllDirectories returns false if is not dir
func AllDirectories(paths ...string) bool {
	for _, path := range paths {
		if !Directory(path) {
			return false
		}
	}

	return true
}
