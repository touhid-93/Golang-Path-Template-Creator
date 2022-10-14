package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/normalize"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyVerifier(
	isNormalize,
	isRecursiveCheck,
	isSkipCheckingOnInvalid,
	isContinueOnError bool,
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	locations []string,
) (isSuccess bool) {
	if verifier == nil || len(locations) == 0 {
		return true
	}

	stateTracker := errCollection.StateTracker()
	locationsWithError := recursiveOrNonRecursiveLocations(
		isContinueOnError,
		isRecursiveCheck,
		isNormalize,
		locations,
	)

	errCollection.AddWrapperPtr(locationsWithError.ErrorWrapper)

	if !isContinueOnError && locationsWithError.HasError() {
		return false
	}

	existingPathsFileInfoMap := normalize.GetFilterPathsInfoMap(
		false,
		isSkipCheckingOnInvalid,
		locationsWithError.SafeValues())

	applyVerifierInternal(
		isContinueOnError,
		verifier,
		errCollection,
		existingPathsFileInfoMap)

	return stateTracker.IsSuccess()
}
