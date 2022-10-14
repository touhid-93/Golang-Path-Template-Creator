package pathinsfmtexectests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"
	"gitlab.com/evatix-go/errorwrapper/errverify"
	"gitlab.com/evatix-go/errorwrapper/errwrappers"

	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifier"
	"gitlab.com/evatix-go/pathhelper/pathinsfmtexec/pathmodifierverify"
	"gitlab.com/evatix-go/pathhelper/tests/testwrappers/pathinsfmtexectestwrappers"
)

func Test_PathWithModifierAndVerifierErrors(t *testing.T) {
	coretests.SkipOnWindows(t)
	DefaultPathsSetup()

	for i, testCase := range pathinsfmtexectestwrappers.PathWithModifyAndVerifyTestCases {
		// Arrange
		errCollection := errwrappers.Empty()
		// Act
		pathmodifier.ApplyPathWithModifier(
			errCollection,
			testCase.Modifier)

		pathmodifierverify.ApplyPathWithVerifier(
			false,
			errCollection,
			testCase.Verifier,
		)

		errVerifyParams := &errverify.VerifyCollectionParams{
			CaseIndex:       i,
			FuncName:        testCase.ErrorVerify.FunctionName,
			TestCaseName:    testCase.ErrorVerify.Header,
			ErrorCollection: errCollection,
		}

		// Assert
		Convey(testCase.Header, t, func() {
			isSuccess := testCase.ErrorVerify.IsMatch(errVerifyParams)

			So(isSuccess, ShouldBeTrue)
		})
	}
}
