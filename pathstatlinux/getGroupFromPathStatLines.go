package pathstatlinux

func getGroupFromPathStatLines(splits []string) *Group {
	// (    0/    root)   Gid
	idNameValidation := pathStatIndexedLineIntIdNameValidation(
		splits, pathStatGroupIndex)

	return &Group{
		IntIdNameValidation: *idNameValidation,
	}
}
