package envpath

import (
	"os"
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func ReadEnvPathsPtr() *[]string {
	pathString := os.Getenv(constants.Path)
	slice := strings.Split(
		pathString,
		GetEnvSeparator())

	return &slice
}
