package ispath

// AllExist returns false if any of the path not exists
func AllExist(paths ...string) bool {
	for _, path := range paths {
		if !Exists(path) {
			return false
		}
	}

	return true
}
