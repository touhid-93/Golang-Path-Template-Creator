package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyUsingFlatPathsDirectReturn(
	isContinueOnError bool,
	verifiers *pathinsfmt.PathVerifiers,
	locations []string,
) *errorwrapper.Wrapper {
	errCollection := errwrappers.Empty()

	isSuccess := ApplyUsingFlatPaths(
		isContinueOnError,
		verifiers,
		errCollection,
		locations)

	if !isSuccess {
		return errCollection.GetAsErrorWrapperPtr()
	}

	return nil
}
