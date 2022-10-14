package fs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func WriteJsonResult(
	isCreateParentDir,
	isSkipErrorOnNilOrEmpty bool,
	jsonResult *corejson.Result,
	location string,
) *errorwrapper.Wrapper {
	if isSkipErrorOnNilOrEmpty && jsonResult == nil {
		return nil
	}

	isSkipBytesError := isSkipErrorOnNilOrEmpty &&
		jsonResult.Bytes == nil &&
		jsonResult.Error == nil

	if isSkipBytesError {
		return nil
	}

	hasExistingError := jsonResult != nil &&
		jsonResult.Error != nil

	if hasExistingError {
		//goland:noinspection GoNilness
		return errnew.
			Path.
			Error(
				errtype.WriteFailed,
				jsonResult.Error,
				location)
	}

	if isSkipErrorOnNilOrEmpty &&
		jsonResult != nil &&
		jsonResult.IsEmptyJsonBytes() {
		return WriteFile(
			isCreateParentDir,
			location,
			[]byte{})
	}

	return WriteFile(
		isCreateParentDir,
		location,
		jsonResult.Bytes)
}
