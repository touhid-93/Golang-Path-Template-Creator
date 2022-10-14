package pathfilter

import (
	"sync"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
	"gitlab.com/evatix-go/errorwrapper/errdata/errstr"
)

func GetExcept(
	separator string,
	isNormalize bool,
	pathTranspiler *corestr.Hashmap,
	rootPath string,
	filter, excepts *Query,
) *errstr.ResultsWithErrorCollection {
	var selectedFilesHashset,
		exceptFilesHashset *corestr.Hashset
	var selectedFilesResults,
		exceptFilesResults *errstr.ResultsWithErrorCollection

	wg := &sync.WaitGroup{}
	wg.Add(constants.Two)

	rootPath = GetTranspiledPathForDollarVariables(
		pathTranspiler,
		rootPath)

	go func() {
		defer wg.Done()

		selectedFilesResults = Get(
			separator,
			isNormalize,
			nil,
			rootPath,
			filter)

		if selectedFilesResults.HasIssuesOrEmpty() {
			return
		}

		selectedFilesHashset = corestr.New.Hashset.StringsPtr(
			selectedFilesResults.SafeValuesPtr())
	}()

	go func() {
		defer wg.Done()

		exceptFilesResults = Get(
			separator,
			isNormalize,
			nil,
			rootPath,
			excepts)

		if !exceptFilesResults.HasSafeItems() {
			return
		}

		exceptFilesHashset = corestr.New.Hashset.StringsPtr(
			exceptFilesResults.SafeValuesPtr())
	}()

	wg.Wait()

	selectedFilesResults.ErrorWrappers.AddCollections(
		exceptFilesResults.ErrorWrappers)

	if selectedFilesResults.HasError() {
		return selectedFilesResults
	}

	finalResult := selectedFilesHashset.GetAllExceptHashset(
		exceptFilesHashset)

	return &errstr.ResultsWithErrorCollection{
		Values:        finalResult,
		ErrorWrappers: selectedFilesResults.ErrorWrappers,
	}
}
