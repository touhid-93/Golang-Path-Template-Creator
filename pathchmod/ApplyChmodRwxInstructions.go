package pathchmod

import (
	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ApplyChmodRwxInstructions(
	instructions *chmodins.BaseRwxInstructions,
	paths []string,
) *errorwrapper.Wrapper {
	if instructions == nil ||
		instructions.RwxInstructions == nil ||
		len(paths) == 0 {
		return nil
	}

	executors, err := chmodhelper.ParseRwxInstructionsToExecutors(
		instructions.RwxInstructions)

	if err != nil {
		return errnew.Type.Error(
			errtype.ParsingFailed,
			err)
	}

	err2 := executors.ApplyOnPathsPtr(
		&paths)

	if err2 != nil {
		return errnew.Type.Error(
			errtype.ChmodApplyFailed,
			err2)
	}

	return nil
}
