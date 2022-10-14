package expandpath

// ExpandVariables
//
// function takes a string input and replaces any word that starts with "$" in the input
// with its expanded path (if exists) and returns the new string.
//
// Acceptable Env paths:
//  ${Java_home} $java_home %{java_home} %java_home all will be expanded
func ExpandVariables(pathContainsEnvVariables string) string {
	if pathContainsEnvVariables == "" {
		return pathContainsEnvVariables
	}

	envInfoItems := GetDollarOrPercentSymbolIdentifierEnvInfoItems(
		pathContainsEnvVariables)

	if len(envInfoItems) == 0 {
		return pathContainsEnvVariables
	}

	replacementMap := ExpandEnvironmentVariable(envInfoItems)

	// must replace exact and as it is
	return GetCompiledPath(
		pathContainsEnvVariables,
		replacementMap)
}
