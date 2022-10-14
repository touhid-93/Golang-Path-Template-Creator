package hashas

func RawSumOfFileErrorIf(
	isReadChecksum,
	isVerifyPathExistence bool,
	method Variant,
	filePath string,
) ([]byte, error) {
	if !isReadChecksum {
		return nil, nil
	}

	return RawSumOfFileError(
		isVerifyPathExistence,
		method,
		filePath)
}
