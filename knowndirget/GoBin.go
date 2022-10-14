package knowndirget

import (
	"os"

	"gitlab.com/evatix-go/core/constants"
)

// Returns env go bin path
func GoBin() string {
	return os.Getenv(constants.GoBinPath)
}
