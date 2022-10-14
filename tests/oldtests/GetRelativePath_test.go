package oldtests

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type getRelativePathTestCaseWrapper struct {
	inputFullpath, inputBasepath, expected, expectedMessage string
}

var getRelativePathTestCaseWrappers = []getRelativePathTestCaseWrapper{
	{
		inputFullpath:   "c:\\Windows\\whatevr",
		inputBasepath:   "ex",
		expected:        "",
		expectedMessage: "should return empty string",
	},
	{
		inputFullpath:   "c:\\Windows\\whatevr",
		inputBasepath:   "c:\\Windows\\whatevr",
		expected:        "Both paths are same",
		expectedMessage: "Both paths are same",
	},
	{
		inputFullpath:   "c:\\Windows\\whatevr",
		inputBasepath:   "c:\\Windows",
		expected:        "\\whatevr",
		expectedMessage: "\\whatevr",
	},
}

func TestGetRelativePath(t *testing.T) {
	for i, testCase := range getRelativePathTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetRelativePath] inputs (%s, %s) expects (%s)", testCase.inputFullpath, testCase.inputBasepath, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetRelativePath(testCase.inputFullpath, testCase.inputBasepath)

			// Arrange
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				if !strings.Contains(testCase.inputFullpath, testCase.inputBasepath) {
					So(actual, ShouldBeEmpty)
				}

				if strings.Contains(testCase.inputFullpath, testCase.inputBasepath) {
					So(actual, ShouldNotBeEmpty)
					So(actual, ShouldEqual, testCase.expected)
				}
			})
		})
	}
}
