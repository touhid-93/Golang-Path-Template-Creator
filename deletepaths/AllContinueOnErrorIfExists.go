package deletepaths

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

func AllContinueOnErrorIfExists(
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
			recursiveErr := SingleOnExist(location)

			errorCollection.AddWrapperPtr(recursiveErr)
		}

		return errCount == errorCollection.Length()
	}

	return AllContinueOnError(
		errorCollection,
		locations)
}
