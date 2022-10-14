package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type combinedOfNonEmptyPathsTestCaseWrapper struct {
	inputSeparator, expected, expectedMessage string
	inputPaths                                []string
}

var combinedOfNonEmptyPathsTestCaseWrappers = []combinedOfNonEmptyPathsTestCaseWrapper{
	{
		inputSeparator:  "/",
		inputPaths:      []string{},
		expected:        "",
		expectedMessage: "empty string",
	},
	{
		inputSeparator:  "/",
		inputPaths:      []string{"home", "user", "", "etc"},
		expected:        "home/user/etc",
		expectedMessage: "home/user/etc",
	},
}

func TestGetCombinedOfNonEmptyPaths(t *testing.T) {
	for i, testCase := range combinedOfNonEmptyPathsTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetCombinedOfNonEmptyPaths] inputs (Separator: %s, paths: %s) expects (%s)", testCase.inputSeparator, testCase.inputPaths, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetCombinedOfNonEmptyPaths(testCase.inputSeparator, testCase.inputPaths)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
