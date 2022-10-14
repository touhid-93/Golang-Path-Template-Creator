package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type compliedPathTestCaseWrapper struct {
	inputPath, expectedMessage, expected string
	inputMap                             map[string]string
}

var compliedPathTestCaseWrappers = []compliedPathTestCaseWrapper{
	{
		inputPath:       "",
		inputMap:        map[string]string{"/": "\\"},
		expected:        "",
		expectedMessage: "non-nil return",
	},
	{
		inputPath:       "c//something//etc",
		inputMap:        map[string]string{"/": "\\"},
		expected:        "c\\\\something\\\\etc",
		expectedMessage: "non-nil return",
	},
}

func TestGetCompiledPath(t *testing.T) {
	for i, testCase := range compliedPathTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetCompiledPath] inputs (path: %s, map: %s) expects (%s)", testCase.inputPath, testCase.inputMap, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetCompiledPath(testCase.inputPath, &testCase.inputMap)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
