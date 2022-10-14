package hexchecksum

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfFilesListIf(
	isGenerate bool,
	hashMethod hashas.Variant,
	files ...string,
) *errstr.Result {
	if !isGenerate {
		return errstr.Empty.Result()
	}

	jsonResult := corejson.NewPtr(files)

	return hashMethod.
		HexOfJsonResult(jsonResult)
}
