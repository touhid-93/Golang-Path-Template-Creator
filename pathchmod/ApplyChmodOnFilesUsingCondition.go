package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
)

func ApplyChmodOnFilesUsingCondition(
	condition *chmodins.Condition,
	changeFileMode os.FileMode,
	locations ...string,
) (*chmodins.RwxInstruction, *errorwrapper.Wrapper) {
	if len(locations) == 0 || condition == nil {
		return &chmodins.RwxInstruction{}, nil
	}

	return ApplyChmodOnFiles(
		condition.IsRecursive,
		condition.IsSkipOnInvalid,
		condition.IsContinueOnError,
		changeFileMode,
		locations...)
}
