package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ApplyOnMismatchSkipOnInvalid(
	changeFileMode os.FileMode,
	location string,
) *errorwrapper.Wrapper {
	err := chmodhelper.ChmodApply.OnMismatch(
		true,
		changeFileMode,
		location)

	if err == nil {
		return nil
	}

	return errnew.Error.Default(
		errtype.ChmodApplyFailed,
		err)
}
