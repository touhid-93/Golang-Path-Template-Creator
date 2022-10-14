package pathchmod

func GetSimpleStats(
	locations []string,
) *SimpleStatMap {
	length := len(locations)
	statMap := NewSimpleStatMap(length)

	if length == 0 {
		return statMap
	}

	return statMap.Adds(locations...)
}
