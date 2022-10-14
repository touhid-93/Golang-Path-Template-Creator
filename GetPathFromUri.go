package pathhelper

// GetPathFromUri function takes a path (string) and a bool input and returns a string
// after removing Uri prefixes (file:///, file://). If bool is set to true any double
// separator is also removed from returned string
func GetPathFromUri(path string, isNormalizePath bool) string {
	return RemoveFromPath(path, &uriRemovePrefixes, isNormalizePath)
}
