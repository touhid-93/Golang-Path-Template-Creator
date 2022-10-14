package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPaths(
	isContinueOnError bool,
	verifiers *pathinsfmt.PathVerifiers,
	errCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if verifiers == nil || verifiers.IsEmpty() {
		return true
	}

	stateTracker := errCollection.StateTracker()
	locationsWithError := recursiveOrNonRecursiveLocations(
		isContinueOnError,
		verifiers.IsRecursiveCheck,
		verifiers.IsNormalize,
		locations,
	)

	errCollection.AddWrapperPtr(locationsWithError.ErrorWrapper)

	if !isContinueOnError && locationsWithError.HasError() {
		return false
	}

	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		verifiers.IsSkipCheckingOnInvalid,
		locationsWithError.SafeValues())

	// exit immediately
	if !isContinueOnError {
		for i := range verifiers.PathVerifiers {
			isSuccess = applyVerifierInternal(
				isContinueOnError,
				&verifiers.PathVerifiers[i],
				errCollection,
				existingPathsFileInfoMap)

			if !isSuccess {
				return false
			}
		}
	}

	// continue on error
	for i := range verifiers.PathVerifiers {
		applyVerifierInternal(
			isContinueOnError,
			&verifiers.PathVerifiers[i],
			errCollection,
			existingPathsFileInfoMap)
	}

	return stateTracker.IsSuccess()
}
