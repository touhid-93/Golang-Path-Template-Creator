package envpath

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func ReadEnvPaths() []string {
	pathString := os.Getenv(constants.Path)

	return strings.Split(
		pathString,
		GetEnvSeparator())
}
