package pathstatlinux

func InvalidGroup() *Group {
	return &Group{
		IntIdNameValidation: *InvalidIntIdNameValidation(),
	}
}
