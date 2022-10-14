package hashas

func HexChecksumOfFilePathError(
	isVerifyPathExistence bool,
	method Variant,
	fullFilePath string,
) (hexChecksum string, err error) {
	byteResults, err := RawSumOfFileError(
		isVerifyPathExistence,
		method,
		fullFilePath)

	if err != nil {
		return "", err
	}

	if len(byteResults) == 0 {
		return "", nil
	}

	toString := convertBytesResultsToEncodedHexString(byteResults)

	return toString, nil
}
