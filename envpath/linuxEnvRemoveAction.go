package envpath

// linuxEnvRemoveAction - is a linuxEnvPathCrudFunc
// Returns the compiled raw env path without `PATH=`
func linuxEnvRemoveAction(
	envPaths []string,
	linuxExistingEnvPathRaw string,
) string {
	hashset := getLinuxRawEnvPathToHashset(
		linuxExistingEnvPathRaw)

	for _, envPath := range envPaths {
		hashset.Remove(
			envPath)
	}

	compiledJoinedPath := hashsetEnvPathToSingleString(
		hashset)

	return compiledJoinedPath
}
