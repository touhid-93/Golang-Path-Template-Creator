package fs

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteErrorJsonResult(
	isCreateParentDir,
	isSkipErrorOnNilOrEmpty bool,
	errJsonResult *errjson.Result,
	location string,
) *errorwrapper.Wrapper {
	if isSkipErrorOnNilOrEmpty && errJsonResult == nil {
		return nil
	}

	hasExistingErrorWrapper := errJsonResult.ErrorWrapper != nil &&
		errJsonResult.HasError()

	if hasExistingErrorWrapper {
		return errJsonResult.ErrorWrapper
	}

	hasExistingError := errJsonResult.Result != nil &&
		errJsonResult.Result.Error != nil

	if hasExistingError {
		return errnew.
			Path.
			Error(
				errtype.WriteFailed,
				errJsonResult.Result.Error,
				location)
	}

	return WriteJsonResult(
		isCreateParentDir,
		isSkipErrorOnNilOrEmpty,
		errJsonResult.Result,
		location)
}
