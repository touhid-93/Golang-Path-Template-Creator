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

func applyVerifierContinueOnErrorInternal(
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	filteredPathFileInfoMap *chmodhelper.FilteredPathFileInfoMap,
) (isSuccess bool) {
	if verifier == nil || filteredPathFileInfoMap == nil {
		return true
	}

	existingErrorCount := errCollection.Length()
	if filteredPathFileInfoMap.Error != nil {
		errCollection.AddTypeError(
			errtype.PathMissingOrInvalid,
			filteredPathFileInfoMap.Error)
	}

	validFileLocations := mics.GetLocationsUsingFilteredPathFileInfoMap(
		filteredPathFileInfoMap)

	// Rwx verify
	if verifier.HasRwxInstructions() {
		executors, err := chmodhelper.ParseBaseRwxInstructionsToExecutors(
			&verifier.BaseRwxInstructions)

		if err != nil {
			errCollection.AddTypeError(
				errtype.ParsingFailed,
				err)
		}

		errRwxInstructions := executors.VerifyRwxModifiers(
			true,
			true, // ignore recursive error
			validFileLocations)

		if errRwxInstructions != nil {
			errCollection.AddTypeError(
				errtype.RwxMismatch,
				errRwxInstructions)
		}
	}

	// immediately exit on error, user+groups verify.
	hasAnyUserGroupValidation := len(validFileLocations) > 0 &&
		verifier.UserGroupName.HasUserNameOrGroup()

	if hasAnyUserGroupValidation && osconsts.IsWindows {
		errCollection.AddUsingMsg(
			errtype.NotSupportInWindows,
			errtype.ChownUserOrGroupApplyIssue.VariantStructure().String()+
				" Cannot verify valid locations:"+
				strings.Join(validFileLocations, constants.CommaSpace))

		return false
	}

	// continue on error, user+groups
	for _, location := range validFileLocations {
		applyVerifierSinglePathNonRecursiveUserGroupVerify(
			verifier,
			errCollection,
			location,
		)
	}

	isSuccess = existingErrorCount ==
		errCollection.Length()

	return isSuccess
}
