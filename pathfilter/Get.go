package pathfilter

import (
	"sync"

	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/normalize"
)

//goland:noinspection ALL
func Get(
	separator string,
	isNormalize bool,
	pathTranspiler *corestr.Hashmap,
	rootPath string,
	filter *Query,
) *errstr.ResultsWithErrorCollection {
	if filter == nil {
		return errstr.
			New.ResultsWithErrorCollection.ErrorType(
			errtype.NullOrEmptyReference)
	}

	if filter.ExtensionsLength() == 0 {
		return errstr.
			Empty.
			ResultsWithErrorCollection()
	}

	rootPath = GetTranspiledPathForDollarVariables(
		pathTranspiler,
		rootPath)

	rootPath2 := normalize.PathUsingSeparatorUsingSingleIf(
		isNormalize,
		separator,
		rootPath)

	filterLength := filter.FilterLength()

	if filterLength == 0 {
		return errstr.
			Empty.
			ResultsWithErrorCollection()
	}

	linkedCollections := corestr.
		Empty.
		LinkedCollections()
	wg := &sync.WaitGroup{}
	wg.Add(filterLength)
	errWrappers := errwrappers.NewCap2()

	for _, f := range *filter.AdditionalFilters {
		eachPath := rootPath2 +
			separator +
			f

		results := getFilesForEachPath(
			separator,
			eachPath,
			filter)

		if results.HasError() {
			errWrappers.AddWrapperPtr(
				results.ErrorWrapper)
			wg.Done()
			continue
		}

		linkedCollections.AddStringsPtrAsync(
			wg,
			false,
			results.SafeValuesPtr(),
		)
	}

	wg.Wait()

	return &errstr.ResultsWithErrorCollection{
		Values:        *linkedCollections.ListPtr(),
		ErrorWrappers: errWrappers,
	}
}
