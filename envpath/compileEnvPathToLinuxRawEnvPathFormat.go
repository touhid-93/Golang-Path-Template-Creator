package envpath

import (
	"gitlab.com/evatix-go/core/simplewrap"
)

// compileEnvPathToLinuxRawEnvPathFormat Location="...."
//
// Adds double quotation and prepends `PATH="..."`
// Given string gx will become `PATH="gx"` without any further checking.
func compileEnvPathToLinuxRawEnvPathFormat(compiledJoinedPath string) string {
	pathsQuotation := simplewrap.WithDoubleQuote(
		compiledJoinedPath)
	finalPath := pathEqual + pathsQuotation

	return finalPath
}
