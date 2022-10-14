package fs

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errjson"
)

func ReadErrorJsonResult(filePath string) *errjson.Result {
	errBytes := ReadFile(filePath)

	if errBytes.IsFailed() {
		return errjson.New.Result.ErrorWrapper(
			errBytes.ErrorWrapper)
	}

	return errjson.New.Result.BytesWithError(
		errBytes.Values,
		errBytes.ErrorWrapper)
}
