package pathmodifierverify

import (
	"gitlab.com/evatix-go/core/errcore"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"
	"gitlab.com/evatix-go/pathhelper/pathinsfmt"
	"gitlab.com/evatix-go/pathhelper/pathstatlinux"
)

// applyVerifierSinglePathNonRecursiveUserGroupVerify
func applyVerifierSinglePathNonRecursiveUserGroupVerify(
	verifier *pathinsfmt.PathVerifier,
	errCollection *errwrappers.Collection,
	location string,
) (isSuccess bool) {
	if verifier.IsGroupNameUserNameBothEmpty() {
		return true
	}

	errCount := errCollection.Length()
	pathStat := pathstatlinux.Get(location)

	if pathStat == nil || !pathStat.IsValidParsing {
		//goland:noinspection GoNilness
		errCollection.AddWrapperPtr(
			pathStat.ErrorWrapper)

		return false
	}

	verifyUsername := pathStat.User.Name
	if verifier.HasUserName() && !verifier.IsUsername(verifyUsername) {
		errCollection.AddPathIssueMessages(
			errtype.Unexpected,
			location,
			errcore.ExpectingSimpleNoType(
				"Username expectation doesn't meet",
				verifier.UserNameSimple(),
				verifyUsername))
	}

	verifyGroupName := pathStat.Group.Name
	if verifier.HasGroupName() && !verifier.IsGroupName(verifyGroupName) {
		errCollection.AddPathIssueMessages(
			errtype.Unexpected,
			location,
			errcore.ExpectingSimpleNoType(
				"Group expectation doesn't meet",
				verifier.GroupName,
				verifyGroupName))
	}

	return errCount == errCollection.Length()
}
