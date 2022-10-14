package pathgetter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func All(
	separator string,
	isNormalize bool,
	exploringPaths []string,
) *errstr.Results {
	length := len(exploringPaths)

	if length == 0 {
		return errstr.Empty.Results()
	}

	if length == 1 {
		return AllOfSinglePath(
			isNormalize,
			separator,
			exploringPaths[0])
	}

	linkedCollection := corestr.Empty.LinkedCollections()
	wg := &sync.WaitGroup{}

	for _, expPath := range exploringPaths {
		wg.Add(1)
		allPaths := AllOfSinglePath(
			isNormalize,
			separator,
			expPath)

		if allPaths.HasError() {
			return &errstr.Results{
				Values:       *linkedCollection.ListPtr(),
				ErrorWrapper: allPaths.ErrorWrapper,
			}
		}

		linkedCollection.AddStringsPtrAsync(
			wg,
			false,
			allPaths.SafeValuesPtr(),
		)
	}

	wg.Wait()

	return &errstr.Results{
		Values:       *linkedCollection.ListPtr(),
		ErrorWrapper: nil,
	}
}
