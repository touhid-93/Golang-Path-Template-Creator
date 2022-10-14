package normalize

import (
	"sync"

	"gitlab.com/evatix-go/core/osconsts"
)

func PathsUsingOptionsAsync(
	options *Options,
	locations ...string,
) []string {
	length := len(locations)
	if length == 0 {
		return []string{}
	}

	if options.IsAllOptionsDisabled() {
		return locations
	}

	newItems := make([]string, length)
	wg := &sync.WaitGroup{}
	wg.Add(length)

	processor := func(index int, location string) {
		newItems[index] = PathUsingSeparatorIf(
			options.IsForceLongPathFix,
			options.IsLongPathFix,
			options.IsNormalize,
			osconsts.PathSeparator,
			location)

		wg.Done()
	}

	for i, location := range locations {
		go processor(i, location)
	}

	wg.Wait()

	return newItems
}
