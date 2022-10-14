package envpath

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
)

func getLinuxRawEnvPathToHashset(existingEnvPathsWithPathEqualColonSeparator string) *corestr.Hashset {
	trimmedEnvPath := strings.TrimSpace(
		existingEnvPathsWithPathEqualColonSeparator)
	splits := strings.Split(
		trimmedEnvPath,
		constants.EqualSymbol)
	envPathsStringWithQuotation := splits[coreindexes.Second]
	envPathsStringWithoutQuotation := strings.ReplaceAll(
		envPathsStringWithQuotation,
		constants.DoubleQuoteStringSymbol,
		constants.EmptyString)
	splits2 := strings.Split(
		envPathsStringWithoutQuotation,
		unixEnvPathSplitter)

	// finalEnvPaths := sliceinternal.MergeSlices(&splits2, envPaths)
	// compiledToSinge := strings.Join(finalEnvPaths, unixEnvPathSplitter)
	hashset := corestr.New.Hashset.StringsPtr(
		&splits2)

	return hashset
}
