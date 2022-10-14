package deletepaths

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

func AllRecursiveContinueOnErrorIfExist(
	isRemoveExistOnly bool,
	errorCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if len(locations) == 0 {
		return true
	}

	if isRemoveExistOnly {
		errCount := errorCollection.Length()

		for _, location := range locations {
			recursiveErr := RecursiveOnExist(location)

			errorCollection.AddWrapperPtr(recursiveErr)
		}

		return errCount == errorCollection.Length()
	}

	// no exist check, just remove
	return AllRecursiveContinueOnError(
		errorCollection,
		locations)
}
