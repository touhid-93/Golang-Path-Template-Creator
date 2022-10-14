package pathstatlinux

func getUserFromPathStatLines(splits []string) *User {
	// (    0/    root)   Gid
	idNameValidation := pathStatIndexedLineIntIdNameValidation(
		splits, pathStatUserIndex)

	return &User{
		IntIdNameValidation: *idNameValidation,
	}
}
