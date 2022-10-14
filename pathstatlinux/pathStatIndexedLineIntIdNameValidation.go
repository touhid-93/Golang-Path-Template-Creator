package pathstatlinux

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/converters"
)

func pathStatIndexedLineIntIdNameValidation(splits []string, index int) *IntIdNameValidation {
	// (    0/    root)   Gid
	leftRight := getLeftRightOfBracketsPatternData(
		splits[index])

	if !leftRight.IsValid {
		return InvalidIntIdNameValidation()
	}

	idInt, isSuccess := converters.StringToIntegerWithDefault(
		leftRight.Left,
		constants.Zero)

	return &IntIdNameValidation{
		Id:         idInt,
		Name:       leftRight.Right,
		HasValidId: isSuccess,
	}
}
