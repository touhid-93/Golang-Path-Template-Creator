package envpath

// linuxEnvAddOrUpdateAction - is a linuxEnvPathCrudFunc
func linuxEnvAddOrUpdateAction(
	envPaths []string,
	linuxExistingEnvPathRaw string,
) string {
	hashset := getLinuxRawEnvPathToHashset(
		linuxExistingEnvPathRaw)
	hashset.AddStrings(
		envPaths)
	compiledJoinedPath := hashsetEnvPathToSingleString(
		hashset)

	return compiledJoinedPath
}
