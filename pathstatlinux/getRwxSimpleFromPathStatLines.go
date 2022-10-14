package pathstatlinux

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func getRwxSimpleFromPathStatLines(splits []string, filePath string) *RwxSimple {
	// Access: (0755/drwxr-xr-x)
	leftRight := getLeftRightOfBracketsPatternData(splits[pathStatRwxIndex])

	if !leftRight.IsValid {
		return InvalidRwxSimple()
	}

	rwxWrapper, err := chmodhelper.
		New.
		RwxWrapper.
		RwxFullString(leftRight.Right)

	if err == nil {
		return &RwxSimple{
			HyphenedRwxValue: leftRight.Right,
			RwxWrapper:       &rwxWrapper,
			ErrorWrapper:     errorwrapper.StaticEmptyPtr,
			IsRwxValid:       true,
		}
	}

	return &RwxSimple{
		HyphenedRwxValue: leftRight.Right,
		RwxWrapper:       &rwxWrapper,
		ErrorWrapper: errorwrapper.NewPath(
			codestack.SkipNone,
			errtype.ChmodInvalid,
			err,
			filePath),
		IsRwxValid: false,
	}
}
