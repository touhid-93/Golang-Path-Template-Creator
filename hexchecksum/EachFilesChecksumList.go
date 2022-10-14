package hexchecksum

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func EachFilesChecksumList(
	method hashas.Variant,
	filesPaths ...string,
) *errstr.Results {
	checkSumSlice := make(
		[]string,
		len(filesPaths))

	for i, filePath := range filesPaths {
		hexFileChecksumResult := method.
			HexSumOfFile(filePath)

		if hexFileChecksumResult.IsSuccess() {
			checkSumSlice[i] = hexFileChecksumResult.Value
		}

		if hexFileChecksumResult.HasError() {
			return errstr.New.Results.ErrorWrapper(
				hexFileChecksumResult.ErrorWrapper)
		}
	}

	return errstr.New.Results.Strings(checkSumSlice)
}
