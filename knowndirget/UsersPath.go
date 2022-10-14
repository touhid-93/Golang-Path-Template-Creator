package knowndirget

import "path/filepath"

// Returns path to Users directory
func UsersPath() string {
	return filepath.Dir(UserPath())
}
