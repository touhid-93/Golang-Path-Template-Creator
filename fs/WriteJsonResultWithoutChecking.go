package fs

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper"
)

func WriteJsonResultWithoutChecking(
	jsonResult *corejson.Result,
	location string,
) *errorwrapper.Wrapper {
	return WriteFile(
		true,
		location,
		jsonResult.Bytes)
}
