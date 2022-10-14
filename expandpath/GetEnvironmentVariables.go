package expandpath

import (
	"gitlab.com/evatix-go/core/coredata/stringslice"
)

// getVariables function takes a string input and identifies every word
// that begins with "$" or every word within two "%"
// in that input string then returns an array of those words.
// If input is empty or has no such word then returns nil.
func GetEnvironmentVariables(
	pathContainsEnvVarStartingDollarSymbol string,
) []string {
	if len(pathContainsEnvVarStartingDollarSymbol) == 0 {
		return []string{}
	}

	envVariableRawKeys :=
		GetDollarOrPercentSymbolIdentifierEnvInfoItems(
			pathContainsEnvVarStartingDollarSymbol)

	simpleVars := stringslice.MakeLen(len(envVariableRawKeys))

	for i, envInfo := range envVariableRawKeys {
		simpleVars[i] = envInfo.SimplifiedName
	}

	return simpleVars
}
