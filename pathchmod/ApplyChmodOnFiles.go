package pathchmod

import (
	"os"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/chmodhelper/chmodins"
	"gitlab.com/evatix-go/core/codestack"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/ref"
)

func ApplyChmodOnFiles(
	isRecursive,
	isSkipOnInvalid,
	isContinueOnError bool,
	changeFileMode os.FileMode,
	locations ...string,
) (*chmodins.RwxInstruction, *errorwrapper.Wrapper) {
	if len(locations) == 0 {
		return &chmodins.RwxInstruction{}, nil
	}

	changingChmodRwxWrapper := chmodhelper.New.RwxWrapper.UsingFileMode(changeFileMode)
	rwxOwnerGroupOther := changingChmodRwxWrapper.ToRwxOwnerGroupOther()
	rwxInstruction := &chmodins.RwxInstruction{
		RwxOwnerGroupOther: *rwxOwnerGroupOther,
		Condition: chmodins.Condition{
			IsSkipOnInvalid:   isSkipOnInvalid,
			IsContinueOnError: isContinueOnError,
			IsRecursive:       isRecursive,
		},
	}

	executor, err := chmodhelper.ParseRwxInstructionToExecutor(rwxInstruction)

	if err != nil {
		return rwxInstruction, errnew.Type.Error(errtype.Conversion, err)
	}

	finalErr := executor.ApplyOnPaths(locations)

	if finalErr == nil {
		return rwxInstruction, nil
	}

	return rwxInstruction, errorwrapper.NewRefs(
		codestack.SkipNone,
		errtype.ChmodApplyFailed,
		finalErr,
		ref.Value{
			Variable: "locations",
			Value:    locations,
		})
}
