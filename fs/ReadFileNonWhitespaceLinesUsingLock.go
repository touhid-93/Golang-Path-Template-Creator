package fs

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func ReadFileNonWhitespaceLinesUsingLock(filePath string) *errstr.Results {
	errString := ReadFileStringUsingLock(filePath)

	if errString.Value == "" {
		return errstr.New.Results.ErrorWrapper(
			errString.ErrorWrapper)
	}

	lines := strings.Split(
		errString.Value,
		constants.NewLineUnix)

	nonEmptyLines := stringslice.NonWhitespaceSlice(lines)

	return &errstr.Results{
		Values:       nonEmptyLines,
		ErrorWrapper: errString.ErrorWrapper,
	}
}
