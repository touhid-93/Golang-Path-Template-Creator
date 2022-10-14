package pathmodifierverify

import (
	"strings"

	"gitlab.com/evatix-go/core/chmodhelper"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/osconsts"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/internal/mics"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func applyVerifierInternal(
	isContinueOnError bool,
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	filteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap,
) (isSuccess bool) {
	if verifier == nil || filteredPathFileInfoMap == nil {
		return true
	}

	if isContinueOnError {
		return applyVerifierContinueOnErrorInternal(
			verifier,
			errCollection,
			filteredPathFileInfoMap)
	}

	stateTracker := errCollection.StateTracker()
	if filteredPathFileInfoMap.Error != nil {
		errCollection.AddTypeError(
			errtype.PathMissingOrInvalid,
			filteredPathFileInfoMap.Error)

		return false
	}

	validFileLocations := mics.GetLocationsUsingFilteredPathFileInfoMap(
		filteredPathFileInfoMap)

	// Rwx verify
	if verifier.HasRwxInstructions() {
		executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(
			&verifier.BaseRwxInstructions)

		if err != nil {
			errCollection.AddTypeError(
				errtype.ParsingFailed, err)

			return false
		}

		errRwxInstructions := executors.VerifyRwxModifiers(
			isContinueOnError,
			true,
			validFileLocations)

		if errRwxInstructions != nil {
			errCollection.AddTypeError(
				errtype.RwxMismatch,
				errRwxInstructions)

			return false
		}
	}

	// immediately exit on error, user+groups verify.
	hasAnyUserGroupValidation := len(validFileLocations) > 0 &&
		verifier.UserGroupName.HasUserNameOrGroup()

	if hasAnyUserGroupValidation && osconsts.IsWindows {
		errCollection.AddUsingMessages(
			errtype.NotSupportInWindows,
			errtype.ChownUserOrGroupApplyIssue.String(),
			"Cannot verify valid locations:",
			strings.Join(validFileLocations, constants.CommaSpace))

		return false
	}

	// immediately exit on error, user+groups verify.
	for _, location := range validFileLocations {
		isSuccess = applyVerifierSinglePathNonRecursiveUserGroupVerify(
			verifier,
			errCollection,
			location)

		if !isSuccess {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
