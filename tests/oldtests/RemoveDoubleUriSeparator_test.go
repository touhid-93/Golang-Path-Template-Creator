package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/pathhelper"
)

type RemoveDoubleUriSepratorTestCaseWrapper struct {
	inputPath, inputSeparator, expected, expectedMessage string
}

var RemoveDoubleUriSepratorTestCaseWrappers = []RemoveDoubleUriSepratorTestCaseWrapper{
	{
		inputPath:       "__hello___",
		inputSeparator:  constants.Underscore,
		expected:        "_hello_",
		expectedMessage: "non-empty, non-nil return of (_hello_)",
	},
	{
		inputPath:       "--hello---",
		inputSeparator:  constants.Underscore,
		expected:        "_hello_-",
		expectedMessage: "non-empty, non-nil return of (_hello_-)",
	},
	{
		inputPath:       "--hello---",
		inputSeparator:  constants.Dash,
		expected:        "-hello-",
		expectedMessage: "non-empty, non-nil return of (-hello-)",
	},
}

func TestRemoveDoubleUriSeparator(t *testing.T) {
	for i, testCase := range RemoveDoubleUriSepratorTestCaseWrappers {
		testCaseMessage := fmt.Sprintf("[RemoveDoubleUriSeparator] inputs (%s, %s) expects (%s)", testCase.inputPath, testCase.inputSeparator, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.RemoveDoubleUriSeparator(testCase.inputPath, testCase.inputSeparator)

			// Arrange
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldNotBeEmpty)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
