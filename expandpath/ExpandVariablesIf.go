package expandpath

// ExpandVariablesIf
//
// Acceptable Env paths:
// ${Java_home} $java_home %{java_home} %java_home all will be expand e
func ExpandVariablesIf(
	isExpand bool,
	pathContainsEnvVariables string,
) string {
	if !isExpand {
		return pathContainsEnvVariables
	}

	if pathContainsEnvVariables == "" {
		return pathContainsEnvVariables
	}

	return ExpandVariables(pathContainsEnvVariables)
}
