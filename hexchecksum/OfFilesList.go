package hexchecksum

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfFilesList(
	hashMethod hashas.Variant,
	files ...string,
) *errstr.Result {
	jsonResult := corejson.NewPtr(files)

	return hashMethod.
		HexOfJsonResult(jsonResult)
}
