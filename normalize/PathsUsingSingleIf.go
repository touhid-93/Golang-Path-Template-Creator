package normalize

func PathsUsingSingleIf(
	isNormalizeLongPathForce bool,
	locations []string,
) []string {
	if len(locations) == 0 {
		return []string{}
	}

	if !isNormalizeLongPathForce {
		return locations
	}

	options := &Options{
		IsNormalize:        isNormalizeLongPathForce,
		IsLongPathFix:      isNormalizeLongPathForce,
		IsForceLongPathFix: isNormalizeLongPathForce,
	}

	return PathsUsingOptions(options, locations)
}
