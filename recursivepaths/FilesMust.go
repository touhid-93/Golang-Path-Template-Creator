package recursivepaths

func FilesMust(rootPath string) []string {
	allResults := Files(rootPath)
	allResults.ErrorWrapper.HandleError()

	return allResults.SafeValues()
}
