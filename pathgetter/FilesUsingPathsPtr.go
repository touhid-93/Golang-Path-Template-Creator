package pathgetter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/core/coreindexes"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
)

func FilesUsingPathsPtr(
	isNormalize bool,
	separator string,
	exploringPaths []string,
) *errstr.Results {
	length := len(exploringPaths)

	if length == 0 {
		return errstr.Empty.Results()
	}

	if length == 1 {
		exploringPath :=
			exploringPaths[coreindexes.First]

		return Files(
			isNormalize,
			separator,
			exploringPath,
		)
	}

	linkedCollections :=
		corestr.Empty.LinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(length)
	errWrappers := errwrappers.NewCap2()

	for _, eachExploringPath := range exploringPaths {
		allPaths := Files(
			isNormalize,
			separator,
			eachExploringPath,
		)

		if allPaths.HasError() {
			errWrappers.AddWrapperPtr(allPaths.ErrorWrapper)
			wg.Done()

			continue
		}

		linkedCollections.AddStringsPtrAsync(
			wg,
			false,
			allPaths.SafeValuesPtr(),
		)
	}

	wg.Wait()

	return &errstr.Results{
		Values:       *linkedCollections.ListPtr(),
		ErrorWrapper: errWrappers.GetAsErrorWrapperPtr(),
	}
}
