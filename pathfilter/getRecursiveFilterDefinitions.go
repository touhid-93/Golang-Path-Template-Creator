package pathfilter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
)

func getRecursiveFilterDefinitions(
	arg *recursiveFilterGetterParam,
) []string {
	linkedCollection :=
		corestr.Empty.LinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(arg.additionalFiltersLength)

	for _, filterPath := range arg.additionalFilters {
		rootPathPlusFilterPath :=
			arg.rootPathPlusSeparator +
				filterPath
		arg.eachFilterPath =
			rootPathPlusFilterPath

		newFilters :=
			getRecursiveFilterForEachFilterPath(arg)

		linkedCollection.AddStringsPtrAsync(
			wg,
			false,
			&newFilters,
		)
	}

	wg.Wait()

	return *linkedCollection.ListPtr()
}
