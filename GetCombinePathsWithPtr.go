package pathhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// Returns path as string after combining all provided paths
// By default isIgnorePath: true, isNormalize: true, and Location separator depends on the OS
func GetCombinePathsWithPtr(paths *[]string) string {
	return GetCombinedPath(
		constants.PathSeparator,
		true,
		false,
		true,
		strings.Join(*paths, constants.PathSeparator))
}
