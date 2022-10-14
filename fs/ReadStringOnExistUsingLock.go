package fs

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func ReadStringOnExistUsingLock(filePath string) *errstr.Result {
	if IsPathExistsUsingLock(filePath) {
		return ReadFileStringUsingLock(filePath)
	}

	return &errstr.Result{
		Value:        constants.EmptyString,
		ErrorWrapper: nil,
	}
}
