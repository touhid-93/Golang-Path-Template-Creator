package pathhelper

// Represents the directory where the application is running from.
// Returns without slash
func GetExecutableDirectory() string {
	exePath := GetExecutablePath()
	exeDir, _ := SplitWithoutSlash(exePath)

	return exeDir
}
