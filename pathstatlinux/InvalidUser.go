package pathstatlinux

func InvalidUser() *User {
	return &User{
		IntIdNameValidation: *InvalidIntIdNameValidation(),
	}
}
