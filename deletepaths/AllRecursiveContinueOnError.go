package deletepaths

import "gitlab.com/evatix-go/errorwrapper/errwrappers"

func AllRecursiveContinueOnError(
	errorCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if len(locations) == 0 {
		return true
	}

	errCount := errorCollection.Length()

	for _, location := range locations {
		recursiveErr := Recursive(location)

		errorCollection.AddWrapperPtr(recursiveErr)
	}

	return errCount == errorCollection.Length()
}
