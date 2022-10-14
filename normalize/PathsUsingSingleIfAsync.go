package normalize

func PathsUsingSingleIfAsync(
	isNormalizeLongPathForce bool,
	locations []string,
) []string {
	if !isNormalizeLongPathForce {
		return locations
	}

	options := &Options{
		IsNormalize:        isNormalizeLongPathForce,
		IsLongPathFix:      isNormalizeLongPathForce,
		IsForceLongPathFix: isNormalizeLongPathForce,
	}

	return PathsUsingOptionsAsync(
		options,
		locations...)
}
