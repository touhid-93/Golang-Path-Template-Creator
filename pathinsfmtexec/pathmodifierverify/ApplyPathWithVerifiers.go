package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyPathWithVerifiers(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	pathsWithVerifiers *pathinsfmt.PathsWithVerifiers,
) (isSuccess bool) {
	if pathsWithVerifiers.IsPathWithVerifiersUndefined() {
		return true
	}

	stateTracker := errCollection.StateTracker()

	for i := range pathsWithVerifiers.Verifiers {
		isSuccess = ApplyPathWithVerifier(
			isContinueOnError,
			errCollection,
			&pathsWithVerifiers.Verifiers[i])

		if !isContinueOnError && !isSuccess {
			return false
		}
	}

	return stateTracker.IsSuccess()
}
