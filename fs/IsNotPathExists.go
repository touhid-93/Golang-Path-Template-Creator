package fs

func IsNotPathExists(location string) bool {
	return !IsPathExists(location)
}
