package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func ReadFileLinesUsingLock(filePath string) *errstr.Results {
	errString := ReadFileStringUsingLock(filePath)

	if errString.Value == "" {
		return errstr.New.Results.ErrorWrapper(
			errString.ErrorWrapper)
	}

	lines := strings.Split(
		errString.Value,
		constants.NewLineUnix)

	return &errstr.Results{
		Values:       lines,
		ErrorWrapper: errString.ErrorWrapper,
	}
}
