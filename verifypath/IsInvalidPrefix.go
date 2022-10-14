package verifypath

func IsInvalidPrefix(
	isVerifyExistence bool,
	homeDirPrefix, currentFullPath string,
) (fixedPath string, isInvalid bool) {
	fixedPath, isValid := IsPrefixValid(
		isVerifyExistence,
		homeDirPrefix,
		currentFullPath)

	return fixedPath, !isValid
}
