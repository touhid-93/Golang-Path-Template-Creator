package normalize

// Path By default apply long path fix and regular normalize using os.PathSeparator
func Path(givenPath string) string {
	return LongPathFixPlusClean(
		false,
		givenPath)
}
