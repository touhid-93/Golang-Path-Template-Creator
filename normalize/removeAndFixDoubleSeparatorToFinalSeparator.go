package normalize

import "gitlab.com/evatix-go/core/constants"

// Replace both double slashes to single slash (// -> /, \\ -> \) and finally all slashes to finalSeparator
func removeAndFixDoubleSeparatorToFinalSeparator(finalSeparator, path string) string {
	pathUsingBackSlash := GetCompiledPath(
		path,
		removeAndFixDoubleSeparatorToFinalSeparatorMap)
	doubleSeparatorPath := ChangeSeparator(
		pathUsingBackSlash, constants.TripleBackSlash,
		constants.BackSlash)
	singleBackSlashesPath := ChangeDoubleBackSlash(
		doubleSeparatorPath,
		constants.BackSlash)
	finalPath := ChangeSeparator(
		singleBackSlashesPath,
		constants.BackSlash,
		finalSeparator)

	return finalPath
}
