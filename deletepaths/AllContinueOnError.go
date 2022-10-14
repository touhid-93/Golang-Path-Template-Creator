package deletepaths

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

func AllContinueOnError(
	errorCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if len(locations) == 0 {
		return true
	}

	errCount := errorCollection.Length()

	for _, location := range locations {
		recursiveErr := Single(location)

		errorCollection.AddWrapperPtr(recursiveErr)
	}

	return errCount == errorCollection.Length()
}
