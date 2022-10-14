package knowndirget

import (
	"os"
)

// Returns home directory. Panics if directory doesn't exist.
func UserPath() string {
	homedir, err := os.UserHomeDir()

	if err != nil {
		panic("An error occurred getting 'os.UserHomeDir'")
	}

	return homedir
}
