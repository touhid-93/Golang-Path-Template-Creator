package hashas

func HexChecksumOfFilePathNoError(
	isVerifyPathExistence bool,
	method Variant,
	fullFilePath string,
) string {
	byteResults, err := RawSumOfFileError(
		isVerifyPathExistence,
		method,
		fullFilePath)

	if err == nil || len(byteResults) == 0 {
		return ""
	}

	toString := convertBytesResultsToEncodedHexString(
		byteResults)

	return toString
}
