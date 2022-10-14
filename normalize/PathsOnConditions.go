package normalize

func PathsOnConditions(isNormalize bool, locations []string) []string {
	length := len(locations)
	if length == 0 {
		return []string{}
	}

	if !isNormalize {
		return locations
	}

	slice := make([]string, length)

	for i, currentPath := range locations {
		slice[i] = Path(currentPath)
	}

	return slice
}
