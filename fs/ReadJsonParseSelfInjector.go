package fs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ReadJsonParseSelfInjector(
	filePath string,
	jsonParseSelfInjector corejson.JsonParseSelfInjector,
) *errorwrapper.Wrapper {
	errJsonResult := ReadErrorJsonResult(filePath)

	if errJsonResult == nil || jsonParseSelfInjector == nil {
		return errnew.Messages.Many(
			errtype.Unmarshalling,
			"Cannot unmarhsal nil result or to nil pointer",
			"Failed to unmarshal jsonParseSelfInjector",
			"ReadJsonParseSelfInjector",
			errorwrapper.SimpleReferencesCompile(
				errtype.FileRead, filePath))
	}

	if errJsonResult.ErrorWrapper != nil && errJsonResult.HasError() {
		return errJsonResult.ErrorWrapper
	}

	if errJsonResult.Error != nil {
		return errnew.
			Path.
			Error(
				errtype.Unmarshalling,
				errJsonResult.Error,
				filePath)
	}

	err := jsonParseSelfInjector.JsonParseSelfInject(
		errJsonResult.Result)

	return errnew.
		Path.
		Error(
			errtype.Unmarshalling,
			err,
			filePath)
}
