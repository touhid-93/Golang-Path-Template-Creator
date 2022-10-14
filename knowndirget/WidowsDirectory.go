package knowndirget

import (
	"os"

	"gitlab.com/evatix-go/pathhelper/knowndir"
)

// WidowsDirectory
//
// Returns windows directory path.
func WidowsDirectory() string {
	return os.Getenv(knowndir.WindowsDirectory.Value())
}
