package pathfuncs

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

func SimpleFilterFullPathsUsingErrorCollectionAsync(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	filter SimpleFilter,
	fullPaths ...string,
) []string {
	if len(fullPaths) == 0 {
		return []string{}
	}

	results := SimpleFilterFullPathsAsync(isContinueOnError,
		filter,
		fullPaths...)

	if results.HasError() {
		errCollection.AddWrapperPtr(results.ErrorWrapper)
	}

	return results.SafeValues()
}
