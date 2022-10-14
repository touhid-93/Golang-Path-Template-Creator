package hashas

import "gitlab.com/evatix-go/errorwrapper/errdata/errstr"

func HexChecksumOfRawBytes(method Variant, inputBytes []byte) *errstr.Result {
	outputBytesResults := method.SumOf(inputBytes)
	toString := outputBytesResults.NonEmptyString(convertBytesResultsToEncodedHexString)

	return &errstr.Result{
		Value:        toString,
		ErrorWrapper: outputBytesResults.ErrorWrapper,
	}
}
