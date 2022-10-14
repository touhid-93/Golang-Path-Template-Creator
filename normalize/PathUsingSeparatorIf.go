package normalize

func PathUsingSeparatorIf(
	isForceLongPath,
	isLongPathFix,
	isNormalize bool,
	pathSeparator,
	givenPath string,
) string {
	isApplyLongPathFixOnly := !isNormalize &&
		(isLongPathFix || isForceLongPath)

	if isNormalize || isLongPathFix || isForceLongPath {
		givenPath = TrimPrefixUncPath(
			givenPath)
	}

	if isApplyLongPathFixOnly {
		return getLongPathFixedUsingSeparator(
			isForceLongPath,
			pathSeparator,
			givenPath,
		)
	}

	if !isNormalize {
		return givenPath
	}

	return pathUsingSeparator(
		isLongPathFix,
		isForceLongPath,
		pathSeparator,
		givenPath,
	)
}
