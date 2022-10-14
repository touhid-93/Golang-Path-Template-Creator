package pathjoin

import "sync"

func TrimReplacePrefixRoots(
	isNormalize,
	isLongPathForceFix bool,
	baseLocations []string,
	trimRootPrefix,
	rootReplacer string,
) []string {
	if len(baseLocations) == 0 || trimRootPrefix == "" {
		return baseLocations
	}

	length := len(baseLocations)
	replacedPaths := make([]string, length)
	wg := &sync.WaitGroup{}

	replacerFunc := func(index int, source string) {
		replacedPaths[index] = TrimReplacePrefixRoot(
			isNormalize,
			isLongPathForceFix,
			source,
			trimRootPrefix,
			rootReplacer)

		wg.Done()
	}

	wg.Add(length)
	for i, location := range baseLocations {
		go replacerFunc(i, location)
	}

	wg.Wait()

	return replacedPaths
}
