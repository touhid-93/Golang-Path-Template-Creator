package recursivepaths

func AllMust(rootPath string) []string {
	allResults := All(rootPath)
	allResults.ErrorWrapper.HandleError()

	return allResults.SafeValues()
}
