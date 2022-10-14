package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/errorwrapper"
)

func ChmodChangeExecuteRevertUsingRwxWrapper(
	isRecursive,
	isSkipOnInvalid bool,
	changeFileModeRwxWrapper *chmodhelper.RwxWrapper,
	location string,
	executor func(location string) *errorwrapper.Wrapper,
) *errorwrapper.Wrapper {
	return ChmodChangeExecuteRevert(
		isRecursive,
		isSkipOnInvalid,
		changeFileModeRwxWrapper.ToFileMode(),
		location,
		executor)
}
