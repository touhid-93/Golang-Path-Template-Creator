package hashas

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func HexChecksumOfFilePath(method Variant, fullFilePath string) *errstr.Result {
	byteResults := SumOfFile(method, fullFilePath)
	if byteResults.HasError() {
		return &errstr.Result{
			ErrorWrapper: byteResults.ErrorWrapper,
		}
	}

	toString := byteResults.NonEmptyString(
		convertBytesResultsToEncodedHexString)

	return errstr.New.Result.ValueOnly(toString)
}
