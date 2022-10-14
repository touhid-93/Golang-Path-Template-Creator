package oldtests

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/pathhelper"
)

type pathAsUriTestCaseWrapper struct {
	input, expected, expectedMessage string
	inputBool                        bool
}

var pathAsUriTestCaseWrappers = []pathAsUriTestCaseWrapper{
	{
		input:           "c:\\windows",
		inputBool:       true,
		expected:        "file:///c:/windows",
		expectedMessage: "file:///c:/windows",
	},
	{
		input:           "c:/windows",
		inputBool:       true,
		expected:        "file:///c:/windows",
		expectedMessage: "file:///c:/windows",
	},
}

func TestGetPathAsUri(t *testing.T) {
	for i, testCase := range pathAsUriTestCaseWrappers {
		// Arrange
		testCaseMessage := fmt.Sprintf("[GetPathAsUri] inputs (%s, %v) expects (%s)", testCase.input, testCase.inputBool, testCase.expectedMessage)

		Convey(testCaseMessage, t, func() {
			// Act
			actual := pathhelper.GetPathAsUri(testCase.input, testCase.inputBool)

			// Assert
			Convey(GetAssertMessage(actual, testCase.expected, i), func() {
				So(actual, ShouldNotBeNil)
				So(actual, ShouldEqual, testCase.expected)
			})
		})
	}
}
