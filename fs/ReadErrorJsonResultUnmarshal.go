package fs

import (
	"encoding/json"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

func ReadErrorJsonResultUnmarshal(
	filePath string,
	unmarshalObject interface{},
) *errorwrapper.Wrapper {
	errJson := ReadErrorJsonResult(filePath)

	if errJson.ErrorWrapper != nil && errJson.HasError() {
		return errJson.ErrorWrapper
	}

	if errJson.Error != nil {
		return errnew.
			Path.
			Error(
				errtype.Unmarshalling,
				errJson.Error,
				filePath)
	}

	if errJson.Bytes == nil {
		return errnew.
			Path.
			Messages(
				errtype.Unmarshalling,
				filePath,
				"Read as nil or empty data cannot unmarshall properly.")
	}

	err := json.Unmarshal(errJson.Bytes, unmarshalObject)

	return errnew.
		Path.
		Error(errtype.Unmarshalling, err, filePath)
}
