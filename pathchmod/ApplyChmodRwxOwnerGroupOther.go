package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
)

func ApplyChmodRwxOwnerGroupOther(
	isRecursive,
	isSkipOnInvalid,
	isContinueOnError bool,
	rwxOwnerGroupOther *chmodins.RwxOwnerGroupOther,
	paths []string,
) *errorwrapper.Wrapper {
	if rwxOwnerGroupOther == nil ||
		len(paths) == 0 {
		return nil
	}

	condition := &chmodins.Condition{
		IsSkipOnInvalid:   isSkipOnInvalid,
		IsContinueOnError: isContinueOnError,
		IsRecursive:       isRecursive,
	}

	fileMode, err := ParseRwxOwnerGroupOtherToFileMode(rwxOwnerGroupOther)

	if err.HasError() {
		return err
	}

	_, errWrap := ApplyChmodOnFilesUsingCondition(
		condition,
		fileMode,
		paths...)

	return errWrap
}
