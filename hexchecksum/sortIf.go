package hexchecksum

import "sort"

func sortIf(
	isSort bool,
	stringItems []string,
) {
	if !isSort {
		return
	}

	sort.Strings(stringItems)
}
