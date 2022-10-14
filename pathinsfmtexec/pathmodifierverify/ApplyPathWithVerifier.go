package pathmodifierverify

import (
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
)

func ApplyPathWithVerifier(
	isContinueOnError bool,
	errCollection *errwrappers.Collection,
	pathWithVerifier *pathinsfmt.PathWithVerifier,
) (isSuccess bool) {
	if pathWithVerifier.IsVerifierUndefined() {
		return true
	}

	return ApplyVerifier(
		pathWithVerifier.IsNormalize,
		pathWithVerifier.IsRecursive,
		pathWithVerifier.IsSkipInvalid,
		isContinueOnError,
		pathWithVerifier.Verifier,
		errCollection,
		pathWithVerifier.CompiledPathAsSlice())
}
