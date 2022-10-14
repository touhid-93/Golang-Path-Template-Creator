package knowndirget

import "os"

func TempDirectory() string {
	return os.TempDir()
}
