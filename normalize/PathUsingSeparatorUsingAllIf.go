package normalize

func PathUsingSeparatorUsingSingleIf(
	isNormalizeLongPathForce bool,
	pathSeparator,
	givenPath string,
) string {
	if !isNormalizeLongPathForce {
		return givenPath
	}

	return PathUsingSeparatorIf(
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		isNormalizeLongPathForce,
		pathSeparator,
		givenPath)
}
