package hexchecksum

import (
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/pathhelper/hashas"
)

func OfFilesContents(
	isSortFilePaths bool,
	method hashas.Variant,
	filesPaths ...string,
) *errstr.Result {
	if len(filesPaths) == 0 {
		return errstr.Empty.Result()
	}

	sortIf(isSortFilePaths, filesPaths)

	checkSumSlice := make([]string, len(filesPaths))

	for i, filePath := range filesPaths {
		hexFileChecksumResult := method.
			HexSumOfFile(filePath)

		if hexFileChecksumResult.IsSuccess() {
			checkSumSlice[i] = hexFileChecksumResult.Value
		} else if hexFileChecksumResult.HasError() {
			return errstr.New.Result.ErrorWrapper(
				hexFileChecksumResult.ErrorWrapper)
		}
	}

	return method.HexSumOfAny(
		checkSumSlice)
}
