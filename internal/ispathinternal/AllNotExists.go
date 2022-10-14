package ispathinternal

// AllNotExists if any not exist return false
func AllNotExists(paths ...string) bool {
	for _, path := range paths {
		if Exists(path) {
			return false
		}
	}

	return true
}
