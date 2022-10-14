package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ApplyChmod(
	isRecursive bool,
	isSkipOnInvalid bool,
	changeFileMode os.FileMode,
	location string,
) (*chmodhelper.RwxWrapper, *errorwrapper.Wrapper) {
	changingChmodRwxWrapper := chmodhelper.New.RwxWrapper.UsingFileMode(changeFileMode)

	if isRecursive {
		err := changingChmodRwxWrapper.LinuxApplyRecursive(isSkipOnInvalid, location)

		return &changingChmodRwxWrapper, errnew.
			Path.
			Error(
				errtype.ChmodApplyFailed,
				err,
				location)
	}

	err := changingChmodRwxWrapper.ApplyChmod(isSkipOnInvalid, location)

	return &changingChmodRwxWrapper, errnew.
		Path.
		Error(
			errtype.ChmodApplyFailed,
			err,
			location)
}
